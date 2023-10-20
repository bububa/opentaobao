package baoxian

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlipaybaoxianclaimreturngoodsstatusupdateAPIResponse 更新理赔单退货货物状态 API返回值
// alipay.baoxian.claim.returngoodsstatus.update
//
// 更新理赔单退货货物状态
type AlipaybaoxianclaimreturngoodsstatusupdateAPIResponse struct {
	model.CommonResponse
	AlipaybaoxianclaimreturngoodsstatusupdateAPIResponseModel
}

// AlipaybaoxianclaimreturngoodsstatusupdateAPIResponseModel is 更新理赔单退货货物状态 成功返回结果
type AlipaybaoxianclaimreturngoodsstatusupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alipay_baoxian_claim_returngoodsstatus_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlipaybaoxianclaimreturngoodsstatusupdateMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}
