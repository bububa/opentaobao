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

// New
