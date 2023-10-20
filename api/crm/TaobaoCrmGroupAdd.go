package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmGroupAdd 卖家创建一个分组
// taobao.crm.group.add
//
// 卖家创建一个新的分组，接口返回一个创建成功的分组的id
func TaobaoCrmGroupAdd(clt *core.SDKClient, req *crm.TaobaoCrmGroupAddAPIRequest, resp *crm.TaobaoCrmGroupAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
