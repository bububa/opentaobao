package baoxian

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
资料上传接口 APIResponse
alipay.baoxian.claim.uploadattachment

给合作伙伴上传申请理赔材料
*/
type AlipayBaoxianClaimUploadattachmentAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alipay_baoxian_claim_uploadattachment_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    UploadResult   *UploadResult `json:"upload_result,omitempty" xml:"