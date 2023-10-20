package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpActivitiesListGetAPIResponse 营销活动列表查询 API返回值
// taobao.ump.activities.list.get
//
// 按照营销活动id的列表ids，查询对应的营销活动列表。
type TaobaoUmpActivitiesListGetAPIResponse struct {
	model.CommonResponse
	TaobaoUmpActivitiesListGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUmpActivitiesListGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUmpActivitiesListGetAPIResponseModel).Reset()
}

// TaobaoUmpActivitiesListGetAPIResponseModel is 营销活动列表查询 成功返回结果
type TaobaoUmpActivitiesListGetAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_activities_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 营销活动列表！
	Activities []string `json:"activities,omitempty" xml:"activities>string,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUmpActivitiesListGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Activities = m.Activities[:0]
}

var poolTaobaoUmpActivitiesListGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUmpActivitiesListGetAPIResponse)
	},
}

// GetTaobaoUmpActivitiesListGetAPIResponse 从 sync.Pool 获取 TaobaoUmpActivitiesListGetAPIResponse
func GetTaobaoUmpActivitiesListGetAPIResponse() *TaobaoUmpActivitiesListGetAPIResponse {
	return poolTaobaoUmpActivitiesListGetAPIResponse.Get().(*TaobaoUmpActivitiesListGetAPIResponse)
}

// ReleaseTaobaoUmpActivitiesListGetAPIResponse 将 TaobaoUmpActivitiesListGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUmpActivitiesListGetAPIResponse(v *TaobaoUmpActivitiesListGetAPIResponse) {
	v.Reset()
	poolTaobaoUmpActivitiesListGetAPIResponse.Put(v)
}
