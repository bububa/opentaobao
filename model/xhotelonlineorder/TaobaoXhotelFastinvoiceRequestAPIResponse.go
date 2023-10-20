package xhotelonlineorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelFastinvoiceRequestAPIResponse 极速开票开票请求回传 API返回值
// taobao.xhotel.fastinvoice.request
//
// 极速开票开票请求回传,用于记录航信开票请求数据
type TaobaoXhotelFastinvoiceRequestAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelFastinvoiceRequestAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelFastinvoiceRequestAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelFastinvoiceRequestAPIResponseModel).Reset()
}

// TaobaoXhotelFastinvoiceRequestAPIResponseModel is 极速开票开票请求回传 成功返回结果
type TaobaoXhotelFastinvoiceRequestAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_fastinvoice_request_response"`
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
func (m *TaobaoXhotelFastinvoiceRequestAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errcode = ""
	m.Errmsg = ""
	m.Issuccess = false
}

var poolTaobaoXhotelFastinvoiceRequestAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelFastinvoiceRequestAPIResponse)
	},
}

// GetTaobaoXhotelFastinvoiceRequestAPIResponse 从 sync.Pool 获取 TaobaoXhotelFastinvoiceRequestAPIResponse
func GetTaobaoXhotelFastinvoiceRequestAPIResponse() *TaobaoXhotelFastinvoiceRequestAPIResponse {
	return poolTaobaoXhotelFastinvoiceRequestAPIResponse.Get().(*TaobaoXhotelFastinvoiceRequestAPIResponse)
}

// ReleaseTaobaoXhotelFastinvoiceRequestAPIResponse 将 TaobaoXhotelFastinvoiceRequestAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelFastinvoiceRequestAPIResponse(v *TaobaoXhotelFastinvoiceRequestAPIResponse) {
	v.Reset()
	poolTaobaoXhotelFastinvoiceRequestAPIResponse.Put(v)
}
