package baoxian

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlipaybaoxianclaimuploadattachmentAPIRequest 资料上传接口 API请求
// alipay.baoxian.claim.uploadattachment
//
// 给合作伙伴上传申请理赔材料
type AlipaybaoxianclaimuploadattachmentAPIRequest struct {
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

// NewAlipaybaoxianclaimuploadattachmentRequest 初始化AlipaybaoxianclaimuploadattachmentAPIRequest对象
func NewAlipaybaoxianclaimuploadattachmentRequest() *AlipaybaoxianclaimuploadattachmentAPIRequest {
	return &AlipaybaoxianclaimuploadattachmentAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlipaybaoxianclaimuploadattachmentAPIRequest) GetApiMethodName() string {
	return "alipay.baoxian.claim.uploadattachment"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlipaybaoxianclaimuploadattachmentAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlipaybaoxianclaimuploadattachmentAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutBizNo is OutBizNo Setter
// 外部业务号，唯一
func (r *AlipaybaoxianclaimuploadattachmentAPIRequest) SetOutBizNo(_outBizNo string) error {
	r._outBizNo = _outBizNo
	r.Set("out_biz_no", _outBizNo)
	return nil
}

// GetOutBizNo OutBizNo Getter
func (r AlipaybaoxianclaimuploadattachmentAPIRequest) GetOutBizNo() string {
	return r._outBizNo
}

// SetBizSource is BizSource Setter
// 业务来源
func (r *AlipaybaoxianclaimuploadattachmentAPIRequest) SetBizSource(_bizSource string) error {
	r._bizSource = _bizSource
	r.Set("biz_source", _bizSource)
	return nil
}

// GetBizSource BizSource Getter
func (r AlipaybaoxianclaimuploadattachmentAPIRequest) GetBizSource() string {
	return r._bizSource
}

// SetSpNo is SpNo Setter
// 标准产品ID
func (r *AlipaybaoxianclaimuploadattachmentAPIRequest) SetSpNo(_spNo string) error {
	r._spNo = _spNo
	r.Set("sp_no", _spNo)
	return nil
}

// GetSpNo SpNo Getter
func (r AlipaybaoxianclaimuploadattachmentAPIRequest) GetSpNo() string {
	return r._spNo
}

// SetAttachmentKey is AttachmentKey Setter
// 文件名,必须带后缀名。例如：test.png,test.doc,test.pdf
func (r *AlipaybaoxianclaimuploadattachmentAPIRequest) SetAttachmentKey(_attachmentKey string) error {
	r._attachmentKey = _attachmentKey
	r.Set("attachment_key", _attachmentKey)
	return nil
}

// GetAttachmentKey AttachmentKey Getter
func (r AlipaybaoxianclaimuploadattachmentAPIRequest) GetAttachmentKey() string {
	return r._attachmentKey
}

// SetPolicyBizNo is PolicyBizNo Setter
// 保单外部业务单号
func (r *AlipaybaoxianclaimuploadattachmentAPIRequest) SetPolicyBizNo(_policyBizNo string) error {
	r._policyBizNo = _policyBizNo
	r.Set("policy_biz_no", _policyBizNo)
	return nil
}

// GetPolicyBizNo PolicyBizNo Getter
func (r AlipaybaoxianclaimuploadattachmentAPIRequest) GetPolicyBizNo() string {
	return r._policyBizNo
}

// SetUploadUser is UploadUser Setter
// 上传者用户标识
func (r *AlipaybaoxianclaimuploadattachmentAPIRequest) SetUploadUser(_uploadUser string) error {
	r._uploadUser = _uploadUser
	r.Set("upload_user", _uploadUser)
	return nil
}

// GetUploadUser UploadUser Getter
func (r AlipaybaoxianclaimuploadattachmentAPIRequest) GetUploadUser() string {
	return r._uploadUser
}

// SetAttachmentByte is AttachmentByte Setter
// 文件字节数组
func (r *AlipaybaoxianclaimuploadattachmentAPIRequest) SetAttachmentByte(_attachmentByte *model.File) error {
	r._attachmentByte = _attachmentByte
	r.Set("attachment_byte", _attachmentByte)
	return nil
}

// GetAttachmentByte AttachmentByte Getter
func (r AlipaybaoxianclaimuploadattachmentAPIRequest) GetAttachmentByte() *model.File {
	return r._attachmentByte
}

// SetBase64Bytes is Base64Bytes Setter
// 是否base格式的字节数组
func (r *AlipaybaoxianclaimuploadattachmentAPIRequest) SetBase64Bytes(_base64Bytes bool) error {
	r._base64Bytes = _base64Bytes
	r.Set("base64_bytes", _base64Bytes)
	return nil
}

// GetBase64Bytes Base64Bytes Getter
func (r AlipaybaoxianclaimuploadattachmentAPIRequest) GetBase64Bytes() bool {
	return r._base64Bytes
}
