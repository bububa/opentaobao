package tmallhk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallHkClearanceOrderGetAPIResponse 天猫国际订单清关信息 API返回值
// tmall.hk.clearance.order.get
//
// 天猫国际订单清关信息
type TmallHkClearanceOrderGetAPIResponse struct {
	model.CommonResponse
	TmallHkClearanceOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallHkClearanceOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallHkClearanceOrderGetAPIResponseModel).Reset()
}

// TmallHkClearanceOrderGetAPIResponseModel is 天猫国际订单清关信息 成功返回结果
type TmallHkClearanceOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_hk_clearance_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果描述
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// 结果码
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 清关数据封装
	Obj *ClearanceDataDo `json:"obj,omitempty" xml:"obj,omitempty"`
	// 是否正常
	Succeeded bool `json:"succeeded,omitempty" xml:"succeeded,omitempty"`
}

// Reset 清空结构体
func (m *TmallHkClearanceOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResponseMessage = ""
	m.ResponseCode = ""
	m.Obj = nil
	m.Succeeded = false
}

var poolTmallHkClearanceOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallHkClearanceOrderGetAPIResponse)
	},
}

// GetTmallHkClearanceOrderGetAPIResponse 从 sync.Pool 获取 TmallHkClearanceOrderGetAPIResponse
func GetTmallHkClearanceOrderGetAPIResponse() *TmallHkClearanceOrderGetAPIResponse {
	return poolTmallHkClearanceOrderGetAPIResponse.Get().(*TmallHkClearanceOrderGetAPIResponse)
}

// ReleaseTmallHkClearanceOrderGetAPIResponse 将 TmallHkClearanceOrderGetAPIResponse 保存到 sync.Pool
func ReleaseTmallHkClearanceOrderGetAPIResponse(v *TmallHkClearanceOrderGetAPIResponse) {
	v.Reset()
	poolTmallHkClearanceOrderGetAPIResponse.Put(v)
}
