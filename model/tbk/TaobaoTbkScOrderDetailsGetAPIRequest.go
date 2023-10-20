package tbk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScOrderDetailsGetAPIRequest 淘宝客-服务商-所有订单查询 API请求
// taobao.tbk.sc.order.details.get
//
// 服务商使用。可通过该接口查询推广者账号下对应的推广订单明细。
type TaobaoTbkScOrderDetailsGetAPIRequest struct {
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
	// 淘客订单状态，12-付款，13-关闭，14-确认收货，3-结算成功;不传，表示所有状态
	_tkStatus int64
	// 跳转类型，当向前或者向后翻页必须提供,-1: 向前翻页,1：向后翻页
	_jumpType int64
	// 第几页，默认1，1~100
	_pageNo int64
	// 筛选订单类型，1:所有订单，2:渠道订单，3:会员运营订单，默认为1
	_orderScene int64
	// member组ID
	_memberGroupId int64
}

// NewTaobaoTbkScOrderDetailsGetRequest 初始化TaobaoTbkScOrderDetailsGetAPIRequest对象
func NewTaobaoTbkScOrderDetailsGetRequest() *TaobaoTbkScOrderDetailsGetAPIRequest {
	return &TaobaoTbkScOrderDetailsGetAPIRequest{
		Params: model.NewParams(11),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkScOrderDetailsGetAPIRequest) Reset() {
	r._positionIndex = ""
	r._endTime = ""
	r._startTime = ""
	r._queryType = 0
	r._pageSize = 0
	r._memberType = 0
	r._tkStatus = 0
	r._jumpType = 0
	r._pageNo = 0
	r._orderScene = 0
	r._memberGroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScOrderDetailsGetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.order.details.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScOrderDetailsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkScOrderDetailsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPositionIndex is PositionIndex Setter
// 位点，除第一页之外，都需要传递；前端原样返回。
func (r *TaobaoTbkScOrderDetailsGetAPIRequest) SetPositionIndex(_positionIndex string) error {
	r._positionIndex = _positionIndex
	r.Set("position_index", _positionIndex)
	return nil
}

// GetPositionIndex PositionIndex Getter
func (r TaobaoTbkScOrderDetailsGetAPIRequest) GetPositionIndex() string {
	return r._positionIndex
}

// SetEndTime is EndTime Setter
// 订单查询结束时间，订单开始时间至订单结束时间，中间时间段日常要求不超过3个小时，但如618、双11、年货节等大促期间预估时间段不可超过20分钟，超过会提示错误，调用时请务必注意时间段的选择，以保证亲能正常调用！
func (r *TaobaoTbkScOrderDetailsGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoTbkScOrderDetailsGetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetStartTime is StartTime Setter
// 订单查询开始时间
func (r *TaobaoTbkScOrderDetailsGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoTbkScOrderDetailsGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetQueryType is QueryType Setter
// 查询时间类型，1：按照订单淘客创建时间查询，2:按照订单淘客付款时间查询，3:按照订单淘客结算时间查询，4:按照订单更新时间；
func (r *TaobaoTbkScOrderDetailsGetAPIRequest) SetQueryType(_queryType int64) error {
	r._queryType = _queryType
	r.Set("query_type", _queryType)
	return nil
}

// GetQueryType QueryType Getter
func (r TaobaoTbkScOrderDetailsGetAPIRequest) GetQueryType() int64 {
	return r._queryType
}

// SetPageSize is PageSize Setter
// 页大小，默认20，1~100
func (r *TaobaoTbkScOrderDetailsGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoTbkScOrderDetailsGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetMemberType is MemberType Setter
// 推广者角色类型,2:二方，3:三方，不传，表示所有角色
func (r *TaobaoTbkScOrderDetailsGetAPIRequest) SetMemberType(_memberType int64) error {
	r._memberType = _memberType
	r.Set("member_type", _memberType)
	return nil
}

// GetMemberType MemberType Getter
func (r TaobaoTbkScOrderDetailsGetAPIRequest) GetMemberType() int64 {
	return r._memberType
}

// SetTkStatus is TkStatus Setter
// 淘客订单状态，12-付款，13-关闭，14-确认收货，3-结算成功;不传，表示所有状态
func (r *TaobaoTbkScOrderDetailsGetAPIRequest) SetTkStatus(_tkStatus int64) error {
	r._tkStatus = _tkStatus
	r.Set("tk_status", _tkStatus)
	return nil
}

// GetTkStatus TkStatus Getter
func (r TaobaoTbkScOrderDetailsGetAPIRequest) GetTkStatus() int64 {
	return r._tkStatus
}

// SetJumpType is JumpType Setter
// 跳转类型，当向前或者向后翻页必须提供,-1: 向前翻页,1：向后翻页
func (r *TaobaoTbkScOrderDetailsGetAPIRequest) SetJumpType(_jumpType int64) error {
	r._jumpType = _jumpType
	r.Set("jump_type", _jumpType)
	return nil
}

// GetJumpType JumpType Getter
func (r TaobaoTbkScOrderDetailsGetAPIRequest) GetJumpType() int64 {
	return r._jumpType
}

// SetPageNo is PageNo Setter
// 第几页，默认1，1~100
func (r *TaobaoTbkScOrderDetailsGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoTbkScOrderDetailsGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetOrderScene is OrderScene Setter
// 筛选订单类型，1:所有订单，2:渠道订单，3:会员运营订单，默认为1
func (r *TaobaoTbkScOrderDetailsGetAPIRequest) SetOrderScene(_orderScene int64) error {
	r._orderScene = _orderScene
	r.Set("order_scene", _orderScene)
	return nil
}

// GetOrderScene OrderScene Getter
func (r TaobaoTbkScOrderDetailsGetAPIRequest) GetOrderScene() int64 {
	return r._orderScene
}

// SetMemberGroupId is MemberGroupId Setter
// member组ID
func (r *TaobaoTbkScOrderDetailsGetAPIRequest) SetMemberGroupId(_memberGroupId int64) error {
	r._memberGroupId = _memberGroupId
	r.Set("member_group_id", _memberGroupId)
	return nil
}

// GetMemberGroupId MemberGroupId Getter
func (r TaobaoTbkScOrderDetailsGetAPIRequest) GetMemberGroupId() int64 {
	return r._memberGroupId
}

var poolTaobaoTbkScOrderDetailsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkScOrderDetailsGetRequest()
	},
}

// GetTaobaoTbkScOrderDetailsGetRequest 从 sync.Pool 获取 TaobaoTbkScOrderDetailsGetAPIRequest
func GetTaobaoTbkScOrderDetailsGetAPIRequest() *TaobaoTbkScOrderDetailsGetAPIRequest {
	return poolTaobaoTbkScOrderDetailsGetAPIRequest.Get().(*TaobaoTbkScOrderDetailsGetAPIRequest)
}

// ReleaseTaobaoTbkScOrderDetailsGetAPIRequest 将 TaobaoTbkScOrderDetailsGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkScOrderDetailsGetAPIRequest(v *TaobaoTbkScOrderDetailsGetAPIRequest) {
	v.Reset()
	poolTaobaoTbkScOrderDetailsGetAPIRequest.Put(v)
}
