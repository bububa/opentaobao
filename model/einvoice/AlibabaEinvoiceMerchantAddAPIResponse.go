package einvoice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceMerchantAddAPIResponse
发票中台-同平台授权税号适用商户 API返回值
alibaba.einvoice.merchant.add

适用于以下场景：
业务税号入驻成功后，需要将税号授权给同平台下其他商户，使得其他商户也具备开票能力 */
type AlibabaEinvoiceMerchantAddAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceMerchantAddAPIResponseModel
}

// AlibabaEinvoiceMerchantAddAPIResponseModel is 发票中台-同平台授权税号适用商户 成功返回结果
type AlibabaEinvoiceMerchantAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_merchant_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 新增成功的业务平台门店ID
	MerchantUserId string `json:"merchant_user_id,omitempty" xml:"merchant_user_id,omitempty"`
}
