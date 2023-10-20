package interactvip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interactvip"
)

// Alibabainteractvipget 会员淘气值获取
// alibaba.interact.vip.get
//
// 提供用户淘气值&amp;用户角色身份查询
func Alibabainteractvipget(clt *core.SDKClient, req *interactvip.AlibabainteractvipgetAPIRequest, session string) (*interactvip.AlibabainteractvipgetAPIResponse, error) {
	var resp interactvip.AlibabainteractvipgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
