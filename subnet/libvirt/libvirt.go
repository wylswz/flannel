package libvirt

import (
	"context"

	"github.com/flannel-io/flannel/pkg/ip"
	"github.com/flannel-io/flannel/subnet"
)

type libvirtSubnetManager struct {
}

func NewSubnetManager(ctx context.Context, libvirtSock string) (subnet.Manager, error) {
	return newLibvirtSubnetManager(libvirtSock), nil
}

func (sm libvirtSubnetManager) GetNetworkConfig(ctx context.Context) (*subnet.Config, error)
func (sm libvirtSubnetManager) AcquireLease(ctx context.Context, attrs *subnet.LeaseAttrs) (*subnet.Lease, error)
func (sm libvirtSubnetManager) RenewLease(ctx context.Context, lease *subnet.Lease) error
func (sm libvirtSubnetManager) WatchLease(ctx context.Context, sn ip.IP4Net, cursor interface{}) (subnet.LeaseWatchResult, error)
func (sm libvirtSubnetManager) WatchLeases(ctx context.Context, cursor interface{}) (subnet.LeaseWatchResult, error)

func (sm libvirtSubnetManager) Name() string

func newLibvirtSubnetManager(libvirtSock string) libvirtSubnetManager {
	return libvirtSubnetManager{}
}
