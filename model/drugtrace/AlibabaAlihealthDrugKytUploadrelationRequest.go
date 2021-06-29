package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关联关系上传 API请求
alibaba.alihealth.drug.kyt.uploadrelation

关联关系上传
*/
type AlibabaAlihealthDrugKytUploadrelationRequest struct {
    model.Params
    // 关联关系文件信息
    _saveCodeRelation   *SaveCodeRelationType
    // affirmFlag
    _affirmFlag   string
    // fileContent
    _fileContent   string
    // 加密之后的文件内容字符串
    _fileContentString   string
    // 上传文件的企业ID
    _refEntId   string
    // 客户端类型
    _clientType   string
}

// 初始化AlibabaAlihealthDrugKytUploadrelationRequest对象
func NewAlibabaAlihealthDrugKytUploadrelationRequest() *AlibabaAlihealthDrugKytUploadrelationRequest{
    return &AlibabaAlihealthDrugKytUploadrelationRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytUploadrelationRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.uploadrelation"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytUploadrelationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SaveCodeRelation Setter
// 关联关系文件信息
func (r *AlibabaAlihealthDrugKytUploadrelationRequest) SetSaveCodeRelation(_saveCodeRelation *SaveCodeRelationType) error {
    r._saveCodeRelation = _saveCodeRelation
    r.Set("save_code_relation", _saveCodeRelation)
    return nil
}

// SaveCodeRelation Getter
func (r AlibabaAlihealthDrugKytUploadrelationRequest) GetSaveCodeRelation() *SaveCodeRelationType {
    return r._saveCodeRelation
}
// AffirmFlag Setter
// affirmFlag
func (r *AlibabaAlihealthDrugKytUploadrelationRequest) SetAffirmFlag(_affirmFlag string) error {
    r._affirmFlag = _affirmFlag
    r.Set("affirm_flag", _affirmFlag)
    return nil
}

// AffirmFlag Getter
func (r AlibabaAlihealthDrugKytUploadrelationRequest) GetAffirmFlag() string {
    return r._affirmFlag
}
// FileContent Setter
// fileContent
func (r *AlibabaAlihealthDrugKytUploadrelationRequest) SetFileContent(_fileContent string) error {
    r._fileContent = _fileContent
    r.Set("file_content", _fileContent)
    return nil
}

// FileContent Getter
func (r AlibabaAlihealthDrugKytUploadrelationRequest) GetFileContent() string {
    return r._fileContent
}
// FileContentString Setter
// 加密之后的文件内容字符串
func (r *AlibabaAlihealthDrugKytUploadrelationRequest) SetFileContentString(_fileContentString string) error {
    r._fileContentString = _fileContentString
    r.Set("file_content_string", _fileContentString)
    return nil
}

// FileContentString Getter
func (r AlibabaAlihealthDrugKytUploadrelationRequest) GetFileContentString() string {
    return r._fileContentString
}
// RefEntId Setter
// 上传文件的企业ID
func (r *AlibabaAlihealthDrugKytUploadrelationRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytUploadrelationRequest) GetRefEntId() string {
    return r._refEntId
}
// ClientType Setter
// 客户端类型
func (r *AlibabaAlihealthDrugKytUploadrelationRequest) SetClientType(_clientType string) error {
    r._clientType = _clientType
    r.Set("client_type", _clientType)
    return nil
}

// ClientType Getter
func (r AlibabaAlihealthDrugKytUploadrelationRequest) GetClientType() string {
    return r._clientType
}
