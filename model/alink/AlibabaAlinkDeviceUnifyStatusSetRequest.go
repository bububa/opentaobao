package alink

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设置设备标准属性状态 APIRequest
alibaba.alink.device.unify.status.set

操作用户绑定的设备
*/
type AlibabaAlinkDeviceUnifyStatusSetRequest struct {
    model.Params

    // uuid
    uuid   string 

    // 设备的设置参数数据
    instructions   string 

}

func NewAlibabaAlinkDeviceUnifyStatusSetRequest() *AlibabaAlinkDeviceUnifyStatusSetRequest{
    return &AlibabaAlinkDeviceUnifyStatusSetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlinkDeviceUnifyStatusSetRequest) GetApiMethodName() string {
    return "alibaba.alink.device.unify.status.set"
}

func (r AlibabaAlinkDeviceUnifyStatusSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlinkDeviceUnifyStatusSetRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r AlibabaAlinkDeviceUnifyStatusSetRequest) GetUuid() string {
    return r.uuid
}

func (r *AlibabaAlinkDeviceUnifyStatusSetRequest) SetInstructions(instructions string) error {
    r.instructions = instructions
    r.Set("instructions", instructions)
    return nil
}

func (r AlibabaAlinkDeviceUnifyStatusSetRequest) GetInstructions() string {
    return r.instructions
}

