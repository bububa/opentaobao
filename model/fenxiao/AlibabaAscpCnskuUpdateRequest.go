package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应链中台货品修改接口 APIRequest
alibaba.ascp.cnsku.update

供应链中台货品修改接口
*/
type AlibabaAscpCnskuUpdateRequest struct {
    model.Params

    // 待新增的货品
    cnsku   *CnskuDto 

    // 修改选项
    option   *UpdateCnskuOption 

}

func NewAlibabaAscpCnskuUpdateRequest() *AlibabaAscpCnskuUpdateRequest{
    return &AlibabaAscpCnskuUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpCnskuUpdateRequest) GetApiMethodName() string {
    return "alibaba.ascp.cnsku.update"
}

func (r AlibabaAscpCnskuUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpCnskuUpdateRequest) SetCnsku(cnsku *CnskuDto) error {
    r.cnsku = cnsku
    r.Set("cnsku", cnsku)
    return nil
}

func (r AlibabaAscpCnskuUpdateRequest) GetCnsku() *CnskuDto {
    return r.cnsku
}

func (r *AlibabaAscpCnskuUpdateRequest) SetOption(option *UpdateCnskuOption) error {
    r.option = option
    r.Set("option", option)
    return nil
}

func (r AlibabaAscpCnskuUpdateRequest) GetOption() *UpdateCnskuOption {
    return r.option
}

