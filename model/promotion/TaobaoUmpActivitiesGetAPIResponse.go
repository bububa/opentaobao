package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpActivitiesGetAPIResponse 查询活动列表 API返回值
// taobao.ump.activities.get
//
// 查询活动列表
type TaobaoUmpActivitiesGetAPIResponse struct {
	model.CommonResponse
	TaobaoUmpActivitiesGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUmpActivitiesGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUmpActivitiesGetAPIResponseModel).Reset()
}

// TaobaoUmpActivitiesGetAPIResponseModel is 查询活动列表 成功返回结果
type TaobaoUmpActivitiesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_activities_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 营销活动内容，可以通过ump sdk来进行处理
	Contents []string `json:"contents,omitempty" xml:"contents>string,omitempty"`
	// 记录总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUmpActivitiesGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Contents = m.Contents[:0]
	m.TotalCount = 0
}

var poolTaobaoUmpActivitiesGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUmpActivitiesGetAPIResponse)
	},
}

// GetTaobaoUmpActivitiesGetAPIResponse 从 sync.Pool 获取 TaobaoUmpActivitiesGetAPIResponse
func GetTaobaoUmpActivitiesGetAPIResponse() *TaobaoUmpActivitiesGetAPIResponse {
	return poolTaobaoUmpActivitiesGetAPIResponse.Get().(*TaobaoUmpActivitiesGetAPIResponse)
}

// ReleaseTaobaoUmpActivitiesGetAPIResponse 将 TaobaoUmpActivitiesGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUmpActivitiesGetAPIResponse(v *TaobaoUmpActivitiesGetAPIResponse) {
	v.Reset()
	poolTaobaoUmpActivitiesGetAPIResponse.Put(v)
}
