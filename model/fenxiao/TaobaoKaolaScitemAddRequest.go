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
    cnsku   *CnskuDto
    // 新增选项
    option   *AddCnskuOption
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
func (r *TaobaoKaolaScitemAddRequest) SetCnsku(cnsku *CnskuDto) error {
    r.cnsku = cnsku
    r.Set("cnsku", cnsku)
    return nil
}

// Cnsku Getter
func (r TaobaoKaolaScitemAddRequest) GetCnsku() *CnskuDto {
    return r.cnsku
}
// Option Setter
// 新增选项
func (r *TaobaoKaolaScitemAddRequest) SetOption(option *AddCnskuOption) error {
    r.option = option
    r.Set("option", option)
    return nil
}

// Option Getter
func (r TaobaoKaolaScitemAddRequest) GetOption() *AddCnskuOption {
    return r.option
}
