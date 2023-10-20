package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayKeywordOfflineFindAPIResponse 查询关键词离线多日汇总报表 API返回值
// taobao.subway.keyword.offline.find
//
// 获取关键词离线报表
type TaobaoSubwayKeywordOfflineFindAPIResponse struct {
	model.CommonResponse
	TaobaoSubwayKeywordOfflineFindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSubwayKeywordOfflineFindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSubwayKeywordOfflineFindAPIResponseModel).Reset()
}

// TaobaoSubwayKeywordOfflineFindAPIResponseModel is 查询关键词离线多日汇总报表 成功返回结果
type TaobaoSubwayKeywordOfflineFindAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_keyword_offline_find_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result []ReportResultTopDto `json:"result,omitempty" xml:"result>report_result_top_dto,omitempty"`
	// 提示
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSubwayKeywordOfflineFindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = m.Result[:0]
	m.Message = ""
	m.TotalCount = 0
}

var poolTaobaoSubwayKeywordOfflineFindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSubwayKeywordOfflineFindAPIResponse)
	},
}

// GetTaobaoSubwayKeywordOfflineFindAPIResponse 从 sync.Pool 获取 TaobaoSubwayKeywordOfflineFindAPIResponse
func GetTaobaoSubwayKeywordOfflineFindAPIResponse() *TaobaoSubwayKeywordOfflineFindAPIResponse {
	return poolTaobaoSubwayKeywordOfflineFindAPIResponse.Get().(*TaobaoSubwayKeywordOfflineFindAPIResponse)
}

// ReleaseTaobaoSubwayKeywordOfflineFindAPIResponse 将 TaobaoSubwayKeywordOfflineFindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSubwayKeywordOfflineFindAPIResponse(v *TaobaoSubwayKeywordOfflineFindAPIResponse) {
	v.Reset()
	poolTaobaoSubwayKeywordOfflineFindAPIResponse.Put(v)
}
