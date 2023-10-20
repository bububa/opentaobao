package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayKeywordOfflineLayeredfindAPIResponse 查询关键词离线报表30天转化周期 API返回值
// taobao.subway.keyword.offline.layeredfind
//
// 获取关键词离线报表
type TaobaoSubwayKeywordOfflineLayeredfindAPIResponse struct {
	model.CommonResponse
	TaobaoSubwayKeywordOfflineLayeredfindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSubwayKeywordOfflineLayeredfindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSubwayKeywordOfflineLayeredfindAPIResponseModel).Reset()
}

// TaobaoSubwayKeywordOfflineLayeredfindAPIResponseModel is 查询关键词离线报表30天转化周期 成功返回结果
type TaobaoSubwayKeywordOfflineLayeredfindAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_keyword_offline_layeredfind_response"`
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
func (m *TaobaoSubwayKeywordOfflineLayeredfindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = m.Result[:0]
	m.Message = ""
	m.TotalCount = 0
}

var poolTaobaoSubwayKeywordOfflineLayeredfindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSubwayKeywordOfflineLayeredfindAPIResponse)
	},
}

// GetTaobaoSubwayKeywordOfflineLayeredfindAPIResponse 从 sync.Pool 获取 TaobaoSubwayKeywordOfflineLayeredfindAPIResponse
func GetTaobaoSubwayKeywordOfflineLayeredfindAPIResponse() *TaobaoSubwayKeywordOfflineLayeredfindAPIResponse {
	return poolTaobaoSubwayKeywordOfflineLayeredfindAPIResponse.Get().(*TaobaoSubwayKeywordOfflineLayeredfindAPIResponse)
}

// ReleaseTaobaoSubwayKeywordOfflineLayeredfindAPIResponse 将 TaobaoSubwayKeywordOfflineLayeredfindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSubwayKeywordOfflineLayeredfindAPIResponse(v *TaobaoSubwayKeywordOfflineLayeredfindAPIResponse) {
	v.Reset()
	poolTaobaoSubwayKeywordOfflineLayeredfindAPIResponse.Put(v)
}
