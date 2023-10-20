package xhotelonlineorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelFastinvoiceCompleteAPIResponse 极速开票开票请求完成 API返回值
// taobao.xhotel.fastinvoice.complete
//
// 极速开票开票请求回传,用于更新航信开票请求数据
type TaobaoXhotelFastinvoiceCompleteAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelFastinvoiceCompleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelFastinvoiceCompleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelFastinvoiceCompleteAPIResponseModel).Reset()
}

// TaobaoXhotelFastinvoiceCompleteAPIResponseModel is 极速开票开票请求完成 成功返回结果
type TaobaoXhotelFastinvoiceCompleteAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_fastinvoice_complete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errorCode
	Errcode string `json:"errcode,omitempty" xml:"errcode,omitempty"`
	// errorMsg
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// success
	Issuccess bool `json:"issuccess,omitempty" xml:"issuccess,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelFastinvoiceCompleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errcode = ""
	m.Errmsg = ""
	m.Issuccess = false
}

var poolTaobaoXhotelFastinvoiceCompleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelFastinvoiceCompleteAPIResponse)
	},
}

// GetTaobaoXhotelFastinvoiceCompleteAPIResponse 从 sync.Pool 获取 TaobaoXhotelFastinvoiceCompleteAPIResponse
func GetTaobaoXhotelFastinvoiceCompleteAPIResponse() *TaobaoXhotelFastinvoiceCompleteAPIResponse {
	return poolTaobaoXhotelFastinvoiceCompleteAPIResponse.Get().(*TaobaoXhotelFastinvoiceCompleteAPIResponse)
}

// ReleaseTaobaoXhotelFastinvoiceCompleteAPIResponse 将 TaobaoXhotelFastinvoiceCompleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelFastinvoiceCompleteAPIResponse(v *TaobaoXhotelFastinvoiceCompleteAPIResponse) {
	v.Reset()
	poolTaobaoXhotelFastinvoiceCompleteAPIResponse.Put(v)
}
