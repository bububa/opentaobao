package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

/* TaobaoOpentradeGroupMemberInfo
组团购获取用户参团信息
taobao.opentrade.group.member.info

组团购场景下，获取用户参团信息 */
func TaobaoOpentradeGroupMemberInfo(clt *core.SDKClient, req *opentrade.TaobaoOpentradeGroupMemberInfoAPIRequest, session string) (*opentrade.TaobaoOpentradeGroupMemberInfoAPIResponse, error) {
	var resp opentrade.TaobaoOpentradeGroupMemberInfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
