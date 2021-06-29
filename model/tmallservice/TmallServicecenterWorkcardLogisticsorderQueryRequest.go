package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流单信息查询 API请求
tmall.servicecenter.workcard.logisticsorder.query

物流订单信息查询API
*/
type TmallServicecenterWorkcardLogisticsorderQueryRequest struct {
    model.Params
    // 物流单号
    _logisticsOrderId   int64
}

// 初始化TmallServicecenterWorkcardLogisticsorderQueryRequest对象
func NewTmallServicecenterWorkcardLogisticsorderQueryRequest() *TmallServicecenterWorkcardLogisticsorderQueryRequest{
    return &TmallServicecenterWorkcardLogisticsorderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardLogisticsorderQueryRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.logisticsorder.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardLogisticsorderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LogisticsOrderId Setter
// 物流单号
func (r *TmallServicecenterWorkcardLogisticsorderQueryRequest) SetLogisticsOrderId(_logisticsOrderId int64) error {
    r._logisticsOrderId = _logisticsOrderId
    r.Set("logistics_order_id", _logisticsOrderId)
    return nil
}

// LogisticsOrderId Getter
func (r TmallServicecenterWorkcardLogisticsorderQueryRequest) GetLogisticsOrderId() int64 {
    return r._logisticsOrderId
}
