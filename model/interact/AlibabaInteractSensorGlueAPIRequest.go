package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
视频播放 API请求
alibaba.interact.sensor.glue

视频播放
*/
type AlibabaInteractSensorGlueAPIRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorGlueAPIRequest对象
func NewAlibabaInteractSensorGlueRequest() *AlibabaInteractSensorGlueAPIRequest{
    return &AlibabaInteractSensorGlueAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorGlueAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.glue"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorGlueAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
