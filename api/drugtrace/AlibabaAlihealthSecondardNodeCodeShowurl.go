package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthSecondardNodeCodeShowurl 查询码信息url
// alibaba.alihealth.secondard.node.code.showurl
//
// 二级节点查询码信息url
func AlibabaAlihealthSecondardNodeCodeShowurl(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthSecondardNodeCodeShowurlAPIRequest, resp *drugtrace.AlibabaAlihealthSecondardNodeCodeShowurlAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
