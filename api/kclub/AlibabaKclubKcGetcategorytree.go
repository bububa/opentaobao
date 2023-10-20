package kclub

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/kclub"
)

// AlibabaKclubKcGetcategorytree 知识云-查询租户下类目树
// alibaba.kclub.kc.getcategorytree
//
// 知识云-查询租户下类目树。通过租户id、类型(外部类目、帮助中心类目、内部类目)。
func AlibabaKclubKcGetcategorytree(clt *core.SDKClient, req *kclub.AlibabaKclubKcGetcategorytreeAPIRequest, resp *kclub.AlibabaKclubKcGetcategorytreeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
