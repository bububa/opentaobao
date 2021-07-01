package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytDrugrescodeAPIRequest
查询药品码段信息 API请求
alibaba.alihealth.drug.kyt.drugrescode

查询药品码段信息 */
type AlibabaAlihealthDrugKytDrugrescodeAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 药品通用名
	_physicName string
	// 批准文号
	_approvalLicenceNo string
	// 开始日期
	_startDate string
	// 结束日期
	_endDate string
	// 页大小
	_pageSize int64
	// 页码
	_page int64
	// 企业名称
	_entName string
	// 包装规格
	_packageSpec string
	// 制剂规格
	_prepnSpec string
}

// New
