package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceItemUpdateAPIResponse 修改商品开票信息 API返回值
// alibaba.einvoice.item.update
//
// ERP通过接口将商品的开票信息同步给阿里发票平台，自动开票时将读取这些开票信息，需要联系阿里小二开通对应的权限
type AlibabaEinvoiceItemUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceItemUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceItemUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceItemUpdateAPIResponseModel).Reset()
}

// AlibabaEinvoiceItemUpdateAPIResponseModel is 修改商品开票信息 成功返回结果
type AlibabaEinvoiceItemUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_item_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改结果
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceItemUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlibabaEinvoiceItemUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceItemUpdateAPIResponse)
	},
}

// GetAlibabaEinvoiceItemUpdateAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceItemUpdateAPIResponse
func GetAlibabaEinvoiceItemUpdateAPIResponse() *AlibabaEinvoiceItemUpdateAPIResponse {
	return poolAlibabaEinvoiceItemUpdateAPIResponse.Get().(*AlibabaEinvoiceItemUpdateAPIResponse)
}

// ReleaseAlibabaEinvoiceItemUpdateAPIResponse 将 AlibabaEinvoiceItemUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceItemUpdateAPIResponse(v *AlibabaEinvoiceItemUpdateAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceItemUpdateAPIResponse.Put(v)
}
