package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

/* AlibabaInteractIsvadminGetinteractbysellernick
根据sellerNick获取互动实例列表
alibaba.interact.isvadmin.getinteractbysellernick

根据sellerNick获取互动实例列表 */
func AlibabaInteractIsvadminGetinteractbysellernick(clt *core.SDKClient, req *interact.AlibabaInteractIsvadminGetinteractbysellernickAPIRequest, session string) (*interact.AlibabaInteractIsvadminGetinteractbysellernickAPIResponse, error) {
	var resp interact.AlibabaInteractIsvadminGetinteractbysellernickAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
