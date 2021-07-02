package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmPointCalAPIResponse 计算积分可以抵扣的金额 API返回值
// alibaba.alsc.crm.point.cal
//
// 计算积分可以抵扣的金额
// 积分的抵扣为区间型
// 如抵扣规则为100积分抵扣50元，则输入消费120积分的话，回返回消费100积分抵扣50元
//  这里为纯计算逻辑，不会校验用户是否有足够的可用积分进行抵扣
type AlibabaAlscCrmPointCalAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmPointCalAPIResponseModel
}

// AlibabaAlscCrmPointCalAPIResponseModel is 计算积分可以抵扣的金额 成功返回结果
type AlibabaAlscCrmPointCalAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_point_cal_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
