package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgNewuserOrderGetAPIRequest 淘宝客-推广者-新用户订单明细查询 API请求
// taobao.tbk.dg.newuser.order.get
//
// 拉新API
type TaobaoTbkDgNewuserOrderGetAPIRequest struct {
	model.Params
	// 开始时间，当活动为淘宝活动，表示最早注册时间；当活动为支付宝活动，表示最早绑定时间；当活动为天猫活动，表示最早领取红包时间
	_startTime string
	// 结束时间，当活动为淘宝活动，表示最晚结束时间；当活动为支付宝活动，表示最晚绑定时间；当活动为天猫活动，表示最晚领取红包的时间
	_endTime string
	// 活动id， 活动名称与活动ID列表（该字段已废弃）
	_activityId string
	// 页大小，默认20，1~100
	_pageSize int64
	// 页码，默认1
	_pageNo int64
	// mm_xxx_xxx_xxx的第三位
	_adzoneId int64
}

// NewTaobaoTbkDgNewuserOrderGetRequest 初始化TaobaoTbkDgNewuserOrderGetAPIRequest对象
func NewTaobaoTbkDgNewuserOrderGetRequest() *TaobaoTbkDgNewuserOrderGetAPIRequest {
	return &TaobaoTbkDgNewuserOrderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkDgNewuserOrderGetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.newuser.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkDgNewuserOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkDgNewuserOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartTime is StartTime Setter
// 开始时间，当活动为淘宝活动，表示最早注册时间；当活动为支付宝活动，表示最早绑定时间；当活动为天猫活动，表示最早领取红包时间
func (r *TaobaoTbkDgNewuserOrderGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoTbkDgNewuserOrderGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间，当活动为淘宝活动，表示最晚结束时间；当活动为支付宝活动，表示最晚绑定时间；当活动为天猫活动，表示最晚领取红包的时间
func (r *TaobaoTbkDgNewuserOrderGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoTbkDgNewuserOrderGetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetActivityId is ActivityId Setter
// 活动id， 活动名称与活动ID列表（该字段已废弃）
func (r *TaobaoTbkDgNewuserOrderGetAPIRequest) SetActivityId(_activityId string) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoTbkDgNewuserOrderGetAPIRequest) GetActivityId() string {
	return r._activityId
}

// SetPageSize is PageSize Setter
// 页大小，默认20，1~100
func (r *TaobaoTbkDgNewuserOrderGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoTbkDgNewuserOrderGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 页码，默认1
func (r *TaobaoTbkDgNewuserOrderGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoTbkDgNewuserOrderGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetAdzoneId is AdzoneId Setter
// mm_xxx_xxx_xxx的第三位
func (r *TaobaoTbkDgNewuserOrderGetAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaoTbkDgNewuserOrderGetAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}
