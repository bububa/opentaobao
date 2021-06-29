package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
takePhoto API请求
alibaba.interact.sensor.takephoto

客户端takePhoto
*/
type AlibabaInteractSensorTakephotoRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorTakephotoRequest对象
func NewAlibabaInteractSensorTakephotoRequest() *AlibabaInteractSensorTakephotoRequest{
    return &AlibabaInteractSensorTakephotoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorTakephotoRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.takephoto"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorTakephotoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
