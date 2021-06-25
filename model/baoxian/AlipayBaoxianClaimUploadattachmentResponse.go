package baoxian

import (
    "github.com/bububa/opentaobao/model"
)

/* 
资料上传接口 APIResponse
alipay.baoxian.claim.uploadattachment

给合作伙伴上传申请理赔材料
*/
type AlipayBaoxianClaimUploadattachmentAPIResponse struct {
    model.CommonResponse
    Response *AlipayBaoxianClaimUploadattachmentResponse `json:"alipay_baoxian_claim_uploadattachment_response,omitempty"`
}

type AlipayBaoxianClaimUploadattachmentResponse struct {

    // result
    UploadResult   *UploadResult `json:"upload_result,omitempty"`

}
