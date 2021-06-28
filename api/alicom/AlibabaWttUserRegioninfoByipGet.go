package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
根据ip获取省市信息 
alibaba.wtt.user.regioninfo.byip.get

通过ip获取省市信息
*/
func AlibabaWttUserRegioninfoByipGet(clt *core.SDKClient, req *alicom.AlibabaWttUserRegioninfoByipGetRequest, session string) (*alicom.AlibabaWttUserRegioninfoByipGetAPIResponse, error) {
    var resp alicom.AlibabaWttUserRegioninfoByipGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
