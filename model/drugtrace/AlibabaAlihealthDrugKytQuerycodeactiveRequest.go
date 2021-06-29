package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询码是否激活 API请求
alibaba.alihealth.drug.kyt.querycodeactive

查询码是否激活
*/
type AlibabaAlihealthDrugKytQuerycodeactiveRequest struct {
    model.Params
    // 企业
    refEntId   string
    // 码
    codes   []string
}

// 初始化AlibabaAlihealthDrugKytQuerycodeactiveRequest对象
func NewAlibabaAlihealthDrugKytQuerycodeactiveRequest() *AlibabaAlihealthDrugKytQuerycodeactiveRequest{
    return &AlibabaAlihealthDrugKytQuerycodeactiveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytQuerycodeactiveRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.querycodeactive"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytQuerycodeactiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业
func (r *AlibabaAlihealthDrugKytQuerycodeactiveRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytQuerycodeactiveRequest) GetRefEntId() string {
    return r.refEntId
}
// Codes Setter
// 码
func (r *AlibabaAlihealthDrugKytQuerycodeactiveRequest) SetCodes(codes []string) error {
    r.codes = codes
    r.Set("codes", codes)
    return nil
}

// Codes Getter
func (r AlibabaAlihealthDrugKytQuerycodeactiveRequest) GetCodes() []string {
    return r.codes
}
