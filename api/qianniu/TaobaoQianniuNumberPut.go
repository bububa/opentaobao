package qianniu

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/qianniu"
)

/* 
ISV上传数据接口 
taobao.qianniu.number.put

ISV提供给卖家使用的业务数据，需要通过这个接口上传到千牛数据中心。
*/
func TaobaoQianniuNumberPut(clt *core.SDKClient, req *qianniu.TaobaoQianniuNumberPutAPIRequest, session string) (*qianniu.TaobaoQianniuNumberPutAPIResponse, error) {
    var resp qianniu.TaobaoQianniuNumberPutAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
