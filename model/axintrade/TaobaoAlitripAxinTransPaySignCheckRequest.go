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
type TaobaoAlitripAxinTransPaySignCheckRequest struct {
    model.Params
    // 验签对象
    axinPayCheckSignDto   *AxinPayCheckSignDto
}

// 初始化TaobaoAlitripAxinTransPaySignCheckRequest对象
func NewTaobaoAlitripAxinTransPaySignCheckRequest() *TaobaoAlitripAxinTransPaySignCheckRequest{
    return &TaobaoAlitripAxinTransPaySignCheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransPaySignCheckRequest) GetApiMethodName() string {
    return "taobao.alitrip.axin.trans.pay.sign.check"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransPaySignCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AxinPayCheckSignDto Setter
// 验签对象
func (r *TaobaoAlitripAxinTransPaySignCheckRequest) SetAxinPayCheckSignDto(axinPayCheckSignDto *AxinPayCheckSignDto) error {
    r.axinPayCheckSignDto = axinPayCheckSignDto
    r.Set("axin_pay_check_sign_dto", axinPayCheckSignDto)
    return nil
}

// AxinPayCheckSignDto Getter
func (r TaobaoAlitripAxinTransPaySignCheckRequest) GetAxinPayCheckSignDto() *AxinPayCheckSignDto {
    return r.axinPayCheckSignDto
}
