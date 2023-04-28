package types

import (
	"fmt"
	"strings"
)

type BdevDriverSpecificLvol struct {
	LvolStoreUUID string   `json:"lvol_store_uuid"`
	BaseBdev      string   `json:"base_bdev"`
	BaseSnapshot  string   `json:"base_snapshot,omitempty"`
	ThinProvision bool     `json:"thin_provision"`
	Snapshot      bool     `json:"snapshot"`
	Clone         bool     `json:"clone"`
	Clones        []string `json:"clones,omitempty"`
}

type LvstoreInfo struct {
	UUID              string `json:"uuid"`
	Name              string `json:"name"`
	BaseBdev          string `json:"base_bdev"`
	TotalDataClusters uint64 `json:"total_data_clusters"`
	FreeClusters      uint64 `json:"free_clusters"`
	BlockSize         uint64 `json:"block_size"`
	ClusterSize       uint64 `json:"cluster_size"`
}

type BdevLvolCreateLvstoreRequest struct {
	BdevName string `json:"bdev_name"`
	LvsName  string `json:"lvs_name"`

	ClusterSz uint32 `json:"cluster_sz,omitempty"`
	// ClearMethod               string `json:"clear_method,omitempty"`
	// NumMdPagesPerClusterRatio uint32 `json:"num_md_pages_per_cluster_ratio,omitempty"`
}

type BdevLvolDeleteLvstoreRequest struct {
	UUID    string `json:"uuid,omitempty"`
	LvsName string `json:"lvs_name,omitempty"`
}

type BdevLvolRenameLvstoreRequest struct {
	OldName string `json:"old_name"`
	NewName string `json:"new_name"`
}

type BdevLvolGetLvstoreRequest struct {
	UUID    string `json:"uuid,omitempty"`
	LvsName string `json:"lvs_name,omitempty"`
}

type BdevLvolClearMethod string

const (
	BdevLvolClearMethodNone        = "none"
	BdevLvolClearMethodUnmap       = "unmap"
	BdevLvolClearMethodWriteZeroes = "write_zeroes"
)

type BdevLvolCreateRequest struct {
	LvsName   string `json:"lvs_name"`
	LvolName  string `json:"lvol_name"`
	SizeInMib uint64 `json:"size_in_mib"`

	UUID          string              `json:"uuid,omitempty"`
	ClearMethod   BdevLvolClearMethod `json:"clear_method,omitempty"`
	ThinProvision bool                `json:"thin_provision,omitempty"`
}

type BdevLvolDeleteRequest struct {
	Name string `json:"name"`
}

type BdevLvolSnapshotRequest struct {
	LvolName     string `json:"lvol_name"`
	SnapshotName string `json:"snapshot_name"`
}

type BdevLvolCloneRequest struct {
	SnapshotName string `json:"snapshot_name"`
	CloneName    string `json:"clone_name"`
}

type BdevLvolDecoupleParentRequest struct {
	Name string `json:"name"`
}

type BdevLvolResizeRequest struct {
	Name string `json:"name"`
	Size uint64 `json:"size"`
}

func GetLvolAlias(lvsName, lvolName string) string {
	return fmt.Sprintf("%s/%s", lvsName, lvolName)
}

func GetLvsNameFromAlias(alias string) string {
	splitRes := strings.Split(alias, "/")
	if len(splitRes) != 2 {
		return ""
	}
	return splitRes[0]
}

func GetLvolNameFromAlias(alias string) string {
	splitRes := strings.Split(alias, "/")
	if len(splitRes) != 2 {
		return ""
	}
	return splitRes[1]
}
