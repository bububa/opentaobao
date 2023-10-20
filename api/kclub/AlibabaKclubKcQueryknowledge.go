package kclub

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/kclub"
)

// Alibabakclubkcqueryknowledge 知识云-通用知识查询服务
// alibaba.kclub.kc.queryknowledge
//
// 知识云-通用知识查询服务。通过租户id、类目id、知识类型、知识状态等条件查询类目。
func Alibabakclubkcqueryknowledge(clt *core.SDKClient, req *kclub.AlibabakclubkcqueryknowledgeAPIRequest, session string) (*kclub.AlibabakclubkcqueryknowledgeAPIResponse, error) {
	var resp kclub.AlibabakclubkcqueryknowledgeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
