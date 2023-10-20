package wlbimports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbImportsOrderCancelAPIResponse 一般进口取消物流订单 API返回值
// taobao.wlb.imports.order.cancel
//
// 商家在发货后，需要对订单进行取消，如果仓库已经收货则无法取消。
type TaobaoWlbImportsOrderCancelAPIResponse struct {
	model.CommonResponse
	TaobaoWlbImportsOrderCancelAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbImportsOrderCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbImportsOrderCancelAPIResponseModel).Reset()
}

// TaobaoWlbImportsOrderCancelAPIResponseModel is 一般进口取消物流订单 成功返回结果
type TaobaoWlbImportsOrderCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_imports_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务错误描述
	ResultErrorMsg string `json:"result_error_msg,omitempty" xml:"result_error_msg,omitempty"`
	// 业务错误编码
	ResultErrorCode string `json:"result_error_code,omitempty" xml:"result_error_code,omitempty"`
	// 是否取消订单成功，true：成功，false：失败
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbImportsOrderCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultErrorMsg = ""
	m.ResultErrorCode = ""
	m.IsSuccess = false
}

var poolTaobaoWlbImportsOrderCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbImportsOrderCancelAPIResponse)
	},
}

// GetTaobaoWlbImportsOrderCancelAPIResponse 从 sync.Pool 获取 TaobaoWlbImportsOrderCancelAPIResponse
func GetTaobaoWlbImportsOrderCancelAPIResponse() *TaobaoWlbImportsOrderCancelAPIResponse {
	return poolTaobaoWlbImportsOrderCancelAPIResponse.Get().(*TaobaoWlbImportsOrderCancelAPIResponse)
}

// ReleaseTaobaoWlbImportsOrderCancelAPIResponse 将 TaobaoWlbImportsOrderCancelAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbImportsOrderCancelAPIResponse(v *TaobaoWlbImportsOrderCancelAPIResponse) {
	v.Reset()
	poolTaobaoWlbImportsOrderCancelAPIResponse.Put(v)
}
