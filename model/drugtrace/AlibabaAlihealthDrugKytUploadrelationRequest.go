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
    saveCodeRelation   *SaveCodeRelationType
    // affirmFlag
    affirmFlag   string
    // fileContent
    fileContent   string
    // 加密之后的文件内容字符串
    fileContentString   string
    // 上传文件的企业ID
    refEntId   string
    // 客户端类型
    clientType   string
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
func (r *AlibabaAlihealthDrugKytUploadrelationRequest) SetSaveCodeRelation(saveCodeRelation *SaveCodeRelationType) error {
    r.saveCodeRelation = saveCodeRelation
    r.Set("save_code_relation", saveCodeRelation)
    return nil
}

// SaveCodeRelation Getter
func (r AlibabaAlihealthDrugKytUploadrelationRequest) GetSaveCodeRelation() *SaveCodeRelationType {
    return r.saveCodeRelation
}
// AffirmFlag Setter
// affirmFlag
func (r *AlibabaAlihealthDrugKytUploadrelationRequest) SetAffirmFlag(affirmFlag string) error {
    r.affirmFlag = affirmFlag
    r.Set("affirm_flag", affirmFlag)
    return nil
}

// AffirmFlag Getter
func (r AlibabaAlihealthDrugKytUploadrelationRequest) GetAffirmFlag() string {
    return r.affirmFlag
}
// FileContent Setter
// fileContent
func (r *AlibabaAlihealthDrugKytUploadrelationRequest) SetFileContent(fileContent string) error {
    r.fileContent = fileContent
    r.Set("file_content", fileContent)
    return nil
}

// FileContent Getter
func (r AlibabaAlihealthDrugKytUploadrelationRequest) GetFileContent() string {
    return r.fileContent
}
// FileContentString Setter
// 加密之后的文件内容字符串
func (r *AlibabaAlihealthDrugKytUploadrelationRequest) SetFileContentString(fileContentString string) error {
    r.fileContentString = fileContentString
    r.Set("file_content_string", fileContentString)
    return nil
}

// FileContentString Getter
func (r AlibabaAlihealthDrugKytUploadrelationRequest) GetFileContentString() string {
    return r.fileContentString
}
// RefEntId Setter
// 上传文件的企业ID
func (r *AlibabaAlihealthDrugKytUploadrelationRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytUploadrelationRequest) GetRefEntId() string {
    return r.refEntId
}
// ClientType Setter
// 客户端类型
func (r *AlibabaAlihealthDrugKytUploadrelationRequest) SetClientType(clientType string) error {
    r.clientType = clientType
    r.Set("client_type", clientType)
    return nil
}

// ClientType Getter
func (r AlibabaAlihealthDrugKytUploadrelationRequest) GetClientType() string {
    return r.clientType
}
