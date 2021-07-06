package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMarketingLotterySchemaCreateAPIResponse 抽奖平台抽奖方案创建接口 API返回值
// alibaba.marketing.lottery.schema.create
//
// 抽奖平台抽奖方案创建接口
type AlibabaMarketingLotterySchemaCreateAPIResponse struct {
	model.CommonResponse
	AlibabaMarketingLotterySchemaCreateAPIResponseModel
}

// AlibabaMarketingLotterySchemaCreateAPIResponseModel is 抽奖平台抽奖方案创建接口 成功返回结果
type AlibabaMarketingLotterySchemaCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_marketing_lottery_schema_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msg
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 抽奖方案对象
	LotterySchema *LotterySchemaDto `json:"lottery_schema,omitempty" xml:"lottery_schema,omitempty"`
	// code
	MsgCode int64 `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
