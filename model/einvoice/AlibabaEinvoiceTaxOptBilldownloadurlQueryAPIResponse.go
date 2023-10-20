package einvoice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoicetaxoptbilldownloadurlqueryAPIResponse 税筹业务账单文件下载URL查询 API返回值
// alibaba.einvoice.tax.opt.billdownloadurl.query
//
// 税筹业务账单文件下载的URL查询
type AlibabaeinvoicetaxoptbilldownloadurlqueryAPIResponse struct {
	model.CommonResponse
	AlibabaeinvoicetaxoptbilldownloadurlqueryAPIResponseModel
}

// AlibabaeinvoicetaxoptbilldownloadurlqueryAPIResponseModel is 税筹业务账单文件下载URL查询 成功返回结果
type AlibabaeinvoicetaxoptbilldownloadurlqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_tax_opt_billdownloadurl_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
