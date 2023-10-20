package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinAxbVendorHeartBeatAPIResponse 供应商心跳上报接口 API返回值
// alibaba.aliqin.axb.vendor.heart.beat
//
// 供应商上报自己的心跳信息
type AlibabaAliqinAxbVendorHeartBeatAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinAxbVendorHeartBeatAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinAxbVendorHeartBeatAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinAxbVendorHeartBeatAPIResponseModel).Reset()
}

// AlibabaAliqinAxbVendorHeartBeatAPIResponseModel is 供应商心跳上报接口 成功返回结果
type AlibabaAliqinAxbVendorHeartBeatAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_axb_vendor_heart_beat_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAliqinAxbVendorHeartBeatResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinAxbVendorHeartBeatAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliqinAxbVendorHeartBeatAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinAxbVendorHeartBeatAPIResponse)
	},
}

// GetAlibabaAliqinAxbVendorHeartBeatAPIResponse 从 sync.Pool 获取 AlibabaAliqinAxbVendorHeartBeatAPIResponse
func GetAlibabaAliqinAxbVendorHeartBeatAPIResponse() *AlibabaAliqinAxbVendorHeartBeatAPIResponse {
	return poolAlibabaAliqinAxbVendorHeartBeatAPIResponse.Get().(*AlibabaAliqinAxbVendorHeartBeatAPIResponse)
}

// ReleaseAlibabaAliqinAxbVendorHeartBeatAPIResponse 将 AlibabaAliqinAxbVendorHeartBeatAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinAxbVendorHeartBeatAPIResponse(v *AlibabaAliqinAxbVendorHeartBeatAPIResponse) {
	v.Reset()
	poolAlibabaAliqinAxbVendorHeartBeatAPIResponse.Put(v)
}
