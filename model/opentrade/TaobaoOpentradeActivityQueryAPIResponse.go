package opentrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeActivityQueryAPIResponse
查询尖货活动信息 API返回值
taobao.opentrade.activity.query

尖货交易活动信息配置，查询尖货活动信息 */
type TaobaoOpentradeActivityQueryAPIResponse struct {
	model.CommonResponse
	TaobaoOpentradeActivityQueryAPIResponseModel
}

// TaobaoOpentradeActivityQueryAPIResponseModel is 查询尖货活动信息 成功返回结果
type TaobaoOpentradeActivityQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"opentrade_activity_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 活动信息记录
	Results []McSceneActivityDto `json:"results,omitempty" xml:"results>mc_scene_activity_dto,omitempty"`
}
