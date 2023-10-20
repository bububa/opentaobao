package alicom

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaAliqinAxbVendorExceptionNoSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinAxbVendorExceptionNoSyncAPIResponseModel).Reset()
}

// AlibabaAliqinAxbVendorExceptionNoSyncAPIResponseModel is 中心化供应商异常号码状态同步接口 成功返回结果
type AlibabaAliqinAxbVendorExceptionNoSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_axb_vendor_exception_no_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAliqinAxbVendorExceptionNoSyncResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinAxbVendorExceptionNoSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliqinAxbVendorExceptionNoSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinAxbVendorExceptionNoSyncAPIResponse)
	},
}

// GetAlibabaAliqinAxbVendorExceptionNoSyncAPIResponse 从 sync.Pool 获取 AlibabaAliqinAxbVendorExceptionNoSyncAPIResponse
func GetAlibabaAliqinAxbVendorExceptionNoSyncAPIResponse() *AlibabaAliqinAxbVendorExceptionNoSyncAPIResponse {
	return poolAlibabaAliqinAxbVendorExceptionNoSyncAPIResponse.Get().(*AlibabaAliqinAxbVendorExceptionNoSyncAPIResponse)
}

// ReleaseAlibabaAliqinAxbVendorExceptionNoSyncAPIResponse 将 AlibabaAliqinAxbVendorExceptionNoSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinAxbVendorExceptionNoSyncAPIResponse(v *AlibabaAliqinAxbVendorExceptionNoSyncAPIResponse) {
	v.Reset()
	poolAlibabaAliqinAxbVendorExceptionNoSyncAPIResponse.Put(v)
}
