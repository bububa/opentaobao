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
type AlibabaAlihealthDrugKytYyGetentinfoRequest struct {
    model.Params
    // 公司名称
    entName   string
}

// 初始化AlibabaAlihealthDrugKytYyGetentinfoRequest对象
func NewAlibabaAlihealthDrugKytYyGetentinfoRequest() *AlibabaAlihealthDrugKytYyGetentinfoRequest{
    return &AlibabaAlihealthDrugKytYyGetentinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytYyGetentinfoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.yy.getentinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytYyGetentinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EntName Setter
// 公司名称
func (r *AlibabaAlihealthDrugKytYyGetentinfoRequest) SetEntName(entName string) error {
    r.entName = entName
    r.Set("ent_name", entName)
    return nil
}

// EntName Getter
func (r AlibabaAlihealthDrugKytYyGetentinfoRequest) GetEntName() string {
    return r.entName
}
