package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
需求订单查询接口 API请求
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

// 初始化TaobaoWeikeSubscinfoGetRequest对象
func NewTaobaoWeikeSubscinfoGetRequest() *TaobaoWeikeSubscinfoGetRequest{
    return &TaobaoWeikeSubscinfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWeikeSubscinfoGetRequest) GetApiMethodName() string {
    return "taobao.weike.subscinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWeikeSubscinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SellerName Setter
// 商家旺旺名称
func (r *TaobaoWeikeSubscinfoGetRequest) SetSellerName(sellerName string) error {
    r.sellerName = sellerName
    r.Set("seller_name", sellerName)
    return nil
}

// SellerName Getter
func (r TaobaoWeikeSubscinfoGetRequest) GetSellerName() string {
    return r.sellerName
}
// StartTime Setter
// 时间范围开始时间
func (r *TaobaoWeikeSubscinfoGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoWeikeSubscinfoGetRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 时间范围结束时间
func (r *TaobaoWeikeSubscinfoGetRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoWeikeSubscinfoGetRequest) GetEndTime() string {
    return r.endTime
}
// PageNum Setter
// 页码
func (r *TaobaoWeikeSubscinfoGetRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

// PageNum Getter
func (r TaobaoWeikeSubscinfoGetRequest) GetPageNum() int64 {
    return r.pageNum
}
