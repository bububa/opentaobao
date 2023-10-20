package mei

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mei"
)

// Tmallmeicrmcallbackpointchange 品牌积分变更回调API
// tmall.mei.crm.callback.point.change
//
// 线下品牌积分变更消息回调API，告诉积分扣减或者累加是否成功。
func Tmallmeicrmcallbackpointchange(clt *core.SDKClient, req *mei.TmallmeicrmcallbackpointchangeAPIRequest, session string) (*mei.TmallmeicrmcallbackpointchangeAPIResponse, error) {
	var resp mei.TmallmeicrmcallbackpointchangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
