package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Alibabawdktradeordercreate 外部交易订单创单接口
// alibaba.wdk.trade.order.create
//
// 通过该接口可以再盒马创建交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
func Alibabawdktradeordercreate(clt *core.SDKClient, req *trade.AlibabawdktradeordercreateAPIRequest, session string) (*trade.AlibabawdktradeordercreateAPIResponse, error) {
	var resp trade.AlibabawdktradeordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
