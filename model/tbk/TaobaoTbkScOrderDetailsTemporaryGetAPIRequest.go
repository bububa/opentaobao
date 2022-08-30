package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScOrderDetailsTemporaryGetAPIRequest 淘宝客-服务商-所有订单查询（临时接口） API请求
// taobao.tbk.sc.order.details.temporary.get
//
// 服务商使用。可通过该接口查询推广者账号下对应的推广订单明细。
type TaobaoTbkScOrderDetailsTemporaryGetAPIRequest struct {
	model.Params
	// 位点，除第一页之外，都需要传递；前端原样返回。
	_positionIndex string
	// 订单查询结束时间，订单开始时间至订单结束时间，中间时间段日常要求不超过3个小时，但如618、双11、年货节等大促期间预估时间段不可超过20分钟，超过会提示错误，调用时请务必注意时间段的选择，以保证亲能正常调用！
	_endTime string
	// 订单查询开始时间
	_startTime string
	// member组ID
	_memberGroupId string
	// 查询时间类型，1：按照订单淘客创建时间查询，2:按照订单淘客付款时间查询，3:按照订单淘客结算时间查询，4:按照订单更新时间；
	_queryType int64
	// 页大小，默认20，1~100
	_pageSize int64
	// 推广者角色类型,2:二方，3:三方，不传，表示所有角色
	_memberType int64
	// 淘客订单状态，12-付款，13-关闭，14-确认收货，3-结算成功;不传，表示所有状态
	_tkStatus int64
	// 跳转类型，当向前或者向后翻页必须提供,-1: 向前翻页,1：向后翻页
	_jumpType int64
	// 第几页，默认1，1~100
	_pageNo int64
	// 场景订单场景类型，1:常规订单，2:渠道订单，3:会员运营订单，默认为1
	_orderScene int64
}

// NewTaobaoTbkScOrderDetailsTemporaryGetRequest 初始化TaobaoTbkScOrderDetailsTemporaryGetAPIRequest对象
func NewTaobaoTbkScOrderDetailsTemporaryGetRequest() *TaobaoTbkScOrderDetailsTemporaryGetAPIRequest {
	return &TaobaoTbkScOrderDetailsTemporaryGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.order.details.temporary.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPositionIndex is PositionIndex Setter
// 位点，除第一页之外，都需要传递；前端原样返回。
func (r *TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) SetPositionIndex(_positionIndex string) error {
	r._positionIndex = _positionIndex
	r.Set("position_index", _positionIndex)
	return nil
}

// GetPositionIndex PositionIndex Getter
func (r TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) GetPositionIndex() string {
	return r._positionIndex
}

// SetEndTime is EndTime Setter
// 订单查询结束时间，订单开始时间至订单结束时间，中间时间段日常要求不超过3个小时，但如618、双11、年货节等大促期间预估时间段不可超过20分钟，超过会提示错误，调用时请务必注意时间段的选择，以保证亲能正常调用！
func (r *TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetStartTime is StartTime Setter
// 订单查询开始时间
func (r *TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetMemberGroupId is MemberGroupId Setter
// member组ID
func (r *TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) SetMemberGroupId(_memberGroupId string) error {
	r._memberGroupId = _memberGroupId
	r.Set("member_group_id", _memberGroupId)
	return nil
}

// GetMemberGroupId MemberGroupId Getter
func (r TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) GetMemberGroupId() string {
	return r._memberGroupId
}

// SetQueryType is QueryType Setter
// 查询时间类型，1：按照订单淘客创建时间查询，2:按照订单淘客付款时间查询，3:按照订单淘客结算时间查询，4:按照订单更新时间；
func (r *TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) SetQueryType(_queryType int64) error {
	r._queryType = _queryType
	r.Set("query_type", _queryType)
	return nil
}

// GetQueryType QueryType Getter
func (r TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) GetQueryType() int64 {
	return r._queryType
}

// SetPageSize is PageSize Setter
// 页大小，默认20，1~100
func (r *TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetMemberType is MemberType Setter
// 推广者角色类型,2:二方，3:三方，不传，表示所有角色
func (r *TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) SetMemberType(_memberType int64) error {
	r._memberType = _memberType
	r.Set("member_type", _memberType)
	return nil
}

// GetMemberType MemberType Getter
func (r TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) GetMemberType() int64 {
	return r._memberType
}

// SetTkStatus is TkStatus Setter
// 淘客订单状态，12-付款，13-关闭，14-确认收货，3-结算成功;不传，表示所有状态
func (r *TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) SetTkStatus(_tkStatus int64) error {
	r._tkStatus = _tkStatus
	r.Set("tk_status", _tkStatus)
	return nil
}

// GetTkStatus TkStatus Getter
func (r TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) GetTkStatus() int64 {
	return r._tkStatus
}

// SetJumpType is JumpType Setter
// 跳转类型，当向前或者向后翻页必须提供,-1: 向前翻页,1：向后翻页
func (r *TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) SetJumpType(_jumpType int64) error {
	r._jumpType = _jumpType
	r.Set("jump_type", _jumpType)
	return nil
}

// GetJumpType JumpType Getter
func (r TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) GetJumpType() int64 {
	return r._jumpType
}

// SetPageNo is PageNo Setter
// 第几页，默认1，1~100
func (r *TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetOrderScene is OrderScene Setter
// 场景订单场景类型，1:常规订单，2:渠道订单，3:会员运营订单，默认为1
func (r *TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) SetOrderScene(_orderScene int64) error {
	r._orderScene = _orderScene
	r.Set("order_scene", _orderScene)
	return nil
}

// GetOrderScene OrderScene Getter
func (r TaobaoTbkScOrderDetailsTemporaryGetAPIRequest) GetOrderScene() int64 {
	return r._orderScene
}
