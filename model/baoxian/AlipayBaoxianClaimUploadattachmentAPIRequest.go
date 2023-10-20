package baoxian

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlipayBaoxianClaimUploadattachmentAPIRequest 资料上传接口 API请求
// alipay.baoxian.claim.uploadattachment
//
// 给合作伙伴上传申请理赔材料
type AlipayBaoxianClaimUploadattachmentAPIRequest struct {
	model.Params
	// 外部业务号，唯一
	_outBizNo string
	// 业务来源
	_bizSource string
	// 标准产品ID
	_spNo string
	// 文件名,必须带后缀名。例如：test.png,test.doc,test.pdf
	_attachmentKey string
	// 保单外部业务单号
	_policyBizNo string
	// 上传者用户标识
	_uploadUser string
	// 文件字节数组
	_attachmentByte *model.File
	// 是否base格式的字节数组
	_base64Bytes bool
}

// NewAlipayBaoxianClaimUploadattachmentRequest 初始化AlipayBaoxianClaimUploadattachmentAPIRequest对象
func NewAlipayBaoxianClaimUploadattachmentRequest() *AlipayBaoxianClaimUploadattachmentAPIRequest {
	return &AlipayBaoxianClaimUploadattachmentAPIRequest{
		Params: model.NewParams(8),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlipayBaoxianClaimUploadattachmentAPIRequest) Reset() {
	r._outBizNo = ""
	r._bizSource = ""
	r._spNo = ""
	r._attachmentKey = ""
	r._policyBizNo = ""
	r._uploadUser = ""
	r._attachmentByte = nil
	r._base64Bytes = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlipayBaoxianClaimUploadattachmentAPIRequest) GetApiMethodName() string {
	return "alipay.baoxian.claim.uploadattachment"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlipayBaoxianClaimUploadattachmentAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlipayBaoxianClaimUploadattachmentAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutBizNo is OutBizNo Setter
// 外部业务号，唯一
func (r *AlipayBaoxianClaimUploadattachmentAPIRequest) SetOutBizNo(_outBizNo string) error {
	r._outBizNo = _outBizNo
	r.Set("out_biz_no", _outBizNo)
	return nil
}

// GetOutBizNo OutBizNo Getter
func (r AlipayBaoxianClaimUploadattachmentAPIRequest) GetOutBizNo() string {
	return r._outBizNo
}

// SetBizSource is BizSource Setter
// 业务来源
func (r *AlipayBaoxianClaimUploadattachmentAPIRequest) SetBizSource(_bizSource string) error {
	r._bizSource = _bizSource
	r.Set("biz_source", _bizSource)
	return nil
}

// GetBizSource BizSource Getter
func (r AlipayBaoxianClaimUploadattachmentAPIRequest) GetBizSource() string {
	return r._bizSource
}

// SetSpNo is SpNo Setter
// 标准产品ID
func (r *AlipayBaoxianClaimUploadattachmentAPIRequest) SetSpNo(_spNo string) error {
	r._spNo = _spNo
	r.Set("sp_no", _spNo)
	return nil
}

// GetSpNo SpNo Getter
func (r AlipayBaoxianClaimUploadattachmentAPIRequest) GetSpNo() string {
	return r._spNo
}

// SetAttachmentKey is AttachmentKey Setter
// 文件名,必须带后缀名。例如：test.png,test.doc,test.pdf
func (r *AlipayBaoxianClaimUploadattachmentAPIRequest) SetAttachmentKey(_attachmentKey string) error {
	r._attachmentKey = _attachmentKey
	r.Set("attachment_key", _attachmentKey)
	return nil
}

// GetAttachmentKey AttachmentKey Getter
func (r AlipayBaoxianClaimUploadattachmentAPIRequest) GetAttachmentKey() string {
	return r._attachmentKey
}

// SetPolicyBizNo is PolicyBizNo Setter
// 保单外部业务单号
func (r *AlipayBaoxianClaimUploadattachmentAPIRequest) SetPolicyBizNo(_policyBizNo string) error {
	r._policyBizNo = _policyBizNo
	r.Set("policy_biz_no", _policyBizNo)
	return nil
}

// GetPolicyBizNo PolicyBizNo Getter
func (r AlipayBaoxianClaimUploadattachmentAPIRequest) GetPolicyBizNo() string {
	return r._policyBizNo
}

// SetUploadUser is UploadUser Setter
// 上传者用户标识
func (r *AlipayBaoxianClaimUploadattachmentAPIRequest) SetUploadUser(_uploadUser string) error {
	r._uploadUser = _uploadUser
	r.Set("upload_user", _uploadUser)
	return nil
}

// GetUploadUser UploadUser Getter
func (r AlipayBaoxianClaimUploadattachmentAPIRequest) GetUploadUser() string {
	return r._uploadUser
}

// SetAttachmentByte is AttachmentByte Setter
// 文件字节数组
func (r *AlipayBaoxianClaimUploadattachmentAPIRequest) SetAttachmentByte(_attachmentByte *model.File) error {
	r._attachmentByte = _attachmentByte
	r.Set("attachment_byte", _attachmentByte)
	return nil
}

// GetAttachmentByte AttachmentByte Getter
func (r AlipayBaoxianClaimUploadattachmentAPIRequest) GetAttachmentByte() *model.File {
	return r._attachmentByte
}

// SetBase64Bytes is Base64Bytes Setter
// 是否base格式的字节数组
func (r *AlipayBaoxianClaimUploadattachmentAPIRequest) SetBase64Bytes(_base64Bytes bool) error {
	r._base64Bytes = _base64Bytes
	r.Set("base64_bytes", _base64Bytes)
	return nil
}

// GetBase64Bytes Base64Bytes Getter
func (r AlipayBaoxianClaimUploadattachmentAPIRequest) GetBase64Bytes() bool {
	return r._base64Bytes
}

var poolAlipayBaoxianClaimUploadattachmentAPIRequest = sync.Pool{
	New: func() any {
		return NewAlipayBaoxianClaimUploadattachmentRequest()
	},
}

// GetAlipayBaoxianClaimUploadattachmentRequest 从 sync.Pool 获取 AlipayBaoxianClaimUploadattachmentAPIRequest
func GetAlipayBaoxianClaimUploadattachmentAPIRequest() *AlipayBaoxianClaimUploadattachmentAPIRequest {
	return poolAlipayBaoxianClaimUploadattachmentAPIRequest.Get().(*AlipayBaoxianClaimUploadattachmentAPIRequest)
}

// ReleaseAlipayBaoxianClaimUploadattachmentAPIRequest 将 AlipayBaoxianClaimUploadattachmentAPIRequest 放入 sync.Pool
func ReleaseAlipayBaoxianClaimUploadattachmentAPIRequest(v *AlipayBaoxianClaimUploadattachmentAPIRequest) {
	v.Reset()
	poolAlipayBaoxianClaimUploadattachmentAPIRequest.Put(v)
}
