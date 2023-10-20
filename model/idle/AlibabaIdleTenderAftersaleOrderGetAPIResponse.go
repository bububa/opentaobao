package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleTenderAftersaleOrderGetAPIResponse 闲鱼帮卖售后服务单查询 API返回值
// alibaba.idle.tender.aftersale.order.get
//
// 闲鱼帮卖售后服务单查询
type AlibabaIdleTenderAftersaleOrderGetAPIResponse struct {
	model.CommonResponse
	AlibabaIdleTenderAftersaleOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleTenderAftersaleOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleTenderAftersaleOrderGetAPIResponseModel).Reset()
}

// AlibabaIdleTenderAftersaleOrderGetAPIResponseModel is 闲鱼帮卖售后服务单查询 成功返回结果
type AlibabaIdleTenderAftersaleOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_tender_aftersale_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	Module *AlibabaIdleTenderAftersaleOrderGetModule `json:"module,omitempty" xml:"module,omitempty"`
	// 查询是否成功
	QuerySuccess bool `json:"query_success,omitempty" xml:"query_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleTenderAftersaleOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Module = nil
	m.QuerySuccess = false
}

var poolAlibabaIdleTenderAftersaleOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleTenderAftersaleOrderGetAPIResponse)
	},
}

// GetAlibabaIdleTenderAftersaleOrderGetAPIResponse 从 sync.Pool 获取 AlibabaIdleTenderAftersaleOrderGetAPIResponse
func GetAlibabaIdleTenderAftersaleOrderGetAPIResponse() *AlibabaIdleTenderAftersaleOrderGetAPIResponse {
	return poolAlibabaIdleTenderAftersaleOrderGetAPIResponse.Get().(*AlibabaIdleTenderAftersaleOrderGetAPIResponse)
}

// ReleaseAlibabaIdleTenderAftersaleOrderGetAPIResponse 将 AlibabaIdleTenderAftersaleOrderGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleTenderAftersaleOrderGetAPIResponse(v *AlibabaIdleTenderAftersaleOrderGetAPIResponse) {
	v.Reset()
	poolAlibabaIdleTenderAftersaleOrderGetAPIResponse.Put(v)
}
