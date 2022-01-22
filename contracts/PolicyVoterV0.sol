// SPDX-License-Identifier: MIT
pragma solidity 0.8.11;

import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

contract PolicyVoterV0 is Ownable, ReentrancyGuard {
	using Address for address;

	struct Policy {
		address creator;
		string bodyHash;
	}

	address public policyCreator;

	mapping(string => Policy) public policies; //policyID - Policy
	mapping(string => uint256) public votes; //policyID - votes
	mapping(address => uint256) public lastVoted; //prevents some spamming

	event Voted(string indexed policyID);
	event UnVoted(string indexed policyID);

	constructor() {}

	/**
		@notice creates a new policy
		@param creator - the creator of the policy
		@param policyID - the policy ID from the database
		@param bodyHash - the hash of this policy body
	 */
	function createNewPolicy(
		address creator,
		string memory policyID,
		string memory bodyHash
	) external nonReentrant {
		require(msg.sender == policyCreator, "only policy creator can create them");
		policies[policyID] = Policy({ creator: creator, bodyHash: bodyHash });
	}

	/**
		@notice votes on a policy
	 */
	function vote(string memory policyID) external {
		require(tx.origin == msg.sender, "no contracts allowed");
		require(lastVoted[msg.sender] < block.timestamp - 30, "30 seconds between votes");

		//require not voted already
		//TODO:

		votes[policyID] = votes[policyID] + 1;
		lastVoted[msg.sender] = block.timestamp;

		emit Voted(policyID);
	}

	/**
		@notice removes a vote from a policy
	 */
	function removeVote(string memory policyID) external nonReentrant {
		require(tx.origin == msg.sender, "no contracts allowed");

		//require not voted already
		//TODO:

		votes[policyID] = votes[policyID] - 1;

		emit UnVoted(policyID);
	}

	/**
		@notice updates the policy creator address
	 */
	function setPolicyCreator(address newPolicyCreator) external onlyOwner {
		policyCreator = newPolicyCreator;
	}
}
