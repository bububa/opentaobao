package aliqin

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliqin"
)

/* 
物联卡信息查询 
alibaba.aliqin.fc.iot.cardInfo

物联卡信息查询
*/
func AlibabaAliqinFcIotCardInfo(clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIotCardInfoRequest, session string) (*aliqin.AlibabaAliqinFcIotCardInfoResponse, error) {
    var resp aliqin.AlibabaAliqinFcIotCardInfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
