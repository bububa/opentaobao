package kclub

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/kclub"
)

// AlibabaKclubKcQaSearchPage 知识云-知识检索(分页)
// alibaba.kclub.kc.qa.search.page
//
// 知识云-知识搜索服务
func AlibabaKclubKcQaSearchPage(clt *core.SDKClient, req *kclub.AlibabaKclubKcQaSearchPageAPIRequest, resp *kclub.AlibabaKclubKcQaSearchPageAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
