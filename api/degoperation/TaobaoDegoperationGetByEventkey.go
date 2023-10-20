package degoperation

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/degoperation"
)

// Taobaodegoperationgetbyeventkey 通用用户抽奖次数限制
// taobao.degoperation.get.by.eventkey
//
// 通用用户抽奖次数限制
func Taobaodegoperationgetbyeventkey(clt *core.SDKClient, req *degoperation.TaobaodegoperationgetbyeventkeyAPIRequest, session string) (*degoperation.TaobaodegoperationgetbyeventkeyAPIResponse, error) {
	var resp degoperation.TaobaodegoperationgetbyeventkeyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
