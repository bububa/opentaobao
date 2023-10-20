package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlicomOrderCheckorderinfoAPIResponse 金融购机订单信息校验 API返回值
// alibaba.alicom.order.checkorderinfo
//
// 金融购机订单信息校验
type AlibabaAlicomOrderCheckorderinfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlicomOrderCheckorderinfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlicomOrderCheckorderinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlicomOrderCheckorderinfoAPIResponseModel).Reset()
}

// AlibabaAlicomOrderCheckorderinfoAPIResponseModel is 金融购机订单信息校验 成功返回结果
type AlibabaAlicomOrderCheckorderinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alicom_order_checkorderinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlicomOrderCheckorderinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlicomOrderCheckorderinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlicomOrderCheckorderinfoAPIResponse)
	},
}

// GetAlibabaAlicomOrderCheckorderinfoAPIResponse 从 sync.Pool 获取 AlibabaAlicomOrderCheckorderinfoAPIResponse
func GetAlibabaAlicomOrderCheckorderinfoAPIResponse() *AlibabaAlicomOrderCheckorderinfoAPIResponse {
	return poolAlibabaAlicomOrderCheckorderinfoAPIResponse.Get().(*AlibabaAlicomOrderCheckorderinfoAPIResponse)
}

// ReleaseAlibabaAlicomOrderCheckorderinfoAPIResponse 将 AlibabaAlicomOrderCheckorderinfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlicomOrderCheckorderinfoAPIResponse(v *AlibabaAlicomOrderCheckorderinfoAPIResponse) {
	v.Reset()
	poolAlibabaAlicomOrderCheckorderinfoAPIResponse.Put(v)
}
