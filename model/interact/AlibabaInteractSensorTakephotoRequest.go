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
type AlibabaInteractSensorTakephotoAPIRequest struct {
    model.Params
}

// 初始化AlibabaInteractSensorTakephotoAPIRequest对象
func NewAlibabaInteractSensorTakephotoRequest() *AlibabaInteractSensorTakephotoAPIRequest{
    return &AlibabaInteractSensorTakephotoAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSensorTakephotoAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.sensor.takephoto"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSensorTakephotoAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
