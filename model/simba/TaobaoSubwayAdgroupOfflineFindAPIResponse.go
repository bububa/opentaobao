package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayAdgroupOfflineFindAPIResponse 查询单元离线多日汇总列表 API返回值
// taobao.subway.adgroup.offline.find
//
// 查询单元离线列表
type TaobaoSubwayAdgroupOfflineFindAPIResponse struct {
	model.CommonResponse
	TaobaoSubwayAdgroupOfflineFindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSubwayAdgroupOfflineFindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSubwayAdgroupOfflineFindAPIResponseModel).Reset()
}

// TaobaoSubwayAdgroupOfflineFindAPIResponseModel is 查询单元离线多日汇总列表 成功返回结果
type TaobaoSubwayAdgroupOfflineFindAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_adgroup_offline_find_response"`
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
func (m *TaobaoSubwayAdgroupOfflineFindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = m.Result[:0]
	m.Message = ""
	m.TotalCount = 0
}

var poolTaobaoSubwayAdgroupOfflineFindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSubwayAdgroupOfflineFindAPIResponse)
	},
}

// GetTaobaoSubwayAdgroupOfflineFindAPIResponse 从 sync.Pool 获取 TaobaoSubwayAdgroupOfflineFindAPIResponse
func GetTaobaoSubwayAdgroupOfflineFindAPIResponse() *TaobaoSubwayAdgroupOfflineFindAPIResponse {
	return poolTaobaoSubwayAdgroupOfflineFindAPIResponse.Get().(*TaobaoSubwayAdgroupOfflineFindAPIResponse)
}

// ReleaseTaobaoSubwayAdgroupOfflineFindAPIResponse 将 TaobaoSubwayAdgroupOfflineFindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSubwayAdgroupOfflineFindAPIResponse(v *TaobaoSubwayAdgroupOfflineFindAPIResponse) {
	v.Reset()
	poolTaobaoSubwayAdgroupOfflineFindAPIResponse.Put(v)
}
