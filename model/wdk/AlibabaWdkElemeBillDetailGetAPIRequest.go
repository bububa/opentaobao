package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkElemeBillDetailGetAPIRequest
饿了么对账单查询，带订单明细 API请求
alibaba.wdk.eleme.bill.detail.get

查询饿了么对账单信息，带订单明细 */
type AlibabaWdkElemeBillDetailGetAPIRequest struct {
	model.Params
	// 对账单查询参数
	_eleBillRequest *EleBillRequest
}

// New
