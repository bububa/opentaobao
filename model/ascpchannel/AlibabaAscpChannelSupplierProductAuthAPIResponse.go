package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpchannelsupplierproductauthAPIResponse 供应商授权渠道产品到市场或分销商 API返回值
// alibaba.ascp.channel.supplier.product.auth
//
// 供应商授权渠道产品到市场或分销商
type AlibabaascpchannelsupplierproductauthAPIResponse struct {
	model.CommonResponse
	AlibabaascpchannelsupplierproductauthAPIResponseModel
}

// AlibabaascpchannelsupplierproductauthAPIResponseModel is 供应商授权渠道产品到市场或分销商 成功返回结果
type AlibabaascpchannelsupplierproductauthAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_supplier_product_auth_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}
