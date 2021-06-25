package xiami

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xiami"
)

/* 
我的电台 
alibaba.xiami.api.radio.myself.get

我的电台
*/
func AlibabaXiamiApiRadioMyselfGet(clt *core.SDKClient, req *xiami.AlibabaXiamiApiRadioMyselfGetRequest, session string) (*xiami.AlibabaXiamiApiRadioMyselfGetResponse, error) {
    var resp xiami.AlibabaXiamiApiRadioMyselfGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
