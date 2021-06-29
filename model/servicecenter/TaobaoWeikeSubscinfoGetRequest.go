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
    _sellerName   string
    // 时间范围开始时间
    _startTime   string
    // 时间范围结束时间
    _endTime   string
    // 页码
    _pageNum   int64
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
func (r *TaobaoWeikeSubscinfoGetRequest) SetSellerName(_sellerName string) error {
    r._sellerName = _sellerName
    r.Set("seller_name", _sellerName)
    return nil
}

// SellerName Getter
func (r TaobaoWeikeSubscinfoGetRequest) GetSellerName() string {
    return r._sellerName
}
// StartTime Setter
// 时间范围开始时间
func (r *TaobaoWeikeSubscinfoGetRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoWeikeSubscinfoGetRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 时间范围结束时间
func (r *TaobaoWeikeSubscinfoGetRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoWeikeSubscinfoGetRequest) GetEndTime() string {
    return r._endTime
}
// PageNum Setter
// 页码
func (r *TaobaoWeikeSubscinfoGetRequest) SetPageNum(_pageNum int64) error {
    r._pageNum = _pageNum
    r.Set("page_num", _pageNum)
    return nil
}

// PageNum Getter
func (r TaobaoWeikeSubscinfoGetRequest) GetPageNum() int64 {
    return r._pageNum
}
