package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectTradeOrder 旺铺楼盘和交易商品排序
// alibaba.alihouse.newhome.project.trade.order
//
// 旺铺楼盘和交易商品排序
func AlibabaAlihouseNewhomeProjectTradeOrder(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectTradeOrderAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeProjectTradeOrderAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeProjectTradeOrderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
