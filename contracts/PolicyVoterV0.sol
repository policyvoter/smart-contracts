// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

contract PolicyVoterV0 is ReentrancyGuard, AccessControl {
	using Address for address;
	using EnumerableSet for EnumerableSet.AddressSet;

	//blacklist can't vote or propose new policies
	bytes32 public constant POLICY_CREATOR_ROLE = keccak256("POLICY_CREATOR_ROLE");
	bytes32 public constant ADDRESS_REGISTER_ROLE = keccak256("ADDRESS_REGISTER_ROLE");
	bytes32 public constant POLICY_BLACKLISTER_ROLE = keccak256("POLICY_BLACKLISTER_ROLE");

	string public version = "0.0.4";

	struct Policy {
		string bodyHash;
		uint256 createdAt;
	}

	mapping(string => Policy) public policies; //policyID - Policy
	mapping(string => uint256) public votes; //policyID - votes
	mapping(address => uint256) public lastVoted; //prevents spamming
	mapping(bytes32 => bool) public votedOnProposal; //voted on proposal or not
	mapping(address => bool) public registered; //registers an address with the system

	uint256 public minIntervalBetweenVotes = 3 seconds;

	//blacklists
	mapping(string => bool) public isPolicyBlacklisted; //can't vote on it
	mapping(string => string) public policyBlacklistedReason; //reason why it was blacklisted

	event NewPolicy(string policyID);
	event Voted(string policyID, uint256 totalVotes);
	event UnVoted(string policyID, uint256 totalVotes);
	event Blacklisted(address user, bool blacklisted);
	event PolicyBlacklisted(string policyID, bool blacklisted, string reason);

	constructor() {
		_setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
		_setupRole(POLICY_CREATOR_ROLE, msg.sender);
		_setupRole(POLICY_BLACKLISTER_ROLE, msg.sender);
		_setupRole(ADDRESS_REGISTER_ROLE, msg.sender);
	}

	/**
		@notice registers/unregisters an address with the system
		@param newAddress - user's address
	 */
	function registerAddress(address newAddress, bool isRegistered)
		external
		onlyRole(ADDRESS_REGISTER_ROLE)
	{
		registered[newAddress] = isRegistered;
	}

	/**
		@notice registers/unregisters an address with the system
		@param newMinInterval - new interval in seconds
	 */
	function setMinIntervalBetweenVotes(uint256 newMinInterval)
		external
		onlyRole(DEFAULT_ADMIN_ROLE)
	{
		minIntervalBetweenVotes = newMinInterval;
	}

	/**
		@notice creates a new policy
		@param policyID - the policy ID from the database
		@param bodyHash - the hash of this policy body
	 */
	function createNewPolicy(string memory policyID, string memory bodyHash)
		external
		onlyRole(POLICY_CREATOR_ROLE)
	{
		require(bytes(policies[policyID].bodyHash).length == 0, "policyID already exists");
		policies[policyID] = Policy({ bodyHash: bodyHash, createdAt: block.timestamp });
		emit NewPolicy(policyID);
	}

	/**
		@notice votes on a policy
		@param policyID - the policy ID from the database
	 */
	function vote(string memory policyID) external nonReentrant {
		require(tx.origin == msg.sender, "no contracts allowed");
		require(lastVoted[msg.sender] < block.timestamp - minIntervalBetweenVotes, "spam vote");
		require(!isPolicyBlacklisted[policyID], "policy is blacklisted");

		require(policies[policyID].createdAt != 0, "policy not found");
		require(registered[msg.sender], "address not registered");

		bytes32 keyVoted = computeKeyVotedOnProposal(msg.sender, policyID);
		require(votedOnProposal[keyVoted] == false, "already voted");

		votedOnProposal[keyVoted] = true;

		votes[policyID] = votes[policyID] + 1;
		lastVoted[msg.sender] = block.timestamp;

		emit Voted(policyID, votes[policyID]);
	}

	/**
		@notice removes a vote from a policy
		@param policyID - the policy ID from the database
	 */
	function unVote(string memory policyID) external nonReentrant {
		require(tx.origin == msg.sender, "no contracts allowed");
		require(!isPolicyBlacklisted[policyID], "policy is blacklisted");
		require(registered[msg.sender], "address not registered");

		bytes32 keyVoted = computeKeyVotedOnProposal(msg.sender, policyID);
		require(votedOnProposal[keyVoted] == true, "didn't voted");

		votes[policyID] = votes[policyID] - 1;
		delete votedOnProposal[keyVoted];

		emit UnVoted(policyID, votes[policyID]);
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

	/**
	 * @dev Computes the key for votedOnProposal
	 */
	function computeKeyVotedOnProposal(address voter, string memory policyID)
		internal
		pure
		virtual
		returns (bytes32)
	{
		return keccak256(abi.encodePacked(voter, policyID));
	}
}
