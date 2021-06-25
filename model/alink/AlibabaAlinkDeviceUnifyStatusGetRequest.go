package alink

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询设备标准属性最新状态 APIRequest
alibaba.alink.device.unify.status.get

查询设备最新标准属性状态
*/
type AlibabaAlinkDeviceUnifyStatusGetRequest struct {
    model.Params

    // 设备id
    uuid   string 

}

func NewAlibabaAlinkDeviceUnifyStatusGetRequest() *AlibabaAlinkDeviceUnifyStatusGetRequest{
    return &AlibabaAlinkDeviceUnifyStatusGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlinkDeviceUnifyStatusGetRequest) GetApiMethodName() string {
    return "alibaba.alink.device.unify.status.get"
}

func (r AlibabaAlinkDeviceUnifyStatusGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlinkDeviceUnifyStatusGetRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r AlibabaAlinkDeviceUnifyStatusGetRequest) GetUuid() string {
    return r.uuid
}

