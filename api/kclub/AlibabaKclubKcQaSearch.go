package kclub

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/kclub"
)

// Alibabakclubkcqasearch 知识云-知识检索
// alibaba.kclub.kc.qa.search
//
// 知识云-知识搜索服务
func Alibabakclubkcqasearch(clt *core.SDKClient, req *kclub.AlibabakclubkcqasearchAPIRequest, session string) (*kclub.AlibabakclubkcqasearchAPIResponse, error) {
	var resp kclub.AlibabakclubkcqasearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
