package traveltrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单服务信息写入接口 API请求
alitrip.travel.trade.serviceinfo.write

订单服务信息写入接口
*/
type AlitripTravelTradeServiceinfoWriteAPIRequest struct {
    model.Params
    // 根据模版要求传递对应的订单服务信息，当涉及多值时，用英文分号隔开";"，目前api暂时不支持文件上传，只支持链接；链接必须外网能访问
    _tipValue   string
    // 对应的订单信息
    _subTcOrderId   int64
}

// 初始化AlitripTravelTradeServiceinfoWriteAPIRequest对象
func NewAlitripTravelTradeServiceinfoWriteRequest() *AlitripTravelTradeServiceinfoWriteAPIRequest{
    return &AlitripTravelTradeServiceinfoWriteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTravelTradeServiceinfoWriteAPIRequest) GetApiMethodName() string {
    return "alitrip.travel.trade.serviceinfo.write"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTravelTradeServiceinfoWriteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TipValue Setter
// 根据模版要求传递对应的订单服务信息，当涉及多值时，用英文分号隔开";"，目前api暂时不支持文件上传，只支持链接；链接必须外网能访问
func (r *AlitripTravelTradeServiceinfoWriteAPIRequest) SetTipValue(_tipValue string) error {
    r._tipValue = _tipValue
    r.Set("tip_value", _tipValue)
    return nil
}

// TipValue Getter
func (r AlitripTravelTradeServiceinfoWriteAPIRequest) GetTipValue() string {
    return r._tipValue
}
// SubTcOrderId Setter
// 对应的订单信息
func (r *AlitripTravelTradeServiceinfoWriteAPIRequest) SetSubTcOrderId(_subTcOrderId int64) error {
    r._subTcOrderId = _subTcOrderId
    r.Set("sub_tc_order_id", _subTcOrderId)
    return nil
}

// SubTcOrderId Getter
func (r AlitripTravelTradeServiceinfoWriteAPIRequest) GetSubTcOrderId() int64 {
    return r._subTcOrderId
}
