package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascppresalespackageconsignAPIResponse 预售预包尾款推单发货 API返回值
// alibaba.ascp.presalespackage.consign
//
// 预售预包尾款发货后推单处理
type AlibabaascppresalespackageconsignAPIResponse struct {
	model.CommonResponse
	AlibabaascppresalespackageconsignAPIResponseModel
}

// AlibabaascppresalespackageconsignAPIResponseModel is 预售预包尾款推单发货 成功返回结果
type AlibabaascppresalespackageconsignAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_presalespackage_consign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *ErpPresaleFinalPayResult `json:"result,omitempty" xml:"result,omitempty"`
}
