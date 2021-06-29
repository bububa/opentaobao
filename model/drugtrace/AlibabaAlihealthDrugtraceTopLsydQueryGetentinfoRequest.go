package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过企业名得到唯一标识ref_ent_id及企业ent_id API请求
alibaba.alihealth.drugtrace.top.lsyd.query.getentinfo

根据企业名称查询ID
*/
type AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoRequest struct {
    model.Params
    // 公司名称(全称)
    entName   string
}

// 初始化AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoRequest对象
func NewAlibabaAlihealthDrugtraceTopLsydQueryGetentinfoRequest() *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoRequest{
    return &AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.lsyd.query.getentinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EntName Setter
// 公司名称(全称)
func (r *AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoRequest) SetEntName(entName string) error {
    r.entName = entName
    r.Set("ent_name", entName)
    return nil
}

// EntName Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryGetentinfoRequest) GetEntName() string {
    return r.entName
}
