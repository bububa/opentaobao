package opentrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeActivityQueryAPIResponse 查询尖货活动信息 API返回值
// taobao.opentrade.activity.query
//
// 尖货交易活动信息配置，查询尖货活动信息
type TaobaoOpentradeActivityQueryAPIResponse struct {
	model.CommonResponse
	TaobaoOpentradeActivityQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpentradeActivityQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpentradeActivityQueryAPIResponseModel).Reset()
}

// TaobaoOpentradeActivityQueryAPIResponseModel is 查询尖货活动信息 成功返回结果
type TaobaoOpentradeActivityQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"opentrade_activity_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 活动信息记录
	Results []McSceneActivityDto `json:"results,omitempty" xml:"results>mc_scene_activity_dto,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpentradeActivityQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
	m.TotalCount = 0
}

var poolTaobaoOpentradeActivityQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpentradeActivityQueryAPIResponse)
	},
}

// GetTaobaoOpentradeActivityQueryAPIResponse 从 sync.Pool 获取 TaobaoOpentradeActivityQueryAPIResponse
func GetTaobaoOpentradeActivityQueryAPIResponse() *TaobaoOpentradeActivityQueryAPIResponse {
	return poolTaobaoOpentradeActivityQueryAPIResponse.Get().(*TaobaoOpentradeActivityQueryAPIResponse)
}

// ReleaseTaobaoOpentradeActivityQueryAPIResponse 将 TaobaoOpentradeActivityQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpentradeActivityQueryAPIResponse(v *TaobaoOpentradeActivityQueryAPIResponse) {
	v.Reset()
	poolTaobaoOpentradeActivityQueryAPIResponse.Put(v)
}
