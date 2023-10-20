package alihealthcrm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthBabyBaseinfoOrderSyncAPIResponse alibaba.alihealth.baby.baseinfo.order.sync API返回值
// alibaba.alihealth.baby.baseinfo.order.sync
//
// 育学园将订单信息回传给我们
type AlibabaAlihealthBabyBaseinfoOrderSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthBabyBaseinfoOrderSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthBabyBaseinfoOrderSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthBabyBaseinfoOrderSyncAPIResponseModel).Reset()
}

// AlibabaAlihealthBabyBaseinfoOrderSyncAPIResponseModel is alibaba.alihealth.baby.baseinfo.order.sync 成功返回结果
type AlibabaAlihealthBabyBaseinfoOrderSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_baby_baseinfo_order_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作提示
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 操作码
	ReturnCode int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthBabyBaseinfoOrderSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ReturnMsg = ""
	m.ReturnCode = 0
}

var poolAlibabaAlihealthBabyBaseinfoOrderSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthBabyBaseinfoOrderSyncAPIResponse)
	},
}

// GetAlibabaAlihealthBabyBaseinfoOrderSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihealthBabyBaseinfoOrderSyncAPIResponse
func GetAlibabaAlihealthBabyBaseinfoOrderSyncAPIResponse() *AlibabaAlihealthBabyBaseinfoOrderSyncAPIResponse {
	return poolAlibabaAlihealthBabyBaseinfoOrderSyncAPIResponse.Get().(*AlibabaAlihealthBabyBaseinfoOrderSyncAPIResponse)
}

// ReleaseAlibabaAlihealthBabyBaseinfoOrderSyncAPIResponse 将 AlibabaAlihealthBabyBaseinfoOrderSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthBabyBaseinfoOrderSyncAPIResponse(v *AlibabaAlihealthBabyBaseinfoOrderSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthBabyBaseinfoOrderSyncAPIResponse.Put(v)
}
