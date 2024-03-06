// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ITeleporterMessenger, TeleporterFeeInfo, TeleporterMessageInput} from "@teleporter/ITeleporterMessenger.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts@4.8.1/security/ReentrancyGuard.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @dev TeleporterPing is an example contract that demonstrates how to send and receive
 * messages cross chain.
 */
contract TeleporterPing is ReentrancyGuard {
    // Messages sent to this contract.
    struct Message {
        address sender;
        string message;
    }

    ITeleporterMessenger public constant teleporter =
        ITeleporterMessenger(0x253b2784c75e510dD0fF1da844684a1aC0aa5fcf);

    mapping(bytes32 sourceBlockchainID => Message message) private _messages;

    /**
     * @dev Emitted when a new message is received from a given chain ID.
     */
    event ReceiveMessage(
        bytes32 indexed sourceBlockchainID, address indexed originSenderAddress, string message
    );

    /**
     * @dev Emitted when a message is submited to be sent.
     */
    event SendMessage(
        bytes32 indexed destinationBlockchainID,
        address indexed destinationAddress,
        uint256 requiredGasLimit,
        string message
    );

    /**
     * @dev Sends a message to another chain.
     * @return The message ID of the newly sent message.
     */
    function sendMessage(
        bytes32 destinationBlockchainID,
        address destinationAddress,
        uint256 requiredGasLimit,
        string calldata message,
        string calldata response
    ) external nonReentrant returns (bytes32) {

        emit SendMessage({
            destinationBlockchainID: destinationBlockchainID,
            destinationAddress: destinationAddress,
            requiredGasLimit: requiredGasLimit,
            message: message
        });
        return teleporter.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationBlockchainID: destinationBlockchainID,
                destinationAddress: destinationAddress,
                feeInfo: TeleporterFeeInfo({feeTokenAddress: address(0), amount: 0}),
                requiredGasLimit: requiredGasLimit,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(message, response)
            })
        );
    }

    /**
     * @dev Returns the message sent from another chain.
     * @return The sender of the message, and the message itself.
     */
    function getMessage(bytes32 sourceBlockchainID)
        external
        view
        returns (address, string memory)
    {
        Message memory messageInfo = _messages[sourceBlockchainID];
        return (messageInfo.sender, messageInfo.message);
    }

    /**
     * @dev Receive a Teleporter message from another chain. Can only be called once per source chain.
     */
    function receiveTeleporterMessage(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        bytes memory message
    ) external nonReentrant {
        require(msg.sender == address(teleporter), "Only the Teleporter contract can send messages");
        require(_messages[sourceBlockchainID].sender == address(0), "Source blockchain already sent a message");
        // Store the message.
        string memory messageString = abi.decode(message, (string));
        _messages[sourceBlockchainID] = Message(originSenderAddress, messageString);
        emit ReceiveMessage(sourceBlockchainID, originSenderAddress, messageString);
    }
}
