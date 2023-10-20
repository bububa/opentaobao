package tvupadmin

import (
	"sync"
)

// TvAppVersion 结构体
type TvAppVersion struct {
	// 适配设备列表
	DeviceAdapterList []Deviceadapterlist `json:"device_adapter_list,omitempty" xml:"device_adapter_list>deviceadapterlist,omitempty"`
	// 重启类型
	RestartType string `json:"restart_type,omitempty" xml:"restart_type,omitempty"`
	// 系统版本
	Version string `json:"version,omitempty" xml:"version,omitempty"`
	// 下载地址
	DownloadPath string `json:"download_path,omitempty" xml:"download_path,omitempty"`
	// 发布说明
	ReleaseNote string `json:"release_note,omitempty" xml:"release_note,omitempty"`
	// 状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 升级包类型
	AppZipType string `json:"app_zip_type,omitempty" xml:"app_zip_type,omitempty"`
	// 是否需要重启
	NeedRestart string `json:"need_restart,omitempty" xml:"need_restart,omitempty"`
	// downloadmd5
	Downloadmd5 string `json:"downloadmd5,omitempty" xml:"downloadmd5,omitempty"`
	// gmtCreate
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// gmtModify
	GmtModify string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	// size
	Size string `json:"size,omitempty" xml:"size,omitempty"`
	// 是否通知
	HasNotified string `json:"has_notified,omitempty" xml:"has_notified,omitempty"`
	// 重启应用参数
	RestartAppParam string `json:"restart_app_param,omitempty" xml:"restart_app_param,omitempty"`
	// 重启应用类型
	RestartAppType string `json:"restart_app_type,omitempty" xml:"restart_app_type,omitempty"`
	// 应用版本
	VersionCode int64 `json:"version_code,omitempty" xml:"version_code,omitempty"`
	// 应用信息
	App *AppDto `json:"app,omitempty" xml:"app,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolTvAppVersion = sync.Pool{
	New: func() any {
		return new(TvAppVersion)
	},
}

// GetTvAppVersion() 从对象池中获取TvAppVersion
func GetTvAppVersion() *TvAppVersion {
	return poolTvAppVersion.Get().(*TvAppVersion)
}

// ReleaseTvAppVersion 释放TvAppVersion
func ReleaseTvAppVersion(v *TvAppVersion) {
	v.DeviceAdapterList = v.DeviceAdapterList[:0]
	v.RestartType = ""
	v.Version = ""
	v.DownloadPath = ""
	v.ReleaseNote = ""
	v.Status = ""
	v.AppZipType = ""
	v.NeedRestart = ""
	v.Downloadmd5 = ""
	v.GmtCreate = ""
	v.GmtModify = ""
	v.Size = ""
	v.HasNotified = ""
	v.RestartAppParam = ""
	v.RestartAppType = ""
	v.VersionCode = 0
	v.App = nil
	v.Id = 0
	poolTvAppVersion.Put(v)
}
