package elife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌惠卡券核销 API请求
taobao.elife.lifecard.consume

用户线上购买生活汇品牌惠虚拟消费卡，线下购物时，商家码枪核销，涉及用户虚拟卡余额扣减操作
*/
type TaobaoElifeLifecardConsumeAPIRequest struct {
    model.Params
    // 交易请求参数
    _consumeRequest   *ConsumeRequest
}

// 初始化TaobaoElifeLifecardConsumeAPIRequest对象
func NewTaobaoElifeLifecardConsumeRequest() *TaobaoElifeLifecardConsumeAPIRequest{
    return &TaobaoElifeLifecardConsumeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoElifeLifecardConsumeAPIRequest) GetApiMethodName() string {
    return "taobao.elife.lifecard.consume"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoElifeLifecardConsumeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ConsumeRequest Setter
// 交易请求参数
func (r *TaobaoElifeLifecardConsumeAPIRequest) SetConsumeRequest(_consumeRequest *ConsumeRequest) error {
    r._consumeRequest = _consumeRequest
    r.Set("consume_request", _consumeRequest)
    return nil
}

// ConsumeRequest Getter
func (r TaobaoElifeLifecardConsumeAPIRequest) GetConsumeRequest() *ConsumeRequest {
    return r._consumeRequest
}
