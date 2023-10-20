package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanItemsUnsubscribeByConditionAPIResponse 根据条件删除订阅关系 API返回值
// taobao.baichuan.items.unsubscribe.by.condition
//
// 根据条件删除订阅关系
type TaobaoBaichuanItemsUnsubscribeByConditionAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanItemsUnsubscribeByConditionAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaichuanItemsUnsubscribeByConditionAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaichuanItemsUnsubscribeByConditionAPIResponseModel).Reset()
}

// TaobaoBaichuanItemsUnsubscribeByConditionAPIResponseModel is 根据条件删除订阅关系 成功返回结果
type TaobaoBaichuanItemsUnsubscribeByConditionAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_items_unsubscribe_by_condition_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoBaichuanItemsUnsubscribeByConditionResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaichuanItemsUnsubscribeByConditionAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoBaichuanItemsUnsubscribeByConditionAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanItemsUnsubscribeByConditionAPIResponse)
	},
}

// GetTaobaoBaichuanItemsUnsubscribeByConditionAPIResponse 从 sync.Pool 获取 TaobaoBaichuanItemsUnsubscribeByConditionAPIResponse
func GetTaobaoBaichuanItemsUnsubscribeByConditionAPIResponse() *TaobaoBaichuanItemsUnsubscribeByConditionAPIResponse {
	return poolTaobaoBaichuanItemsUnsubscribeByConditionAPIResponse.Get().(*TaobaoBaichuanItemsUnsubscribeByConditionAPIResponse)
}

// ReleaseTaobaoBaichuanItemsUnsubscribeByConditionAPIResponse 将 TaobaoBaichuanItemsUnsubscribeByConditionAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaichuanItemsUnsubscribeByConditionAPIResponse(v *TaobaoBaichuanItemsUnsubscribeByConditionAPIResponse) {
	v.Reset()
	poolTaobaoBaichuanItemsUnsubscribeByConditionAPIResponse.Put(v)
}
