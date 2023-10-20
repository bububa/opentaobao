package xhotelonlineorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelCommoninvoiceListVtwoAPIResponse 用户常用发票信息查询接口 API返回值
// taobao.xhotel.commoninvoice.list.vtwo
//
// 获取用户常用发票信息接口
type TaobaoXhotelCommoninvoiceListVtwoAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelCommoninvoiceListVtwoAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelCommoninvoiceListVtwoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelCommoninvoiceListVtwoAPIResponseModel).Reset()
}

// TaobaoXhotelCommoninvoiceListVtwoAPIResponseModel is 用户常用发票信息查询接口 成功返回结果
type TaobaoXhotelCommoninvoiceListVtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_commoninvoice_list_vtwo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果集
	Result *TaobaoXhotelCommoninvoiceListVtwoResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelCommoninvoiceListVtwoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoXhotelCommoninvoiceListVtwoAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelCommoninvoiceListVtwoAPIResponse)
	},
}

// GetTaobaoXhotelCommoninvoiceListVtwoAPIResponse 从 sync.Pool 获取 TaobaoXhotelCommoninvoiceListVtwoAPIResponse
func GetTaobaoXhotelCommoninvoiceListVtwoAPIResponse() *TaobaoXhotelCommoninvoiceListVtwoAPIResponse {
	return poolTaobaoXhotelCommoninvoiceListVtwoAPIResponse.Get().(*TaobaoXhotelCommoninvoiceListVtwoAPIResponse)
}

// ReleaseTaobaoXhotelCommoninvoiceListVtwoAPIResponse 将 TaobaoXhotelCommoninvoiceListVtwoAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelCommoninvoiceListVtwoAPIResponse(v *TaobaoXhotelCommoninvoiceListVtwoAPIResponse) {
	v.Reset()
	poolTaobaoXhotelCommoninvoiceListVtwoAPIResponse.Put(v)
}
