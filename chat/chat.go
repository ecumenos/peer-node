package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/ecumenos/golang-toolkit/randomtools"
	"github.com/ecumenos/peer-node/host"
	peer "github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-libp2p/core/protocol"
)

type Chat interface{}

type chatImpl struct {
	host       host.Host
	protocolID protocol.ID
}

func New(ctx context.Context, h host.Host) (Chat, error) {
	nano, err := randomtools.GetNanoString(12)
	if err != nil {
		return nil, err
	}
	c := &chatImpl{
		host:       h,
		protocolID: protocol.ID(fmt.Sprintf("/p2p/chat/%s", nano)),
	}

	// Start listening for incoming chat messages
	go c.ListenForMessages(ctx)

	return c, nil
}

func (c *chatImpl) SendMessage(ctx context.Context, peerID peer.ID, message string) error {
	data, err := json.Marshal(Message{
		Data:     Data{Text: message},
		Metadata: Metadata{Timestamp: time.Now().Unix()},
	})
	if err != nil {
		return err
	}

	stream, err := c.host.NewStream(ctx, peerID, c.protocolID)
	if err != nil {
		return err
	}
	defer stream.Close()
	_, err = stream.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (c *chatImpl) ListenForMessages(ctx context.Context) {
	for {
		stream, err := c.host.NewStream(ctx, "", "")
		if err != nil {
			log.Println("Failed to open stream:", err)
			continue
		}
		go func() {
			// Read the incoming message data
			data, err := ioutil.ReadAll(stream)
			if err != nil {
				log.Println("Failed to read data:", err)
				return
			}

			// Parse the message data into a ChatMessage struct
			var message Message
			err = json.Unmarshal(data, &message)
			if err != nil {
				log.Println("Failed to unmarshal message:", err)
				return
			}

			// Print the message to the console
			log.Printf("[%s]: %s", stream.Conn().RemotePeer(), message.Data.Text)
		}()
	}
}
