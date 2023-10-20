package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpActivityUpdateAPIResponse 修改活动信息 API返回值
// taobao.ump.activity.update
//
// 修改营销活动
type TaobaoUmpActivityUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoUmpActivityUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUmpActivityUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUmpActivityUpdateAPIResponseModel).Reset()
}

// TaobaoUmpActivityUpdateAPIResponseModel is 修改活动信息 成功返回结果
type TaobaoUmpActivityUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_activity_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUmpActivityUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoUmpActivityUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUmpActivityUpdateAPIResponse)
	},
}

// GetTaobaoUmpActivityUpdateAPIResponse 从 sync.Pool 获取 TaobaoUmpActivityUpdateAPIResponse
func GetTaobaoUmpActivityUpdateAPIResponse() *TaobaoUmpActivityUpdateAPIResponse {
	return poolTaobaoUmpActivityUpdateAPIResponse.Get().(*TaobaoUmpActivityUpdateAPIResponse)
}

// ReleaseTaobaoUmpActivityUpdateAPIResponse 将 TaobaoUmpActivityUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUmpActivityUpdateAPIResponse(v *TaobaoUmpActivityUpdateAPIResponse) {
	v.Reset()
	poolTaobaoUmpActivityUpdateAPIResponse.Put(v)
}
