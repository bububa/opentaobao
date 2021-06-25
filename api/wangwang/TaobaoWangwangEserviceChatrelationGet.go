package wangwang

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wangwang"
)

/* 
聊天关系获取接口 
taobao.wangwang.eservice.chatrelation.get

获取指定时间区间内的聊天关系（查询账号，和谁，在哪天说过话）。如：
A 和 B 在2016-09-01 和 2016-09-02 都说过话。以A为查询账号，则该接口将返回：
2016-09-01， B
2016-09-02， B
*/
func TaobaoWangwangEserviceChatrelationGet(clt *core.SDKClient, req *wangwang.TaobaoWangwangEserviceChatrelationGetRequest, session string) (*wangwang.TaobaoWangwangEserviceChatrelationGetResponse, error) {
    var resp wangwang.TaobaoWangwangEserviceChatrelationGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
