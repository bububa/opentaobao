package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardQuerybysellerAPIRequest
工单查询接口（面向商家） API请求
tmall.servicecenter.workcard.querybyseller

查询工单 */
type TmallServicecenterWorkcardQuerybysellerAPIRequest struct {
	model.Params
	// 门店/网点id
	_serviceStoreId int64
	// 核销码
	_identifyCode string
	// 工单id
	_id int64
	// 工单创建开始时间
	_gmtCreateStart string
	// 工单创建结束时间，必须与工单创建开始时间一起传入，且间隔不超过15分钟
	_gmtCreateEnd string
	// 淘宝交易订单号。主订单或子订单均可
	_bizOrderId int64
	// 当前页数
	_currentPage int64
	// 每页大小
	_pageSize int64
}

// New
