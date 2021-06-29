package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
终端(医疗机构|零售药店)ID生成接口 API请求
alibaba.alihealth.drug.kyt.idgenerate

终端(医疗机构|零售药店)ID生成接口
*/
type AlibabaAlihealthDrugKytIdgenerateRequest struct {
    model.Params
    // 行政区（省市区）
    _regionCode   string
    // 零售药店、医疗机构名称
    _terminalName   string
}

// 初始化AlibabaAlihealthDrugKytIdgenerateRequest对象
func NewAlibabaAlihealthDrugKytIdgenerateRequest() *AlibabaAlihealthDrugKytIdgenerateRequest{
    return &AlibabaAlihealthDrugKytIdgenerateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytIdgenerateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.idgenerate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytIdgenerateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RegionCode Setter
// 行政区（省市区）
func (r *AlibabaAlihealthDrugKytIdgenerateRequest) SetRegionCode(_regionCode string) error {
    r._regionCode = _regionCode
    r.Set("region_code", _regionCode)
    return nil
}

// RegionCode Getter
func (r AlibabaAlihealthDrugKytIdgenerateRequest) GetRegionCode() string {
    return r._regionCode
}
// TerminalName Setter
// 零售药店、医疗机构名称
func (r *AlibabaAlihealthDrugKytIdgenerateRequest) SetTerminalName(_terminalName string) error {
    r._terminalName = _terminalName
    r.Set("terminal_name", _terminalName)
    return nil
}

// TerminalName Getter
func (r AlibabaAlihealthDrugKytIdgenerateRequest) GetTerminalName() string {
    return r._terminalName
}