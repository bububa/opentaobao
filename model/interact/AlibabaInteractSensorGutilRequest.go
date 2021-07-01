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
type AlibabaInteractSensorGutilAPIRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorGutilAPIRequest对象
func NewAlibabaInteractSensorGutilRequest() *AlibabaInteractSensorGutilAPIRequest{
    return &AlibabaInteractSensorGutilAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorGutilAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.gutil"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorGutilAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
