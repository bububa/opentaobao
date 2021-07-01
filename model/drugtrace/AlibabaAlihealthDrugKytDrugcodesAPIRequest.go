package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytDrugcodesAPIRequest
药品是否赋码 API请求
alibaba.alihealth.drug.kyt.drugcodes

药品是否赋码 */
type AlibabaAlihealthDrugKytDrugcodesAPIRequest struct {
	model.Params
	// 企业名称
	_refEntName string
	// 药品名称
	_physicName string
	// 生产批号
	_produceBatchNo string
	// 药品类型
	_physicType string
	// 包装规格
	_pkgSpec string
	// 制剂规格
	_prepnSpec string
}

// New
