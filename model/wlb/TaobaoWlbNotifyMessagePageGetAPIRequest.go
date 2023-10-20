package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbnotifymessagepagegetAPIRequest 物流宝通知消息查询接口 API请求
// taobao.wlb.notify.message.page.get
//
// 物流宝提供的消息通知查询接口，消息内容包括;出入库单不一致消息，取消订单成功消息，盘点单消息
type TaobaowlbnotifymessagepagegetAPIRequest struct {
	model.Params
	// 通知消息编码： STOCK_IN_NOT_CONSISTENT---入库单不一致 CANCEL_ORDER_SUCCESS---取消订单成功 INVENTORY_CHECK---盘点 CANCEL_ORDER_FAILURE---取消订单失败 ORDER_REJECT--wms拒单 ORDER_CONFIRMED--订单处理成功
	_msgCode string
	// 记录开始时间
	_startDate string
	// 记录截至时间
	_endDate string
	// 消息状态： 不需要确认：NO_NEED_CONFIRM 已确认：CONFIRMED 待确认：TO_BE_CONFIRM
	_status string
	// 分页查询页数
	_pageNo int64
	// 分页查询的每页页数
	_pageSize int64
}

// NewTaobaowlbnotifymessagepagegetRequest 初始化TaobaowlbnotifymessagepagegetAPIRequest对象
func NewTaobaowlbnotifymessagepagegetRequest() *TaobaowlbnotifymessagepagegetAPIRequest {
	return &TaobaowlbnotifymessagepagegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbnotifymessagepagegetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.notify.message.page.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbnotifymessagepagegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbnotifymessagepagegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMsgCode is MsgCode Setter
// 通知消息编码： STOCK_IN_NOT_CONSISTENT---入库单不一致 CANCEL_ORDER_SUCCESS---取消订单成功 INVENTORY_CHECK---盘点 CANCEL_ORDER_FAILURE---取消订单失败 ORDER_REJECT--wms拒单 ORDER_CONFIRMED--订单处理成功
func (r *TaobaowlbnotifymessagepagegetAPIRequest) SetMsgCode(_msgCode string) error {
	r._msgCode = _msgCode
	r.Set("msg_code", _msgCode)
	return nil
}

// GetMsgCode MsgCode Getter
func (r TaobaowlbnotifymessagepagegetAPIRequest) GetMsgCode() string {
	return r._msgCode
}

// SetStartDate is StartDate Setter
// 记录开始时间
func (r *TaobaowlbnotifymessagepagegetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaowlbnotifymessagepagegetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 记录截至时间
func (r *TaobaowlbnotifymessagepagegetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaowlbnotifymessagepagegetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetStatus is Status Setter
// 消息状态： 不需要确认：NO_NEED_CONFIRM 已确认：CONFIRMED 待确认：TO_BE_CONFIRM
func (r *TaobaowlbnotifymessagepagegetAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaowlbnotifymessagepagegetAPIRequest) GetStatus() string {
	return r._status
}

// SetPageNo is PageNo Setter
// 分页查询页数
func (r *TaobaowlbnotifymessagepagegetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaowlbnotifymessagepagegetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 分页查询的每页页数
func (r *TaobaowlbnotifymessagepagegetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaowlbnotifymessagepagegetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
