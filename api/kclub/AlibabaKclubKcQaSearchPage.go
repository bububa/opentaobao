package kclub

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/kclub"
)

// Alibabakclubkcqasearchpage 知识云-知识检索(分页)
// alibaba.kclub.kc.qa.search.page
//
// 知识云-知识搜索服务
func Alibabakclubkcqasearchpage(clt *core.SDKClient, req *kclub.AlibabakclubkcqasearchpageAPIRequest, session string) (*kclub.AlibabakclubkcqasearchpageAPIResponse, error) {
	var resp kclub.AlibabakclubkcqasearchpageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
