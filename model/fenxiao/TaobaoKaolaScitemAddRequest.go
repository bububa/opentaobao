package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
考拉货品新增接口 APIRequest
taobao.kaola.scitem.add

考拉货品新增接口
*/
type TaobaoKaolaScitemAddRequest struct {
    model.Params

    // 待新增的货品
    cnsku   *CnskuDto 

    // 新增选项
    option   *AddCnskuOption 

}

func NewTaobaoKaolaScitemAddRequest() *TaobaoKaolaScitemAddRequest{
    return &TaobaoKaolaScitemAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoKaolaScitemAddRequest) GetApiMethodName() string {
    return "taobao.kaola.scitem.add"
}

func (r TaobaoKaolaScitemAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoKaolaScitemAddRequest) SetCnsku(cnsku *CnskuDto) error {
    r.cnsku = cnsku
    r.Set("cnsku", cnsku)
    return nil
}

func (r TaobaoKaolaScitemAddRequest) GetCnsku() *CnskuDto {
    return r.cnsku
}

func (r *TaobaoKaolaScitemAddRequest) SetOption(option *AddCnskuOption) error {
    r.option = option
    r.Set("option", option)
    return nil
}

func (r TaobaoKaolaScitemAddRequest) GetOption() *AddCnskuOption {
    return r.option
}

