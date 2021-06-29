package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查企业标识信息 API请求
alibaba.alihealth.drug.kyt.smyx.getentinfo

根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
*/
type AlibabaAlihealthDrugKytSmyxGetentinfoRequest struct {
    model.Params
    // 公司名称
    entName   string
}

// 初始化AlibabaAlihealthDrugKytSmyxGetentinfoRequest对象
func NewAlibabaAlihealthDrugKytSmyxGetentinfoRequest() *AlibabaAlihealthDrugKytSmyxGetentinfoRequest{
    return &AlibabaAlihealthDrugKytSmyxGetentinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytSmyxGetentinfoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.smyx.getentinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytSmyxGetentinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EntName Setter
// 公司名称
func (r *AlibabaAlihealthDrugKytSmyxGetentinfoRequest) SetEntName(entName string) error {
    r.entName = entName
    r.Set("ent_name", entName)
    return nil
}

// EntName Getter
func (r AlibabaAlihealthDrugKytSmyxGetentinfoRequest) GetEntName() string {
    return r.entName
}
