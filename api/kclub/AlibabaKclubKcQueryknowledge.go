package kclub

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/kclub"
)

/* 
知识云-通用知识查询服务 
alibaba.kclub.kc.queryknowledge

知识云-通用知识查询服务。通过租户id、类目id、知识类型、知识状态等条件查询类目。
*/
func AlibabaKclubKcQueryknowledge(clt *core.SDKClient, req *kclub.AlibabaKclubKcQueryknowledgeAPIRequest, session string) (*kclub.AlibabaKclubKcQueryknowledgeAPIResponse, error) {
    var resp kclub.AlibabaKclubKcQueryknowledgeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
