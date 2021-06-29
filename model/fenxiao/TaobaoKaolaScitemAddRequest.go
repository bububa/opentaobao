package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
考拉货品新增接口 API请求
taobao.kaola.scitem.add

考拉货品新增接口
*/
type TaobaoKaolaScitemAddRequest struct {
    model.Params
    // 待新增的货品
    _cnsku   *CnskuDTO
    // 新增选项
    _option   *AddCnskuOption
}

// 初始化TaobaoKaolaScitemAddRequest对象
func NewTaobaoKaolaScitemAddRequest() *TaobaoKaolaScitemAddRequest{
    return &TaobaoKaolaScitemAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoKaolaScitemAddRequest) GetApiMethodName() string {
    return "taobao.kaola.scitem.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoKaolaScitemAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Cnsku Setter
// 待新增的货品
func (r *TaobaoKaolaScitemAddRequest) SetCnsku(_cnsku *CnskuDTO) error {
    r._cnsku = _cnsku
    r.Set("cnsku", _cnsku)
    return nil
}

// Cnsku Getter
func (r TaobaoKaolaScitemAddRequest) GetCnsku() *CnskuDTO {
    return r._cnsku
}
// Option Setter
// 新增选项
func (r *TaobaoKaolaScitemAddRequest) SetOption(_option *AddCnskuOption) error {
    r._option = _option
    r.Set("option", _option)
    return nil
}

// Option Getter
func (r TaobaoKaolaScitemAddRequest) GetOption() *AddCnskuOption {
    return r._option
}
