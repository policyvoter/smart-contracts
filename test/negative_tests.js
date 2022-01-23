const { expect, assert } = require("chai")
const { web3, ethers } = require("hardhat")
const { BN, time, balance, expectEvent, expectRevert } = require("@openzeppelin/test-helpers")
const ether = require("@openzeppelin/test-helpers/src/ether")

let pv
let owner, acc1, acc2, acc3

describe("PolicyVoter Negative Tests", function () {
	beforeEach(async function () {
		let TContract = await ethers.getContractFactory("PolicyVoterV0")

		pv = await TContract.deploy()
		await pv.deployed()

		signers = await ethers.getSigners()
		owner = signers[0]
		acc1 = signers[1]
		acc2 = signers[2]
		acc3 = signers[3]
	})

	it("simple test...", async function () {
		//policy creator role
		let hasRole = await pv.hasRole(
			"0x90fe2ba5da14f172ed5a0a0fec391dbf8f191c9a2f3557d79ede5d6b1c1c9ffb",
			owner.address
		)
		expect(hasRole === true)
	})

	it("can't policies if you're not POLICY_CREATOR_ROLE", async function () {
		await expect(
			pv
				.connect(acc1)
				.createNewPolicy(
					acc1.address,
					"a1b2c3",
					"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
				)
		).to.to.be.revertedWith(
			"account 0x17d4f1c23e45312b4ad8251b293b1afc3d3271a4 is missing role 0x90fe2ba5da14f172ed5a0a0fec391dbf8f191c9a2f3557d79ede5d6b1c1c9ffb"
		)
	})

	it("can't create the same policy two times", async function () {
		await expect(
			pv.createNewPolicy(
				acc1.address,
				"a1b2c3",
				"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
			)
		).to.emit(pv, "NewPolicy")

		await expect(
			pv.createNewPolicy(
				acc1.address,
				"a1b2c3",
				"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
			)
		).to.to.be.revertedWith("policyID already exists")
	})

	it("can't create the a policy by a blacklisted creator", async function () {
		await pv.setBlacklist(acc1.address, true)
		await expect(
			pv.createNewPolicy(
				acc1.address,
				"a1b2c3",
				"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
			)
		).to.be.revertedWith("account is blacklisted")
	})

	// function vote(string memory policyID) external nonReentrant {
	// 	require(tx.origin == msg.sender, "no contracts allowed");
	// 	require(lastVoted[msg.sender] < block.timestamp - 30, "30 seconds between votes");
	// 	require(!isBlacklisted[msg.sender], "account is blacklisted");
	// 	require(bytes(votedOnProposal[msg.sender]).length == 0, "already voted");

	// 	votedOnProposal[msg.sender] = policyID;

	// 	votes[policyID] = votes[policyID] + 1;
	// 	lastVoted[msg.sender] = block.timestamp;

	// 	emit Voted(policyID);
	// }
	it("can't vote too fast", async function () {
		await expect(
			pv.createNewPolicy(
				acc1.address,
				"a1b2c3",
				"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
			)
		).to.emit(pv, "NewPolicy")

		await expect(pv.connect(acc1).vote("a1b2c3")).to.emit(pv, "Voted")
		await expect(pv.connect(acc1).vote("a1b2c3")).to.be.revertedWith("30 seconds between votes")
	})

	it("can't vote two times", async function () {
		await expect(
			pv.createNewPolicy(
				acc1.address,
				"a1b2c3",
				"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
			)
		).to.emit(pv, "NewPolicy")

		await expect(pv.connect(acc1).vote("a1b2c3")).to.emit(pv, "Voted")
		await waitForSeconds(31)
		await expect(pv.connect(acc1).vote("a1b2c3")).to.be.revertedWith("already voted")
	})

	it("can't vote if blacklisted", async function () {
		await expect(
			pv.createNewPolicy(
				acc1.address,
				"a1b2c3",
				"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
			)
		).to.emit(pv, "NewPolicy")

		await pv.setBlacklist(acc2.address, true)
		await expect(pv.connect(acc2).vote("a1b2c3")).to.be.revertedWith("account is blacklisted")
	})

	// 	/**
	// 	@notice removes a vote from a policy
	//  */
	// function unVote(string memory policyID) external nonReentrant {
	// 	require(tx.origin == msg.sender, "no contracts allowed");
	// 	require(bytes(votedOnProposal[msg.sender]).length > 0, "not voted");

	// 	votes[policyID] = votes[policyID] - 1;
	// 	delete votedOnProposal[msg.sender];

	// 	emit UnVoted(policyID);
	// }

	it("can't unvote a not voted project", async function () {
		await expect(
			pv.createNewPolicy(
				acc1.address,
				"a1b2c3",
				"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
			)
		).to.emit(pv, "NewPolicy")
		await expect(pv.connect(acc2).unVote("a1b2c3")).to.be.revertedWith("not voted")
	})

	it("can't *UNVOTE* multiple times", async function () {
		await expect(
			pv.createNewPolicy(
				acc1.address,
				"a1b2c3",
				"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
			)
		).to.emit(pv, "NewPolicy")

		await expect(pv.connect(acc1).vote("a1b2c3")).to.emit(pv, "Voted")
		await expect(pv.connect(acc1).unVote("a1b2c3")).to.emit(pv, "UnVoted")

		await expect(pv.connect(acc2).unVote("a1b2c3")).to.be.revertedWith("not voted")
	})

	it("can't unvote from a blacklisted account", async function () {
		await expect(
			pv.createNewPolicy(
				acc1.address,
				"a1b2c3",
				"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
			)
		).to.emit(pv, "NewPolicy")
		await expect(pv.connect(acc2).vote("a1b2c3")).to.emit(pv, "Voted")
		await pv.setBlacklist(acc2.address, true)
		await expect(pv.connect(acc2).unVote("a1b2c3")).to.be.revertedWith("account is blacklisted")
	})

	it("can't vote or unvote a blacklisted policy", async function () {
		await expect(
			pv.createNewPolicy(
				acc1.address,
				"a1b2c3",
				"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
			)
		).to.emit(pv, "NewPolicy")
		await expect(pv.connect(acc2).vote("a1b2c3")).to.emit(pv, "Voted")
		await pv.setPolicyBlacklist("a1b2c3", true, "illegal policy")

		await expect(pv.connect(acc1).vote("a1b2c3")).to.be.revertedWith("policy is blacklisted")
		await expect(pv.connect(acc2).unVote("a1b2c3")).to.be.revertedWith("policy is blacklisted")
	})

	async function waitForSeconds(seconds) {
		await ethers.provider.send("evm_increaseTime", [seconds])
		await ethers.provider.send("evm_mine")
	}

	async function getBlockTimestamp() {
		block = await web3.eth.getBlock(await web3.eth.getBlockNumber())
		return block["timestamp"]
	}
})
