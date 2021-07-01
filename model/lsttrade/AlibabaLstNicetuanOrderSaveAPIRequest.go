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
type AlibabaLstNicetuanOrderSaveAPIRequest struct {
    model.Params
    // 订单数据
    _param   *NicetuanMainOrderParam
}

// 初始化AlibabaLstNicetuanOrderSaveAPIRequest对象
func NewAlibabaLstNicetuanOrderSaveRequest() *AlibabaLstNicetuanOrderSaveAPIRequest{
    return &AlibabaLstNicetuanOrderSaveAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstNicetuanOrderSaveAPIRequest) GetApiMethodName() string {
    return "alibaba.lst.nicetuan.order.save"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstNicetuanOrderSaveAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 订单数据
func (r *AlibabaLstNicetuanOrderSaveAPIRequest) SetParam(_param *NicetuanMainOrderParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaLstNicetuanOrderSaveAPIRequest) GetParam() *NicetuanMainOrderParam {
    return r._param
}
