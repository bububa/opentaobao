package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

/* AlibabaInteractIsvadminAllponds
获取天猫互动奖池列表
alibaba.interact.isvadmin.allponds

获取天猫互动奖池列表 */
func AlibabaInteractIsvadminAllponds(clt *core.SDKClient, req *interact.AlibabaInteractIsvadminAllpondsAPIRequest, session string) (*interact.AlibabaInteractIsvadminAllpondsAPIResponse, error) {
	var resp interact.AlibabaInteractIsvadminAllpondsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
