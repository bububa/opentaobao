package baoxian

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
资料上传接口 API请求
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
    attachmentByte   []*model.File
    // 是否base格式的字节数组
    base64Bytes   bool
    // 保单外部业务单号
    policyBizNo   string
    // 上传者用户标识
    uploadUser   string
}

// 初始化AlipayBaoxianClaimUploadattachmentRequest对象
func NewAlipayBaoxianClaimUploadattachmentRequest() *AlipayBaoxianClaimUploadattachmentRequest{
    return &AlipayBaoxianClaimUploadattachmentRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlipayBaoxianClaimUploadattachmentRequest) GetApiMethodName() string {
    return "alipay.baoxian.claim.uploadattachment"
}

// IRequest interface 方法, 获取API参数
func (r AlipayBaoxianClaimUploadattachmentRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutBizNo Setter
// 外部业务号，唯一
func (r *AlipayBaoxianClaimUploadattachmentRequest) SetOutBizNo(outBizNo string) error {
    r.outBizNo = outBizNo
    r.Set("out_biz_no", outBizNo)
    return nil
}

// OutBizNo Getter
func (r AlipayBaoxianClaimUploadattachmentRequest) GetOutBizNo() string {
    return r.outBizNo
}
// BizSource Setter
// 业务来源
func (r *AlipayBaoxianClaimUploadattachmentRequest) SetBizSource(bizSource string) error {
    r.bizSource = bizSource
    r.Set("biz_source", bizSource)
    return nil
}

// BizSource Getter
func (r AlipayBaoxianClaimUploadattachmentRequest) GetBizSource() string {
    return r.bizSource
}
// SpNo Setter
// 标准产品ID
func (r *AlipayBaoxianClaimUploadattachmentRequest) SetSpNo(spNo string) error {
    r.spNo = spNo
    r.Set("sp_no", spNo)
    return nil
}

// SpNo Getter
func (r AlipayBaoxianClaimUploadattachmentRequest) GetSpNo() string {
    return r.spNo
}
// AttachmentKey Setter
// 文件名,必须带后缀名。例如：test.png,test.doc,test.pdf
func (r *AlipayBaoxianClaimUploadattachmentRequest) SetAttachmentKey(attachmentKey string) error {
    r.attachmentKey = attachmentKey
    r.Set("attachment_key", attachmentKey)
    return nil
}

// AttachmentKey Getter
func (r AlipayBaoxianClaimUploadattachmentRequest) GetAttachmentKey() string {
    return r.attachmentKey
}
// AttachmentByte Setter
// 文件字节数组
func (r *AlipayBaoxianClaimUploadattachmentRequest) SetAttachmentByte(attachmentByte []*model.File) error {
    r.attachmentByte = attachmentByte
    r.Set("attachment_byte", attachmentByte)
    return nil
}

// AttachmentByte Getter
func (r AlipayBaoxianClaimUploadattachmentRequest) GetAttachmentByte() []*model.File {
    return r.attachmentByte
}
// Base64Bytes Setter
// 是否base格式的字节数组
func (r *AlipayBaoxianClaimUploadattachmentRequest) SetBase64Bytes(base64Bytes bool) error {
    r.base64Bytes = base64Bytes
    r.Set("base64_bytes", base64Bytes)
    return nil
}

// Base64Bytes Getter
func (r AlipayBaoxianClaimUploadattachmentRequest) GetBase64Bytes() bool {
    return r.base64Bytes
}
// PolicyBizNo Setter
// 保单外部业务单号
func (r *AlipayBaoxianClaimUploadattachmentRequest) SetPolicyBizNo(policyBizNo string) error {
    r.policyBizNo = policyBizNo
    r.Set("policy_biz_no", policyBizNo)
    return nil
}

// PolicyBizNo Getter
func (r AlipayBaoxianClaimUploadattachmentRequest) GetPolicyBizNo() string {
    return r.policyBizNo
}
// UploadUser Setter
// 上传者用户标识
func (r *AlipayBaoxianClaimUploadattachmentRequest) SetUploadUser(uploadUser string) error {
    r.uploadUser = uploadUser
    r.Set("upload_user", uploadUser)
    return nil
}

// UploadUser Getter
func (r AlipayBaoxianClaimUploadattachmentRequest) GetUploadUser() string {
    return r.uploadUser
}
