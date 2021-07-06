package elife

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoElifeLifecardConsumeAPIResponse 品牌惠卡券核销 API返回值
// taobao.elife.lifecard.consume
//
// 用户线上购买生活汇品牌惠虚拟消费卡，线下购物时，商家码枪核销，涉及用户虚拟卡余额扣减操作
type TaobaoElifeLifecardConsumeAPIResponse struct {
	model.CommonResponse
	TaobaoElifeLifecardConsumeAPIResponseModel
}

// TaobaoElifeLifecardConsumeAPIResponseModel is 品牌惠卡券核销 成功返回结果
type TaobaoElifeLifecardConsumeAPIResponseModel struct {
	XMLName xml.Name `xml:"elife_lifecard_consume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 结果描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 本金
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 膨胀金
	InflateAmount int64 `json:"inflate_amount,omitempty" xml:"inflate_amount,omitempty"`
	// 是否成功
	Successed bool `json:"successed,omitempty" xml:"successed,omitempty"`
}
