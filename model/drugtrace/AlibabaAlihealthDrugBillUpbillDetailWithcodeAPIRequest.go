package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugBillUpbillDetailWithcodeAPIRequest
查询上游出库单明细(带追溯码信息) API请求
alibaba.alihealth.drug.bill.upbill.detail.withcode

查询上游出库单明细(带追溯码信息) */
type AlibabaAlihealthDrugBillUpbillDetailWithcodeAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
	// 单据编码
	_billCode string
	// 发货企业renEntId
	_fromRefUserId string
	// 收货企业refEntId
	_toRefUserId string
}

// New
