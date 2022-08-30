package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeTradeitemRelation 货独立绑定货品
// alibaba.alihouse.newhome.tradeitem.relation
//
// 货独立绑定货品
func AlibabaAlihouseNewhomeTradeitemRelation(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeTradeitemRelationAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeTradeitemRelationAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeTradeitemRelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
