require("@nomiclabs/hardhat-waffle")
require("dotenv").config()
require("hardhat-gas-reporter")
require("@nomiclabs/hardhat-web3")
require("@nomiclabs/hardhat-etherscan")

module.exports = {
	solidity: {
		version: "0.8.11",
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
			blockGasLimit: 199022552,
			gas: 1500000,
			gasPrice: 100,
			allowUnlimitedContractSize: false,
			throwOnTransactionFailures: false,
			throwOnCallFailures: true,
		},
		ganache: {
			url: "http://localhost:7545",
			blockGasLimit: 10000000,
		},
		arbitrum_test: {
			url: process.env.ARBITRUM_TEST,
			network_id: 421611,
			gas: 2500000,
			gasPrice: 30000000, //0.03 gwei
			timeout: 25000,
			accounts: [process.env.PRIVATE_KEY],
		},
	},
	gasReporter: {
		enabled: !!process.env.REPORT_GAS === true,
		currency: "USD",
		gasPrice: 100,
		showTimeSpent: true,
		coinmarketcap: process.env.COINMARKETCAP_API,
	},
	mocha: {
		timeout: 25000,
	},
	etherscan: {
		apiKey: process.env.ARBISCAN_KEY,
	},
}
