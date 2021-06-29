package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家回调接口 API请求
taobao.wt.trade.order.resultcallback

阿里通信定制服务，商家发货后进行调用该接口，用于自动发货并确认收货
*/
type TaobaoWtTradeOrderResultcallbackRequest struct {
    model.Params
    // 系统自动生成
    _param0   *OrderResultDto
}

// 初始化TaobaoWtTradeOrderResultcallbackRequest对象
func NewTaobaoWtTradeOrderResultcallbackRequest() *TaobaoWtTradeOrderResultcallbackRequest{
    return &TaobaoWtTradeOrderResultcallbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWtTradeOrderResultcallbackRequest) GetApiMethodName() string {
    return "taobao.wt.trade.order.resultcallback"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWtTradeOrderResultcallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 系统自动生成
func (r *TaobaoWtTradeOrderResultcallbackRequest) SetParam0(_param0 *OrderResultDto) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TaobaoWtTradeOrderResultcallbackRequest) GetParam0() *OrderResultDto {
    return r._param0
}
