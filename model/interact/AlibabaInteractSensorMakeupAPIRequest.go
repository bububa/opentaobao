package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
美妆虚拟试装 API请求
alibaba.interact.sensor.makeup

手机淘宝美妆类目虚拟试妆权限，客户端能力（JS－API）
*/
type AlibabaInteractSensorMakeupAPIRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorMakeupAPIRequest对象
func NewAlibabaInteractSensorMakeupRequest() *AlibabaInteractSensorMakeupAPIRequest{
    return &AlibabaInteractSensorMakeupAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorMakeupAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.makeup"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorMakeupAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
