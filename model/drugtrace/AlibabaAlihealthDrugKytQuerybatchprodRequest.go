package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批次产品查询(根据企业名和批次号查询产品信息) APIRequest
alibaba.alihealth.drug.kyt.querybatchprod

根据企业名和批次号查询药品信息，支持使用更名之前的老企业名查询，支持批次号大小写模糊，应用于药店或医院入库环节，通过在入库环节获取赋码的产品目录，可强制要求对相应的产品必须进行扫码入库；
*/
type AlibabaAlihealthDrugKytQuerybatchprodRequest struct {
    model.Params

    // 生产企业名称；支持新老企业名称匹配
    productEntName   string 

    // 生产批号;支持大小写模糊匹配
    productBatchNo   string 

    // 社保局(所属地市名称)
    bureauName   string 

    // 请求终端名称
    terminalName   string 

    // 终端类型：1005100-零售，1005200-医疗
    terminalType   string 

    // 调用方式：formal-正式、test-测试
    invocation   string 

}

func NewAlibabaAlihealthDrugKytQuerybatchprodRequest() *AlibabaAlihealthDrugKytQuerybatchprodRequest{
    return &AlibabaAlihealthDrugKytQuerybatchprodRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytQuerybatchprodRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.querybatchprod"
}

func (r AlibabaAlihealthDrugKytQuerybatchprodRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytQuerybatchprodRequest) SetProductEntName(productEntName string) error {
    r.productEntName = productEntName
    r.Set("product_ent_name", productEntName)
    return nil
}

func (r AlibabaAlihealthDrugKytQuerybatchprodRequest) GetProductEntName() string {
    return r.productEntName
}

func (r *AlibabaAlihealthDrugKytQuerybatchprodRequest) SetProductBatchNo(productBatchNo string) error {
    r.productBatchNo = productBatchNo
    r.Set("product_batch_no", productBatchNo)
    return nil
}

func (r AlibabaAlihealthDrugKytQuerybatchprodRequest) GetProductBatchNo() string {
    return r.productBatchNo
}

func (r *AlibabaAlihealthDrugKytQuerybatchprodRequest) SetBureauName(bureauName string) error {
    r.bureauName = bureauName
    r.Set("bureau_name", bureauName)
    return nil
}

func (r AlibabaAlihealthDrugKytQuerybatchprodRequest) GetBureauName() string {
    return r.bureauName
}

func (r *AlibabaAlihealthDrugKytQuerybatchprodRequest) SetTerminalName(terminalName string) error {
    r.terminalName = terminalName
    r.Set("terminal_name", terminalName)
    return nil
}

func (r AlibabaAlihealthDrugKytQuerybatchprodRequest) GetTerminalName() string {
    return r.terminalName
}

func (r *AlibabaAlihealthDrugKytQuerybatchprodRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

func (r AlibabaAlihealthDrugKytQuerybatchprodRequest) GetTerminalType() string {
    return r.terminalType
}

func (r *AlibabaAlihealthDrugKytQuerybatchprodRequest) SetInvocation(invocation string) error {
    r.invocation = invocation
    r.Set("invocation", invocation)
    return nil
}

func (r AlibabaAlihealthDrugKytQuerybatchprodRequest) GetInvocation() string {
    return r.invocation
}

