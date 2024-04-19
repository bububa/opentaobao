package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmPointAvailableGet CRM会员积分查询开放接口
// taobao.crm.point.available.get
//
// 查询用户在某个商家的可用积分数
func TaobaoCrmPointAvailableGet(clt *core.SDKClient, req *crm.TaobaoCrmPointAvailableGetAPIRequest, resp *crm.TaobaoCrmPointAvailableGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
