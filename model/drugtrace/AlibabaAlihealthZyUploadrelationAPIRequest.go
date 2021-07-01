package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
中药片关联关系上传 API请求
alibaba.alihealth.zy.uploadrelation

中药片关联关系上传
*/
type AlibabaAlihealthZyUploadrelationAPIRequest struct {
    model.Params
    // 关联关系文件信息
    _saveCodeRelation   *SaveCodeRelationType
    // affirmFlag
    _affirmFlag   string
    // fileContent
    _fileContent   string
    // 加密之后的上传的关联关系文件内容
    _fileContentString   string
    // 企业ID
    _refEntId   string
    // 客户端类型
    _clientType   string
}

// 初始化AlibabaAlihealthZyUploadrelationAPIRequest对象
func NewAlibabaAlihealthZyUploadrelationRequest() *AlibabaAlihealthZyUploadrelationAPIRequest{
    return &AlibabaAlihealthZyUploadrelationAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthZyUploadrelationAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.zy.uploadrelation"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthZyUploadrelationAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SaveCodeRelation Setter
// 关联关系文件信息
func (r *AlibabaAlihealthZyUploadrelationAPIRequest) SetSaveCodeRelation(_saveCodeRelation *SaveCodeRelationType) error {
    r._saveCodeRelation = _saveCodeRelation
    r.Set("save_code_relation", _saveCodeRelation)
    return nil
}

// SaveCodeRelation Getter
func (r AlibabaAlihealthZyUploadrelationAPIRequest) GetSaveCodeRelation() *SaveCodeRelationType {
    return r._saveCodeRelation
}
// AffirmFlag Setter
// affirmFlag
func (r *AlibabaAlihealthZyUploadrelationAPIRequest) SetAffirmFlag(_affirmFlag string) error {
    r._affirmFlag = _affirmFlag
    r.Set("affirm_flag", _affirmFlag)
    return nil
}

// AffirmFlag Getter
func (r AlibabaAlihealthZyUploadrelationAPIRequest) GetAffirmFlag() string {
    return r._affirmFlag
}
// FileContent Setter
// fileContent
func (r *AlibabaAlihealthZyUploadrelationAPIRequest) SetFileContent(_fileContent string) error {
    r._fileContent = _fileContent
    r.Set("file_content", _fileContent)
    return nil
}

// FileContent Getter
func (r AlibabaAlihealthZyUploadrelationAPIRequest) GetFileContent() string {
    return r._fileContent
}
// FileContentString Setter
// 加密之后的上传的关联关系文件内容
func (r *AlibabaAlihealthZyUploadrelationAPIRequest) SetFileContentString(_fileContentString string) error {
    r._fileContentString = _fileContentString
    r.Set("file_content_string", _fileContentString)
    return nil
}

// FileContentString Getter
func (r AlibabaAlihealthZyUploadrelationAPIRequest) GetFileContentString() string {
    return r._fileContentString
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthZyUploadrelationAPIRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthZyUploadrelationAPIRequest) GetRefEntId() string {
    return r._refEntId
}
// ClientType Setter
// 客户端类型
func (r *AlibabaAlihealthZyUploadrelationAPIRequest) SetClientType(_clientType string) error {
    r._clientType = _clientType
    r.Set("client_type", _clientType)
    return nil
}

// ClientType Getter
func (r AlibabaAlihealthZyUploadrelationAPIRequest) GetClientType() string {
    return r._clientType
}
