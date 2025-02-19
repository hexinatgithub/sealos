package common

import "time"

const (
	InfraInstancesLabel   = "infra.sealos.io/instances/label"
	InfraInstancesUUID    = "infra.sealos.io/infra/uuid"
	InfraInstancesIndex   = "infra.sealos.io/instances/index"
	InfraVolumesLabel     = "infra.sealos.io/volumes/label"
	IPTypePublic          = "public"
	IPTypePrivate         = "private"
	DataVolumeLabel       = "data"
	RootVolumeLabel       = "root"
	TryTimes              = 8
	TrySleepTime          = time.Second
	TRUELable             = "true"
	SealosInfraFinalizer  = "infra.sealos.io/finalizers"
	InfraSecretPrefix     = "infra-secret"
	InfraVolumeIndex      = "infra.sealos.io/volumes/index"
	VolumeInfraID         = "infra.sealos.io/volumes/infraID"
	InstanceState         = "instance-state-name"
	DefaultRegion         = "cn-north-1b"
	KeyPairUser           = "sealos.io/keypair/user"
	ArchAmd64             = "amd64"
	ArchArm64             = "arm64"
	InstanceStatusRunning = "running"
	AliyunKeyPairPrefix   = "infra"
	DefaultNamespace      = "default"
	MasterO               = "master0"
)

var DefaultRootVolumeSize = int32(40)
var DriverList = []string{"aliyun", "aws"}
