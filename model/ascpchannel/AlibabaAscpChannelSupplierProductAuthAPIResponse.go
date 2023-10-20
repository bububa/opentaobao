package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelSupplierProductAuthAPIResponse 供应商授权渠道产品到市场或分销商 API返回值
// alibaba.ascp.channel.supplier.product.auth
//
// 供应商授权渠道产品到市场或分销商
type AlibabaAscpChannelSupplierProductAuthAPIResponse struct {
	model.CommonResponse
	AlibabaAscpChannelSupplierProductAuthAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpChannelSupplierProductAuthAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpChannelSupplierProductAuthAPIResponseModel).Reset()
}

// AlibabaAscpChannelSupplierProductAuthAPIResponseModel is 供应商授权渠道产品到市场或分销商 成功返回结果
type AlibabaAscpChannelSupplierProductAuthAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_supplier_product_auth_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpChannelSupplierProductAuthAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpChannelSupplierProductAuthAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpChannelSupplierProductAuthAPIResponse)
	},
}

// GetAlibabaAscpChannelSupplierProductAuthAPIResponse 从 sync.Pool 获取 AlibabaAscpChannelSupplierProductAuthAPIResponse
func GetAlibabaAscpChannelSupplierProductAuthAPIResponse() *AlibabaAscpChannelSupplierProductAuthAPIResponse {
	return poolAlibabaAscpChannelSupplierProductAuthAPIResponse.Get().(*AlibabaAscpChannelSupplierProductAuthAPIResponse)
}

// ReleaseAlibabaAscpChannelSupplierProductAuthAPIResponse 将 AlibabaAscpChannelSupplierProductAuthAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpChannelSupplierProductAuthAPIResponse(v *AlibabaAscpChannelSupplierProductAuthAPIResponse) {
	v.Reset()
	poolAlibabaAscpChannelSupplierProductAuthAPIResponse.Put(v)
}
