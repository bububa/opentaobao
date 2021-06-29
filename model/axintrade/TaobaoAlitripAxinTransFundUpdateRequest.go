package axintrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改资金单接口 APIRequest
taobao.alitrip.axin.trans.fund.update

阿信供销平台-修改资金单接口
*/
type TaobaoAlitripAxinTransFundUpdateRequest struct {
    model.Params

    // 更新资金单接口入参
    axinFundUpdateDTO   *AxinFundUpdateDto 

}

func NewTaobaoAlitripAxinTransFundUpdateRequest() *TaobaoAlitripAxinTransFundUpdateRequest{
    return &TaobaoAlitripAxinTransFundUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripAxinTransFundUpdateRequest) GetApiMethodName() string {
    return "taobao.alitrip.axin.trans.fund.update"
}

func (r TaobaoAlitripAxinTransFundUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripAxinTransFundUpdateRequest) SetAxinFundUpdateDTO(axinFundUpdateDTO *AxinFundUpdateDto) error {
    r.axinFundUpdateDTO = axinFundUpdateDTO
    r.Set("axin_fund_update_d_t_o", axinFundUpdateDTO)
    return nil
}

func (r TaobaoAlitripAxinTransFundUpdateRequest) GetAxinFundUpdateDTO() *AxinFundUpdateDto {
    return r.axinFundUpdateDTO
}

