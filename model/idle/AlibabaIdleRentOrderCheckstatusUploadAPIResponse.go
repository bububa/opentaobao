package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRentOrderCheckstatusUploadAPIResponse 上传验收结果 API返回值
// alibaba.idle.rent.order.checkstatus.upload
//
// 上传验收结果
type AlibabaIdleRentOrderCheckstatusUploadAPIResponse struct {
	model.CommonResponse
	AlibabaIdleRentOrderCheckstatusUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleRentOrderCheckstatusUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleRentOrderCheckstatusUploadAPIResponseModel).Reset()
}

// AlibabaIdleRentOrderCheckstatusUploadAPIResponseModel is 上传验收结果 成功返回结果
type AlibabaIdleRentOrderCheckstatusUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_rent_order_checkstatus_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabaIdleRentOrderCheckstatusUploadTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleRentOrderCheckstatusUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleRentOrderCheckstatusUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRentOrderCheckstatusUploadAPIResponse)
	},
}

// GetAlibabaIdleRentOrderCheckstatusUploadAPIResponse 从 sync.Pool 获取 AlibabaIdleRentOrderCheckstatusUploadAPIResponse
func GetAlibabaIdleRentOrderCheckstatusUploadAPIResponse() *AlibabaIdleRentOrderCheckstatusUploadAPIResponse {
	return poolAlibabaIdleRentOrderCheckstatusUploadAPIResponse.Get().(*AlibabaIdleRentOrderCheckstatusUploadAPIResponse)
}

// ReleaseAlibabaIdleRentOrderCheckstatusUploadAPIResponse 将 AlibabaIdleRentOrderCheckstatusUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleRentOrderCheckstatusUploadAPIResponse(v *AlibabaIdleRentOrderCheckstatusUploadAPIResponse) {
	v.Reset()
	poolAlibabaIdleRentOrderCheckstatusUploadAPIResponse.Put(v)
}
