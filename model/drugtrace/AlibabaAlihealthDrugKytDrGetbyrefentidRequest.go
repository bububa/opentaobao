package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
多融通过一个企业唯一标识查询企业详细信息 API请求
alibaba.alihealth.drug.kyt.dr.getbyrefentid

根据企业唯一标识查看企业详细信息
*/
type AlibabaAlihealthDrugKytDrGetbyrefentidRequest struct {
    model.Params
    // 接口调用企业的唯一标识（接口调用者）
    _refEntId   string
    // 准备要查询的企业唯一标识（返回该唯一标识企业的详细信息）
    _destRefEntId   string
}

// 初始化AlibabaAlihealthDrugKytDrGetbyrefentidRequest对象
func NewAlibabaAlihealthDrugKytDrGetbyrefentidRequest() *AlibabaAlihealthDrugKytDrGetbyrefentidRequest{
    return &AlibabaAlihealthDrugKytDrGetbyrefentidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrGetbyrefentidRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.getbyrefentid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrGetbyrefentidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugKytDrGetbyrefentidRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDrGetbyrefentidRequest) GetRefEntId() string {
    return r._refEntId
}
// DestRefEntId Setter
// 准备要查询的企业唯一标识（返回该唯一标识企业的详细信息）
func (r *AlibabaAlihealthDrugKytDrGetbyrefentidRequest) SetDestRefEntId(_destRefEntId string) error {
    r._destRefEntId = _destRefEntId
    r.Set("dest_ref_ent_id", _destRefEntId)
    return nil
}

// DestRefEntId Getter
func (r AlibabaAlihealthDrugKytDrGetbyrefentidRequest) GetDestRefEntId() string {
    return r._destRefEntId
}