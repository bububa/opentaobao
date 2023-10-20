package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbOrderJzConsignAPIResponse 家装发货接口 API返回值
// taobao.wlb.order.jz.consign
//
// 家装类订单使用该接口发货
type TaobaoWlbOrderJzConsignAPIResponse struct {
	model.CommonResponse
	TaobaoWlbOrderJzConsignAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbOrderJzConsignAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbOrderJzConsignAPIResponseModel).Reset()
}

// TaobaoWlbOrderJzConsignAPIResponseModel is 家装发货接口 成功返回结果
type TaobaoWlbOrderJzConsignAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_order_jz_consign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultErrorCode string `json:"result_error_code,omitempty" xml:"result_error_code,omitempty"`
	// 错误信息描述
	ResultErrorMsg string `json:"result_error_msg,omitempty" xml:"result_error_msg,omitempty"`
	// 是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbOrderJzConsignAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultErrorCode = ""
	m.ResultErrorMsg = ""
	m.ResultSuccess = false
}

var poolTaobaoWlbOrderJzConsignAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbOrderJzConsignAPIResponse)
	},
}

// GetTaobaoWlbOrderJzConsignAPIResponse 从 sync.Pool 获取 TaobaoWlbOrderJzConsignAPIResponse
func GetTaobaoWlbOrderJzConsignAPIResponse() *TaobaoWlbOrderJzConsignAPIResponse {
	return poolTaobaoWlbOrderJzConsignAPIResponse.Get().(*TaobaoWlbOrderJzConsignAPIResponse)
}

// ReleaseTaobaoWlbOrderJzConsignAPIResponse 将 TaobaoWlbOrderJzConsignAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbOrderJzConsignAPIResponse(v *TaobaoWlbOrderJzConsignAPIResponse) {
	v.Reset()
	poolTaobaoWlbOrderJzConsignAPIResponse.Put(v)
}
