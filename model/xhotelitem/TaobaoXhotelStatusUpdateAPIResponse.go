package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelStatusUpdateAPIResponse 酒店状态更新 API返回值
// taobao.xhotel.status.update
//
// 酒店状态更新
type TaobaoXhotelStatusUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelStatusUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelStatusUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelStatusUpdateAPIResponseModel).Reset()
}

// TaobaoXhotelStatusUpdateAPIResponseModel is 酒店状态更新 成功返回结果
type TaobaoXhotelStatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_status_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否出错
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelStatusUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
}

var poolTaobaoXhotelStatusUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelStatusUpdateAPIResponse)
	},
}

// GetTaobaoXhotelStatusUpdateAPIResponse 从 sync.Pool 获取 TaobaoXhotelStatusUpdateAPIResponse
func GetTaobaoXhotelStatusUpdateAPIResponse() *TaobaoXhotelStatusUpdateAPIResponse {
	return poolTaobaoXhotelStatusUpdateAPIResponse.Get().(*TaobaoXhotelStatusUpdateAPIResponse)
}

// ReleaseTaobaoXhotelStatusUpdateAPIResponse 将 TaobaoXhotelStatusUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelStatusUpdateAPIResponse(v *TaobaoXhotelStatusUpdateAPIResponse) {
	v.Reset()
	poolTaobaoXhotelStatusUpdateAPIResponse.Put(v)
}
