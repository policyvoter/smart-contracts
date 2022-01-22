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

		await nft.toggleSaleStatus()
		await nft.togglePresaleStatus()
	})

	it("simple test...", async function () {
		await nft.setPolicyCreator(acc1.address)
		let pCreator = await nft.policyCreator()
		expect(pCreator == acc1.address)
	})

	// function publicBuy(uint256 qty) external payable {
	// 	require(saleLive, "sale not live");
	// 	require(qty <= maxQtyForPublicSale, "no more than 15 at once");
	// 	require(totalSupply() + qty < maxSupply, "out of stock");
	// 	require(pricePerToken * qty == msg.value, "exact amount needed");
	// 	for (uint256 i = 0; i < qty; i++) {
	// 		_safeMint(msg.sender, totalSupply() + 1);
	// 	}
	// }
	// it("purchasing tokens works", async function () {
	// 	await expect(
	// 		nft.connect(acc1).publicBuy(1, { value: web3.utils.toWei("0.09", "ether") })
	// 	).to.emit(nft, "Transfer")
	// 	await expect(
	// 		nft.connect(acc1).publicBuy(3, { value: web3.utils.toWei("0.27", "ether") })
	// 	).to.emit(nft, "Transfer")
	// })

	// it("burning tokens works", async function () {
	// 	await expect(
	// 		nft.connect(acc1).publicBuy(3, { value: web3.utils.toWei("0.27", "ether") })
	// 	).to.emit(nft, "Transfer")
	// 	expect(await nft.balanceOf(acc1.address)).to.equal(3)
	// 	await nft.connect(acc1).burn(1)
	// 	expect(await nft.balanceOf(acc1.address)).to.equal(2)
	// })

	// // it("purchasing 3 tokens works", async function () {
	// // 	await nft.startSale()
	// // 	await expect(
	// // 		nft.connect(acc1).acquire(3, { value: web3.utils.toWei("0.21", "ether") })
	// // 	).to.emit(nft, "Transfer")
	// // })

	// // it("burning a token works", async function () {
	// // 	await nft.startSale()
	// // 	await expect(
	// // 		nft.connect(acc1).acquire(1, { value: web3.utils.toWei("0.07", "ether") })
	// // 	).to.emit(nft, "Transfer")

	// // 	expect(await nft.balanceOf(acc1.address)).to.equal(1)

	// // 	await nft.connect(acc1).burn(1)
	// // 	expect(await nft.balanceOf(acc1.address)).to.equal(0)
	// // })

	// // it("can't burn other tokens than your own", async function () {
	// // 	await nft.startSale()
	// // 	await expect(
	// // 		nft.connect(acc1).acquire(1, { value: web3.utils.toWei("0.07", "ether") })
	// // 	).to.emit(nft, "Transfer")
	// // 	await expect(
	// // 		nft.connect(acc2).acquire(1, { value: web3.utils.toWei("0.07", "ether") })
	// // 	).to.emit(nft, "Transfer")

	// // 	expect(await nft.balanceOf(acc1.address)).to.equal(1)

	// // 	await expect(nft.connect(acc1).burn(2)).to.be.revertedWith(
	// // 		"revert caller is not owner nor approved"
	// // 	)
	// // })

	// // it("custom thing emits", async function () {
	// // 	await nft.startSale()
	// // 	await expect(
	// // 		nft.connect(acc1).acquire(1, { value: web3.utils.toWei("0.07", "ether") })
	// // 	).to.emit(nft, "Transfer")

	// // 	await expect(
	// // 		nft.connect(acc1).customThing(1, 100, "hello", { value: web3.utils.toWei("0.1", "ether") })
	// // 	).to.emit(nft, "CustomThing")
	// // })

	// // it("withdraw money works", async function () {
	// // 	const tracker = await balance.tracker(owner.address)
	// // 	let ownerInitialBalance = Number(await tracker.get("wei"))
	// // 	await nft.startSale()
	// // 	await expect(
	// // 		nft.connect(acc1).acquire(1, { value: web3.utils.toWei("0.07", "ether") })
	// // 	).to.emit(nft, "Transfer")

	// // 	await nft.withdrawEarnings()

	// // 	let ownerFinalBalance = Number(await tracker.get("wei"))
	// // 	expect(ownerFinalBalance - ownerInitialBalance).to.be.greaterThan(
	// // 		Number(web3.utils.toWei("0.06", "ether")) //some gas costs are lost
	// // 	)
	// // })

	// // it("should transfer accidentally sent ERC20 tokens to this contract", async function () {
	// // 	//deploy an erc20 token for
	// // 	let ERC20MockContract = await ethers.getContractFactory("ERC20Mock")
	// // 	erc20 = await ERC20MockContract.connect(acc1).deploy("ERCToken", "ERC", "10000")
	// // 	await erc20.deployed()

	// // 	// Transfer some tokens to this contract
	// // 	await erc20.connect(acc1).transfer(nft.address, 100)
	// // 	expect(await erc20.balanceOf(nft.address)).to.equal(100)

	// // 	// get them
	// // 	await nft.reclaimERC20(erc20.address)
	// // 	expect(await erc20.balanceOf(owner.address)).to.equal(100)
	// // })
})
