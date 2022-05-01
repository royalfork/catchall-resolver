// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "./openzeppelin/contracts/utils/introspection/IERC165.sol";
import "./ens-contracts/contracts/registry/ENS.sol";
import "./ens-contracts/contracts/resolvers/profiles/IABIResolver.sol";
import "./ens-contracts/contracts/resolvers/profiles/IAddressResolver.sol";
import "./ens-contracts/contracts/resolvers/profiles/IAddrResolver.sol";
import "./ens-contracts/contracts/resolvers/profiles/IContentHashResolver.sol";
import "./ens-contracts/contracts/resolvers/profiles/IDNSRecordResolver.sol";
import "./ens-contracts/contracts/resolvers/profiles/IDNSZoneResolver.sol";
import "./ens-contracts/contracts/resolvers/profiles/IInterfaceResolver.sol";
import "./ens-contracts/contracts/resolvers/profiles/INameResolver.sol";
import "./ens-contracts/contracts/resolvers/profiles/IPubkeyResolver.sol";
import "./ens-contracts/contracts/resolvers/profiles/ITextResolver.sol";
import "./ens-contracts/contracts/resolvers/profiles/IExtendedResolver.sol";

interface Resolver is
    IERC165,
    IABIResolver,
    IAddressResolver,
    IAddrResolver,
    IContentHashResolver,
    IDNSRecordResolver,
    IDNSZoneResolver,
    IInterfaceResolver,
    INameResolver,
    IPubkeyResolver,
    ITextResolver {}

/**
 * @title Catch-all ENSIP-10 resolver.
 * @author royalfork.eth
 * @notice ENS resolver which proxies all resolver functions for any
 *         subdomain of a node to a set resolver.
 */
contract CatchallResolver is IExtendedResolver, Resolver {
    ENS public immutable registry;

    mapping(bytes32=>Resolver) resolvers;

    event NewCatchallResolver(bytes32 indexed node, address resolver);

    constructor(address _registry) {
        registry = ENS(_registry);
    }

    /**
     * @notice Set catchall resolver for node.  Node and all of its
     *         subdomains will use this resolver.  Message sender must
     *         be owner or operator for node.
     * @param _node Catchall resolver will be set for this node.
     * @param _resolver Catchall reoslver proxies all resolver
     *        functions to this address.
	 */
	function setResolver(bytes32 _node, Resolver _resolver) public {
        address owner = registry.owner(_node);
        require(owner == msg.sender || registry.isApprovedForAll(owner, msg.sender));
		resolvers[_node] = _resolver;
		emit NewCatchallResolver(_node, address(_resolver));
	}

    /**
	 * @notice ENSIP-10 defined wildcard resolution function.
     * @dev Resolve only works with resolver functions where the first
     *      argument is a bytes32 node (as of ENSIP-12, all resolver
     *      functions meet this criteria).
	 * @param _name DNS-encoded name to be resolved.
	 * @param _data ABI-encoded calldata for a resolver function.
	 * @return output ABI-encoded return output from function encoded
	 *         by _data.
	 */
    function resolve(bytes calldata _name, bytes memory _data) external override view returns(bytes memory) {
		(address r,,bytes32 node,) = resolver(_name, 0);

		// Replace node argument in data with parentNode
		for (uint8 i = 0; i < 32; i++) {
			_data[i+4] = node[i];
		}

		(bool ok, bytes memory out) = r.staticcall(_data);
		if (!ok) {
			revert("invalid call");
		}
		return out;
    }

    /**
     * @notice Returns ENSIP-10 resolver for name.
     * @param _name The name to resolve, in normalised and DNS-encoded form (eg: sub.example.eth)
     * @return resolverAddr Found resolver for name.
	 * @return resolverOwner DNS-encoded name which set the resolver (eg: example.eth).
	 */
	function resolver(bytes calldata _name) public view returns(address, bytes memory) {
		(address r,uint256 o,,) = resolver(_name, 0);
		return (r, _name[o:]);
	}

    function supportsInterface(bytes4 interfaceID) public pure override returns(bool) {
        return interfaceID == type(IExtendedResolver).interfaceId ||
			interfaceID == type(IERC165).interfaceId ||
			interfaceID == type(IABIResolver).interfaceId ||
			interfaceID == type(IAddressResolver).interfaceId ||
			interfaceID == type(IAddrResolver).interfaceId ||
			interfaceID == type(IContentHashResolver).interfaceId ||
			interfaceID == type(IDNSRecordResolver).interfaceId ||
			interfaceID == type(IDNSZoneResolver).interfaceId ||
			interfaceID == type(IInterfaceResolver).interfaceId ||
			interfaceID == type(INameResolver).interfaceId ||
			interfaceID == type(IPubkeyResolver).interfaceId ||
			interfaceID == type(ITextResolver).interfaceId;
    }

	function ABI(bytes32 node, uint256 contentTypes) external override view returns (uint256, bytes memory) {
		return resolvers[node].ABI(node, contentTypes);
	}
	function addr(bytes32 node) external override view returns (address payable) {
		return resolvers[node].addr(node);
	}
	function addr(bytes32 node, uint coinType) external override view returns(bytes memory) {
		return resolvers[node].addr(node, coinType);
	}
	function contenthash(bytes32 node) external override view returns (bytes memory) {
		return resolvers[node].contenthash(node);
	}
	function dnsRecord(bytes32 node, bytes32 _name, uint16 resource) external override view returns (bytes memory) {
		return resolvers[node].dnsRecord(node, _name, resource);
	}
	function interfaceImplementer(bytes32 node, bytes4 interfaceID) external override view returns (address) {
		return resolvers[node].interfaceImplementer(node, interfaceID);
	}
	function name(bytes32 node) external override view returns (string memory) {
		return resolvers[node].name(node);
	}
	function pubkey(bytes32 node) external override view returns (bytes32 x, bytes32 y) {
		return resolvers[node].pubkey(node);
	}
	function text(bytes32 node, string calldata key) external override view returns (string memory) {
		return resolvers[node].text(node, key);
	}
	function zonehash(bytes32 node) external override view returns (bytes memory) {
		return resolvers[node].zonehash(node);
	}

    /**
     * @dev Performs recursive ENSIP-10 lookup for a catchall resolver.
     * @param _name The name to resolve, in normalised and DNS-encoded
     *        form (eg: sub.example.eth)
     * @param _offset The offset within name on which a catchall
     *        resolver lookup is performed.
     * @return resolverAddr Found resolver for name.
     * @return resolverNameOffset Domain at this offset within name
     *         which set resolverAddr (eg: offset of example.eth).
     * @return resolverNode Namehash of name[resolverNameOffset:].
     * @return node Namehash of name.
     */
    function resolver(bytes calldata _name, uint256 _offset) internal view returns(address, uint256, bytes32, bytes32) {
        uint256 labelLength = uint256(uint8(_name[_offset]));
        if(labelLength == 0) {
            return (address(0), _name.length, bytes32(0), bytes32(0));
        }
        uint256 nextLabel = _offset + labelLength + 1;
        bytes32 labelHash = keccak256(_name[_offset + 1: nextLabel]);
        (address r, uint256 roffset, bytes32 rnode, bytes32 parentnode) = resolver(_name, nextLabel);
        bytes32 node = keccak256(abi.encodePacked(parentnode, labelHash));
        address newr = address(resolvers[node]);
        if(newr != address(0)) {
            return (newr, _offset, node, node);
        }
        return (r, roffset, rnode, node);
    }
}
