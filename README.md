# PolicyVoter Smart Contracts V2

deloyed: 0xF804c129e91A9e2Bad8F919B0D0fC503788b365E

## Instructions to verify authenticity

- The blockchain RPC URL is: https://rpc.policyvoter.tech
- At this moment we don't have a blockchain explorer (we plan to have on in the future at some point) so the only way to check if the smart contract advertised contains the source code from the repo is to compile it yourself (`npm run compile`) and compare the bytecode manually

New votes:

- remember that voting/unvoting doesn't pass though our backend and goes directly to the smart contract
- checkout the smart contract code in this repo under `/contracts`
- the smart contract's latest address is published above
- take note of your address (you can find it in the browser's localstorage)
- download the blocks when you pass a vote and check if your vote is registered correctly (for example an easy way is to check if the event ` emit Voted(policyID, policies[policyID].groupID, votes[policyID]);` or `Unvoted` is triggered.
- if you would like to host a node please contact us and we'll give you an enode address

Old votes:

- for the old pilots there were several (small) smart contract changes. The old ones didn't have a "groupID" parameter on the votes.
- there are 2 ways to do this. The easy way, check the github history here for old smart contrats address. compile the smart contract at that particular time in history and check the bytecode coresponds. get the tx address of the vote(s), check what policy they voted and calculate the authenticity of the votes/policy
- Another way is to download all the blockchain, parse the txes since the time you want to check the votes and do the calculation like this. Contact us for an enode if you would like one.

For any questions don't hesitate to contact us 

## License: Mozilla Public License Version 2.0 (see file LICENSE)
