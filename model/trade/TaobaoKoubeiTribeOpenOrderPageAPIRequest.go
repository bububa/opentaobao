package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoKoubeiTribeOpenOrderPageAPIRequest
口碑综合体订单列表信息查询 API请求
taobao.koubei.tribe.open.order.page

查询口碑商圈用户的订单列表信息 */
type TaobaoKoubeiTribeOpenOrderPageAPIRequest struct {
	model.Params
	// 订单状态；ALL（全部），WAIT_PAY（代付款），WAIT_CONSUME（代消费）
	_orderStatus string
	// 每页大小
	_pageSize int64
	// 起始页
	_pageNo int64
	// 数据集Id
	_dataSetId string
	// 用户openId
	_openId string
}

// New
