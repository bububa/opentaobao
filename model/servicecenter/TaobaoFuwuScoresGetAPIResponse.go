package servicecenter

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFuwuScoresGetAPIResponse 服务平台评价查询接口 API返回值
// taobao.fuwu.scores.get
//
// 根据日期、查询appkey对应服务评价，每次调用只能查询某一天服务评价信息，可设置分页查询，页大小最大为100，非实时接口，延迟时间为30分钟
type TaobaoFuwuScoresGetAPIResponse struct {
	model.CommonResponse
	TaobaoFuwuScoresGetAPIResponseModel
}

// TaobaoFuwuScoresGetAPIResponseModel is 服务平台评价查询接口 成功返回结果
type TaobaoFuwuScoresGetAPIResponseModel struct {
	XMLName xml.Name `xml:"fuwu_scores_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 评价流水记录
	ScoreResult []ScoreResult `json:"score_result,omitempty" xml:"score_result>score_result,omitempty"`
}
