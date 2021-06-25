package alink

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alink"
)

/* 
设置设备标准属性状态 
alibaba.alink.device.unify.status.set

操作用户绑定的设备
*/
func AlibabaAlinkDeviceUnifyStatusSet(clt *core.SDKClient, req *alink.AlibabaAlinkDeviceUnifyStatusSetRequest, session string) (*alink.AlibabaAlinkDeviceUnifyStatusSetResponse, error) {
    var resp alink.AlibabaAlinkDeviceUnifyStatusSetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
