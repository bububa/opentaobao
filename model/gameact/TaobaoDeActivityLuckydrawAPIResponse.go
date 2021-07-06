package gameact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDeActivityLuckydrawAPIResponse 抽奖 API返回值
// taobao.de.activity.luckydraw
//
// 用于激励平台对外提供抽奖功能，包括但不限于集分宝、红包、宝点、淘金币、淘彩票等
type TaobaoDeActivityLuckydrawAPIResponse struct {
	model.CommonResponse
	TaobaoDeActivityLuckydrawAPIResponseModel
}

// TaobaoDeActivityLuckydrawAPIResponseModel is 抽奖 成功返回结果
type TaobaoDeActivityLuckydrawAPIResponseModel struct {
	XMLName xml.Name `xml:"de_activity_luckydraw_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 中奖奖品
	AwardVO *AwardVO `json:"award_v_o,omitempty" xml:"award_v_o,omitempty"`
	// 数娱积分/金牌余额
	TotalAmount int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 可用抽奖次数
	AccessAmount int64 `json:"access_amount,omitempty" xml:"access_amount,omitempty"`
	// 123
	IsLucky bool `json:"is_lucky,omitempty" xml:"is_lucky,omitempty"`
}
