package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLogisticsOrdersDetailGetAPIRequest
批量查询物流订单,返回详细信息 API请求
taobao.logistics.orders.detail.get

查询物流订单的详细信息，涉及用户隐私字段。 */
type TaobaoLogisticsOrdersDetailGetAPIRequest struct {
	model.Params
	// 需返回的字段列表.可选值:Shipping 物流数据结构中所有字段.fileds中可以指定返回以上任意一个或者多个字段,以","分隔.
	_fields string
	// 交易ID.如果加入tid参数的话,不用传其他的参数,但是仅会返回一条物流订单信息.
	_tid int64
	// 买家昵称
	_buyerNick string
	// 物流状态.可查看数据结构 Shipping 中的status字段.
	_status string
	// 卖家是否发货.可选值:yes(是),no(否).如:yes.
	_sellerConfirm string
	// 收货人姓名
	_receiverName string
	// 创建时间开始.格式:yyyy-MM-dd HH:mm:ss
	_startCreated string
	// 创建时间结束.格式:yyyy-MM-dd HH:mm:ss
	_endCreated string
	// 谁承担运费.可选值:buyer(买家),seller(卖家).如:buyer
	_freightPayer string
	// 物流方式.可选值:post(平邮),express(快递),ems(EMS).如:post
	_type string
	// 页码.该字段没传 或 值<1 ,则默认page_no为1
	_pageNo int64
	// 每页条数.该字段没传 或 值<1 ，则默认page_size为40
	_pageSize int64
}

// New
