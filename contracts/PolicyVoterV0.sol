// SPDX-License-Identifier: MIT
pragma solidity 0.8.11;

import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

contract PolicyVoterV0 is ReentrancyGuard, AccessControl {
	using Address for address;
	using EnumerableSet for EnumerableSet.AddressSet;

	//blacklist can't vote or propose new policies
	bytes32 public constant POLICY_CREATOR_ROLE = keccak256("POLICY_CREATOR_ROLE");
	bytes32 public constant BLACKLISTER_ROLE = keccak256("BLACKLISTER_ROLE");
	bytes32 public constant POLICY_BLACKLISTER_ROLE = keccak256("POLICY_BLACKLISTER_ROLE");

	string public version = "0.0.1";

	struct Policy {
		address creator;
		string bodyHash;
		uint256 createdAt;
	}

	mapping(string => Policy) public policies; //policyID - Policy
	mapping(string => uint256) public votes; //policyID - votes
	mapping(address => uint256) public lastVoted; //prevents some spamming
	mapping(address => string) public votedOnProposal; //voted on proposal or not

	//blacklists
	mapping(address => bool) public isBlacklisted; //can't vote or propose
	mapping(string => bool) public isPolicyBlacklisted; //can't vote on it
	mapping(string => string) public policyBlacklistedReason; //reason why it was blacklisted

	event NewPolicy(string indexed policyID);
	event Voted(string indexed policyID, uint256 totalVotes);
	event UnVoted(string indexed policyID, uint256 totalVotes);
	event Blacklisted(address indexed policyID, bool blacklisted);
	event PolicyBlacklisted(string indexed policyID, bool blacklisted, string reason);

	constructor() {
		_setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
		_setupRole(POLICY_CREATOR_ROLE, msg.sender);
		_setupRole(BLACKLISTER_ROLE, msg.sender);
		_setupRole(POLICY_BLACKLISTER_ROLE, msg.sender);
	}

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
	) external onlyRole(POLICY_CREATOR_ROLE) {
		require(!isBlacklisted[creator], "account is blacklisted");
		require(bytes(policies[policyID].bodyHash).length == 0, "policyID already exists");
		policies[policyID] = Policy({
			creator: creator,
			bodyHash: bodyHash,
			createdAt: block.timestamp
		});
		emit NewPolicy(policyID);
	}

	/**
		@notice votes on a policy
	 */
	function vote(string memory policyID) external nonReentrant {
		require(tx.origin == msg.sender, "no contracts allowed");
		require(lastVoted[msg.sender] < block.timestamp - 30, "30 seconds between votes");
		require(!isBlacklisted[msg.sender], "account is blacklisted");
		require(!isPolicyBlacklisted[policyID], "policy is blacklisted");
		require(bytes(votedOnProposal[msg.sender]).length == 0, "already voted");
		require(policies[policyID].creator != address(0), "policy not found");

		votedOnProposal[msg.sender] = policyID;

		votes[policyID] = votes[policyID] + 1;
		lastVoted[msg.sender] = block.timestamp;

		emit Voted(policyID, votes[policyID]);
	}

	/**
		@notice removes a vote from a policy
	 */
	function unVote(string memory policyID) external nonReentrant {
		require(tx.origin == msg.sender, "no contracts allowed");
		require(bytes(votedOnProposal[msg.sender]).length > 0, "not voted");
		require(!isBlacklisted[msg.sender], "account is blacklisted");
		require(!isPolicyBlacklisted[policyID], "policy is blacklisted");

		votes[policyID] = votes[policyID] - 1;
		delete votedOnProposal[msg.sender];

		emit UnVoted(policyID, votes[policyID]);
	}

	/**
		@notice adds/removes an address from blacklist
	 */
	function setBlacklist(address account, bool blacklisted) external onlyRole(BLACKLISTER_ROLE) {
		isBlacklisted[account] = blacklisted;
		emit Blacklisted(account, blacklisted);
	}

	/**
		@notice adds/removes a policy from blacklist
	 */
	function setPolicyBlacklist(
		string memory policyID,
		bool blacklisted,
		string memory reason
	) external onlyRole(POLICY_BLACKLISTER_ROLE) {
		isPolicyBlacklisted[policyID] = blacklisted;
		policyBlacklistedReason[policyID] = reason;
		emit PolicyBlacklisted(policyID, blacklisted, reason);
	}
}
