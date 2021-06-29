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
    _pageSize   int64
    // mm_xxx_xxx_xxx的第三位
    _adzoneId   int64
    // 页码，默认1
    _pageNo   int64
    // mm_xxx_xxx_xxx的第二位
    _siteId   int64
    // 活动id， 活动名称与活动ID列表，请参见https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8599277
    _activityId   string
    // 结算月份
    _settleMonth   string
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
func (r *TaobaoTbkDgNewuserOrderSumRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTbkDgNewuserOrderSumRequest) GetPageSize() int64 {
    return r._pageSize
}
// AdzoneId Setter
// mm_xxx_xxx_xxx的第三位
func (r *TaobaoTbkDgNewuserOrderSumRequest) SetAdzoneId(_adzoneId int64) error {
    r._adzoneId = _adzoneId
    r.Set("adzone_id", _adzoneId)
    return nil
}

// AdzoneId Getter
func (r TaobaoTbkDgNewuserOrderSumRequest) GetAdzoneId() int64 {
    return r._adzoneId
}
// PageNo Setter
// 页码，默认1
func (r *TaobaoTbkDgNewuserOrderSumRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoTbkDgNewuserOrderSumRequest) GetPageNo() int64 {
    return r._pageNo
}
// SiteId Setter
// mm_xxx_xxx_xxx的第二位
func (r *TaobaoTbkDgNewuserOrderSumRequest) SetSiteId(_siteId int64) error {
    r._siteId = _siteId
    r.Set("site_id", _siteId)
    return nil
}

// SiteId Getter
func (r TaobaoTbkDgNewuserOrderSumRequest) GetSiteId() int64 {
    return r._siteId
}
// ActivityId Setter
// 活动id， 活动名称与活动ID列表，请参见https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8599277
func (r *TaobaoTbkDgNewuserOrderSumRequest) SetActivityId(_activityId string) error {
    r._activityId = _activityId
    r.Set("activity_id", _activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoTbkDgNewuserOrderSumRequest) GetActivityId() string {
    return r._activityId
}
// SettleMonth Setter
// 结算月份
func (r *TaobaoTbkDgNewuserOrderSumRequest) SetSettleMonth(_settleMonth string) error {
    r._settleMonth = _settleMonth
    r.Set("settle_month", _settleMonth)
    return nil
}

// SettleMonth Getter
func (r TaobaoTbkDgNewuserOrderSumRequest) GetSettleMonth() string {
    return r._settleMonth
}
