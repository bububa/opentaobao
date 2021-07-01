package xhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelDataServiceOrderDetailAPIRequest
服务订单详情 API请求
taobao.xhotel.data.service.order.detail

服务订单详情top接口构建 */
type TaobaoXhotelDataServiceOrderDetailAPIRequest struct {
	model.Params
	// 查询开始日期
	_startDate string
	// 是否查无订单 1:是 0:否
	_isCallNoOrder int64
	// 酒店id
	_hid int64
	// 是否特殊时段订单 1:是 0:否
	_isSpecTimeOrder int64
	// 渠道商名称
	_vendor string
	// 分页大小
	_pageSize int64
	// 查询结束时间
	_endDate string
	// 是否到店无房 1:是
	_isNoRoomCompen int64
	// 分页参数
	_startRow int64
	// 是否拒单 1:是 0:否
	_isSellerDeny int64
	// 是否卖家原因退款 1:是 0:否
	_isSellerRefund int64
	// 供应商名称
	_supplier string
}

// New
