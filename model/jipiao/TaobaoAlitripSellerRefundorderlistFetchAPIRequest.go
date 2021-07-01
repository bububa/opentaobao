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
type TaobaoAlitripSellerRefundorderlistFetchAPIRequest struct {
    model.Params
    // 提取数据的开始时间
    _startDate   string
    // 1：初始，2：接受，3：确认，4：失败，5：买家处理，6：卖家处理，7：等待小二回填，8：退款成功
    _status   int64
    // 提取数据的结束时间
    _endDate   string
}

// 初始化TaobaoAlitripSellerRefundorderlistFetchAPIRequest对象
func NewTaobaoAlitripSellerRefundorderlistFetchRequest() *TaobaoAlitripSellerRefundorderlistFetchAPIRequest{
    return &TaobaoAlitripSellerRefundorderlistFetchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripSellerRefundorderlistFetchAPIRequest) GetApiMethodName() string {
    return "taobao.alitrip.seller.refundorderlist.fetch"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripSellerRefundorderlistFetchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartDate Setter
// 提取数据的开始时间
func (r *TaobaoAlitripSellerRefundorderlistFetchAPIRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r TaobaoAlitripSellerRefundorderlistFetchAPIRequest) GetStartDate() string {
    return r._startDate
}
// Status Setter
// 1：初始，2：接受，3：确认，4：失败，5：买家处理，6：卖家处理，7：等待小二回填，8：退款成功
func (r *TaobaoAlitripSellerRefundorderlistFetchAPIRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoAlitripSellerRefundorderlistFetchAPIRequest) GetStatus() int64 {
    return r._status
}
// EndDate Setter
// 提取数据的结束时间
func (r *TaobaoAlitripSellerRefundorderlistFetchAPIRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r TaobaoAlitripSellerRefundorderlistFetchAPIRequest) GetEndDate() string {
    return r._endDate
}
