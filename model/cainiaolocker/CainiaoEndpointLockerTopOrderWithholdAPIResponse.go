package cainiaolocker

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoendpointlockertoporderwithholdAPIResponse 代扣支付 API返回值
// cainiao.endpoint.locker.top.order.withhold
//
// 提供代扣，允许有一笔欠款。
type CainiaoendpointlockertoporderwithholdAPIResponse struct {
	model.CommonResponse
	CainiaoendpointlockertoporderwithholdAPIResponseModel
}

// CainiaoendpointlockertoporderwithholdAPIResponseModel is 代扣支付 成功返回结果
type CainiaoendpointlockertoporderwithholdAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_endpoint_locker_top_order_withhold_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *SingleResult `json:"result,omitempty" xml:"result,omitempty"`
}
