package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkElemeBillGetAPIRequest
饿了么日维度对账单查询 API请求
alibaba.wdk.eleme.bill.get

查询饿了么日维度对账单信息 */
type AlibabaWdkElemeBillGetAPIRequest struct {
	model.Params
	// 对账单查询参数
	_eleBillRequest *EleBillRequest
}

// New
