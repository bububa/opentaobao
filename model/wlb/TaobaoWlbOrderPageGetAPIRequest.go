package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbOrderPageGetAPIRequest 分页查询物流宝订单 API请求
// taobao.wlb.order.page.get
//
// 分页查询物流宝订单
type TaobaoWlbOrderPageGetAPIRequest struct {
	model.Params
	// 物流订单编号
	_orderCode string
	// 查询截止时间
	_endTime string
	// 查询开始时间
	_startTime string
	// 订单类型: （1）NORMAL_OUT ：正常出库 （2）NORMAL_IN：正常入库 （3）RETURN_IN：退货入库 （4）EXCHANGE_OUT：换货出库
	_orderType string
	// 订单子类型： （1）OTHER： 其他 （2）TAOBAO_TRADE： 淘宝交易 （3）OTHER_TRADE：其他交易 （4）ALLCOATE： 调拨 （5）CHECK: 盘点单 （6）PURCHASE: 采购单
	_orderSubType string
	// 每页多少条
	_pageSize int64
	// 分页的第几页
	_pageNo int64
	// TMS拒签：-100 接收方拒签：-200
	_orderStatus int64
}

// NewTaobaoWlbOrderPageGetRequest 初始化TaobaoWlbOrderPageGetAPIRequest对象
func NewTaobaoWlbOrderPageGetRequest() *TaobaoWlbOrderPageGetAPIRequest {
	return &TaobaoWlbOrderPageGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbOrderPageGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.order.page.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbOrderPageGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbOrderPageGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// 物流订单编号
func (r *TaobaoWlbOrderPageGetAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r TaobaoWlbOrderPageGetAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetEndTime is EndTime Setter
// 查询截止时间
func (r *TaobaoWlbOrderPageGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoWlbOrderPageGetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetStartTime is StartTime Setter
// 查询开始时间
func (r *TaobaoWlbOrderPageGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoWlbOrderPageGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetOrderType is OrderType Setter
// 订单类型: （1）NORMAL_OUT ：正常出库 （2）NORMAL_IN：正常入库 （3）RETURN_IN：退货入库 （4）EXCHANGE_OUT：换货出库
func (r *TaobaoWlbOrderPageGetAPIRequest) SetOrderType(_orderType string) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// GetOrderType OrderType Getter
func (r TaobaoWlbOrderPageGetAPIRequest) GetOrderType() string {
	return r._orderType
}

// SetOrderSubType is OrderSubType Setter
// 订单子类型： （1）OTHER： 其他 （2）TAOBAO_TRADE： 淘宝交易 （3）OTHER_TRADE：其他交易 （4）ALLCOATE： 调拨 （5）CHECK: 盘点单 （6）PURCHASE: 采购单
func (r *TaobaoWlbOrderPageGetAPIRequest) SetOrderSubType(_orderSubType string) error {
	r._orderSubType = _orderSubType
	r.Set("order_sub_type", _orderSubType)
	return nil
}

// GetOrderSubType OrderSubType Getter
func (r TaobaoWlbOrderPageGetAPIRequest) GetOrderSubType() string {
	return r._orderSubType
}

// SetPageSize is PageSize Setter
// 每页多少条
func (r *TaobaoWlbOrderPageGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoWlbOrderPageGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 分页的第几页
func (r *TaobaoWlbOrderPageGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoWlbOrderPageGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetOrderStatus is OrderStatus Setter
// TMS拒签：-100 接收方拒签：-200
func (r *TaobaoWlbOrderPageGetAPIRequest) SetOrderStatus(_orderStatus int64) error {
	r._orderStatus = _orderStatus
	r.Set("order_status", _orderStatus)
	return nil
}

// GetOrderStatus OrderStatus Getter
func (r TaobaoWlbOrderPageGetAPIRequest) GetOrderStatus() int64 {
	return r._orderStatus
}
