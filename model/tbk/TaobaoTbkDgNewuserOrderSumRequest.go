package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-拉新活动对应数据查询 API请求
taobao.tbk.dg.newuser.order.sum

拉新活动汇总API
*/
type TaobaoTbkDgNewuserOrderSumRequest struct {
    model.Params
    // 页大小，默认20，1~100
    pageSize   int64
    // mm_xxx_xxx_xxx的第三位
    adzoneId   int64
    // 页码，默认1
    pageNo   int64
    // mm_xxx_xxx_xxx的第二位
    siteId   int64
    // 活动id， 活动名称与活动ID列表，请参见https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8599277
    activityId   string
    // 结算月份
    settleMonth   string
}

// 初始化TaobaoTbkDgNewuserOrderSumRequest对象
func NewTaobaoTbkDgNewuserOrderSumRequest() *TaobaoTbkDgNewuserOrderSumRequest{
    return &TaobaoTbkDgNewuserOrderSumRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTbkDgNewuserOrderSumRequest) GetApiMethodName() string {
    return "taobao.tbk.dg.newuser.order.sum"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTbkDgNewuserOrderSumRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageSize Setter
// 页大小，默认20，1~100
func (r *TaobaoTbkDgNewuserOrderSumRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTbkDgNewuserOrderSumRequest) GetPageSize() int64 {
    return r.pageSize
}
// AdzoneId Setter
// mm_xxx_xxx_xxx的第三位
func (r *TaobaoTbkDgNewuserOrderSumRequest) SetAdzoneId(adzoneId int64) error {
    r.adzoneId = adzoneId
    r.Set("adzone_id", adzoneId)
    return nil
}

// AdzoneId Getter
func (r TaobaoTbkDgNewuserOrderSumRequest) GetAdzoneId() int64 {
    return r.adzoneId
}
// PageNo Setter
// 页码，默认1
func (r *TaobaoTbkDgNewuserOrderSumRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoTbkDgNewuserOrderSumRequest) GetPageNo() int64 {
    return r.pageNo
}
// SiteId Setter
// mm_xxx_xxx_xxx的第二位
func (r *TaobaoTbkDgNewuserOrderSumRequest) SetSiteId(siteId int64) error {
    r.siteId = siteId
    r.Set("site_id", siteId)
    return nil
}

// SiteId Getter
func (r TaobaoTbkDgNewuserOrderSumRequest) GetSiteId() int64 {
    return r.siteId
}
// ActivityId Setter
// 活动id， 活动名称与活动ID列表，请参见https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8599277
func (r *TaobaoTbkDgNewuserOrderSumRequest) SetActivityId(activityId string) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoTbkDgNewuserOrderSumRequest) GetActivityId() string {
    return r.activityId
}
// SettleMonth Setter
// 结算月份
func (r *TaobaoTbkDgNewuserOrderSumRequest) SetSettleMonth(settleMonth string) error {
    r.settleMonth = settleMonth
    r.Set("settle_month", settleMonth)
    return nil
}

// SettleMonth Getter
func (r TaobaoTbkDgNewuserOrderSumRequest) GetSettleMonth() string {
    return r.settleMonth
}
