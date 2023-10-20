package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoReceiptStateUpdateAPIResponse 服务工单状态更新 API返回值
// tmall.aliauto.receipt.state.update
//
// 二轮车服务工单状态更新
type TmallAliautoReceiptStateUpdateAPIResponse struct {
	model.CommonResponse
	TmallAliautoReceiptStateUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallAliautoReceiptStateUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallAliautoReceiptStateUpdateAPIResponseModel).Reset()
}

// TmallAliautoReceiptStateUpdateAPIResponseModel is 服务工单状态更新 成功返回结果
type TmallAliautoReceiptStateUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_receipt_state_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallAliautoReceiptStateUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallAliautoReceiptStateUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallAliautoReceiptStateUpdateAPIResponse)
	},
}

// GetTmallAliautoReceiptStateUpdateAPIResponse 从 sync.Pool 获取 TmallAliautoReceiptStateUpdateAPIResponse
func GetTmallAliautoReceiptStateUpdateAPIResponse() *TmallAliautoReceiptStateUpdateAPIResponse {
	return poolTmallAliautoReceiptStateUpdateAPIResponse.Get().(*TmallAliautoReceiptStateUpdateAPIResponse)
}

// ReleaseTmallAliautoReceiptStateUpdateAPIResponse 将 TmallAliautoReceiptStateUpdateAPIResponse 保存到 sync.Pool
func ReleaseTmallAliautoReceiptStateUpdateAPIResponse(v *TmallAliautoReceiptStateUpdateAPIResponse) {
	v.Reset()
	poolTmallAliautoReceiptStateUpdateAPIResponse.Put(v)
}
