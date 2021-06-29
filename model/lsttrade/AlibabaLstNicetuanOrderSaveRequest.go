package lsttrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
十荟团订单同步至零售通 APIRequest
alibaba.lst.nicetuan.order.save

十荟团订单同步至零售通，十荟团单向写到零售通
*/
type AlibabaLstNicetuanOrderSaveRequest struct {
    model.Params

    // 订单数据
    param   *NicetuanMainOrderParam 

}

func NewAlibabaLstNicetuanOrderSaveRequest() *AlibabaLstNicetuanOrderSaveRequest{
    return &AlibabaLstNicetuanOrderSaveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstNicetuanOrderSaveRequest) GetApiMethodName() string {
    return "alibaba.lst.nicetuan.order.save"
}

func (r AlibabaLstNicetuanOrderSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstNicetuanOrderSaveRequest) SetParam(param *NicetuanMainOrderParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaLstNicetuanOrderSaveRequest) GetParam() *NicetuanMainOrderParam {
    return r.param
}

