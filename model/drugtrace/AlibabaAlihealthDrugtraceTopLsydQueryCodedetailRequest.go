package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据码查询码信息 API请求
alibaba.alihealth.drugtrace.top.lsyd.query.codedetail

服务描述
此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息。
核查平台优先过滤非8开头的，长度非20位数字的码信息。
*/
type AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIRequest struct {
    model.Params
    // 企业唯一标识（或appkey）
    _refEntId   string
    // 码列表
    _codes   []string
}

// 初始化AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIRequest对象
func NewAlibabaAlihealthDrugtraceTopLsydQueryCodedetailRequest() *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIRequest{
    return &AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.lsyd.query.codedetail"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业唯一标识（或appkey）
func (r *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIRequest) GetRefEntId() string {
    return r._refEntId
}
// Codes Setter
// 码列表
func (r *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIRequest) SetCodes(_codes []string) error {
    r._codes = _codes
    r.Set("codes", _codes)
    return nil
}

// Codes Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryCodedetailAPIRequest) GetCodes() []string {
    return r._codes
}
