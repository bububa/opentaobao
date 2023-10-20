package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinAxbVendorPushCallReleaseAPIResponse 供应商推送通话结束事件 API返回值
// alibaba.aliqin.axb.vendor.push.call.release
//
// 通话结束挂断的时候，供应商推送通话结束事件给阿里侧
type AlibabaAliqinAxbVendorPushCallReleaseAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinAxbVendorPushCallReleaseAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinAxbVendorPushCallReleaseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinAxbVendorPushCallReleaseAPIResponseModel).Reset()
}

// AlibabaAliqinAxbVendorPushCallReleaseAPIResponseModel is 供应商推送通话结束事件 成功返回结果
type AlibabaAliqinAxbVendorPushCallReleaseAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_axb_vendor_push_call_release_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAliqinAxbVendorPushCallReleaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinAxbVendorPushCallReleaseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliqinAxbVendorPushCallReleaseAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinAxbVendorPushCallReleaseAPIResponse)
	},
}

// GetAlibabaAliqinAxbVendorPushCallReleaseAPIResponse 从 sync.Pool 获取 AlibabaAliqinAxbVendorPushCallReleaseAPIResponse
func GetAlibabaAliqinAxbVendorPushCallReleaseAPIResponse() *AlibabaAliqinAxbVendorPushCallReleaseAPIResponse {
	return poolAlibabaAliqinAxbVendorPushCallReleaseAPIResponse.Get().(*AlibabaAliqinAxbVendorPushCallReleaseAPIResponse)
}

// ReleaseAlibabaAliqinAxbVendorPushCallReleaseAPIResponse 将 AlibabaAliqinAxbVendorPushCallReleaseAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinAxbVendorPushCallReleaseAPIResponse(v *AlibabaAliqinAxbVendorPushCallReleaseAPIResponse) {
	v.Reset()
	poolAlibabaAliqinAxbVendorPushCallReleaseAPIResponse.Put(v)
}
