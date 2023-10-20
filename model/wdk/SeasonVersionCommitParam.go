package wdk

import (
	"sync"
)

// SeasonVersionCommitParam 结构体
type SeasonVersionCommitParam struct {
	// generate的版本号
	SeasonVersion int64 `json:"season_version,omitempty" xml:"season_version,omitempty"`
}

var poolSeasonVersionCommitParam = sync.Pool{
	New: func() any {
		return new(SeasonVersionCommitParam)
	},
}

// GetSeasonVersionCommitParam() 从对象池中获取SeasonVersionCommitParam
func GetSeasonVersionCommitParam() *SeasonVersionCommitParam {
	return poolSeasonVersionCommitParam.Get().(*SeasonVersionCommitParam)
}

// ReleaseSeasonVersionCommitParam 释放SeasonVersionCommitParam
func ReleaseSeasonVersionCommitParam(v *SeasonVersionCommitParam) {
	v.SeasonVersion = 0
	poolSeasonVersionCommitParam.Put(v)
}
