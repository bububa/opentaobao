package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytDestbillListAPIRequest
直调单据查询 API请求
alibaba.alihealth.drug.kyt.destbill.list

为药企提供直调单据查询功能 */
type AlibabaAlihealthDrugKytDestbillListAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 开始时间，格式yyyy-MM-dd
	_beginDate string
	// 结束时间，格式yyyy-MM-dd
	_endDate string
	// 单据编号
	_billCode string
	// 审核状态，1：未审核；2：审核通过；3：审核失败
	_approvalStatus string
}

// New
