package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据企业唯一标识查看企业详细信息 API请求
alibaba.alihealth.drug.kyt.getbyrefentid

根据企业唯一标识查看企业详细信息
*/
type AlibabaAlihealthDrugKytGetbyrefentidRequest struct {
    model.Params
    // 接口调用企业的唯一标识（接口调用者）
    _refEntId   string
    // 准备要查询的企业唯一标识（返回该唯一标识企业的详细信息）
    _destRefEntId   string
}

// 初始化AlibabaAlihealthDrugKytGetbyrefentidRequest对象
func NewAlibabaAlihealthDrugKytGetbyrefentidRequest() *AlibabaAlihealthDrugKytGetbyrefentidRequest{
    return &AlibabaAlihealthDrugKytGetbyrefentidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytGetbyrefentidRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.getbyrefentid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytGetbyrefentidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugKytGetbyrefentidRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytGetbyrefentidRequest) GetRefEntId() string {
    return r._refEntId
}
// DestRefEntId Setter
// 准备要查询的企业唯一标识（返回该唯一标识企业的详细信息）
func (r *AlibabaAlihealthDrugKytGetbyrefentidRequest) SetDestRefEntId(_destRefEntId string) error {
    r._destRefEntId = _destRefEntId
    r.Set("dest_ref_ent_id", _destRefEntId)
    return nil
}

// DestRefEntId Getter
func (r AlibabaAlihealthDrugKytGetbyrefentidRequest) GetDestRefEntId() string {
    return r._destRefEntId
}
