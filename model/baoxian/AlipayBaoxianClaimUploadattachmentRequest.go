package baoxian

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
资料上传接口 APIRequest
alipay.baoxian.claim.uploadattachment

给合作伙伴上传申请理赔材料
*/
type AlipayBaoxianClaimUploadattachmentRequest struct {
    model.Params

    // 外部业务号，唯一
    outBizNo   string 

    // 业务来源
    bizSource   string 

    // 标准产品ID
    spNo   string 

    // 文件名,必须带后缀名。例如：test.png,test.doc,test.pdf
    attachmentKey   string 

    // 文件字节数组
    attachmentByte   []byte 

    // 是否base格式的字节数组
    base64Bytes   bool 

    // 保单外部业务单号
    policyBizNo   string 

    // 上传者用户标识
    uploadUser   string 

}

func NewAlipayBaoxianClaimUploadattachmentRequest() *AlipayBaoxianClaimUploadattachmentRequest{
    return &AlipayBaoxianClaimUploadattachmentRequest{
        Params: model.NewParams(),
    }
}

func (r AlipayBaoxianClaimUploadattachmentRequest) GetApiMethodName() string {
    return "alipay.baoxian.claim.uploadattachment"
}

func (r AlipayBaoxianClaimUploadattachmentRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlipayBaoxianClaimUploadattachmentRequest) SetOutBizNo(outBizNo string) error {
    r.outBizNo = outBizNo
    r.Set("out_biz_no", outBizNo)
    return nil
}

func (r AlipayBaoxianClaimUploadattachmentRequest) GetOutBizNo() string {
    return r.outBizNo
}

func (r *AlipayBaoxianClaimUploadattachmentRequest) SetBizSource(bizSource string) error {
    r.bizSource = bizSource
    r.Set("biz_source", bizSource)
    return nil
}

func (r AlipayBaoxianClaimUploadattachmentRequest) GetBizSource() string {
    return r.bizSource
}

func (r *AlipayBaoxianClaimUploadattachmentRequest) SetSpNo(spNo string) error {
    r.spNo = spNo
    r.Set("sp_no", spNo)
    return nil
}

func (r AlipayBaoxianClaimUploadattachmentRequest) GetSpNo() string {
    return r.spNo
}

func (r *AlipayBaoxianClaimUploadattachmentRequest) SetAttachmentKey(attachmentKey string) error {
    r.attachmentKey = attachmentKey
    r.Set("attachment_key", attachmentKey)
    return nil
}

func (r AlipayBaoxianClaimUploadattachmentRequest) GetAttachmentKey() string {
    return r.attachmentKey
}

func (r *AlipayBaoxianClaimUploadattachmentRequest) SetAttachmentByte(attachmentByte []byte) error {
    r.attachmentByte = attachmentByte
    r.Set("attachment_byte", attachmentByte)
    return nil
}

func (r AlipayBaoxianClaimUploadattachmentRequest) GetAttachmentByte() []byte {
    return r.attachmentByte
}

func (r *AlipayBaoxianClaimUploadattachmentRequest) SetBase64Bytes(base64Bytes bool) error {
    r.base64Bytes = base64Bytes
    r.Set("base64_bytes", base64Bytes)
    return nil
}

func (r AlipayBaoxianClaimUploadattachmentRequest) GetBase64Bytes() bool {
    return r.base64Bytes
}

func (r *AlipayBaoxianClaimUploadattachmentRequest) SetPolicyBizNo(policyBizNo string) error {
    r.policyBizNo = policyBizNo
    r.Set("policy_biz_no", policyBizNo)
    return nil
}

func (r AlipayBaoxianClaimUploadattachmentRequest) GetPolicyBizNo() string {
    return r.policyBizNo
}

func (r *AlipayBaoxianClaimUploadattachmentRequest) SetUploadUser(uploadUser string) error {
    r.uploadUser = uploadUser
    r.Set("upload_user", uploadUser)
    return nil
}

func (r AlipayBaoxianClaimUploadattachmentRequest) GetUploadUser() string {
    return r.uploadUser
}

