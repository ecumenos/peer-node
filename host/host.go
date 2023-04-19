package host

import (
	"context"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
)

type Host interface {
	NewStream(ctx context.Context, peerID peer.ID, protocolID protocol.ID) (network.Stream, error)
}

type hostImpl struct {
	host host.Host
}

func New(ctx context.Context) (Host, error) {
	h, err := libp2p.New()
	if err != nil {
		return nil, err
	}

	return &hostImpl{host: h}, nil
}

func (h *hostImpl) NewStream(ctx context.Context, peerID peer.ID, protocolID protocol.ID) (network.Stream, error) {
	return h.host.NewStream(ctx, peerID, protocolID)
}
