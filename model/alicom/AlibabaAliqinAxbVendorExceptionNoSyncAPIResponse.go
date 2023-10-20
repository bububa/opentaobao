package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinAxbVendorExceptionNoSyncAPIResponse 中心化供应商异常号码状态同步接口 API返回值
// alibaba.aliqin.axb.vendor.exception.no.sync
//
// 用于中心化供应商同步异常号码
type AlibabaAliqinAxbVendorExceptionNoSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinAxbVendorExceptionNoSyncAPIResponseModel
}

// AlibabaAliqinAxbVendorExceptionNoSyncAPIResponseModel is 中心化供应商异常号码状态同步接口 成功返回结果
type AlibabaAliqinAxbVendorExceptionNoSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_axb_vendor_exception_no_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAliqinAxbVendorExceptionNoSyncResponse `json:"result,omitempty" xml:"result,omitempty"`
}
