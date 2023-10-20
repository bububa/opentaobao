package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayAdgroupOfflineLayeredfindAPIResponse 查询单元离线列表30天转化周期 API返回值
// taobao.subway.adgroup.offline.layeredfind
//
// 查询单元离线列表
type TaobaoSubwayAdgroupOfflineLayeredfindAPIResponse struct {
	model.CommonResponse
	TaobaoSubwayAdgroupOfflineLayeredfindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSubwayAdgroupOfflineLayeredfindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSubwayAdgroupOfflineLayeredfindAPIResponseModel).Reset()
}

// TaobaoSubwayAdgroupOfflineLayeredfindAPIResponseModel is 查询单元离线列表30天转化周期 成功返回结果
type TaobaoSubwayAdgroupOfflineLayeredfindAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_adgroup_offline_layeredfind_response"`
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
func (m *TaobaoSubwayAdgroupOfflineLayeredfindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = m.Result[:0]
	m.Message = ""
	m.TotalCount = 0
}

var poolTaobaoSubwayAdgroupOfflineLayeredfindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSubwayAdgroupOfflineLayeredfindAPIResponse)
	},
}

// GetTaobaoSubwayAdgroupOfflineLayeredfindAPIResponse 从 sync.Pool 获取 TaobaoSubwayAdgroupOfflineLayeredfindAPIResponse
func GetTaobaoSubwayAdgroupOfflineLayeredfindAPIResponse() *TaobaoSubwayAdgroupOfflineLayeredfindAPIResponse {
	return poolTaobaoSubwayAdgroupOfflineLayeredfindAPIResponse.Get().(*TaobaoSubwayAdgroupOfflineLayeredfindAPIResponse)
}

// ReleaseTaobaoSubwayAdgroupOfflineLayeredfindAPIResponse 将 TaobaoSubwayAdgroupOfflineLayeredfindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSubwayAdgroupOfflineLayeredfindAPIResponse(v *TaobaoSubwayAdgroupOfflineLayeredfindAPIResponse) {
	v.Reset()
	poolTaobaoSubwayAdgroupOfflineLayeredfindAPIResponse.Put(v)
}
