package baoxian

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse
更新理赔单退货货物状态 API返回值
alipay.baoxian.claim.returngoodsstatus.update

更新理赔单退货货物状态 */
type AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponse struct {
	model.CommonResponse
	AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponseModel
}

// AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponseModel is 更新理赔单退货货物状态 成功返回结果
type AlipayBaoxianClaimReturngoodsstatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alipay_baoxian_claim_returngoodsstatus_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlipayBaoxianClaimReturngoodsstatusUpdateMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}
