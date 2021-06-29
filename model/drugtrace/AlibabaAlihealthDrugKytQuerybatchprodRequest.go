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
    _productEntName   string
    // 生产批号;支持大小写模糊匹配
    _productBatchNo   string
    // 社保局(所属地市名称)
    _bureauName   string
    // 请求终端名称
    _terminalName   string
    // 终端类型：1005100-零售，1005200-医疗
    _terminalType   string
    // 调用方式：formal-正式、test-测试
    _invocation   string
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
func (r *AlibabaAlihealthDrugKytQuerybatchprodRequest) SetProductEntName(_productEntName string) error {
    r._productEntName = _productEntName
    r.Set("product_ent_name", _productEntName)
    return nil
}

// ProductEntName Getter
func (r AlibabaAlihealthDrugKytQuerybatchprodRequest) GetProductEntName() string {
    return r._productEntName
}
// ProductBatchNo Setter
// 生产批号;支持大小写模糊匹配
func (r *AlibabaAlihealthDrugKytQuerybatchprodRequest) SetProductBatchNo(_productBatchNo string) error {
    r._productBatchNo = _productBatchNo
    r.Set("product_batch_no", _productBatchNo)
    return nil
}

// ProductBatchNo Getter
func (r AlibabaAlihealthDrugKytQuerybatchprodRequest) GetProductBatchNo() string {
    return r._productBatchNo
}
// BureauName Setter
// 社保局(所属地市名称)
func (r *AlibabaAlihealthDrugKytQuerybatchprodRequest) SetBureauName(_bureauName string) error {
    r._bureauName = _bureauName
    r.Set("bureau_name", _bureauName)
    return nil
}

// BureauName Getter
func (r AlibabaAlihealthDrugKytQuerybatchprodRequest) GetBureauName() string {
    return r._bureauName
}
// TerminalName Setter
// 请求终端名称
func (r *AlibabaAlihealthDrugKytQuerybatchprodRequest) SetTerminalName(_terminalName string) error {
    r._terminalName = _terminalName
    r.Set("terminal_name", _terminalName)
    return nil
}

// TerminalName Getter
func (r AlibabaAlihealthDrugKytQuerybatchprodRequest) GetTerminalName() string {
    return r._terminalName
}
// TerminalType Setter
// 终端类型：1005100-零售，1005200-医疗
func (r *AlibabaAlihealthDrugKytQuerybatchprodRequest) SetTerminalType(_terminalType string) error {
    r._terminalType = _terminalType
    r.Set("terminal_type", _terminalType)
    return nil
}

// TerminalType Getter
func (r AlibabaAlihealthDrugKytQuerybatchprodRequest) GetTerminalType() string {
    return r._terminalType
}
// Invocation Setter
// 调用方式：formal-正式、test-测试
func (r *AlibabaAlihealthDrugKytQuerybatchprodRequest) SetInvocation(_invocation string) error {
    r._invocation = _invocation
    r.Set("invocation", _invocation)
    return nil
}

// Invocation Getter
func (r AlibabaAlihealthDrugKytQuerybatchprodRequest) GetInvocation() string {
    return r._invocation
}
