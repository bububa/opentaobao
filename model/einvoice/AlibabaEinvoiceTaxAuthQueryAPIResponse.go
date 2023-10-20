package einvoice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicetaxauthqueryAPIResponse 发票中台授权信息获取 API返回值
// alibaba.einvoice.tax.auth.query
//
// 发票中台授权信息获取
type AlibabaeinvoicetaxauthqueryAPIResponse struct {
	model.CommonResponse
	AlibabaeinvoicetaxauthqueryAPIResponseModel
}

// AlibabaeinvoicetaxauthqueryAPIResponseModel is 发票中台授权信息获取 成功返回结果
type AlibabaeinvoicetaxauthqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_tax_auth_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
