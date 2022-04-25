// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "./ens-contracts/ISupportsInterface.sol";
import "./ens-contracts/profiles/IAddrResolver.sol";
import "./ens-contracts/profiles/IExtendedResolver.sol";

interface Resolver is IAddrResolver, ISupportsInterface {}

contract CatchallResolver is IExtendedResolver, Resolver {
	Resolver public parent;
	bytes32 public parentNode;

    constructor(Resolver _parent, bytes32 _node) {
		parent = _parent;
		parentNode = _node;
    }

	function addr(bytes32 node) external view returns (address payable) {
		return parent.addr(node);
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
		// TODO add IAddrResolver interfaceID support
        return interfaceID == type(IExtendedResolver).interfaceId;
    }
}
