package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
canvas工具包 API请求
alibaba.interact.sensor.gutil

canvas工具包
*/
type AlibabaInteractSensorGutilRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorGutilRequest对象
func NewAlibabaInteractSensorGutilRequest() *AlibabaInteractSensorGutilRequest{
    return &AlibabaInteractSensorGutilRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorGutilRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.gutil"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorGutilRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
