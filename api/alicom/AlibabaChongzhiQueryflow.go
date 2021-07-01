package alicom

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alicom"
)

/* 
查询指定商家的可用的流量宝贝接口 
alibaba.chongzhi.queryflow

查询指定商家的可用的流量宝贝
*/
func AlibabaChongzhiQueryflow(clt *core.SDKClient, req *alicom.AlibabaChongzhiQueryflowAPIRequest, session string) (*alicom.AlibabaChongzhiQueryflowAPIResponse, error) {
    var resp alicom.AlibabaChongzhiQueryflowAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
