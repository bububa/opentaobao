package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFuwuSpBillreordAddAPIResponse 内购服务确认单明细上传接口 API返回值
// taobao.fuwu.sp.billreord.add
//
// isv能通过该接口上传确认单明细数据
type TaobaoFuwuSpBillreordAddAPIResponse struct {
	model.CommonResponse
	TaobaoFuwuSpBillreordAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFuwuSpBillreordAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFuwuSpBillreordAddAPIResponseModel).Reset()
}

// TaobaoFuwuSpBillreordAddAPIResponseModel is 内购服务确认单明细上传接口 成功返回结果
type TaobaoFuwuSpBillreordAddAPIResponseModel struct {
	XMLName xml.Name `xml:"fuwu_sp_billreord_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回调用结果
	AddResult bool `json:"add_result,omitempty" xml:"add_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFuwuSpBillreordAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AddResult = false
}

var poolTaobaoFuwuSpBillreordAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFuwuSpBillreordAddAPIResponse)
	},
}

// GetTaobaoFuwuSpBillreordAddAPIResponse 从 sync.Pool 获取 TaobaoFuwuSpBillreordAddAPIResponse
func GetTaobaoFuwuSpBillreordAddAPIResponse() *TaobaoFuwuSpBillreordAddAPIResponse {
	return poolTaobaoFuwuSpBillreordAddAPIResponse.Get().(*TaobaoFuwuSpBillreordAddAPIResponse)
}

// ReleaseTaobaoFuwuSpBillreordAddAPIResponse 将 TaobaoFuwuSpBillreordAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFuwuSpBillreordAddAPIResponse(v *TaobaoFuwuSpBillreordAddAPIResponse) {
	v.Reset()
	poolTaobaoFuwuSpBillreordAddAPIResponse.Put(v)
}
