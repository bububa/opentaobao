package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbOrderJzQueryAPIResponse 家装业务查询物流公司api API返回值
// taobao.wlb.order.jz.query
//
// 家装业务查询物流公司api
type TaobaoWlbOrderJzQueryAPIResponse struct {
	model.CommonResponse
	TaobaoWlbOrderJzQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbOrderJzQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbOrderJzQueryAPIResponseModel).Reset()
}

// TaobaoWlbOrderJzQueryAPIResponseModel is 家装业务查询物流公司api 成功返回结果
type TaobaoWlbOrderJzQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_order_jz_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误编码
	ResultErrorCode string `json:"result_error_code,omitempty" xml:"result_error_code,omitempty"`
	// 错误信息
	ResultErrorMsg string `json:"result_error_msg,omitempty" xml:"result_error_msg,omitempty"`
	// 结果信息
	Result *JzTopDto `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbOrderJzQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultErrorCode = ""
	m.ResultErrorMsg = ""
	m.Result = nil
	m.ResultSuccess = false
}

var poolTaobaoWlbOrderJzQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbOrderJzQueryAPIResponse)
	},
}

// GetTaobaoWlbOrderJzQueryAPIResponse 从 sync.Pool 获取 TaobaoWlbOrderJzQueryAPIResponse
func GetTaobaoWlbOrderJzQueryAPIResponse() *TaobaoWlbOrderJzQueryAPIResponse {
	return poolTaobaoWlbOrderJzQueryAPIResponse.Get().(*TaobaoWlbOrderJzQueryAPIResponse)
}

// ReleaseTaobaoWlbOrderJzQueryAPIResponse 将 TaobaoWlbOrderJzQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbOrderJzQueryAPIResponse(v *TaobaoWlbOrderJzQueryAPIResponse) {
	v.Reset()
	poolTaobaoWlbOrderJzQueryAPIResponse.Put(v)
}
