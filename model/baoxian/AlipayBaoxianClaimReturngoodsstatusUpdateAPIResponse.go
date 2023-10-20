package baoxian

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse 更新理赔单退货货物状态 API返回值
// alipay.baoxian.claim.returngoodsstatus.update
//
// 更新理赔单退货货物状态
type AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse struct {
	model.CommonResponse
	AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponseModel).Reset()
}

// AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponseModel is 更新理赔单退货货物状态 成功返回结果
type AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alipay_baoxian_claim_returngoodsstatus_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlipayBaoxianClaimReturngoodsstatusUpdateMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse)
	},
}

// GetAlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse 从 sync.Pool 获取 AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse
func GetAlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse() *AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse {
	return poolAlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse.Get().(*AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse)
}

// ReleaseAlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse 将 AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse(v *AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse) {
	v.Reset()
	poolAlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse.Put(v)
}
