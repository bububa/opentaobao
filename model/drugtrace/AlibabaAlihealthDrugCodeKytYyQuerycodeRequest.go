package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
医院根据码查询码信息 API请求
alibaba.alihealth.drug.code.kyt.yy.querycode

服务描述
此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息。
核查平台优先过滤非8开头的，长度非20位数字的码信息。
*/
type AlibabaAlihealthDrugCodeKytYyQuerycodeRequest struct {
    model.Params
    // 企业唯一标识（或appkey）
    refEntId   string
    // 码列表
    codes   []string
}

// 初始化AlibabaAlihealthDrugCodeKytYyQuerycodeRequest对象
func NewAlibabaAlihealthDrugCodeKytYyQuerycodeRequest() *AlibabaAlihealthDrugCodeKytYyQuerycodeRequest{
    return &AlibabaAlihealthDrugCodeKytYyQuerycodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeKytYyQuerycodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.code.kyt.yy.querycode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeKytYyQuerycodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业唯一标识（或appkey）
func (r *AlibabaAlihealthDrugCodeKytYyQuerycodeRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugCodeKytYyQuerycodeRequest) GetRefEntId() string {
    return r.refEntId
}
// Codes Setter
// 码列表
func (r *AlibabaAlihealthDrugCodeKytYyQuerycodeRequest) SetCodes(codes []string) error {
    r.codes = codes
    r.Set("codes", codes)
    return nil
}

// Codes Getter
func (r AlibabaAlihealthDrugCodeKytYyQuerycodeRequest) GetCodes() []string {
    return r.codes
}
