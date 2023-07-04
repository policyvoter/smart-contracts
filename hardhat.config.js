require("@nomiclabs/hardhat-waffle");
require("dotenv").config();
require("hardhat-gas-reporter");
require("@nomiclabs/hardhat-web3");
require("@nomiclabs/hardhat-etherscan");
require("hardhat-abi-exporter");

module.exports = {
  solidity: {
    version: "0.8.19",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200,
      },
    },
  },
  networks: {
    hardhat: {
      chainId: 127001,
      accounts: {
        mnemonic: "test test test test test test test test test test test junk",
      },
      allowUnlimitedContractSize: false,
      throwOnTransactionFailures: false,
      throwOnCallFailures: true,
    },
    ganache: {
      url: "http://localhost:7545",
      blockGasLimit: 10000000,
    },
    policy_voter: {
      url: process.env.POLICY_VOTER_RPC,
      network_id: 4444,
      gas: 2500000,
      gasPrice: 10000000000, //1 gwei
      timeout: 35000,
      accounts: [process.env.PRIVATE_KEY],
    },
  },
  mocha: {
    timeout: 25000,
  },
  abiExporter: {
    path: "./abi",
    runOnCompile: true,
    clear: true,
    flat: true,
    spacing: 2,
    pretty: false,
  },
};
