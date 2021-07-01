package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallRefundBatchGetAPIRequest
批量获取openmall退款单 API请求
taobao.openmall.refund.batch.get

批量获取openmall退款单
注意：该接口信息存在延迟，如需实时详情请访问taobao.openmall.refund.get */
type TaobaoOpenmallRefundBatchGetAPIRequest struct {
	model.Params
	// 查询范围结束时间，闭区间
	_endCreated string
	// 翻页页码，从1开始
	_pageIndex int64
	// 页面大小，不超过100
	_pageSize int64
	// 查询的渠道商Nick
	_distributor string
	// 查询范围开始时间，闭区间
	_startCreated string
}

// New
