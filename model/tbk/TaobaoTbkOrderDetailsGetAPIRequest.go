package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkOrderDetailsGetAPIRequest 淘宝客-推广者-所有订单查询 API请求
// taobao.tbk.order.details.get
//
// 淘宝客订单查询
type TaobaoTbkOrderDetailsGetAPIRequest struct {
	model.Params
	// 位点，除第一页之外，都需要传递；前端原样返回。
	_positionIndex string
	// 订单查询结束时间，订单开始时间至订单结束时间，中间时间段日常要求不超过3个小时，但如618、双11、年货节等大促期间预估时间段不可超过20分钟，超过会提示错误，调用时请务必注意时间段的选择，以保证亲能正常调用！
	_endTime string
	// 订单查询开始时间
	_startTime string
	// 查询时间类型，1：按照订单淘客创建时间查询，2:按照订单淘客付款时间查询，3:按照订单淘客结算时间查询，4:按照订单更新时间；
	_queryType int64
	// 页大小，默认20，1~100
	_pageSize int64
	// 推广者角色类型,2:二方，3:三方，不传，表示所有角色
	_memberType int64
	// 淘客订单状态，11-拍下未付款，12-付款，13-关闭，14-确认收货，3-结算成功;不传，表示所有状态
	_tkStatus int64
	// 跳转类型，当向前或者向后翻页必须提供,-1: 向前翻页,1：向后翻页
	_jumpType int64
	// 第几页，默认1，1~100
	_pageNo int64
	// 场景订单场景类型，1:常规订单，2:渠道订单，3:会员运营订单，默认为1
	_orderScene int64
}

// NewTaobaoTbkOrderDetailsGetRequest 初始化TaobaoTbkOrderDetailsGetAPIRequest对象
func NewTaobaoTbkOrderDetailsGetRequest() *TaobaoTbkOrderDetailsGetAPIRequest {
	return &TaobaoTbkOrderDetailsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkOrderDetailsGetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.order.details.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkOrderDetailsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkOrderDetailsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPositionIndex is PositionIndex Setter
// 位点，除第一页之外，都需要传递；前端原样返回。
func (r *TaobaoTbkOrderDetailsGetAPIRequest) SetPositionIndex(_positionIndex string) error {
	r._positionIndex = _positionIndex
	r.Set("position_index", _positionIndex)
	return nil
}

// GetPositionIndex PositionIndex Getter
func (r TaobaoTbkOrderDetailsGetAPIRequest) GetPositionIndex() string {
	return r._positionIndex
}

// SetEndTime is EndTime Setter
// 订单查询结束时间，订单开始时间至订单结束时间，中间时间段日常要求不超过3个小时，但如618、双11、年货节等大促期间预估时间段不可超过20分钟，超过会提示错误，调用时请务必注意时间段的选择，以保证亲能正常调用！
func (r *TaobaoTbkOrderDetailsGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoTbkOrderDetailsGetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetStartTime is StartTime Setter
// 订单查询开始时间
func (r *TaobaoTbkOrderDetailsGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoTbkOrderDetailsGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetQueryType is QueryType Setter
// 查询时间类型，1：按照订单淘客创建时间查询，2:按照订单淘客付款时间查询，3:按照订单淘客结算时间查询，4:按照订单更新时间；
func (r *TaobaoTbkOrderDetailsGetAPIRequest) SetQueryType(_queryType int64) error {
	r._queryType = _queryType
	r.Set("query_type", _queryType)
	return nil
}

// GetQueryType QueryType Getter
func (r TaobaoTbkOrderDetailsGetAPIRequest) GetQueryType() int64 {
	return r._queryType
}

// SetPageSize is PageSize Setter
// 页大小，默认20，1~100
func (r *TaobaoTbkOrderDetailsGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoTbkOrderDetailsGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetMemberType is MemberType Setter
// 推广者角色类型,2:二方，3:三方，不传，表示所有角色
func (r *TaobaoTbkOrderDetailsGetAPIRequest) SetMemberType(_memberType int64) error {
	r._memberType = _memberType
	r.Set("member_type", _memberType)
	return nil
}

// GetMemberType MemberType Getter
func (r TaobaoTbkOrderDetailsGetAPIRequest) GetMemberType() int64 {
	return r._memberType
}

// SetTkStatus is TkStatus Setter
// 淘客订单状态，11-拍下未付款，12-付款，13-关闭，14-确认收货，3-结算成功;不传，表示所有状态
func (r *TaobaoTbkOrderDetailsGetAPIRequest) SetTkStatus(_tkStatus int64) error {
	r._tkStatus = _tkStatus
	r.Set("tk_status", _tkStatus)
	return nil
}

// GetTkStatus TkStatus Getter
func (r TaobaoTbkOrderDetailsGetAPIRequest) GetTkStatus() int64 {
	return r._tkStatus
}

// SetJumpType is JumpType Setter
// 跳转类型，当向前或者向后翻页必须提供,-1: 向前翻页,1：向后翻页
func (r *TaobaoTbkOrderDetailsGetAPIRequest) SetJumpType(_jumpType int64) error {
	r._jumpType = _jumpType
	r.Set("jump_type", _jumpType)
	return nil
}

// GetJumpType JumpType Getter
func (r TaobaoTbkOrderDetailsGetAPIRequest) GetJumpType() int64 {
	return r._jumpType
}

// SetPageNo is PageNo Setter
// 第几页，默认1，1~100
func (r *TaobaoTbkOrderDetailsGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoTbkOrderDetailsGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetOrderScene is OrderScene Setter
// 场景订单场景类型，1:常规订单，2:渠道订单，3:会员运营订单，默认为1
func (r *TaobaoTbkOrderDetailsGetAPIRequest) SetOrderScene(_orderScene int64) error {
	r._orderScene = _orderScene
	r.Set("order_scene", _orderScene)
	return nil
}

// GetOrderScene OrderScene Getter
func (r TaobaoTbkOrderDetailsGetAPIRequest) GetOrderScene() int64 {
	return r._orderScene
}
