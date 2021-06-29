package jipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【机票代理商】——退票订单列表提取 API请求
taobao.alitrip.seller.refundorderlist.fetch

代理商纬度退票订单列表提取
*/
type TaobaoAlitripSellerRefundorderlistFetchRequest struct {
    model.Params
    // 提取数据的开始时间
    _startDate   string
    // 1：初始，2：接受，3：确认，4：失败，5：买家处理，6：卖家处理，7：等待小二回填，8：退款成功
    _status   int64
    // 提取数据的结束时间
    _endDate   string
}

// 初始化TaobaoAlitripSellerRefundorderlistFetchRequest对象
func NewTaobaoAlitripSellerRefundorderlistFetchRequest() *TaobaoAlitripSellerRefundorderlistFetchRequest{
    return &TaobaoAlitripSellerRefundorderlistFetchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripSellerRefundorderlistFetchRequest) GetApiMethodName() string {
    return "taobao.alitrip.seller.refundorderlist.fetch"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripSellerRefundorderlistFetchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartDate Setter
// 提取数据的开始时间
func (r *TaobaoAlitripSellerRefundorderlistFetchRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r TaobaoAlitripSellerRefundorderlistFetchRequest) GetStartDate() string {
    return r._startDate
}
// Status Setter
// 1：初始，2：接受，3：确认，4：失败，5：买家处理，6：卖家处理，7：等待小二回填，8：退款成功
func (r *TaobaoAlitripSellerRefundorderlistFetchRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoAlitripSellerRefundorderlistFetchRequest) GetStatus() int64 {
    return r._status
}
// EndDate Setter
// 提取数据的结束时间
func (r *TaobaoAlitripSellerRefundorderlistFetchRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TaobaoAlitripSellerRefundorderlistFetchRequest) GetEndDate() string {
    return r._endDate
}
