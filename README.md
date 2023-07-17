# PolicyVoter Smart Contracts V2

deloyed: 0x7a9818Dc8Fc73855109aEF43d7208711b1f3C83f

## Instructions to verify authenticity

- The blockchain RPC URL is: https://rpc.policyvoter.tech
- Currently, no blockchain explorer is present (we plan to have this in the near future) so the only way to check if the smart contract advertised contains the source code from the repo is to compile it yourself (npm run compile) and compare the Bytecode manually

#### New votes:

- Keep in mind that voting/unvoting bypasses our backend goes directly interacts with the smart contract
- Checkout the smart contract code in this repo under /contracts
- The smart contract's latest address is published above
- Take note of your address (you can find it in the browser's localstorage)
- Download the blocks when you pass a vote and check if your vote is registered correctly (for example, an easy way is to check if the event emit Voted(policyID, policies[policyID].groupID, votes[policyID]); or Unvoted is triggered

##### Old votes:

- For the pilots there were several (small) smart contract changes, they didn't have a "groupID" parameter on the votes
- There are two ways to do this, the easy way is to check the GitHub history here for the old smart contrats address, compile the smart contract at that particular time in history and check the Bytecode coresponds, get the tx address of the vote(s), check what policy they voted and calculate the authenticity of the votes/policy
- Another way is to download all the blockchain, parse the txes since the time you want to check the votes and do the calculation like this

#### If you have any questions, feel free to contact us.

## License: Mozilla Public License Version 2.0 (see file LICENSE)
