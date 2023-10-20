package baoxian

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlipayBaoxianClaimUpdateAPIResponse 更新赔案 API返回值
// alipay.baoxian.claim.update
//
// 更新保险理赔单
type AlipayBaoxianClaimUpdateAPIResponse struct {
	model.CommonResponse
	AlipayBaoxianClaimUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlipayBaoxianClaimUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlipayBaoxianClaimUpdateAPIResponseModel).Reset()
}

// AlipayBaoxianClaimUpdateAPIResponseModel is 更新赔案 成功返回结果
type AlipayBaoxianClaimUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alipay_baoxian_claim_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlipayBaoxianClaimUpdateMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlipayBaoxianClaimUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlipayBaoxianClaimUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlipayBaoxianClaimUpdateAPIResponse)
	},
}

// GetAlipayBaoxianClaimUpdateAPIResponse 从 sync.Pool 获取 AlipayBaoxianClaimUpdateAPIResponse
func GetAlipayBaoxianClaimUpdateAPIResponse() *AlipayBaoxianClaimUpdateAPIResponse {
	return poolAlipayBaoxianClaimUpdateAPIResponse.Get().(*AlipayBaoxianClaimUpdateAPIResponse)
}

// ReleaseAlipayBaoxianClaimUpdateAPIResponse 将 AlipayBaoxianClaimUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlipayBaoxianClaimUpdateAPIResponse(v *AlipayBaoxianClaimUpdateAPIResponse) {
	v.Reset()
	poolAlipayBaoxianClaimUpdateAPIResponse.Put(v)
}
