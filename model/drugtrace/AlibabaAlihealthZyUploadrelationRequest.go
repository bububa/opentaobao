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
type AlibabaAlihealthZyUploadrelationRequest struct {
    model.Params
    // 关联关系文件信息
    saveCodeRelation   *SaveCodeRelationType
    // affirmFlag
    affirmFlag   string
    // fileContent
    fileContent   string
    // 加密之后的上传的关联关系文件内容
    fileContentString   string
    // 企业ID
    refEntId   string
    // 客户端类型
    clientType   string
}

// 初始化AlibabaAlihealthZyUploadrelationRequest对象
func NewAlibabaAlihealthZyUploadrelationRequest() *AlibabaAlihealthZyUploadrelationRequest{
    return &AlibabaAlihealthZyUploadrelationRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthZyUploadrelationRequest) GetApiMethodName() string {
    return "alibaba.alihealth.zy.uploadrelation"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthZyUploadrelationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SaveCodeRelation Setter
// 关联关系文件信息
func (r *AlibabaAlihealthZyUploadrelationRequest) SetSaveCodeRelation(saveCodeRelation *SaveCodeRelationType) error {
    r.saveCodeRelation = saveCodeRelation
    r.Set("save_code_relation", saveCodeRelation)
    return nil
}

// SaveCodeRelation Getter
func (r AlibabaAlihealthZyUploadrelationRequest) GetSaveCodeRelation() *SaveCodeRelationType {
    return r.saveCodeRelation
}
// AffirmFlag Setter
// affirmFlag
func (r *AlibabaAlihealthZyUploadrelationRequest) SetAffirmFlag(affirmFlag string) error {
    r.affirmFlag = affirmFlag
    r.Set("affirm_flag", affirmFlag)
    return nil
}

// AffirmFlag Getter
func (r AlibabaAlihealthZyUploadrelationRequest) GetAffirmFlag() string {
    return r.affirmFlag
}
// FileContent Setter
// fileContent
func (r *AlibabaAlihealthZyUploadrelationRequest) SetFileContent(fileContent string) error {
    r.fileContent = fileContent
    r.Set("file_content", fileContent)
    return nil
}

// FileContent Getter
func (r AlibabaAlihealthZyUploadrelationRequest) GetFileContent() string {
    return r.fileContent
}
// FileContentString Setter
// 加密之后的上传的关联关系文件内容
func (r *AlibabaAlihealthZyUploadrelationRequest) SetFileContentString(fileContentString string) error {
    r.fileContentString = fileContentString
    r.Set("file_content_string", fileContentString)
    return nil
}

// FileContentString Getter
func (r AlibabaAlihealthZyUploadrelationRequest) GetFileContentString() string {
    return r.fileContentString
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthZyUploadrelationRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthZyUploadrelationRequest) GetRefEntId() string {
    return r.refEntId
}
// ClientType Setter
// 客户端类型
func (r *AlibabaAlihealthZyUploadrelationRequest) SetClientType(clientType string) error {
    r.clientType = clientType
    r.Set("client_type", clientType)
    return nil
}

// ClientType Getter
func (r AlibabaAlihealthZyUploadrelationRequest) GetClientType() string {
    return r.clientType
}
