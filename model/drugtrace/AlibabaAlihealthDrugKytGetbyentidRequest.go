package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据企业主键查看企业详细信息 API请求
alibaba.alihealth.drug.kyt.getbyentid

根据企业主键查看企业详细信息
*/
type AlibabaAlihealthDrugKytGetbyentidRequest struct {
    model.Params
    // 接口调用企业的唯一标识（接口调用者）
    refEntId   string
    // 准备要查询的企业ID（返回该企业ID的详细信息）
    entId   string
}

// 初始化AlibabaAlihealthDrugKytGetbyentidRequest对象
func NewAlibabaAlihealthDrugKytGetbyentidRequest() *AlibabaAlihealthDrugKytGetbyentidRequest{
    return &AlibabaAlihealthDrugKytGetbyentidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytGetbyentidRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.getbyentid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytGetbyentidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugKytGetbyentidRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytGetbyentidRequest) GetRefEntId() string {
    return r.refEntId
}
// EntId Setter
// 准备要查询的企业ID（返回该企业ID的详细信息）
func (r *AlibabaAlihealthDrugKytGetbyentidRequest) SetEntId(entId string) error {
    r.entId = entId
    r.Set("ent_id", entId)
    return nil
}

// EntId Getter
func (r AlibabaAlihealthDrugKytGetbyentidRequest) GetEntId() string {
    return r.entId
}
