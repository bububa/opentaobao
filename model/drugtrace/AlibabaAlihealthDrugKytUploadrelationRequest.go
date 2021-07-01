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
type AlibabaAlihealthDrugKytUploadrelationAPIRequest struct {
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

// 初始化AlibabaAlihealthDrugKytUploadrelationAPIRequest对象
func NewAlibabaAlihealthDrugKytUploadrelationRequest() *AlibabaAlihealthDrugKytUploadrelationAPIRequest{
    return &AlibabaAlihealthDrugKytUploadrelationAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytUploadrelationAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.uploadrelation"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytUploadrelationAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SaveCodeRelation Setter
// 关联关系文件信息
func (r *AlibabaAlihealthDrugKytUploadrelationAPIRequest) SetSaveCodeRelation(_saveCodeRelation *SaveCodeRelationType) error {
    r._saveCodeRelation = _saveCodeRelation
    r.Set("save_code_relation", _saveCodeRelation)
    return nil
}

// SaveCodeRelation Getter
func (r AlibabaAlihealthDrugKytUploadrelationAPIRequest) GetSaveCodeRelation() *SaveCodeRelationType {
    return r._saveCodeRelation
}
// AffirmFlag Setter
// affirmFlag
func (r *AlibabaAlihealthDrugKytUploadrelationAPIRequest) SetAffirmFlag(_affirmFlag string) error {
    r._affirmFlag = _affirmFlag
    r.Set("affirm_flag", _affirmFlag)
    return nil
}

// AffirmFlag Getter
func (r AlibabaAlihealthDrugKytUploadrelationAPIRequest) GetAffirmFlag() string {
    return r._affirmFlag
}
// FileContent Setter
// fileContent
func (r *AlibabaAlihealthDrugKytUploadrelationAPIRequest) SetFileContent(_fileContent string) error {
    r._fileContent = _fileContent
    r.Set("file_content", _fileContent)
    return nil
}

// FileContent Getter
func (r AlibabaAlihealthDrugKytUploadrelationAPIRequest) GetFileContent() string {
    return r._fileContent
}
// FileContentString Setter
// 加密之后的文件内容字符串
func (r *AlibabaAlihealthDrugKytUploadrelationAPIRequest) SetFileContentString(_fileContentString string) error {
    r._fileContentString = _fileContentString
    r.Set("file_content_string", _fileContentString)
    return nil
}

// FileContentString Getter
func (r AlibabaAlihealthDrugKytUploadrelationAPIRequest) GetFileContentString() string {
    return r._fileContentString
}
// RefEntId Setter
// 上传文件的企业ID
func (r *AlibabaAlihealthDrugKytUploadrelationAPIRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytUploadrelationAPIRequest) GetRefEntId() string {
    return r._refEntId
}
// ClientType Setter
// 客户端类型
func (r *AlibabaAlihealthDrugKytUploadrelationAPIRequest) SetClientType(_clientType string) error {
    r._clientType = _clientType
    r.Set("client_type", _clientType)
    return nil
}

// ClientType Getter
func (r AlibabaAlihealthDrugKytUploadrelationAPIRequest) GetClientType() string {
    return r._clientType
}
