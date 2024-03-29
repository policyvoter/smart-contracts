const hre = require("hardhat");
require("@nomiclabs/hardhat-web3");
const fs = require("fs-extra");

async function main() {
  fs.removeSync("cache");
  fs.removeSync("artifacts");
  await hre.run("compile");

  console.log(">-> Network is set to " + process.env.NETWORK);

  // ethers is avaialble in the global scope
  const [deployer] = await ethers.getSigners();
  const deployerAddress = await deployer.getAddress();
  const account = await web3.utils.toChecksumAddress(deployerAddress);
  const balance = await web3.eth.getBalance(account);

  console.log(
    "Deployer Account " +
      deployerAddress +
      " has balance: " +
      web3.utils.fromWei(balance, "ether"),
    "ETH"
  );

  let PVContract = await hre.ethers.getContractFactory("PolicyVoterV1");
  console.log("Deploying PolicyVoter Contract...");
  let deployed = await PVContract.deploy();
  let pv = await deployed.deployed();

  console.log("Contract deployed to ", pv.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
