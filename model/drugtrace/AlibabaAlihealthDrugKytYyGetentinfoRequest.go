package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
得到企业信息 API请求
alibaba.alihealth.drug.kyt.yy.getentinfo

根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
*/
type AlibabaAlihealthDrugKytYyGetentinfoAPIRequest struct {
    model.Params
    // 公司名称
    _entName   string
}

// 初始化AlibabaAlihealthDrugKytYyGetentinfoAPIRequest对象
func NewAlibabaAlihealthDrugKytYyGetentinfoRequest() *AlibabaAlihealthDrugKytYyGetentinfoAPIRequest{
    return &AlibabaAlihealthDrugKytYyGetentinfoAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytYyGetentinfoAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.yy.getentinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytYyGetentinfoAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EntName Setter
// 公司名称
func (r *AlibabaAlihealthDrugKytYyGetentinfoAPIRequest) SetEntName(_entName string) error {
    r._entName = _entName
    r.Set("ent_name", _entName)
    return nil
}

// EntName Getter
func (r AlibabaAlihealthDrugKytYyGetentinfoAPIRequest) GetEntName() string {
    return r._entName
}
