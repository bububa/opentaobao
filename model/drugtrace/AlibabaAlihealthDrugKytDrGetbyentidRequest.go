package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
多融通过企业ID得到一个企业的详细信息 API请求
alibaba.alihealth.drug.kyt.dr.getbyentid

根据企业主键查看企业详细信息
*/
type AlibabaAlihealthDrugKytDrGetbyentidRequest struct {
    model.Params
    // 接口调用企业的唯一标识（接口调用者）
    _refEntId   string
    // 准备要查询的企业ID（返回该企业ID的详细信息）
    _entId   string
}

// 初始化AlibabaAlihealthDrugKytDrGetbyentidRequest对象
func NewAlibabaAlihealthDrugKytDrGetbyentidRequest() *AlibabaAlihealthDrugKytDrGetbyentidRequest{
    return &AlibabaAlihealthDrugKytDrGetbyentidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrGetbyentidRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.getbyentid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrGetbyentidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugKytDrGetbyentidRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDrGetbyentidRequest) GetRefEntId() string {
    return r._refEntId
}
// EntId Setter
// 准备要查询的企业ID（返回该企业ID的详细信息）
func (r *AlibabaAlihealthDrugKytDrGetbyentidRequest) SetEntId(_entId string) error {
    r._entId = _entId
    r.Set("ent_id", _entId)
    return nil
}

// EntId Getter
func (r AlibabaAlihealthDrugKytDrGetbyentidRequest) GetEntId() string {
    return r._entId
}
