package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgNewuserOrderGetAPIResponse 淘宝客-推广者-新用户订单明细查询 API返回值
// taobao.tbk.dg.newuser.order.get
//
// 拉新API
type TaobaoTbkDgNewuserOrderGetAPIResponse struct {
	model.CommonResponse
	TaobaoTbkDgNewuserOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkDgNewuserOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkDgNewuserOrderGetAPIResponseModel).Reset()
}

// TaobaoTbkDgNewuserOrderGetAPIResponseModel is 淘宝客-推广者-新用户订单明细查询 成功返回结果
type TaobaoTbkDgNewuserOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_dg_newuser_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Results *TaobaoTbkDgNewuserOrderGetResults `json:"results,omitempty" xml:"results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkDgNewuserOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = nil
}

var poolTaobaoTbkDgNewuserOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkDgNewuserOrderGetAPIResponse)
	},
}

// GetTaobaoTbkDgNewuserOrderGetAPIResponse 从 sync.Pool 获取 TaobaoTbkDgNewuserOrderGetAPIResponse
func GetTaobaoTbkDgNewuserOrderGetAPIResponse() *TaobaoTbkDgNewuserOrderGetAPIResponse {
	return poolTaobaoTbkDgNewuserOrderGetAPIResponse.Get().(*TaobaoTbkDgNewuserOrderGetAPIResponse)
}

// ReleaseTaobaoTbkDgNewuserOrderGetAPIResponse 将 TaobaoTbkDgNewuserOrderGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkDgNewuserOrderGetAPIResponse(v *TaobaoTbkDgNewuserOrderGetAPIResponse) {
	v.Reset()
	poolTaobaoTbkDgNewuserOrderGetAPIResponse.Put(v)
}
