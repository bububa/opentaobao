package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店自送对码进行核销 API请求
taobao.omniorder.dtd.consume

该接口根据传入的码及订单信息，如果码与订单一致，则对门店自送服务进行核销。
*/
type TaobaoOmniorderDtdConsumeRequest struct {
    model.Params
    // 核销信息
    _paramDoor2doorConsumeRequest   *Door2doorConsumeRequest
}

// 初始化TaobaoOmniorderDtdConsumeRequest对象
func NewTaobaoOmniorderDtdConsumeRequest() *TaobaoOmniorderDtdConsumeRequest{
    return &TaobaoOmniorderDtdConsumeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderDtdConsumeRequest) GetApiMethodName() string {
    return "taobao.omniorder.dtd.consume"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderDtdConsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamDoor2doorConsumeRequest Setter
// 核销信息
func (r *TaobaoOmniorderDtdConsumeRequest) SetParamDoor2doorConsumeRequest(_paramDoor2doorConsumeRequest *Door2doorConsumeRequest) error {
    r._paramDoor2doorConsumeRequest = _paramDoor2doorConsumeRequest
    r.Set("param_door2door_consume_request", _paramDoor2doorConsumeRequest)
    return nil
}

// ParamDoor2doorConsumeRequest Getter
func (r TaobaoOmniorderDtdConsumeRequest) GetParamDoor2doorConsumeRequest() *Door2doorConsumeRequest {
    return r._paramDoor2doorConsumeRequest
}
