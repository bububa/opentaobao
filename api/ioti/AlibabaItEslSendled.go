package ioti

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ioti"
)

/* 
厂测LED控制 
alibaba.it.esl.sendled

针对厂测生产的的价签，增加led闪灯的接口，进行led 闪灯测试
*/
func AlibabaItEslSendled(clt *core.SDKClient, req *ioti.AlibabaItEslSendledAPIRequest, session string) (*ioti.AlibabaItEslSendledAPIResponse, error) {
    var resp ioti.AlibabaItEslSendledAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
