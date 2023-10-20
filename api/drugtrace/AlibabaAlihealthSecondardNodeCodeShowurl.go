package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthsecondardnodecodeshowurl 查询码信息url
// alibaba.alihealth.secondard.node.code.showurl
//
// 二级节点查询码信息url
func Alibabaalihealthsecondardnodecodeshowurl(clt *core.SDKClient, req *drugtrace.AlibabaalihealthsecondardnodecodeshowurlAPIRequest, session string) (*drugtrace.AlibabaalihealthsecondardnodecodeshowurlAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthsecondardnodecodeshowurlAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
