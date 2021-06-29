package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关联关系上传 APIRequest
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

func NewAlibabaAlihealthDrugKytUploadrelationRequest() *AlibabaAlihealthDrugKytUploadrelationRequest{
    return &AlibabaAlihealthDrugKytUploadrelationRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytUploadrelationRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.uploadrelation"
}

func (r AlibabaAlihealthDrugKytUploadrelationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytUploadrelationRequest) SetSaveCodeRelation(saveCodeRelation *SaveCodeRelationType) error {
    r.saveCodeRelation = saveCodeRelation
    r.Set("save_code_relation", saveCodeRelation)
    return nil
}

func (r AlibabaAlihealthDrugKytUploadrelationRequest) GetSaveCodeRelation() *SaveCodeRelationType {
    return r.saveCodeRelation
}

func (r *AlibabaAlihealthDrugKytUploadrelationRequest) SetAffirmFlag(affirmFlag string) error {
    r.affirmFlag = affirmFlag
    r.Set("affirm_flag", affirmFlag)
    return nil
}

func (r AlibabaAlihealthDrugKytUploadrelationRequest) GetAffirmFlag() string {
    return r.affirmFlag
}

func (r *AlibabaAlihealthDrugKytUploadrelationRequest) SetFileContent(fileContent string) error {
    r.fileContent = fileContent
    r.Set("file_content", fileContent)
    return nil
}

func (r AlibabaAlihealthDrugKytUploadrelationRequest) GetFileContent() string {
    return r.fileContent
}

func (r *AlibabaAlihealthDrugKytUploadrelationRequest) SetFileContentString(fileContentString string) error {
    r.fileContentString = fileContentString
    r.Set("file_content_string", fileContentString)
    return nil
}

func (r AlibabaAlihealthDrugKytUploadrelationRequest) GetFileContentString() string {
    return r.fileContentString
}

func (r *AlibabaAlihealthDrugKytUploadrelationRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytUploadrelationRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytUploadrelationRequest) SetClientType(clientType string) error {
    r.clientType = clientType
    r.Set("client_type", clientType)
    return nil
}

func (r AlibabaAlihealthDrugKytUploadrelationRequest) GetClientType() string {
    return r.clientType
}

