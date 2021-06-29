package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
扫码营销码查询 API请求
alibaba.alihealth.drug.code.kyt.smyx.querycode

此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
核查平台优先过滤非8开头的，长度非20位数字的码信息。
*/
type AlibabaAlihealthDrugCodeKytSmyxQuerycodeRequest struct {
    model.Params
    // 企业唯一标识
    refEntId   string
    // 码列表
    codes   []string
}

// 初始化AlibabaAlihealthDrugCodeKytSmyxQuerycodeRequest对象
func NewAlibabaAlihealthDrugCodeKytSmyxQuerycodeRequest() *AlibabaAlihealthDrugCodeKytSmyxQuerycodeRequest{
    return &AlibabaAlihealthDrugCodeKytSmyxQuerycodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeKytSmyxQuerycodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.code.kyt.smyx.querycode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeKytSmyxQuerycodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业唯一标识
func (r *AlibabaAlihealthDrugCodeKytSmyxQuerycodeRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugCodeKytSmyxQuerycodeRequest) GetRefEntId() string {
    return r.refEntId
}
// Codes Setter
// 码列表
func (r *AlibabaAlihealthDrugCodeKytSmyxQuerycodeRequest) SetCodes(codes []string) error {
    r.codes = codes
    r.Set("codes", codes)
    return nil
}

// Codes Getter
func (r AlibabaAlihealthDrugCodeKytSmyxQuerycodeRequest) GetCodes() []string {
    return r.codes
}
