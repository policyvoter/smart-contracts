const { expect, assert } = require("chai")
const { web3, ethers } = require("hardhat")
const { BN, time, balance, expectEvent, expectRevert } = require("@openzeppelin/test-helpers")
const ether = require("@openzeppelin/test-helpers/src/ether")

let pv
let owner, acc1, acc2, acc3

describe("PolicyVoter Tests", function () {
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
		let version = await pv.version()
		expect(String(version).length > 2)
	})

	it("simple create policies", async function () {
		await expect(
			pv.createNewPolicy(
				acc1.address,
				"a1b2c3",
				"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
			)
		).to.emit(pv, "NewPolicy")
	})

	it("can create policy if account is removed from blacklist", async function () {
		await pv.setBlacklist(acc1.address, true)
		await expect(
			pv.createNewPolicy(
				acc1.address,
				"a1b2c3",
				"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
			)
		).to.be.revertedWith("account is blacklisted")
		await pv.setBlacklist(acc1.address, false)
		await expect(
			pv.createNewPolicy(
				acc1.address,
				"a1b2c3",
				"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
			)
		).to.emit(pv, "NewPolicy")
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
	it("simple vote", async function () {
		await expect(
			pv.createNewPolicy(
				acc1.address,
				"a1b2c3",
				"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
			)
		).to.emit(pv, "NewPolicy")
		await expect(pv.vote("a1b2c3")).to.emit(pv, "Voted")
	})

	it("votes increase ok", async function () {
		await expect(
			pv.createNewPolicy(
				acc1.address,
				"a1b2c3",
				"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
			)
		).to.emit(pv, "NewPolicy")

		await expect(pv.vote("a1b2c3")).to.emit(pv, "Voted")
		await expect(pv.connect(acc1).vote("a1b2c3")).to.emit(pv, "Voted")
		await expect(pv.connect(acc2).vote("a1b2c3")).to.emit(pv, "Voted")

		expect(Number(await pv.votes("a1b2c3"))).to.equal(3)
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

	it("votes decrease ok", async function () {
		await expect(
			pv.createNewPolicy(
				acc1.address,
				"a1b2c3",
				"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
			)
		).to.emit(pv, "NewPolicy")

		await expect(pv.vote("a1b2c3")).to.emit(pv, "Voted")
		await expect(pv.connect(acc1).vote("a1b2c3")).to.emit(pv, "Voted")
		await expect(pv.connect(acc2).vote("a1b2c3")).to.emit(pv, "Voted")
		await expect(pv.connect(acc1).unVote("a1b2c3")).to.emit(pv, "UnVoted")

		expect(Number(await pv.votes("a1b2c3"))).to.equal(2)
	})

	it("can vote or unvote a removed from blacklisted policy", async function () {
		await expect(
			pv.createNewPolicy(
				acc1.address,
				"a1b2c3",
				"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
			)
		).to.emit(pv, "NewPolicy")
		await expect(pv.vote("a1b2c3")).to.emit(pv, "Voted")

		await expect(pv.connect(acc2).vote("a1b2c3")).to.emit(pv, "Voted")
		await pv.setPolicyBlacklist("a1b2c3", true, "illegal policy")

		await expect(pv.connect(acc1).vote("a1b2c3")).to.be.revertedWith("policy is blacklisted")
		await expect(pv.connect(acc2).unVote("a1b2c3")).to.be.revertedWith("policy is blacklisted")
		await waitForSeconds(31)
		await pv.setPolicyBlacklist("a1b2c3", false, "policy removed from blacklist")
		await expect(pv.connect(acc2).unVote("a1b2c3")).to.emit(pv, "UnVoted")
		await expect(pv.connect(acc1).vote("a1b2c3")).to.emit(pv, "Voted")
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
