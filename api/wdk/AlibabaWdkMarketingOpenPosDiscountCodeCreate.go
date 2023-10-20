package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingopenposdiscountcodecreate pos一物一码创建
// alibaba.wdk.marketing.open.pos.discount.code.create
//
// pos一物一码创建
func Alibabawdkmarketingopenposdiscountcodecreate(clt *core.SDKClient, req *wdk.AlibabawdkmarketingopenposdiscountcodecreateAPIRequest, session string) (*wdk.AlibabawdkmarketingopenposdiscountcodecreateAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingopenposdiscountcodecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
