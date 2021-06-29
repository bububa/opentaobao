package lsttrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
十荟团订单同步至零售通 API请求
alibaba.lst.nicetuan.order.save

十荟团订单同步至零售通，十荟团单向写到零售通
*/
type AlibabaLstNicetuanOrderSaveRequest struct {
    model.Params
    // 订单数据
    param   *NicetuanMainOrderParam
}

// 初始化AlibabaLstNicetuanOrderSaveRequest对象
func NewAlibabaLstNicetuanOrderSaveRequest() *AlibabaLstNicetuanOrderSaveRequest{
    return &AlibabaLstNicetuanOrderSaveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstNicetuanOrderSaveRequest) GetApiMethodName() string {
    return "alibaba.lst.nicetuan.order.save"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstNicetuanOrderSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 订单数据
func (r *AlibabaLstNicetuanOrderSaveRequest) SetParam(param *NicetuanMainOrderParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaLstNicetuanOrderSaveRequest) GetParam() *NicetuanMainOrderParam {
    return r.param
}
