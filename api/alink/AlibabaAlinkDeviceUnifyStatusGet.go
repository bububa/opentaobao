package alink

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alink"
)

/* 
查询设备标准属性最新状态 
alibaba.alink.device.unify.status.get

查询设备最新标准属性状态
*/
func AlibabaAlinkDeviceUnifyStatusGet(clt *core.SDKClient, req *alink.AlibabaAlinkDeviceUnifyStatusGetAPIRequest, session string) (*alink.AlibabaAlinkDeviceUnifyStatusGetAPIResponse, error) {
    var resp alink.AlibabaAlinkDeviceUnifyStatusGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
