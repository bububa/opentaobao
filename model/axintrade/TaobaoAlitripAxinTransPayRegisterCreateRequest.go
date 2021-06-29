package axintrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交支付服务开通 API请求
taobao.alitrip.axin.trans.pay.register.create

阿信供销平台-提交支付服务开通接口
*/
type TaobaoAlitripAxinTransPayRegisterCreateRequest struct {
    model.Params
    // 提交支付服务开通接口入参
    _createDTO   *AxinPayRegisterCreateDto
}

// 初始化TaobaoAlitripAxinTransPayRegisterCreateRequest对象
func NewTaobaoAlitripAxinTransPayRegisterCreateRequest() *TaobaoAlitripAxinTransPayRegisterCreateRequest{
    return &TaobaoAlitripAxinTransPayRegisterCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransPayRegisterCreateRequest) GetApiMethodName() string {
    return "taobao.alitrip.axin.trans.pay.register.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransPayRegisterCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CreateDTO Setter
// 提交支付服务开通接口入参
func (r *TaobaoAlitripAxinTransPayRegisterCreateRequest) SetCreateDTO(_createDTO *AxinPayRegisterCreateDto) error {
    r._createDTO = _createDTO
    r.Set("create_d_t_o", _createDTO)
    return nil
}

// CreateDTO Getter
func (r TaobaoAlitripAxinTransPayRegisterCreateRequest) GetCreateDTO() *AxinPayRegisterCreateDto {
    return r._createDTO
}
