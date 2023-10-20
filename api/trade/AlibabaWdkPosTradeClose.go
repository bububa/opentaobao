package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Alibabawdkpostradeclose 轻pos品牌营销关单接口
// alibaba.wdk.pos.trade.close
//
// 轻pos品牌营销场景，提供关单接口给外部商家
func Alibabawdkpostradeclose(clt *core.SDKClient, req *trade.AlibabawdkpostradecloseAPIRequest, session string) (*trade.AlibabawdkpostradecloseAPIResponse, error) {
	var resp trade.AlibabawdkpostradecloseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
