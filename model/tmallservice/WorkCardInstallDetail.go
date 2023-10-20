package tmallservice

import (
	"sync"
)

// WorkCardInstallDetail 结构体
type WorkCardInstallDetail struct {
	// 机器条码，如果有多个机器码，用英文逗号&#34;,&#34;隔开
	Sn string `json:"sn,omitempty" xml:"sn,omitempty"`
	// 安装图片，多个图片链接用英文逗号&#34;,&#34;隔开
	ImgUrls string `json:"img_urls,omitempty" xml:"img_urls,omitempty"`
	// 备注
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 辅材项目
	AccessoryInfo string `json:"accessory_info,omitempty" xml:"accessory_info,omitempty"`
	// 收费金额
	InstallFee string `json:"install_fee,omitempty" xml:"install_fee,omitempty"`
	// 机器安装状态(1未完成；20暂不安装；5完成；12取消)
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolWorkCardInstallDetail = sync.Pool{
	New: func() any {
		return new(WorkCardInstallDetail)
	},
}

// GetWorkCardInstallDetail() 从对象池中获取WorkCardInstallDetail
func GetWorkCardInstallDetail() *WorkCardInstallDetail {
	return poolWorkCardInstallDetail.Get().(*WorkCardInstallDetail)
}

// ReleaseWorkCardInstallDetail 释放WorkCardInstallDetail
func ReleaseWorkCardInstallDetail(v *WorkCardInstallDetail) {
	v.Sn = ""
	v.ImgUrls = ""
	v.Memo = ""
	v.AccessoryInfo = ""
	v.InstallFee = ""
	v.Status = 0
	poolWorkCardInstallDetail.Put(v)
}
