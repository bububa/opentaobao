package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoServiceReceiptGetAPIResponse isv查询服务工单详情 API返回值
// tmall.aliauto.service.receipt.get
//
// isv查询服务工单详情
type TmallAliautoServiceReceiptGetAPIResponse struct {
	model.CommonResponse
	TmallAliautoServiceReceiptGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallAliautoServiceReceiptGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallAliautoServiceReceiptGetAPIResponseModel).Reset()
}

// TmallAliautoServiceReceiptGetAPIResponseModel is isv查询服务工单详情 成功返回结果
type TmallAliautoServiceReceiptGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_service_receipt_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallAliautoServiceReceiptGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallAliautoServiceReceiptGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallAliautoServiceReceiptGetAPIResponse)
	},
}

// GetTmallAliautoServiceReceiptGetAPIResponse 从 sync.Pool 获取 TmallAliautoServiceReceiptGetAPIResponse
func GetTmallAliautoServiceReceiptGetAPIResponse() *TmallAliautoServiceReceiptGetAPIResponse {
	return poolTmallAliautoServiceReceiptGetAPIResponse.Get().(*TmallAliautoServiceReceiptGetAPIResponse)
}

// ReleaseTmallAliautoServiceReceiptGetAPIResponse 将 TmallAliautoServiceReceiptGetAPIResponse 保存到 sync.Pool
func ReleaseTmallAliautoServiceReceiptGetAPIResponse(v *TmallAliautoServiceReceiptGetAPIResponse) {
	v.Reset()
	poolTmallAliautoServiceReceiptGetAPIResponse.Put(v)
}
