package baoxian

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlipayBaoxianClaimUploadattachmentAPIResponse 资料上传接口 API返回值
// alipay.baoxian.claim.uploadattachment
//
// 给合作伙伴上传申请理赔材料
type AlipayBaoxianClaimUploadattachmentAPIResponse struct {
	model.CommonResponse
	AlipayBaoxianClaimUploadattachmentAPIResponseModel
}

// Reset 清空结构体
func (m *AlipayBaoxianClaimUploadattachmentAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlipayBaoxianClaimUploadattachmentAPIResponseModel).Reset()
}

// AlipayBaoxianClaimUploadattachmentAPIResponseModel is 资料上传接口 成功返回结果
type AlipayBaoxianClaimUploadattachmentAPIResponseModel struct {
	XMLName xml.Name `xml:"alipay_baoxian_claim_uploadattachment_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	UploadResult *UploadResult `json:"upload_result,omitempty" xml:"upload_result,omitempty"`
}

// Reset 清空结构体
func (m *AlipayBaoxianClaimUploadattachmentAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UploadResult = nil
}

var poolAlipayBaoxianClaimUploadattachmentAPIResponse = sync.Pool{
	New: func() any {
		return new(AlipayBaoxianClaimUploadattachmentAPIResponse)
	},
}

// GetAlipayBaoxianClaimUploadattachmentAPIResponse 从 sync.Pool 获取 AlipayBaoxianClaimUploadattachmentAPIResponse
func GetAlipayBaoxianClaimUploadattachmentAPIResponse() *AlipayBaoxianClaimUploadattachmentAPIResponse {
	return poolAlipayBaoxianClaimUploadattachmentAPIResponse.Get().(*AlipayBaoxianClaimUploadattachmentAPIResponse)
}

// ReleaseAlipayBaoxianClaimUploadattachmentAPIResponse 将 AlipayBaoxianClaimUploadattachmentAPIResponse 保存到 sync.Pool
func ReleaseAlipayBaoxianClaimUploadattachmentAPIResponse(v *AlipayBaoxianClaimUploadattachmentAPIResponse) {
	v.Reset()
	poolAlipayBaoxianClaimUploadattachmentAPIResponse.Put(v)
}
