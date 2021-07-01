package lstlogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-异云-使用第三方物流发货 API请求
alibaba.lst.logistics.thirdpart.send

异地云仓的订单，使用第三方物流的发货方式来变更订单发货状态
*/
type AlibabaLstLogisticsThirdpartSendAPIRequest struct {
    model.Params
    // 入参
    _param   *SendOfflineOrderParam
}

// 初始化AlibabaLstLogisticsThirdpartSendAPIRequest对象
func NewAlibabaLstLogisticsThirdpartSendRequest() *AlibabaLstLogisticsThirdpartSendAPIRequest{
    return &AlibabaLstLogisticsThirdpartSendAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstLogisticsThirdpartSendAPIRequest) GetApiMethodName() string {
    return "alibaba.lst.logistics.thirdpart.send"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstLogisticsThirdpartSendAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参
func (r *AlibabaLstLogisticsThirdpartSendAPIRequest) SetParam(_param *SendOfflineOrderParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaLstLogisticsThirdpartSendAPIRequest) GetParam() *SendOfflineOrderParam {
    return r._param
}
