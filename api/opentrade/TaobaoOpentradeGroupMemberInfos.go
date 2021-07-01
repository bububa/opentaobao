package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

/* TaobaoOpentradeGroupMemberInfos
组团购批量获取用户参团信息
taobao.opentrade.group.member.infos

组团购场景下，获取用户参团信息 */
func TaobaoOpentradeGroupMemberInfos(clt *core.SDKClient, req *opentrade.TaobaoOpentradeGroupMemberInfosAPIRequest, session string) (*opentrade.TaobaoOpentradeGroupMemberInfosAPIResponse, error) {
	var resp opentrade.TaobaoOpentradeGroupMemberInfosAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
