package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询追溯码对应的药品信息（药店） API请求
alibaba.alihealth.drug.code.kyt.yd.querycode

此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
核查平台优先过滤非8开头的，长度非20位数字的码信息。
*/
type AlibabaAlihealthDrugCodeKytYdQuerycodeRequest struct {
    model.Params
    // 企业唯一标识
    refEntId   string
    // 码列表
    codes   []string
}

// 初始化AlibabaAlihealthDrugCodeKytYdQuerycodeRequest对象
func NewAlibabaAlihealthDrugCodeKytYdQuerycodeRequest() *AlibabaAlihealthDrugCodeKytYdQuerycodeRequest{
    return &AlibabaAlihealthDrugCodeKytYdQuerycodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeKytYdQuerycodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.code.kyt.yd.querycode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeKytYdQuerycodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业唯一标识
func (r *AlibabaAlihealthDrugCodeKytYdQuerycodeRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugCodeKytYdQuerycodeRequest) GetRefEntId() string {
    return r.refEntId
}
// Codes Setter
// 码列表
func (r *AlibabaAlihealthDrugCodeKytYdQuerycodeRequest) SetCodes(codes []string) error {
    r.codes = codes
    r.Set("codes", codes)
    return nil
}

// Codes Getter
func (r AlibabaAlihealthDrugCodeKytYdQuerycodeRequest) GetCodes() []string {
    return r.codes
}
