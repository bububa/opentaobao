package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractIsvadminBind 创建及绑定互动实例
// alibaba.interact.isvadmin.bind
//
// 创建互动实例，并绑定奖池
func AlibabaInteractIsvadminBind(clt *core.SDKClient, req *interact.AlibabaInteractIsvadminBindAPIRequest, session string) (*interact.AlibabaInteractIsvadminBindAPIResponse, error) {
	var resp interact.AlibabaInteractIsvadminBindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
