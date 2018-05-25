package volumecommands

import (
	"github.com/gluster/glusterd2/pkg/api"
)

// defaultGroupOptions maps from a profile name to a set of options
var defaultGroupOptions = map[string]*api.OptionGroup{
	"profile.gluster-block": {"profile.gluster-block",
		[]api.VolumeOption{{Name: "performance.quick-read", OnValue: "off"},
			{Name: "performance.read-ahead", OnValue: "off"},
			{Name: "performance.io-cache", OnValue: "off"},
			{Name: "performance.stat-prefetch", OnValue: "off"},
			{Name: "performance.open-behind", OnValue: "off"},
			{Name: "performance.readdir-ahead", OnValue: "off"},
			{Name: "performance.strict-o-direct", OnValue: "on"},
			{Name: "network.remote-dio", OnValue: "disable"},
			{Name: "cluster.eager-lock", OnValue: "disable"},
			{Name: "cluster.quorum-type", OnValue: "auto"},
			{Name: "cluster.data-self-heal-algorithm", OnValue: "full"},
			{Name: "cluster.locking-scheme", OnValue: "granular"},
			{Name: "cluster.shd-max-threads", OnValue: "8"},
			{Name: "cluster.shd-wait-qlength", OnValue: "10000"},
			{Name: "features.shard", OnValue: "on"},
			{Name: "features.shard-block-size", OnValue: "64MB"},
			{Name: "user.cifs", OnValue: "off"},
			{Name: "server.allow-insecure", OnValue: "on"}},
		"Enable this profile for optimal results, in block use case"},
	"profile.SMB-small-file": {"profile.SMB-small-file",
		[]api.VolumeOption{{Name: "features.cache-invalidation", OnValue: "on"},
			{Name: "features.cache-invalidation-timeout", OnValue: "600"},
			{Name: "network.inode-lru-limit", OnValue: "200000"},
			{Name: "performance.stat-prefetch", OnValue: "on"},
			{Name: "performance.cache-invalidation", OnValue: "on"},
			{Name: "performance.md-cache-timeout", OnValue: "600"},
			{Name: "performance.cache-samba-metadata", OnValue: "on"},
			{Name: "performance.nl-cache", OnValue: "on"},
			{Name: "performance.nl-cache-timeout", OnValue: "600"},
			{Name: "performance.qr-cache-timeout", OnValue: "600"},
			{Name: "performance.readdir-ahead", OnValue: "on"},
			{Name: "performance.parallel-readdir", OnValue: "on"},
			{Name: "client.event-threads", OnValue: "4"},
			{Name: "server.event-threads", OnValue: "4"},
			{Name: "cluster.lookup-optimize", OnValue: "on"},
			{Name: "cluster.readdir-optimize", OnValue: "on"}},
		"For use cases with dominant small file workload in SMB access, enable this profile"},
	"profile.FUSE-small-file": {"profile.FUSE-small-file",
		[]api.VolumeOption{{Name: "features.cache-invalidation", OnValue: "on"},
			{Name: "features.cache-invalidation-timeout", OnValue: "600"},
			{Name: "network.inode-lru-limit", OnValue: "200000"},
			{Name: "performance.stat-prefetch", OnValue: "on"},
			{Name: "performance.cache-invalidation", OnValue: "on"},
			{Name: "performance.md-cache-timeout", OnValue: "600"},
			{Name: "performance.nl-cache", OnValue: "on"},
			{Name: "performance.nl-cache-timeout", OnValue: "600"},
			{Name: "performance.qr-cache-timeout", OnValue: "600"},
			{Name: "performance.readdir-ahead", OnValue: "on"},
			{Name: "performance.parallel-readdir", OnValue: "on"},
			{Name: "client.event-threads", OnValue: "4"},
			{Name: "server.event-threads", OnValue: "4"},
			{Name: "cluster.lookup-optimize", OnValue: "on"},
			{Name: "cluster.readdir-optimize", OnValue: "on"}},
		"For use cases with dominant small file workload in native FUSE mount access, enable this profile"},
	"profile.SMB-large-file-EC": {"profile.SMB-large-file-EC",
		[]api.VolumeOption{{Name: "features.cache-invalidation", OnValue: "on"},
			{Name: "features.cache-invalidation-timeout", OnValue: "600"},
			{Name: "network.inode-lru-limit", OnValue: "200000"},
			{Name: "performance.stat-prefetch", OnValue: "on"},
			{Name: "performance.cache-invalidation", OnValue: "on"},
			{Name: "performance.md-cache-timeout", OnValue: "600"},
			{Name: "performance.cache-samba-metadata", OnValue: "on"},
			{Name: "performance.nl-cache", OnValue: "on"},
			{Name: "performance.nl-cache-timeout", OnValue: "600"},
			{Name: "performance.readdir-ahead", OnValue: "on"},
			{Name: "performance.parallel-readdir", OnValue: "on"},
			{Name: "performance.quick-read", OnValue: "off"},
			{Name: "performance.open-behind", OnValue: "off"},
			{Name: "performance.write-behind-trickling-writes", OnValue: "off"},
			{Name: "performance.aggregate-size", OnValue: "1MB"},
			{Name: "client.event-threads", OnValue: "4"},
			{Name: "server.event-threads", OnValue: "4"}},
		"Enable this profile for use cases consisting of mostly large files like video surveillance, backup, video streaming and others, with SMB access to an erasure coded volume"},
	"profile.FUSE-large-file-EC": {"profile.FUSE-large-file-EC",
		[]api.VolumeOption{{Name: "features.cache-invalidation", OnValue: "on"},
			{Name: "features.cache-invalidation-timeout", OnValue: "600"},
			{Name: "network.inode-lru-limit", OnValue: "200000"},
			{Name: "performance.stat-prefetch", OnValue: "on"},
			{Name: "performance.cache-invalidation", OnValue: "on"},
			{Name: "performance.md-cache-timeout", OnValue: "600"},
			{Name: "performance.nl-cache", OnValue: "on"},
			{Name: "performance.nl-cache-timeout", OnValue: "600"},
			{Name: "performance.readdir-ahead", OnValue: "on"},
			{Name: "performance.parallel-readdir", OnValue: "on"},
			{Name: "performance.quick-read", OnValue: "off"},
			{Name: "performance.open-behind", OnValue: "off"},
			{Name: "performance.write-behind-trickling-writes", OnValue: "off"},
			{Name: "performance.aggregate-size", OnValue: "1MB"},
			{Name: "client.event-threads", OnValue: "4"},
			{Name: "server.event-threads", OnValue: "4"}},
		"Enable this profile for use cases consisting of mostly large files like video surveillance, backup, video streaming and others, with native FUSE mount access to an erasure coded volume"},
	"profile.nl-cache": {"profile.nl-cache",
		[]api.VolumeOption{{Name: "features.cache-invalidation", OnValue: "on"},
			{Name: "features.cache-invalidation-timeout", OnValue: "600"},
			{Name: "performance.nl-cache", OnValue: "on"},
			{Name: "performance.nl-cache-timeout", OnValue: "600"},
			{Name: "network.inode-lru-limit", OnValue: "200000"}},
		"Enable this for the workloads that generate lots of lookups before creating files, eg: SMB access, git tool and others"},
	"profile.metadata-cache": {"profile.metadata-cache",
		[]api.VolumeOption{{Name: "cache-invalidation", OnValue: "on"},
			{Name: "cache-invalidation-timeout", OnValue: "600"},
			{Name: "performance.stat-prefetch", OnValue: "on"},
			{Name: "performance.cache-invalidation", OnValue: "on"},
			{Name: "performance.md-cache-timeout", OnValue: "600"},
			{Name: "inode-lru-limit", OnValue: "200000"}},
		"This profile enables metadata(stat, xattr) caching on client side. Its recommended to enable this for most workloads other than the ones that require 100% consistency like databases"},
	"profile.virt": {"profile.virt",
		[]api.VolumeOption{{Name: "performance.quick-read", OnValue: "off"},
			{Name: "performance.read-ahead", OnValue: "off"},
			{Name: "performance.io-cache", OnValue: "off"},
			{Name: "performance.low-prio-threads", OnValue: "32"},
			{Name: "network.remote-dio", OnValue: "enable"},
			{Name: "cluster.eager-lock", OnValue: "enable"},
			{Name: "cluster.quorum-type", OnValue: "auto"},
			{Name: "cluster.server-quorum-type", OnValue: "server"},
			{Name: "cluster.data-self-heal-algorithm", OnValue: "full"},
			{Name: "cluster.locking-scheme", OnValue: "granular"},
			{Name: "cluster.shd-max-threads", OnValue: "8"},
			{Name: "cluster.shd-wait-qlength", OnValue: "10000"},
			{Name: "features.shard", OnValue: "on"},
			{Name: "user.cifs", OnValue: "off"}},
		"Enable this profile, if the Gluster Volume is used to store virtaul machines"},
	"profile.test": {"profile.test",
		[]api.VolumeOption{{Name: "afr.eager-lock", OnValue: "on"},
			{Name: "gfproxy.afr.eager-lock", OnValue: "on"}},
		"Test purpose only"},
}
