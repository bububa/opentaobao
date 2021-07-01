package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWttOfflineRecordQueryagentinfoAPIRequest
线下推广充送等业务订单来源 API请求
alibaba.wtt.offline.record.queryagentinfo

线下推广充送等业务订单来源的查询接口 */
type AlibabaWttOfflineRecordQueryagentinfoAPIRequest struct {
	model.Params
	// 淘宝订单号
	_orderId int64
	// 业务号码
	_phone string
}

// New
