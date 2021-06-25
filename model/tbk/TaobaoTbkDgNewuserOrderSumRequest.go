package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-拉新活动对应数据查询 APIRequest
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

func NewTaobaoTbkDgNewuserOrderSumRequest() *TaobaoTbkDgNewuserOrderSumRequest{
    return &TaobaoTbkDgNewuserOrderSumRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTbkDgNewuserOrderSumRequest) GetApiMethodName() string {
    return "taobao.tbk.dg.newuser.order.sum"
}

func (r TaobaoTbkDgNewuserOrderSumRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTbkDgNewuserOrderSumRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoTbkDgNewuserOrderSumRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoTbkDgNewuserOrderSumRequest) SetAdzoneId(adzoneId int64) error {
    r.adzoneId = adzoneId
    r.Set("adzone_id", adzoneId)
    return nil
}

func (r TaobaoTbkDgNewuserOrderSumRequest) GetAdzoneId() int64 {
    return r.adzoneId
}

func (r *TaobaoTbkDgNewuserOrderSumRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoTbkDgNewuserOrderSumRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoTbkDgNewuserOrderSumRequest) SetSiteId(siteId int64) error {
    r.siteId = siteId
    r.Set("site_id", siteId)
    return nil
}

func (r TaobaoTbkDgNewuserOrderSumRequest) GetSiteId() int64 {
    return r.siteId
}

func (r *TaobaoTbkDgNewuserOrderSumRequest) SetActivityId(activityId string) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r TaobaoTbkDgNewuserOrderSumRequest) GetActivityId() string {
    return r.activityId
}

func (r *TaobaoTbkDgNewuserOrderSumRequest) SetSettleMonth(settleMonth string) error {
    r.settleMonth = settleMonth
    r.Set("settle_month", settleMonth)
    return nil
}

func (r TaobaoTbkDgNewuserOrderSumRequest) GetSettleMonth() string {
    return r.settleMonth
}

