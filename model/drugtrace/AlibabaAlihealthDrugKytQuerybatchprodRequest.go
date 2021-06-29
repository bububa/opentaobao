package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批次产品查询(根据企业名和批次号查询产品信息) API请求
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

// 初始化AlibabaAlihealthDrugKytQuerybatchprodRequest对象
func NewAlibabaAlihealthDrugKytQuerybatchprodRequest() *AlibabaAlihealthDrugKytQuerybatchprodRequest{
    return &AlibabaAlihealthDrugKytQuerybatchprodRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytQuerybatchprodRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.querybatchprod"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytQuerybatchprodRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProductEntName Setter
// 生产企业名称；支持新老企业名称匹配
func (r *AlibabaAlihealthDrugKytQuerybatchprodRequest) SetProductEntName(productEntName string) error {
    r.productEntName = productEntName
    r.Set("product_ent_name", productEntName)
    return nil
}

// ProductEntName Getter
func (r AlibabaAlihealthDrugKytQuerybatchprodRequest) GetProductEntName() string {
    return r.productEntName
}
// ProductBatchNo Setter
// 生产批号;支持大小写模糊匹配
func (r *AlibabaAlihealthDrugKytQuerybatchprodRequest) SetProductBatchNo(productBatchNo string) error {
    r.productBatchNo = productBatchNo
    r.Set("product_batch_no", productBatchNo)
    return nil
}

// ProductBatchNo Getter
func (r AlibabaAlihealthDrugKytQuerybatchprodRequest) GetProductBatchNo() string {
    return r.productBatchNo
}
// BureauName Setter
// 社保局(所属地市名称)
func (r *AlibabaAlihealthDrugKytQuerybatchprodRequest) SetBureauName(bureauName string) error {
    r.bureauName = bureauName
    r.Set("bureau_name", bureauName)
    return nil
}

// BureauName Getter
func (r AlibabaAlihealthDrugKytQuerybatchprodRequest) GetBureauName() string {
    return r.bureauName
}
// TerminalName Setter
// 请求终端名称
func (r *AlibabaAlihealthDrugKytQuerybatchprodRequest) SetTerminalName(terminalName string) error {
    r.terminalName = terminalName
    r.Set("terminal_name", terminalName)
    return nil
}

// TerminalName Getter
func (r AlibabaAlihealthDrugKytQuerybatchprodRequest) GetTerminalName() string {
    return r.terminalName
}
// TerminalType Setter
// 终端类型：1005100-零售，1005200-医疗
func (r *AlibabaAlihealthDrugKytQuerybatchprodRequest) SetTerminalType(terminalType string) error {
    r.terminalType = terminalType
    r.Set("terminal_type", terminalType)
    return nil
}

// TerminalType Getter
func (r AlibabaAlihealthDrugKytQuerybatchprodRequest) GetTerminalType() string {
    return r.terminalType
}
// Invocation Setter
// 调用方式：formal-正式、test-测试
func (r *AlibabaAlihealthDrugKytQuerybatchprodRequest) SetInvocation(invocation string) error {
    r.invocation = invocation
    r.Set("invocation", invocation)
    return nil
}

// Invocation Getter
func (r AlibabaAlihealthDrugKytQuerybatchprodRequest) GetInvocation() string {
    return r.invocation
}
