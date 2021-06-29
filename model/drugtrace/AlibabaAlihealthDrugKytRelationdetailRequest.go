package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关联关系处理详情 API请求
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

// 初始化AlibabaAlihealthDrugKytRelationdetailRequest对象
func NewAlibabaAlihealthDrugKytRelationdetailRequest() *AlibabaAlihealthDrugKytRelationdetailRequest{
    return &AlibabaAlihealthDrugKytRelationdetailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytRelationdetailRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.relationdetail"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytRelationdetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CodeActiveInfoId Setter
// 码激活文件上传信息标识
func (r *AlibabaAlihealthDrugKytRelationdetailRequest) SetCodeActiveInfoId(codeActiveInfoId string) error {
    r.codeActiveInfoId = codeActiveInfoId
    r.Set("code_active_info_id", codeActiveInfoId)
    return nil
}

// CodeActiveInfoId Getter
func (r AlibabaAlihealthDrugKytRelationdetailRequest) GetCodeActiveInfoId() string {
    return r.codeActiveInfoId
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytRelationdetailRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytRelationdetailRequest) GetRefEntId() string {
    return r.refEntId
}
// ClientType Setter
// 客户端ID【默认写2】
func (r *AlibabaAlihealthDrugKytRelationdetailRequest) SetClientType(clientType string) error {
    r.clientType = clientType
    r.Set("client_type", clientType)
    return nil
}

// ClientType Getter
func (r AlibabaAlihealthDrugKytRelationdetailRequest) GetClientType() string {
    return r.clientType
}
