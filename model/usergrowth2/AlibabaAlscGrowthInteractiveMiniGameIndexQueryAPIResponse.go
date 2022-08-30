package usergrowth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscGrowthInteractiveMiniGameIndexQueryAPIResponse 小游戏首页查询接口 API返回值
// alibaba.alsc.growth.interactive.mini.game.index.query
//
// 小游戏首页查询接口
type AlibabaAlscGrowthInteractiveMiniGameIndexQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlscGrowthInteractiveMiniGameIndexQueryAPIResponseModel
}

// AlibabaAlscGrowthInteractiveMiniGameIndexQueryAPIResponseModel is 小游戏首页查询接口 成功返回结果
type AlibabaAlscGrowthInteractiveMiniGameIndexQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_growth_interactive_mini_game_index_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 链路id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误码
	BizErrorCode string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 错误信息
	BizErrorMsg string `json:"biz_error_msg,omitempty" xml:"biz_error_msg,omitempty"`
	// 服务端返回数据
	Data *MiniGameIndexDto `json:"data,omitempty" xml:"data,omitempty"`
	// 接口调用是否成功
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
}
