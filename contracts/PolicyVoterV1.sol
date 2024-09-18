// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";

contract PolicyVoterV1 is ReentrancyGuard, AccessControl {
  using Address for address;

  //blacklist can't vote or propose new policies
  bytes32 public constant POLICY_CREATOR_ROLE =
    keccak256("POLICY_CREATOR_ROLE");
  bytes32 public constant ADDRESS_REGISTER_ROLE =
    keccak256("ADDRESS_REGISTER_ROLE");
  bytes32 public constant POLICY_BLACKLISTER_ROLE =
    keccak256("POLICY_BLACKLISTER_ROLE");

  string public version = "0.0.6";

  struct Policy {
    string bodyHash;
    uint256 groupID;
    uint256 createdAt;
  }

  mapping(string => Policy) public policies; //policyID - Policy
  mapping(string => uint256) public votes; //policyID - votes
  mapping(address => uint256) public lastVoted; //prevents spamming
  mapping(bytes32 => bool) public votedOnProposal; //voted on proposal or not
  mapping(address => bool) public registered; //registers an address with the system

  uint256 public minIntervalBetweenVotes = 1 seconds; //block time is also 3 seconds, lower doesn't matter

  //blacklists
  mapping(string => bool) public isPolicyBlacklisted; //can't vote on it
  mapping(string => string) public policyBlacklistedReason; //reason why it was blacklisted
  mapping(address => uint256) public groupForAddress; //the group that this address belongs to (0 is public)

  event NewPolicy(string policyID, uint256 groupID);
  event Voted(string policyID, uint256 groupID, uint256 totalVotes);
  event UnVoted(string policyID, uint256 groupID, uint256 totalVotes);
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
  @param isRegistered - set to true if it is
  @param groupID - the pilot ID, set to 0 for no pilot
  */
  function registerAddress(
    address newAddress,
    bool isRegistered,
    uint256 groupID
  ) external onlyRole(ADDRESS_REGISTER_ROLE) {
    registered[newAddress] = isRegistered;
    groupForAddress[newAddress] = groupID;
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
  @param gID - the group ID
  @param bodyHash - the hash of this policy body
 */
  function createNewPolicy(
    string memory policyID,
    uint256 gID,
    string memory bodyHash
  ) external onlyRole(POLICY_CREATOR_ROLE) {
    require(
      bytes(policies[policyID].bodyHash).length == 0,
      "policyID already exists"
    );

    policies[policyID] = Policy({
      bodyHash: bodyHash,
      groupID: gID,
      createdAt: block.timestamp
    });
    emit NewPolicy(policyID, gID);
  }

  /**
  @notice votes on a policy
  @param policyID - the policy ID from the database
  */
  function vote(string memory policyID) external nonReentrant {
    require(tx.origin == msg.sender, "no contracts allowed");
    require(
      lastVoted[msg.sender] < block.timestamp - minIntervalBetweenVotes,
      "spam vote"
    );
    require(!isPolicyBlacklisted[policyID], "policy is blacklisted");

    require(policies[policyID].createdAt != 0, "policy not found");

    require(
      policies[policyID].groupID == groupForAddress[msg.sender],
      "you can only vote policies in your group"
    );

    require(registered[msg.sender], "address not registered");

    bytes32 keyVoted = computeKeyVotedOnProposal(msg.sender, policyID);
    require(votedOnProposal[keyVoted] == false, "already voted");

    votedOnProposal[keyVoted] = true;

    votes[policyID] = votes[policyID] + 1;
    lastVoted[msg.sender] = block.timestamp;

    emit Voted(policyID, policies[policyID].groupID, votes[policyID]);
  }

  /**
   @notice removes a vote from a policy
   @param policyID - the policy ID from the database
  */
  function unVote(string memory policyID) external nonReentrant {
    require(tx.origin == msg.sender, "no contracts allowed");
    require(!isPolicyBlacklisted[policyID], "policy is blacklisted");

    require(
      policies[policyID].groupID == groupForAddress[msg.sender],
      "you can only unvote policies in your group"
    );

    require(registered[msg.sender], "address not registered");

    bytes32 keyVoted = computeKeyVotedOnProposal(msg.sender, policyID);
    require(votedOnProposal[keyVoted] == true, "didn't voted");

    votes[policyID] = votes[policyID] - 1;
    delete votedOnProposal[keyVoted];

    emit UnVoted(policyID, policies[policyID].groupID, votes[policyID]);
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
