package axintrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建资金单接口 API请求
taobao.alitrip.axin.trans.fund.add

创建资金单
*/
type TaobaoAlitripAxinTransFundAddRequest struct {
    model.Params
    // 创建资金单接口入参
    _axinFundCreateDTO   *AxinFundCreateDto
}

// 初始化TaobaoAlitripAxinTransFundAddRequest对象
func NewTaobaoAlitripAxinTransFundAddRequest() *TaobaoAlitripAxinTransFundAddRequest{
    return &TaobaoAlitripAxinTransFundAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransFundAddRequest) GetApiMethodName() string {
    return "taobao.alitrip.axin.trans.fund.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransFundAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AxinFundCreateDTO Setter
// 创建资金单接口入参
func (r *TaobaoAlitripAxinTransFundAddRequest) SetAxinFundCreateDTO(_axinFundCreateDTO *AxinFundCreateDto) error {
    r._axinFundCreateDTO = _axinFundCreateDTO
    r.Set("axin_fund_create_d_t_o", _axinFundCreateDTO)
    return nil
}

// AxinFundCreateDTO Getter
func (r TaobaoAlitripAxinTransFundAddRequest) GetAxinFundCreateDTO() *AxinFundCreateDto {
    return r._axinFundCreateDTO
}
