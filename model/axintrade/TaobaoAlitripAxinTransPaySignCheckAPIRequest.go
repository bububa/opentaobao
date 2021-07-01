package axintrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿信支付宝验签服务 API请求
taobao.alitrip.axin.trans.pay.sign.check

阿信支付宝验签服务
*/
type TaobaoAlitripAxinTransPaySignCheckAPIRequest struct {
    model.Params
    // 验签对象
    _axinPayCheckSignDto   *AxinPayCheckSignDto
}

// 初始化TaobaoAlitripAxinTransPaySignCheckAPIRequest对象
func NewTaobaoAlitripAxinTransPaySignCheckRequest() *TaobaoAlitripAxinTransPaySignCheckAPIRequest{
    return &TaobaoAlitripAxinTransPaySignCheckAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransPaySignCheckAPIRequest) GetApiMethodName() string {
    return "taobao.alitrip.axin.trans.pay.sign.check"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransPaySignCheckAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AxinPayCheckSignDto Setter
// 验签对象
func (r *TaobaoAlitripAxinTransPaySignCheckAPIRequest) SetAxinPayCheckSignDto(_axinPayCheckSignDto *AxinPayCheckSignDto) error {
    r._axinPayCheckSignDto = _axinPayCheckSignDto
    r.Set("axin_pay_check_sign_dto", _axinPayCheckSignDto)
    return nil
}

// AxinPayCheckSignDto Getter
func (r TaobaoAlitripAxinTransPaySignCheckAPIRequest) GetAxinPayCheckSignDto() *AxinPayCheckSignDto {
    return r._axinPayCheckSignDto
}
