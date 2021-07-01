package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelOrderSearchAPIRequest
酒店产品库订单查询 API请求
taobao.xhotel.order.search

酒店产品库订单查询功能，查询90天内的订单 */
type TaobaoXhotelOrderSearchAPIRequest struct {
	model.Params
	// 酒店订单oids列表，多个oid用英文逗号隔开，一次不超过20个。
	_orderIds string
	// 订单创建时间查询结束时间，格式为：yyyy-MM-dd HH:mm:ss。不能早于created_start或者间隔不能大于30
	_createdEnd string
	// 订单创建时间查询起始时间，格式为：yyyy-MM-dd HH:mm:ss
	_createdStart string
	// 分页页码。取值范围，大于零的整数，默认值1，即返回第一页的数据。页面大小为20
	_pageNo int64
	// 酒店订单tids列表，多个tid用英文逗号隔开，一次不超过20个。oids和tids都传的情况下默认使用tids
	_orderTids string
	// 系统商标识
	_vendor string
	// 外部订单号out_oids列表，多个oid用英文逗号隔开，一次不超过20个。
	_outOids string
}

// New
