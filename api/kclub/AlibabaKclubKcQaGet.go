package kclub

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/kclub"
)

/* 
知识云-查询单个知识详情 
alibaba.kclub.kc.qa.get

知识云-查询单个知识详情。通过租户id、问题id查询问题详情
*/
func AlibabaKclubKcQaGet(clt *core.SDKClient, req *kclub.AlibabaKclubKcQaGetAPIRequest, session string) (*kclub.AlibabaKclubKcQaGetAPIResponse, error) {
    var resp kclub.AlibabaKclubKcQaGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
