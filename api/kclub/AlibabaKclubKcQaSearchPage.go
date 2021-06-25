package kclub

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/kclub"
)

/* 
知识云-知识检索(分页) 
alibaba.kclub.kc.qa.search.page

知识云-知识搜索服务
*/
func AlibabaKclubKcQaSearchPage(clt *core.SDKClient, req *kclub.AlibabaKclubKcQaSearchPageRequest, session string) (*kclub.AlibabaKclubKcQaSearchPageResponse, error) {
    var resp kclub.AlibabaKclubKcQaSearchPageAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
