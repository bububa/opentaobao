package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-新用户订单明细查询 API请求
taobao.tbk.dg.newuser.order.get

拉新API
*/
type TaobaoTbkDgNewuserOrderGetRequest struct {
    model.Params
    // 页大小，默认20，1~100
    _pageSize   int64
    // mm_xxx_xxx_xxx的第三位
    _adzoneId   int64
    // 页码，默认1
    _pageNo   int64
    // 开始时间，当活动为淘宝活动，表示最早注册时间；当活动为支付宝活动，表示最早绑定时间；当活动为天猫活动，表示最早领取红包时间
    _startTime   string
    // 结束时间，当活动为淘宝活动，表示最晚结束时间；当活动为支付宝活动，表示最晚绑定时间；当活动为天猫活动，表示最晚领取红包的时间
    _endTime   string
    // 活动id， 活动名称与活动ID列表，请参见https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8599277
    _activityId   string
}

// 初始化TaobaoTbkDgNewuserOrderGetRequest对象
func NewTaobaoTbkDgNewuserOrderGetRequest() *TaobaoTbkDgNewuserOrderGetRequest{
    return &TaobaoTbkDgNewuserOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTbkDgNewuserOrderGetRequest) GetApiMethodName() string {
    return "taobao.tbk.dg.newuser.order.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTbkDgNewuserOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageSize Setter
// 页大小，默认20，1~100
func (r *TaobaoTbkDgNewuserOrderGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTbkDgNewuserOrderGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// AdzoneId Setter
// mm_xxx_xxx_xxx的第三位
func (r *TaobaoTbkDgNewuserOrderGetRequest) SetAdzoneId(_adzoneId int64) error {
    r._adzoneId = _adzoneId
    r.Set("adzone_id", _adzoneId)
    return nil
}

// AdzoneId Getter
func (r TaobaoTbkDgNewuserOrderGetRequest) GetAdzoneId() int64 {
    return r._adzoneId
}
// PageNo Setter
// 页码，默认1
func (r *TaobaoTbkDgNewuserOrderGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoTbkDgNewuserOrderGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// StartTime Setter
// 开始时间，当活动为淘宝活动，表示最早注册时间；当活动为支付宝活动，表示最早绑定时间；当活动为天猫活动，表示最早领取红包时间
func (r *TaobaoTbkDgNewuserOrderGetRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoTbkDgNewuserOrderGetRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 结束时间，当活动为淘宝活动，表示最晚结束时间；当活动为支付宝活动，表示最晚绑定时间；当活动为天猫活动，表示最晚领取红包的时间
func (r *TaobaoTbkDgNewuserOrderGetRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoTbkDgNewuserOrderGetRequest) GetEndTime() string {
    return r._endTime
}
// ActivityId Setter
// 活动id， 活动名称与活动ID列表，请参见https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8599277
func (r *TaobaoTbkDgNewuserOrderGetRequest) SetActivityId(_activityId string) error {
    r._activityId = _activityId
    r.Set("activity_id", _activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoTbkDgNewuserOrderGetRequest) GetActivityId() string {
    return r._activityId
}
