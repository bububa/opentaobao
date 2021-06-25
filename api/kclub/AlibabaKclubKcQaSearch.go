package kclub

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/kclub"
)

/* 
知识云-知识检索 
alibaba.kclub.kc.qa.search

知识云-知识搜索服务
*/
func AlibabaKclubKcQaSearch(clt *core.SDKClient, req *kclub.AlibabaKclubKcQaSearchRequest, session string) (*kclub.AlibabaKclubKcQaSearchResponse, error) {
    var resp kclub.AlibabaKclubKcQaSearchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
