package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJipiaoAgentOrderSearchAPIRequest
【机票代理商订单】订单搜索 API请求
taobao.jipiao.agent.order.search

卖家根据条件查询淘宝订单id列表 */
type TaobaoJipiaoAgentOrderSearchAPIRequest struct {
	model.Params
	// 创建订单时间范围的开始时间，注意：当前搜索条件开始结束时间范围不能超过三天，默认开始时间为当前时间往前推三天 （具体天数可能调整）
	_beginTime string
	// 创建订单时间范围的结束时间，注意：当前搜索条件开始结束时间范围不能超过三天，默认为当前时间 （具体天数可能调整）
	_endTime string
	// 订单状态，默认为空，查询所有状态的订单<br/>1:待确认<br/>2:待出票<br/>3:强制成功<br/>4:未付款<br/>5:预订成功<br/>6:已失效
	_status int64
	// 航程类型： 0.单程；1.往返
	_tripType int64
	// 是否需要行程单，true表示需要行程单；false表示不许要
	_hasItinerary bool
	// 页码，默认第一页；注意：页大小固定，搜索结果中返回页大小pageSize，和是否包含下一页hasNext
	_page int64
	// 扩展字段:<br/>needFilterAutobook：默认true。待出票状态下，会根据此值过滤掉自动出票状态下订单，以防止重复出票的问题。对于精选票，此值需要设置成false，并由API使用者保证不会重复出票。
	_extra string
}

// New
