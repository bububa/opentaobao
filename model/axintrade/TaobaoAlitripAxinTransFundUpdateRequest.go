package axintrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改资金单接口 API请求
taobao.alitrip.axin.trans.fund.update

阿信供销平台-修改资金单接口
*/
type TaobaoAlitripAxinTransFundUpdateRequest struct {
    model.Params
    // 更新资金单接口入参
    _axinFundUpdateDTO   *AxinFundUpdateDto
}

// 初始化TaobaoAlitripAxinTransFundUpdateRequest对象
func NewTaobaoAlitripAxinTransFundUpdateRequest() *TaobaoAlitripAxinTransFundUpdateRequest{
    return &TaobaoAlitripAxinTransFundUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripAxinTransFundUpdateRequest) GetApiMethodName() string {
    return "taobao.alitrip.axin.trans.fund.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripAxinTransFundUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AxinFundUpdateDTO Setter
// 更新资金单接口入参
func (r *TaobaoAlitripAxinTransFundUpdateRequest) SetAxinFundUpdateDTO(_axinFundUpdateDTO *AxinFundUpdateDto) error {
    r._axinFundUpdateDTO = _axinFundUpdateDTO
    r.Set("axin_fund_update_d_t_o", _axinFundUpdateDTO)
    return nil
}

// AxinFundUpdateDTO Getter
func (r TaobaoAlitripAxinTransFundUpdateRequest) GetAxinFundUpdateDTO() *AxinFundUpdateDto {
    return r._axinFundUpdateDTO
}
