package waybill

import (
	"sync"
)

// WaybillCloudPrintUpdateRequest 结构体
type WaybillCloudPrintUpdateRequest struct {
	// 物流公司CODE
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 物流服务内容&lt;a href=&#34;http://open.taobao.com/doc2/detail.htm?spm=a219a.7629140.0.0.eK8aZm&amp;treeId=17&amp;articleId=26765&amp;docType=2&#34;&gt;链接&lt;/a&gt;
	LogisticsServices string `json:"logistics_services,omitempty" xml:"logistics_services,omitempty"`
	// 模板URL
	TemplateUrl string `json:"template_url,omitempty" xml:"template_url,omitempty"`
	// 面单号
	WaybillCode string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
	// 请求表示id
	ObjectId string `json:"object_id,omitempty" xml:"object_id,omitempty"`
	// 包裹信息
	PackageInfo *PackageInfoDto `json:"package_info,omitempty" xml:"package_info,omitempty"`
	// 收件信息
	Recipient *UserInfoDto `json:"recipient,omitempty" xml:"recipient,omitempty"`
	// 发件信息
	Sender *UserInfoDto `json:"sender,omitempty" xml:"sender,omitempty"`
}

var poolWaybillCloudPrintUpdateRequest = sync.Pool{
	New: func() any {
		return new(WaybillCloudPrintUpdateRequest)
	},
}

// GetWaybillCloudPrintUpdateRequest() 从对象池中获取WaybillCloudPrintUpdateRequest
func GetWaybillCloudPrintUpdateRequest() *WaybillCloudPrintUpdateRequest {
	return poolWaybillCloudPrintUpdateRequest.Get().(*WaybillCloudPrintUpdateRequest)
}

// ReleaseWaybillCloudPrintUpdateRequest 释放WaybillCloudPrintUpdateRequest
func ReleaseWaybillCloudPrintUpdateRequest(v *WaybillCloudPrintUpdateRequest) {
	v.CpCode = ""
	v.LogisticsServices = ""
	v.TemplateUrl = ""
	v.WaybillCode = ""
	v.ObjectId = ""
	v.PackageInfo = nil
	v.Recipient = nil
	v.Sender = nil
	poolWaybillCloudPrintUpdateRequest.Put(v)
}
