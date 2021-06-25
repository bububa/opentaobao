package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
需求订单查询接口 APIRequest
taobao.weike.subscinfo.get

需求订单查询接口
*/
type TaobaoWeikeSubscinfoGetRequest struct {
    model.Params

    // 商家旺旺名称
    sellerName   string 

    // 时间范围开始时间
    startTime   string 

    // 时间范围结束时间
    endTime   string 

    // 页码
    pageNum   int64 

}

func NewTaobaoWeikeSubscinfoGetRequest() *TaobaoWeikeSubscinfoGetRequest{
    return &TaobaoWeikeSubscinfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWeikeSubscinfoGetRequest) GetApiMethodName() string {
    return "taobao.weike.subscinfo.get"
}

func (r TaobaoWeikeSubscinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWeikeSubscinfoGetRequest) SetSellerName(sellerName string) error {
    r.sellerName = sellerName
    r.Set("seller_name", sellerName)
    return nil
}

func (r TaobaoWeikeSubscinfoGetRequest) GetSellerName() string {
    return r.sellerName
}

func (r *TaobaoWeikeSubscinfoGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoWeikeSubscinfoGetRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoWeikeSubscinfoGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoWeikeSubscinfoGetRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoWeikeSubscinfoGetRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

func (r TaobaoWeikeSubscinfoGetRequest) GetPageNum() int64 {
    return r.pageNum
}

