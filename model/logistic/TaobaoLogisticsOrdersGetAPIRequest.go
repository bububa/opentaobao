package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLogisticsOrdersGetAPIRequest
批量查询物流订单 API请求
taobao.logistics.orders.get

批量查询物流订单。 */
type TaobaoLogisticsOrdersGetAPIRequest struct {
	model.Params
	// 需返回的字段列表.可选值:Shipping 物流数据结构中的以下字段: <br><br/>tid,order_code,seller_nick,buyer_nick,delivery_start, delivery_end,out_sid,item_title,receiver_name, created,modified,status,type,freight_payer,seller_confirm,company_name,sub_tids,is_spilt；<br>多个字段之间用","分隔。如tid,seller_nick,buyer_nick,delivery_start。
	_fields string
	// 交易ID.如果加入tid参数的话,不用传其他的参数,若传入tid：非拆单场景，仅会返回一条物流订单信息；拆单场景，会返回多条物流订单信息
	_tid int64
	// 买家昵称
	_buyerNick string
	// 物流状态.查看数据结构 Shipping 中的status字段.
	_status string
	// 卖家是否发货.可选值:yes(是),no(否).如:yes
	_sellerConfirm string
	// 收货人姓名
	_receiverName string
	// 创建时间开始
	_startCreated string
	// 创建时间结束
	_endCreated string
	// 谁承担运费.可选值:buyer(买家),seller(卖家).如:buyer
	_freightPayer string
	// 物流方式.可选值:post(平邮),express(快递),ems(EMS).如:post
	_type string
	// 页码.该字段没传 或 值<1 ,则默认page_no为1
	_pageNo int64
	// 每页条数.该字段没传 或 值<1 ,则默认page_size为40
	_pageSize int64
}

// New
