package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytScqyUploadrelationAPIRequest 关联关系上传 API请求
// alibaba.alihealth.drug.kyt.scqy.uploadrelation
//
// 关联关系上传
type AlibabaAlihealthDrugKytScqyUploadrelationAPIRequest struct {
	model.Params
	// affirmFlag
	_affirmFlag string
	// fileContent(可不添)
	_fileContent string
	// 加密之后的文件内容字符串
	_fileContentString string
	// 上传文件的企业ID
	_refEntId string
	// 客户端类型
	_clientType string
	// 关联关系文件信息
	_saveCodeRelation *SaveCodeRelationType
}

// NewAlibabaAlihealthDrugKytScqyUploadrelationRequest 初始化AlibabaAlihealthDrugKytScqyUploadrelationAPIRequest对象
func NewAlibabaAlihealthDrugKytScqyUploadrelationRequest() *AlibabaAlihealthDrugKytScqyUploadrelationAPIRequest {
	return &AlibabaAlihealthDrugKytScqyUploadrelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytScqyUploadrelationAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.scqy.uploadrelation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytScqyUploadrelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytScqyUploadrelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAffirmFlag is AffirmFlag Setter
// affirmFlag
func (r *AlibabaAlihealthDrugKytScqyUploadrelationAPIRequest) SetAffirmFlag(_affirmFlag string) error {
	r._affirmFlag = _affirmFlag
	r.Set("affirm_flag", _affirmFlag)
	return nil
}

// GetAffirmFlag AffirmFlag Getter
func (r AlibabaAlihealthDrugKytScqyUploadrelationAPIRequest) GetAffirmFlag() string {
	return r._affirmFlag
}

// SetFileContent is FileContent Setter
// fileContent(可不添)
func (r *AlibabaAlihealthDrugKytScqyUploadrelationAPIRequest) SetFileContent(_fileContent string) error {
	r._fileContent = _fileContent
	r.Set("file_content", _fileContent)
	return nil
}

// GetFileContent FileContent Getter
func (r AlibabaAlihealthDrugKytScqyUploadrelationAPIRequest) GetFileContent() string {
	return r._fileContent
}

// SetFileContentString is FileContentString Setter
// 加密之后的文件内容字符串
func (r *AlibabaAlihealthDrugKytScqyUploadrelationAPIRequest) SetFileContentString(_fileContentString string) error {
	r._fileContentString = _fileContentString
	r.Set("file_content_string", _fileContentString)
	return nil
}

// GetFileContentString FileContentString Getter
func (r AlibabaAlihealthDrugKytScqyUploadrelationAPIRequest) GetFileContentString() string {
	return r._fileContentString
}

// SetRefEntId is RefEntId Setter
// 上传文件的企业ID
func (r *AlibabaAlihealthDrugKytScqyUploadrelationAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytScqyUploadrelationAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetClientType is ClientType Setter
// 客户端类型
func (r *AlibabaAlihealthDrugKytScqyUploadrelationAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// GetClientType ClientType Getter
func (r AlibabaAlihealthDrugKytScqyUploadrelationAPIRequest) GetClientType() string {
	return r._clientType
}

// SetSaveCodeRelation is SaveCodeRelation Setter
// 关联关系文件信息
func (r *AlibabaAlihealthDrugKytScqyUploadrelationAPIRequest) SetSaveCodeRelation(_saveCodeRelation *SaveCodeRelationType) error {
	r._saveCodeRelation = _saveCodeRelation
	r.Set("save_code_relation", _saveCodeRelation)
	return nil
}

// GetSaveCodeRelation SaveCodeRelation Getter
func (r AlibabaAlihealthDrugKytScqyUploadrelationAPIRequest) GetSaveCodeRelation() *SaveCodeRelationType {
	return r._saveCodeRelation
}
