// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "forge-std/Test.sol";
import "../src/RandomNumberConsumer.sol";
import "@chainlink/contracts/src/v0.8/vrf/mocks/VRFCoordinatorV2Mock.sol";

contract RandomNumberConsumerTest is Test {
    VRFCoordinatorV2Mock public vrfCoordinator;
    RandomNumberConsumer public consumer;
    uint64 public subscriptionId;
    // Note: This is an example keyHash value (it does not correspond to any live network).
    bytes32 public keyHash = 0x6c3699283bda56ad74f6b855546325b68d482e983852a338b6d0b5843bc42f2d;
    uint96 public constant BASE_FEE = 0.25 ether;
    uint96 public constant GAS_PRICE_LINK = 1e9;

    function setUp() public {
        // Deploy the VRFCoordinatorV2Mock contract.
        vrfCoordinator = new VRFCoordinatorV2Mock(BASE_FEE, GAS_PRICE_LINK);

        // Create a subscription; the first subscription id is 1.
        subscriptionId = vrfCoordinator.createSubscription();
        vrfCoordinator.fundSubscription(subscriptionId, 10 ether);

        // Deploy the consumer contract.
        consumer = new RandomNumberConsumer(
            subscriptionId,
            address(vrfCoordinator),
            keyHash
        );

        // Register the consumer contract with the subscription.
        vrfCoordinator.addConsumer(subscriptionId, address(consumer));
    }

    function testRequestRandomNumber() public {
        // The contract owner requests a random number.
        uint256 requestId = consumer.requestRandomNumber();

        // Simulate the Chainlink VRF callback.
        vrfCoordinator.fulfillRandomWords(requestId, address(consumer));

        // Compute the expected random result based on VRFCoordinatorV2Mock logic.
        uint256 expectedRandomWord = uint256(keccak256(abi.encode(requestId, uint256(0))));
        assertEq(
            consumer.s_randomWord(),
            expectedRandomWord,
            "Random word not set correctly"
        );
    }
} 