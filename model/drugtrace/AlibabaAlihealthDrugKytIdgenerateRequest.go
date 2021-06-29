package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
终端(医疗机构|零售药店)ID生成接口 APIRequest
alibaba.alihealth.drug.kyt.idgenerate

终端(医疗机构|零售药店)ID生成接口
*/
type AlibabaAlihealthDrugKytIdgenerateRequest struct {
    model.Params

    // 行政区（省市区）
    regionCode   string 

    // 零售药店、医疗机构名称
    terminalName   string 

}

func NewAlibabaAlihealthDrugKytIdgenerateRequest() *AlibabaAlihealthDrugKytIdgenerateRequest{
    return &AlibabaAlihealthDrugKytIdgenerateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytIdgenerateRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.idgenerate"
}

func (r AlibabaAlihealthDrugKytIdgenerateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytIdgenerateRequest) SetRegionCode(regionCode string) error {
    r.regionCode = regionCode
    r.Set("region_code", regionCode)
    return nil
}

func (r AlibabaAlihealthDrugKytIdgenerateRequest) GetRegionCode() string {
    return r.regionCode
}

func (r *AlibabaAlihealthDrugKytIdgenerateRequest) SetTerminalName(terminalName string) error {
    r.terminalName = terminalName
    r.Set("terminal_name", terminalName)
    return nil
}

func (r AlibabaAlihealthDrugKytIdgenerateRequest) GetTerminalName() string {
    return r.terminalName
}

