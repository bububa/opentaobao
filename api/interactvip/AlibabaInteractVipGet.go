package interactvip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interactvip"
)

// AlibabaInteractVipGet 会员淘气值获取
// alibaba.interact.vip.get
//
// 提供用户淘气值&用户角色身份查询
func AlibabaInteractVipGet(clt *core.SDKClient, req *interactvip.AlibabaInteractVipGetAPIRequest, session string) (*interactvip.AlibabaInteractVipGetAPIResponse, error) {
	var resp interactvip.AlibabaInteractVipGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
