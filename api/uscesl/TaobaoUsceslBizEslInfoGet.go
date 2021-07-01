package uscesl

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/uscesl"
)

/* 
价签设备信息查询接口 
taobao.uscesl.biz.esl.info.get

价签设备信息查询接口
*/
func TaobaoUsceslBizEslInfoGet(clt *core.SDKClient, req *uscesl.TaobaoUsceslBizEslInfoGetAPIRequest, session string) (*uscesl.TaobaoUsceslBizEslInfoGetAPIResponse, error) {
    var resp uscesl.TaobaoUsceslBizEslInfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
