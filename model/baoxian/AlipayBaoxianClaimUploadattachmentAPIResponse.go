package baoxian

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlipaybaoxianclaimuploadattachmentAPIResponse 资料上传接口 API返回值
// alipay.baoxian.claim.uploadattachment
//
// 给合作伙伴上传申请理赔材料
type AlipaybaoxianclaimuploadattachmentAPIResponse struct {
	model.CommonResponse
	AlipaybaoxianclaimuploadattachmentAPIResponseModel
}

// AlipaybaoxianclaimuploadattachmentAPIResponseModel is 资料上传接口 成功返回结果
type AlipaybaoxianclaimuploadattachmentAPIResponseModel struct {
	XMLName xml.Name `xml:"alipay_baoxian_claim_uploadattachment_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	UploadResult *UploadResult `json:"upload_result,omitempty" xml:"upload_result,omitempty"`
}
