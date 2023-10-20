package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthSecondardNodeCodeShowurl 查询码信息url
// alibaba.alihealth.secondard.node.code.showurl
//
// 二级节点查询码信息url
func AlibabaAlihealthSecondardNodeCodeShowurl(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthSecondardNodeCodeShowurlAPIRequest, session string) (*drugtrace.AlibabaAlihealthSecondardNodeCodeShowurlAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthSecondardNodeCodeShowurlAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
