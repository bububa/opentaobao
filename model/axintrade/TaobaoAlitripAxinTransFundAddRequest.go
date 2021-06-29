package axintrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建资金单接口 APIRequest
taobao.alitrip.axin.trans.fund.add

创建资金单
*/
type TaobaoAlitripAxinTransFundAddRequest struct {
    model.Params

    // 创建资金单接口入参
    axinFundCreateDTO   *AxinFundCreateDto 

}

func NewTaobaoAlitripAxinTransFundAddRequest() *TaobaoAlitripAxinTransFundAddRequest{
    return &TaobaoAlitripAxinTransFundAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripAxinTransFundAddRequest) GetApiMethodName() string {
    return "taobao.alitrip.axin.trans.fund.add"
}

func (r TaobaoAlitripAxinTransFundAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripAxinTransFundAddRequest) SetAxinFundCreateDTO(axinFundCreateDTO *AxinFundCreateDto) error {
    r.axinFundCreateDTO = axinFundCreateDTO
    r.Set("axin_fund_create_d_t_o", axinFundCreateDTO)
    return nil
}

func (r TaobaoAlitripAxinTransFundAddRequest) GetAxinFundCreateDTO() *AxinFundCreateDto {
    return r.axinFundCreateDTO
}

