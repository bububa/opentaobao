package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应链中台货品修改接口 API请求
alibaba.ascp.cnsku.update

供应链中台货品修改接口
*/
type AlibabaAscpCnskuUpdateAPIRequest struct {
    model.Params
    // 待新增的货品
    _cnsku   *CnskuDTO
    // 修改选项
    _option   *UpdateCnskuOption
}

// 初始化AlibabaAscpCnskuUpdateAPIRequest对象
func NewAlibabaAscpCnskuUpdateRequest() *AlibabaAscpCnskuUpdateAPIRequest{
    return &AlibabaAscpCnskuUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpCnskuUpdateAPIRequest) GetApiMethodName() string {
    return "alibaba.ascp.cnsku.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpCnskuUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Cnsku Setter
// 待新增的货品
func (r *AlibabaAscpCnskuUpdateAPIRequest) SetCnsku(_cnsku *CnskuDTO) error {
    r._cnsku = _cnsku
    r.Set("cnsku", _cnsku)
    return nil
}

// Cnsku Getter
func (r AlibabaAscpCnskuUpdateAPIRequest) GetCnsku() *CnskuDTO {
    return r._cnsku
}
// Option Setter
// 修改选项
func (r *AlibabaAscpCnskuUpdateAPIRequest) SetOption(_option *UpdateCnskuOption) error {
    r._option = _option
    r.Set("option", _option)
    return nil
}

// Option Getter
func (r AlibabaAscpCnskuUpdateAPIRequest) GetOption() *UpdateCnskuOption {
    return r._option
}
