package xhotelonlineorder

import (
	"encoding/xml"

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

// TaobaoXhotelFastinvoiceRequestAPIResponseModel is 极速开票开票请求回传 成功返回结果
type TaobaoXhotelFastinvoiceRequestAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_fastinvoice_request_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// success
	Issuccess bool `json:"issuccess,omitempty" xml:"issuccess,omitempty"`
	// errorCode
	Errcode string `json:"errcode,omitempty" xml:"errcode,omitempty"`
	// errorMsg
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
}
