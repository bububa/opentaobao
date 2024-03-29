package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpRangeDeleteAPIResponse 删除活动范围 API返回值
// taobao.ump.range.delete
//
// 去指先前指定在某项活动中，某些类型的物品参加或者不参加活动的设置
type TaobaoUmpRangeDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoUmpRangeDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUmpRangeDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUmpRangeDeleteAPIResponseModel).Reset()
}

// TaobaoUmpRangeDeleteAPIResponseModel is 删除活动范围 成功返回结果
type TaobaoUmpRangeDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_range_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUmpRangeDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoUmpRangeDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUmpRangeDeleteAPIResponse)
	},
}

// GetTaobaoUmpRangeDeleteAPIResponse 从 sync.Pool 获取 TaobaoUmpRangeDeleteAPIResponse
func GetTaobaoUmpRangeDeleteAPIResponse() *TaobaoUmpRangeDeleteAPIResponse {
	return poolTaobaoUmpRangeDeleteAPIResponse.Get().(*TaobaoUmpRangeDeleteAPIResponse)
}

// ReleaseTaobaoUmpRangeDeleteAPIResponse 将 TaobaoUmpRangeDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUmpRangeDeleteAPIResponse(v *TaobaoUmpRangeDeleteAPIResponse) {
	v.Reset()
	poolTaobaoUmpRangeDeleteAPIResponse.Put(v)
}
