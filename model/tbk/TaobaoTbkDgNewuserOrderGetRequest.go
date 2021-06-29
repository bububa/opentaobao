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
    pageSize   int64
    // mm_xxx_xxx_xxx的第三位
    adzoneId   int64
    // 页码，默认1
    pageNo   int64
    // 开始时间，当活动为淘宝活动，表示最早注册时间；当活动为支付宝活动，表示最早绑定时间；当活动为天猫活动，表示最早领取红包时间
    startTime   string
    // 结束时间，当活动为淘宝活动，表示最晚结束时间；当活动为支付宝活动，表示最晚绑定时间；当活动为天猫活动，表示最晚领取红包的时间
    endTime   string
    // 活动id， 活动名称与活动ID列表，请参见https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8599277
    activityId   string
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
func (r *TaobaoTbkDgNewuserOrderGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTbkDgNewuserOrderGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// AdzoneId Setter
// mm_xxx_xxx_xxx的第三位
func (r *TaobaoTbkDgNewuserOrderGetRequest) SetAdzoneId(adzoneId int64) error {
    r.adzoneId = adzoneId
    r.Set("adzone_id", adzoneId)
    return nil
}

// AdzoneId Getter
func (r TaobaoTbkDgNewuserOrderGetRequest) GetAdzoneId() int64 {
    return r.adzoneId
}
// PageNo Setter
// 页码，默认1
func (r *TaobaoTbkDgNewuserOrderGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoTbkDgNewuserOrderGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// StartTime Setter
// 开始时间，当活动为淘宝活动，表示最早注册时间；当活动为支付宝活动，表示最早绑定时间；当活动为天猫活动，表示最早领取红包时间
func (r *TaobaoTbkDgNewuserOrderGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoTbkDgNewuserOrderGetRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 结束时间，当活动为淘宝活动，表示最晚结束时间；当活动为支付宝活动，表示最晚绑定时间；当活动为天猫活动，表示最晚领取红包的时间
func (r *TaobaoTbkDgNewuserOrderGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoTbkDgNewuserOrderGetRequest) GetEndTime() string {
    return r.endTime
}
// ActivityId Setter
// 活动id， 活动名称与活动ID列表，请参见https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8599277
func (r *TaobaoTbkDgNewuserOrderGetRequest) SetActivityId(activityId string) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoTbkDgNewuserOrderGetRequest) GetActivityId() string {
    return r.activityId
}
