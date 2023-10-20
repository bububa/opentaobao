package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayCrowdOfflineLayeredfindAPIResponse 获取人群离线报表30转化周期 API返回值
// taobao.subway.crowd.offline.layeredfind
//
// 获取人群离线报表
type TaobaoSubwayCrowdOfflineLayeredfindAPIResponse struct {
	model.CommonResponse
	TaobaoSubwayCrowdOfflineLayeredfindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSubwayCrowdOfflineLayeredfindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSubwayCrowdOfflineLayeredfindAPIResponseModel).Reset()
}

// TaobaoSubwayCrowdOfflineLayeredfindAPIResponseModel is 获取人群离线报表30转化周期 成功返回结果
type TaobaoSubwayCrowdOfflineLayeredfindAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_crowd_offline_layeredfind_response"`
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
func (m *TaobaoSubwayCrowdOfflineLayeredfindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = m.Result[:0]
	m.Message = ""
	m.TotalCount = 0
}

var poolTaobaoSubwayCrowdOfflineLayeredfindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSubwayCrowdOfflineLayeredfindAPIResponse)
	},
}

// GetTaobaoSubwayCrowdOfflineLayeredfindAPIResponse 从 sync.Pool 获取 TaobaoSubwayCrowdOfflineLayeredfindAPIResponse
func GetTaobaoSubwayCrowdOfflineLayeredfindAPIResponse() *TaobaoSubwayCrowdOfflineLayeredfindAPIResponse {
	return poolTaobaoSubwayCrowdOfflineLayeredfindAPIResponse.Get().(*TaobaoSubwayCrowdOfflineLayeredfindAPIResponse)
}

// ReleaseTaobaoSubwayCrowdOfflineLayeredfindAPIResponse 将 TaobaoSubwayCrowdOfflineLayeredfindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSubwayCrowdOfflineLayeredfindAPIResponse(v *TaobaoSubwayCrowdOfflineLayeredfindAPIResponse) {
	v.Reset()
	poolTaobaoSubwayCrowdOfflineLayeredfindAPIResponse.Put(v)
}
