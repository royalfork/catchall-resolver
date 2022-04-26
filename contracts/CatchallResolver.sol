// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "./openzeppelin/contracts/utils/introspection/IERC165.sol";
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

contract CatchallResolver is IExtendedResolver, Resolver {
	Resolver public parent;
	bytes32 public parentNode;

    constructor(Resolver _parent, bytes32 _node) {
		parent = _parent;
		parentNode = _node;
    }

	// resolve only works with resolver functions where the first argument is a bytes32 node.
    function resolve(bytes calldata, bytes memory data) external override view returns(bytes memory) {
		// Replace node argument in data with parentNode
		for (uint8 i = 0; i < 32; i++) {
			data[i+4] = parentNode[i];
		}

		(bool ok, bytes memory out) = address(parent).staticcall(data);
		if (!ok) {
			revert("invalid call");
		}
		return out;
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

	// Resolver functions are proxied to parent resolver.
	function ABI(bytes32 node, uint256 contentTypes) external view returns (uint256, bytes memory) {
		return parent.ABI(node, contentTypes);
	}
	function addr(bytes32 node) external view returns (address payable) {
		return parent.addr(node);
	}
	function addr(bytes32 node, uint coinType) external view returns(bytes memory) {
		return parent.addr(node, coinType);
	}
	function contenthash(bytes32 node) external view returns (bytes memory) {
		return parent.contenthash(node);
	}
	function dnsRecord(bytes32 node, bytes32 _name, uint16 resource) external view returns (bytes memory) {
		return parent.dnsRecord(node, _name, resource);
	}
	function interfaceImplementer(bytes32 node, bytes4 interfaceID) external view returns (address) {
		return parent.interfaceImplementer(node, interfaceID);
	}
	function name(bytes32 node) external view returns (string memory) {
		return parent.name(node);
	}
	function pubkey(bytes32 node) external view returns (bytes32 x, bytes32 y) {
		return parent.pubkey(node);
	}
	function text(bytes32 node, string calldata key) external view returns (string memory) {
		return parent.text(node, key);
	}
	function zonehash(bytes32 node) external view returns (bytes memory) {
		return parent.zonehash(node);
	}
}
