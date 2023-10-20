package c2m

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/c2m"
)

// TaobaoSebpOrganizationGetorderinfo 淘小铺机构订单信息
// taobao.sebp.organization.getorderinfo
//
// 淘小铺合作机构获取机构订单信息，用于机构结算使用
func TaobaoSebpOrganizationGetorderinfo(clt *core.SDKClient, req *c2m.TaobaoSebpOrganizationGetorderinfoAPIRequest, resp *c2m.TaobaoSebpOrganizationGetorderinfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
