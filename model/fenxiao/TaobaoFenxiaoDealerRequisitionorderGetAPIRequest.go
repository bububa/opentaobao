package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaodealerrequisitionordergetAPIRequest 批量查询采购申请/经销采购单 API请求
// taobao.fenxiao.dealer.requisitionorder.get
//
// 批量查询采购申请/经销采购单，目前支持供应商和分销商查询
type TaobaofenxiaodealerrequisitionordergetAPIRequest struct {
	model.Params
	// 采购申请/经销采购单最早修改时间
	_startDate string
	// 采购申请/经销采购单最迟修改时间。与start_date字段的最大时间间隔不能超过30天
	_endDate string
	// 多个字段用","分隔。 fields 如果为空：返回所有采购申请/经销采购单对象(dealer_orders)字段。 如果不为空：返回指定采购单对象(dealer_orders)字段。 例1： dealer_order_details.product_id 表示只返回product_id 例2： dealer_order_details 表示只返回明细列表
	_fields string
	// 页码（大于0的整数。无值或小于1的值按默认值1计）
	_pageNo int64
	// 每页条数（大于0但小于等于50的整数。无值或大于50或小于1的值按默认值50计）
	_pageSize int64
	// 采购申请/经销采购单状态。<br/>0：全部状态。<br/>1：分销商提交申请，待供应商审核。<br/>2：供应商驳回申请，待分销商确认。<br/>3：供应商修改后，待分销商确认。<br/>4：分销商拒绝修改，待供应商再审核。<br/>5：审核通过下单成功，待分销商付款。<br/>7：付款成功，待供应商发货。<br/>8：供应商发货，待分销商收货。<br/>9：分销商收货，交易成功。<br/>10：采购申请/经销采购单关闭。<br/><br/>注：无值按默认值0计，超出状态范围返回错误信息。
	_orderStatus int64
	// 查询者自己在所要查询的采购申请/经销采购单中的身份。<br/>1：供应商。<br/>2：分销商。<br/>注：填写其他值当做错误处理。
	_identity int64
}

// NewTaobaofenxiaodealerrequisitionordergetRequest 初始化TaobaofenxiaodealerrequisitionordergetAPIRequest对象
func NewTaobaofenxiaodealerrequisitionordergetRequest() *TaobaofenxiaodealerrequisitionordergetAPIRequest {
	return &TaobaofenxiaodealerrequisitionordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaodealerrequisitionordergetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.dealer.requisitionorder.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaodealerrequisitionordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaodealerrequisitionordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartDate is StartDate Setter
// 采购申请/经销采购单最早修改时间
func (r *TaobaofenxiaodealerrequisitionordergetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaofenxiaodealerrequisitionordergetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 采购申请/经销采购单最迟修改时间。与start_date字段的最大时间间隔不能超过30天
func (r *TaobaofenxiaodealerrequisitionordergetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaofenxiaodealerrequisitionordergetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetFields is Fields Setter
// 多个字段用&#34;,&#34;分隔。 fields 如果为空：返回所有采购申请/经销采购单对象(dealer_orders)字段。 如果不为空：返回指定采购单对象(dealer_orders)字段。 例1： dealer_order_details.product_id 表示只返回product_id 例2： dealer_order_details 表示只返回明细列表
func (r *TaobaofenxiaodealerrequisitionordergetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaofenxiaodealerrequisitionordergetAPIRequest) GetFields() string {
	return r._fields
}

// SetPageNo is PageNo Setter
// 页码（大于0的整数。无值或小于1的值按默认值1计）
func (r *TaobaofenxiaodealerrequisitionordergetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaofenxiaodealerrequisitionordergetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页条数（大于0但小于等于50的整数。无值或大于50或小于1的值按默认值50计）
func (r *TaobaofenxiaodealerrequisitionordergetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaofenxiaodealerrequisitionordergetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetOrderStatus is OrderStatus Setter
// 采购申请/经销采购单状态。&lt;br/&gt;0：全部状态。&lt;br/&gt;1：分销商提交申请，待供应商审核。&lt;br/&gt;2：供应商驳回申请，待分销商确认。&lt;br/&gt;3：供应商修改后，待分销商确认。&lt;br/&gt;4：分销商拒绝修改，待供应商再审核。&lt;br/&gt;5：审核通过下单成功，待分销商付款。&lt;br/&gt;7：付款成功，待供应商发货。&lt;br/&gt;8：供应商发货，待分销商收货。&lt;br/&gt;9：分销商收货，交易成功。&lt;br/&gt;10：采购申请/经销采购单关闭。&lt;br/&gt;&lt;br/&gt;注：无值按默认值0计，超出状态范围返回错误信息。
func (r *TaobaofenxiaodealerrequisitionordergetAPIRequest) SetOrderStatus(_orderStatus int64) error {
	r._orderStatus = _orderStatus
	r.Set("order_status", _orderStatus)
	return nil
}

// GetOrderStatus OrderStatus Getter
func (r TaobaofenxiaodealerrequisitionordergetAPIRequest) GetOrderStatus() int64 {
	return r._orderStatus
}

// SetIdentity is Identity Setter
// 查询者自己在所要查询的采购申请/经销采购单中的身份。&lt;br/&gt;1：供应商。&lt;br/&gt;2：分销商。&lt;br/&gt;注：填写其他值当做错误处理。
func (r *TaobaofenxiaodealerrequisitionordergetAPIRequest) SetIdentity(_identity int64) error {
	r._identity = _identity
	r.Set("identity", _identity)
	return nil
}

// GetIdentity Identity Getter
func (r TaobaofenxiaodealerrequisitionordergetAPIRequest) GetIdentity() int64 {
	return r._identity
}
