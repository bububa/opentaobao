package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectTradeitem 新二品同步
// alibaba.alihouse.newhome.project.tradeitem
//
// 新二品同步
func AlibabaAlihouseNewhomeProjectTradeitem(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectTradeitemAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeProjectTradeitemAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeProjectTradeitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
