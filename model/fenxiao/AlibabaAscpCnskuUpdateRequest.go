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
type AlibabaAscpCnskuUpdateRequest struct {
    model.Params
    // 待新增的货品
    _cnsku   *CnskuDTO
    // 修改选项
    _option   *UpdateCnskuOption
}

// 初始化AlibabaAscpCnskuUpdateRequest对象
func NewAlibabaAscpCnskuUpdateRequest() *AlibabaAscpCnskuUpdateRequest{
    return &AlibabaAscpCnskuUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpCnskuUpdateRequest) GetApiMethodName() string {
    return "alibaba.ascp.cnsku.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpCnskuUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Cnsku Setter
// 待新增的货品
func (r *AlibabaAscpCnskuUpdateRequest) SetCnsku(_cnsku *CnskuDTO) error {
    r._cnsku = _cnsku
    r.Set("cnsku", _cnsku)
    return nil
}

// Cnsku Getter
func (r AlibabaAscpCnskuUpdateRequest) GetCnsku() *CnskuDTO {
    return r._cnsku
}
// Option Setter
// 修改选项
func (r *AlibabaAscpCnskuUpdateRequest) SetOption(_option *UpdateCnskuOption) error {
    r._option = _option
    r.Set("option", _option)
    return nil
}

// Option Getter
func (r AlibabaAscpCnskuUpdateRequest) GetOption() *UpdateCnskuOption {
    return r._option
}
