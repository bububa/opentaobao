package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Alibabawdkpostradecreate 轻pos品牌营销下单接口
// alibaba.wdk.pos.trade.create
//
// 提供给石基进行轻pos品牌营销下单
func Alibabawdkpostradecreate(clt *core.SDKClient, req *trade.AlibabawdkpostradecreateAPIRequest, session string) (*trade.AlibabawdkpostradecreateAPIResponse, error) {
	var resp trade.AlibabawdkpostradecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
