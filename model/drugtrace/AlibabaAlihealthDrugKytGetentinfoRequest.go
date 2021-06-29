package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】 API请求
alibaba.alihealth.drug.kyt.getentinfo

根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
*/
type AlibabaAlihealthDrugKytGetentinfoRequest struct {
    model.Params
    // 公司名称
    _entName   string
}

// 初始化AlibabaAlihealthDrugKytGetentinfoRequest对象
func NewAlibabaAlihealthDrugKytGetentinfoRequest() *AlibabaAlihealthDrugKytGetentinfoRequest{
    return &AlibabaAlihealthDrugKytGetentinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytGetentinfoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.getentinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytGetentinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EntName Setter
// 公司名称
func (r *AlibabaAlihealthDrugKytGetentinfoRequest) SetEntName(_entName string) error {
    r._entName = _entName
    r.Set("ent_name", _entName)
    return nil
}

// EntName Getter
func (r AlibabaAlihealthDrugKytGetentinfoRequest) GetEntName() string {
    return r._entName
}
