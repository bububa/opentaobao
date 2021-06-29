package alink

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询设备标准属性最新状态 API请求
alibaba.alink.device.unify.status.get

查询设备最新标准属性状态
*/
type AlibabaAlinkDeviceUnifyStatusGetRequest struct {
    model.Params
    // 设备id
    uuid   string
}

// 初始化AlibabaAlinkDeviceUnifyStatusGetRequest对象
func NewAlibabaAlinkDeviceUnifyStatusGetRequest() *AlibabaAlinkDeviceUnifyStatusGetRequest{
    return &AlibabaAlinkDeviceUnifyStatusGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlinkDeviceUnifyStatusGetRequest) GetApiMethodName() string {
    return "alibaba.alink.device.unify.status.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlinkDeviceUnifyStatusGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Uuid Setter
// 设备id
func (r *AlibabaAlinkDeviceUnifyStatusGetRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

// Uuid Getter
func (r AlibabaAlinkDeviceUnifyStatusGetRequest) GetUuid() string {
    return r.uuid
}
