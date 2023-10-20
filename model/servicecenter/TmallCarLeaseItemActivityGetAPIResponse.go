package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseItemActivityGetAPIResponse 查询汽车租赁活动信息 API返回值
// tmall.car.lease.item.activity.get
//
// 查询汽车租赁活动信息
type TmallCarLeaseItemActivityGetAPIResponse struct {
	model.CommonResponse
	TmallCarLeaseItemActivityGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCarLeaseItemActivityGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCarLeaseItemActivityGetAPIResponseModel).Reset()
}

// TmallCarLeaseItemActivityGetAPIResponseModel is 查询汽车租赁活动信息 成功返回结果
type TmallCarLeaseItemActivityGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_lease_item_activity_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果对象
	Result *TmallCarLeaseItemActivityGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallCarLeaseItemActivityGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallCarLeaseItemActivityGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCarLeaseItemActivityGetAPIResponse)
	},
}

// GetTmallCarLeaseItemActivityGetAPIResponse 从 sync.Pool 获取 TmallCarLeaseItemActivityGetAPIResponse
func GetTmallCarLeaseItemActivityGetAPIResponse() *TmallCarLeaseItemActivityGetAPIResponse {
	return poolTmallCarLeaseItemActivityGetAPIResponse.Get().(*TmallCarLeaseItemActivityGetAPIResponse)
}

// ReleaseTmallCarLeaseItemActivityGetAPIResponse 将 TmallCarLeaseItemActivityGetAPIResponse 保存到 sync.Pool
func ReleaseTmallCarLeaseItemActivityGetAPIResponse(v *TmallCarLeaseItemActivityGetAPIResponse) {
	v.Reset()
	poolTmallCarLeaseItemActivityGetAPIResponse.Put(v)
}
