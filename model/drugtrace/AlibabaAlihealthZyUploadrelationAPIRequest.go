package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthzyuploadrelationAPIRequest 中药片关联关系上传 API请求
// alibaba.alihealth.zy.uploadrelation
//
// 中药片关联关系上传
type AlibabaalihealthzyuploadrelationAPIRequest struct {
	model.Params
	// affirmFlag
	_affirmFlag string
	// fileContent
	_fileContent string
	// 加密之后的上传的关联关系文件内容
	_fileContentString string
	// 企业ID
	_refEntId string
	// 客户端类型
	_clientType string
	// 关联关系文件信息
	_saveCodeRelation *SaveCodeRelationType
}

// NewAlibabaalihealthzyuploadrelationRequest 初始化AlibabaalihealthzyuploadrelationAPIRequest对象
func NewAlibabaalihealthzyuploadrelationRequest() *AlibabaalihealthzyuploadrelationAPIRequest {
	return &AlibabaalihealthzyuploadrelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthzyuploadrelationAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.zy.uploadrelation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthzyuploadrelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthzyuploadrelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAffirmFlag is AffirmFlag Setter
// affirmFlag
func (r *AlibabaalihealthzyuploadrelationAPIRequest) SetAffirmFlag(_affirmFlag string) error {
	r._affirmFlag = _affirmFlag
	r.Set("affirm_flag", _affirmFlag)
	return nil
}

// GetAffirmFlag AffirmFlag Getter
func (r AlibabaalihealthzyuploadrelationAPIRequest) GetAffirmFlag() string {
	return r._affirmFlag
}

// SetFileContent is FileContent Setter
// fileContent
func (r *AlibabaalihealthzyuploadrelationAPIRequest) SetFileContent(_fileContent string) error {
	r._fileContent = _fileContent
	r.Set("file_content", _fileContent)
	return nil
}

// GetFileContent FileContent Getter
func (r AlibabaalihealthzyuploadrelationAPIRequest) GetFileContent() string {
	return r._fileContent
}

// SetFileContentString is FileContentString Setter
// 加密之后的上传的关联关系文件内容
func (r *AlibabaalihealthzyuploadrelationAPIRequest) SetFileContentString(_fileContentString string) error {
	r._fileContentString = _fileContentString
	r.Set("file_content_string", _fileContentString)
	return nil
}

// GetFileContentString FileContentString Getter
func (r AlibabaalihealthzyuploadrelationAPIRequest) GetFileContentString() string {
	return r._fileContentString
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaalihealthzyuploadrelationAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthzyuploadrelationAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetClientType is ClientType Setter
// 客户端类型
func (r *AlibabaalihealthzyuploadrelationAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// GetClientType ClientType Getter
func (r AlibabaalihealthzyuploadrelationAPIRequest) GetClientType() string {
	return r._clientType
}

// SetSaveCodeRelation is SaveCodeRelation Setter
// 关联关系文件信息
func (r *AlibabaalihealthzyuploadrelationAPIRequest) SetSaveCodeRelation(_saveCodeRelation *SaveCodeRelationType) error {
	r._saveCodeRelation = _saveCodeRelation
	r.Set("save_code_relation", _saveCodeRelation)
	return nil
}

// GetSaveCodeRelation SaveCodeRelation Getter
func (r AlibabaalihealthzyuploadrelationAPIRequest) GetSaveCodeRelation() *SaveCodeRelationType {
	return r._saveCodeRelation
}
