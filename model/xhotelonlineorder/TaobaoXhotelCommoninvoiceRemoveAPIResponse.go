package xhotelonlineorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelCommoninvoiceRemoveAPIResponse 常用发票信息删除接口 API返回值
// taobao.xhotel.commoninvoice.remove
//
// 常用发票信息删除接口
type TaobaoXhotelCommoninvoiceRemoveAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelCommoninvoiceRemoveAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelCommoninvoiceRemoveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelCommoninvoiceRemoveAPIResponseModel).Reset()
}

// TaobaoXhotelCommoninvoiceRemoveAPIResponseModel is 常用发票信息删除接口 成功返回结果
type TaobaoXhotelCommoninvoiceRemoveAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_commoninvoice_remove_response"`
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
func (m *TaobaoXhotelCommoninvoiceRemoveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errcode = ""
	m.Errmsg = ""
	m.Issuccess = false
}

var poolTaobaoXhotelCommoninvoiceRemoveAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelCommoninvoiceRemoveAPIResponse)
	},
}

// GetTaobaoXhotelCommoninvoiceRemoveAPIResponse 从 sync.Pool 获取 TaobaoXhotelCommoninvoiceRemoveAPIResponse
func GetTaobaoXhotelCommoninvoiceRemoveAPIResponse() *TaobaoXhotelCommoninvoiceRemoveAPIResponse {
	return poolTaobaoXhotelCommoninvoiceRemoveAPIResponse.Get().(*TaobaoXhotelCommoninvoiceRemoveAPIResponse)
}

// ReleaseTaobaoXhotelCommoninvoiceRemoveAPIResponse 将 TaobaoXhotelCommoninvoiceRemoveAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelCommoninvoiceRemoveAPIResponse(v *TaobaoXhotelCommoninvoiceRemoveAPIResponse) {
	v.Reset()
	poolTaobaoXhotelCommoninvoiceRemoveAPIResponse.Put(v)
}
