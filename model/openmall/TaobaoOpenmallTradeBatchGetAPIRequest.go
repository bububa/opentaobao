package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallTradeBatchGetAPIRequest
批量获取openmall订单 API请求
taobao.openmall.trade.batch.get

批量获取openmall订单
注意：该接口数据存在延迟，实时数据请通过taobao.openmall.trade.get获取 */
type TaobaoOpenmallTradeBatchGetAPIRequest struct {
	model.Params
	// 查询范围结束时间，闭区间
	_endCreated string
	// 查询页码，从1开始
	_pageIndex int64
	// 页面大小，不超过100
	_pageSize int64
	// 渠道商Nick
	_distributor string
	// 查询范围开始时间，闭区间
	_startCreated string
}

// New
