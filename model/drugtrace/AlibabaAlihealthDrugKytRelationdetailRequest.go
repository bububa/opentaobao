package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关联关系处理详情 APIRequest
alibaba.alihealth.drug.kyt.relationdetail

关联关系处理详情
*/
type AlibabaAlihealthDrugKytRelationdetailRequest struct {
    model.Params

    // 码激活文件上传信息标识
    codeActiveInfoId   string 

    // 企业ID
    refEntId   string 

    // 客户端ID【默认写2】
    clientType   string 

}

func NewAlibabaAlihealthDrugKytRelationdetailRequest() *AlibabaAlihealthDrugKytRelationdetailRequest{
    return &AlibabaAlihealthDrugKytRelationdetailRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytRelationdetailRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.relationdetail"
}

func (r AlibabaAlihealthDrugKytRelationdetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytRelationdetailRequest) SetCodeActiveInfoId(codeActiveInfoId string) error {
    r.codeActiveInfoId = codeActiveInfoId
    r.Set("code_active_info_id", codeActiveInfoId)
    return nil
}

func (r AlibabaAlihealthDrugKytRelationdetailRequest) GetCodeActiveInfoId() string {
    return r.codeActiveInfoId
}

func (r *AlibabaAlihealthDrugKytRelationdetailRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytRelationdetailRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytRelationdetailRequest) SetClientType(clientType string) error {
    r.clientType = clientType
    r.Set("client_type", clientType)
    return nil
}

func (r AlibabaAlihealthDrugKytRelationdetailRequest) GetClientType() string {
    return r.clientType
}

