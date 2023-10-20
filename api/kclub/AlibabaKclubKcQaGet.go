package kclub

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/kclub"
)

// Alibabakclubkcqaget 知识云-查询单个知识详情
// alibaba.kclub.kc.qa.get
//
// 知识云-查询单个知识详情。通过租户id、问题id查询问题详情
func Alibabakclubkcqaget(clt *core.SDKClient, req *kclub.AlibabakclubkcqagetAPIRequest, session string) (*kclub.AlibabakclubkcqagetAPIResponse, error) {
	var resp kclub.AlibabakclubkcqagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
