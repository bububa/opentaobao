package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytQuerybatchprodAPIRequest
批次产品查询(根据企业名和批次号查询产品信息) API请求
alibaba.alihealth.drug.kyt.querybatchprod

根据企业名和批次号查询药品信息，支持使用更名之前的老企业名查询，支持批次号大小写模糊，应用于药店或医院入库环节，通过在入库环节获取赋码的产品目录，可强制要求对相应的产品必须进行扫码入库； */
type AlibabaAlihealthDrugKytQuerybatchprodAPIRequest struct {
	model.Params
	// 生产企业名称；支持新老企业名称匹配
	_productEntName string
	// 生产批号;支持大小写模糊匹配
	_productBatchNo string
	// 社保局(所属地市名称)
	_bureauName string
	// 请求终端名称
	_terminalName string
	// 终端类型：1005100-零售，1005200-医疗
	_terminalType string
	// 调用方式：formal-正式、test-测试
	_invocation string
}

// NewAlibabaAlihealthDrugKytQuerybatchprodRequest 初始化AlibabaAlihealthDrugKytQuerybatchprodAPIRequest对象
func NewAlibabaAlihealthDrugKytQuerybatchprodRequest() *AlibabaAlihealthDrugKytQuerybatchprodAPIRequest {
	return &AlibabaAlihealthDrugKytQuerybatchprodAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytQuerybatchprodAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.querybatchprod"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytQuerybatchprodAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ProductEntName Setter
// 生产企业名称；支持新老企业名称匹配
func (r *AlibabaAlihealthDrugKytQuerybatchprodAPIRequest) SetProductEntName(_productEntName string) error {
	r._productEntName = _productEntName
	r.Set("product_ent_name", _productEntName)
	return nil
}

// Get ProductEntName Getter
func (r AlibabaAlihealthDrugKytQuerybatchprodAPIRequest) GetProductEntName() string {
	return r._productEntName
}

// Set is ProductBatchNo Setter
// 生产批号;支持大小写模糊匹配
func (r *AlibabaAlihealthDrugKytQuerybatchprodAPIRequest) SetProductBatchNo(_productBatchNo string) error {
	r._productBatchNo = _productBatchNo
	r.Set("product_batch_no", _productBatchNo)
	return nil
}

// Get ProductBatchNo Getter
func (r AlibabaAlihealthDrugKytQuerybatchprodAPIRequest) GetProductBatchNo() string {
	return r._productBatchNo
}

// Set is BureauName Setter
// 社保局(所属地市名称)
func (r *AlibabaAlihealthDrugKytQuerybatchprodAPIRequest) SetBureauName(_bureauName string) error {
	r._bureauName = _bureauName
	r.Set("bureau_name", _bureauName)
	return nil
}

// Get BureauName Getter
func (r AlibabaAlihealthDrugKytQuerybatchprodAPIRequest) GetBureauName() string {
	return r._bureauName
}

// Set is TerminalName Setter
// 请求终端名称
func (r *AlibabaAlihealthDrugKytQuerybatchprodAPIRequest) SetTerminalName(_terminalName string) error {
	r._terminalName = _terminalName
	r.Set("terminal_name", _terminalName)
	return nil
}

// Get TerminalName Getter
func (r AlibabaAlihealthDrugKytQuerybatchprodAPIRequest) GetTerminalName() string {
	return r._terminalName
}

// Set is TerminalType Setter
// 终端类型：1005100-零售，1005200-医疗
func (r *AlibabaAlihealthDrugKytQuerybatchprodAPIRequest) SetTerminalType(_terminalType string) error {
	r._terminalType = _terminalType
	r.Set("terminal_type", _terminalType)
	return nil
}

// Get TerminalType Getter
func (r AlibabaAlihealthDrugKytQuerybatchprodAPIRequest) GetTerminalType() string {
	return r._terminalType
}

// Set is Invocation Setter
// 调用方式：formal-正式、test-测试
func (r *AlibabaAlihealthDrugKytQuerybatchprodAPIRequest) SetInvocation(_invocation string) error {
	r._invocation = _invocation
	r.Set("invocation", _invocation)
	return nil
}

// Get Invocation Getter
func (r AlibabaAlihealthDrugKytQuerybatchprodAPIRequest) GetInvocation() string {
	return r._invocation
}
