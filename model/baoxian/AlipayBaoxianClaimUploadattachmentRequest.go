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
    _outBizNo   string
    // 业务来源
    _bizSource   string
    // 标准产品ID
    _spNo   string
    // 文件名,必须带后缀名。例如：test.png,test.doc,test.pdf
    _attachmentKey   string
    // 文件字节数组
    _attachmentByte   []*model.File
    // 是否base格式的字节数组
    _base64Bytes   bool
    // 保单外部业务单号
    _policyBizNo   string
    // 上传者用户标识
    _uploadUser   string
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
func (r *AlipayBaoxianClaimUploadattachmentRequest) SetOutBizNo(_outBizNo string) error {
    r._outBizNo = _outBizNo
    r.Set("out_biz_no", _outBizNo)
    return nil
}

// OutBizNo Getter
func (r AlipayBaoxianClaimUploadattachmentRequest) GetOutBizNo() string {
    return r._outBizNo
}
// BizSource Setter
// 业务来源
func (r *AlipayBaoxianClaimUploadattachmentRequest) SetBizSource(_bizSource string) error {
    r._bizSource = _bizSource
    r.Set("biz_source", _bizSource)
    return nil
}

// BizSource Getter
func (r AlipayBaoxianClaimUploadattachmentRequest) GetBizSource() string {
    return r._bizSource
}
// SpNo Setter
// 标准产品ID
func (r *AlipayBaoxianClaimUploadattachmentRequest) SetSpNo(_spNo string) error {
    r._spNo = _spNo
    r.Set("sp_no", _spNo)
    return nil
}

// SpNo Getter
func (r AlipayBaoxianClaimUploadattachmentRequest) GetSpNo() string {
    return r._spNo
}
// AttachmentKey Setter
// 文件名,必须带后缀名。例如：test.png,test.doc,test.pdf
func (r *AlipayBaoxianClaimUploadattachmentRequest) SetAttachmentKey(_attachmentKey string) error {
    r._attachmentKey = _attachmentKey
    r.Set("attachment_key", _attachmentKey)
    return nil
}

// AttachmentKey Getter
func (r AlipayBaoxianClaimUploadattachmentRequest) GetAttachmentKey() string {
    return r._attachmentKey
}
// AttachmentByte Setter
// 文件字节数组
func (r *AlipayBaoxianClaimUploadattachmentRequest) SetAttachmentByte(_attachmentByte []*model.File) error {
    r._attachmentByte = _attachmentByte
    r.Set("attachment_byte", _attachmentByte)
    return nil
}

// AttachmentByte Getter
func (r AlipayBaoxianClaimUploadattachmentRequest) GetAttachmentByte() []*model.File {
    return r._attachmentByte
}
// Base64Bytes Setter
// 是否base格式的字节数组
func (r *AlipayBaoxianClaimUploadattachmentRequest) SetBase64Bytes(_base64Bytes bool) error {
    r._base64Bytes = _base64Bytes
    r.Set("base64_bytes", _base64Bytes)
    return nil
}

// Base64Bytes Getter
func (r AlipayBaoxianClaimUploadattachmentRequest) GetBase64Bytes() bool {
    return r._base64Bytes
}
// PolicyBizNo Setter
// 保单外部业务单号
func (r *AlipayBaoxianClaimUploadattachmentRequest) SetPolicyBizNo(_policyBizNo string) error {
    r._policyBizNo = _policyBizNo
    r.Set("policy_biz_no", _policyBizNo)
    return nil
}

// PolicyBizNo Getter
func (r AlipayBaoxianClaimUploadattachmentRequest) GetPolicyBizNo() string {
    return r._policyBizNo
}
// UploadUser Setter
// 上传者用户标识
func (r *AlipayBaoxianClaimUploadattachmentRequest) SetUploadUser(_uploadUser string) error {
    r._uploadUser = _uploadUser
    r.Set("upload_user", _uploadUser)
    return nil
}

// UploadUser Getter
func (r AlipayBaoxianClaimUploadattachmentRequest) GetUploadUser() string {
    return r._uploadUser
}
