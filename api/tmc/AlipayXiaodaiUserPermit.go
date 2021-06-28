package tmc

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmc"
)

/* 
阿里金融用户授权 
alipay.xiaodai.user.permit

阿里金融为用户开通消息通道接口
*/
func AlipayXiaodaiUserPermit(clt *core.SDKClient, req *tmc.AlipayXiaodaiUserPermitRequest, session string) (*tmc.AlipayXiaodaiUserPermitAPIResponse, error) {
    var resp tmc.AlipayXiaodaiUserPermitAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
