package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscorderdetailsgetAPIRequest 淘宝客-服务商-所有订单查询 API请求
// taobao.tbk.sc.order.details.get
//
// 服务商使用。可通过该接口查询推广者账号下对应的推广订单明细。
type TaobaotbkscorderdetailsgetAPIRequest struct {
	model.Params
	// 位点，除第一页之外，都需要传递；前端原样返回。
	_positionindex string
	// 订单查询结束时间，订单开始时间至订单结束时间，中间时间段日常要求不超过3个小时，但如618、双11、年货节等大促期间预估时间段不可超过20分钟，超过会提示错误，调用时请务必注意时间段的选择，以保证亲能正常调用！
	_endtime string
	// 订单查询开始时间
	_starttime string
	// 查询时间类型，1：按照订单淘客创建时间查询，2:按照订单淘客付款时间查询，3:按照订单淘客结算时间查询，4:按照订单更新时间；
	_querytype int64
	// 页大小，默认20，1~100
	_pagesize int64
	// 推广者角色类型,2:二方，3:三方，不传，表示所有角色
	_membertype int64
	// 淘客订单状态，12-付款，13-关闭，14-确认收货，3-结算成功;不传，表示所有状态
	_tkstatus int64
	// 跳转类型，当向前或者向后翻页必须提供,-1: 向前翻页,1：向后翻页
	_jumptype int64
	// 第几页，默认1，1~100
	_pageno int64
	// 筛选订单类型，1:所有订单，2:渠道订单，3:会员运营订单，默认为1
	_orderscene int64
	// member组ID
	_membergroupid int64
}

// NewTaobaotbkscorderdetailsgetRequest 初始化TaobaotbkscorderdetailsgetAPIRequest对象
func NewTaobaotbkscorderdetailsgetRequest() *TaobaotbkscorderdetailsgetAPIRequest {
	return &TaobaotbkscorderdetailsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkscorderdetailsgetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.order.details.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkscorderdetailsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkscorderdetailsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPositionindex is Positionindex Setter
// 位点，除第一页之外，都需要传递；前端原样返回。
func (r *TaobaotbkscorderdetailsgetAPIRequest) SetPositionindex(_positionindex string) error {
	r._positionindex = _positionindex
	r.Set("position_index", _positionindex)
	return nil
}

// GetPositionindex Positionindex Getter
func (r TaobaotbkscorderdetailsgetAPIRequest) GetPositionindex() string {
	return r._positionindex
}

// SetEndtime is Endtime Setter
// 订单查询结束时间，订单开始时间至订单结束时间，中间时间段日常要求不超过3个小时，但如618、双11、年货节等大促期间预估时间段不可超过20分钟，超过会提示错误，调用时请务必注意时间段的选择，以保证亲能正常调用！
func (r *TaobaotbkscorderdetailsgetAPIRequest) SetEndtime(_endtime string) error {
	r._endtime = _endtime
	r.Set("end_time", _endtime)
	return nil
}

// GetEndtime Endtime Getter
func (r TaobaotbkscorderdetailsgetAPIRequest) GetEndtime() string {
	return r._endtime
}

// SetStarttime is Starttime Setter
// 订单查询开始时间
func (r *TaobaotbkscorderdetailsgetAPIRequest) SetStarttime(_starttime string) error {
	r._starttime = _starttime
	r.Set("start_time", _starttime)
	return nil
}

// GetStarttime Starttime Getter
func (r TaobaotbkscorderdetailsgetAPIRequest) GetStarttime() string {
	return r._starttime
}

// SetQuerytype is Querytype Setter
// 查询时间类型，1：按照订单淘客创建时间查询，2:按照订单淘客付款时间查询，3:按照订单淘客结算时间查询，4:按照订单更新时间；
func (r *TaobaotbkscorderdetailsgetAPIRequest) SetQuerytype(_querytype int64) error {
	r._querytype = _querytype
	r.Set("query_type", _querytype)
	return nil
}

// GetQuerytype Querytype Getter
func (r TaobaotbkscorderdetailsgetAPIRequest) GetQuerytype() int64 {
	return r._querytype
}

// SetPagesize is Pagesize Setter
// 页大小，默认20，1~100
func (r *TaobaotbkscorderdetailsgetAPIRequest) SetPagesize(_pagesize int64) error {
	r._pagesize = _pagesize
	r.Set("page_size", _pagesize)
	return nil
}

// GetPagesize Pagesize Getter
func (r TaobaotbkscorderdetailsgetAPIRequest) GetPagesize() int64 {
	return r._pagesize
}

// SetMembertype is Membertype Setter
// 推广者角色类型,2:二方，3:三方，不传，表示所有角色
func (r *TaobaotbkscorderdetailsgetAPIRequest) SetMembertype(_membertype int64) error {
	r._membertype = _membertype
	r.Set("member_type", _membertype)
	return nil
}

// GetMembertype Membertype Getter
func (r TaobaotbkscorderdetailsgetAPIRequest) GetMembertype() int64 {
	return r._membertype
}

// SetTkstatus is Tkstatus Setter
// 淘客订单状态，12-付款，13-关闭，14-确认收货，3-结算成功;不传，表示所有状态
func (r *TaobaotbkscorderdetailsgetAPIRequest) SetTkstatus(_tkstatus int64) error {
	r._tkstatus = _tkstatus
	r.Set("tk_status", _tkstatus)
	return nil
}

// GetTkstatus Tkstatus Getter
func (r TaobaotbkscorderdetailsgetAPIRequest) GetTkstatus() int64 {
	return r._tkstatus
}

// SetJumptype is Jumptype Setter
// 跳转类型，当向前或者向后翻页必须提供,-1: 向前翻页,1：向后翻页
func (r *TaobaotbkscorderdetailsgetAPIRequest) SetJumptype(_jumptype int64) error {
	r._jumptype = _jumptype
	r.Set("jump_type", _jumptype)
	return nil
}

// GetJumptype Jumptype Getter
func (r TaobaotbkscorderdetailsgetAPIRequest) GetJumptype() int64 {
	return r._jumptype
}

// SetPageno is Pageno Setter
// 第几页，默认1，1~100
func (r *TaobaotbkscorderdetailsgetAPIRequest) SetPageno(_pageno int64) error {
	r._pageno = _pageno
	r.Set("page_no", _pageno)
	return nil
}

// GetPageno Pageno Getter
func (r TaobaotbkscorderdetailsgetAPIRequest) GetPageno() int64 {
	return r._pageno
}

// SetOrderscene is Orderscene Setter
// 筛选订单类型，1:所有订单，2:渠道订单，3:会员运营订单，默认为1
func (r *TaobaotbkscorderdetailsgetAPIRequest) SetOrderscene(_orderscene int64) error {
	r._orderscene = _orderscene
	r.Set("order_scene", _orderscene)
	return nil
}

// GetOrderscene Orderscene Getter
func (r TaobaotbkscorderdetailsgetAPIRequest) GetOrderscene() int64 {
	return r._orderscene
}

// SetMembergroupid is Membergroupid Setter
// member组ID
func (r *TaobaotbkscorderdetailsgetAPIRequest) SetMembergroupid(_membergroupid int64) error {
	r._membergroupid = _membergroupid
	r.Set("member_group_id", _membergroupid)
	return nil
}

// GetMembergroupid Membergroupid Getter
func (r TaobaotbkscorderdetailsgetAPIRequest) GetMembergroupid() int64 {
	return r._membergroupid
}
