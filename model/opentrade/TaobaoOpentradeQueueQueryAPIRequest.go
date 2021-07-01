package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeQueueQueryAPIRequest
尖货交易排队信息查询 API请求
taobao.opentrade.queue.query

尖货交易排队信息查询 */
type TaobaoOpentradeQueueQueryAPIRequest struct {
	model.Params
	// 排队用户状态，新用户为NEW
	_status string
	// 排队活动ID，排队时如传入，这里需要填写；若未传，这里也可以不传
	_activityId string
	// 排队商品SKU ID，不存在传0
	_skuId int64
	// 排队商品ID
	_itemId int64
	// 分页参数，每页大小
	_pageSize int64
	// 分页参数，当前页，以0开始
	_pageIndex int64
}

// New
