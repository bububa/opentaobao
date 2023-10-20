package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpActivityDeleteAPIResponse 删除营销活动 API返回值
// taobao.ump.activity.delete
//
// 删除营销活动。对应的活动详情等将会被全部删除。
type TaobaoUmpActivityDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoUmpActivityDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUmpActivityDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUmpActivityDeleteAPIResponseModel).Reset()
}

// TaobaoUmpActivityDeleteAPIResponseModel is 删除营销活动 成功返回结果
type TaobaoUmpActivityDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_activity_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUmpActivityDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoUmpActivityDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUmpActivityDeleteAPIResponse)
	},
}

// GetTaobaoUmpActivityDeleteAPIResponse 从 sync.Pool 获取 TaobaoUmpActivityDeleteAPIResponse
func GetTaobaoUmpActivityDeleteAPIResponse() *TaobaoUmpActivityDeleteAPIResponse {
	return poolTaobaoUmpActivityDeleteAPIResponse.Get().(*TaobaoUmpActivityDeleteAPIResponse)
}

// ReleaseTaobaoUmpActivityDeleteAPIResponse 将 TaobaoUmpActivityDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUmpActivityDeleteAPIResponse(v *TaobaoUmpActivityDeleteAPIResponse) {
	v.Reset()
	poolTaobaoUmpActivityDeleteAPIResponse.Put(v)
}
