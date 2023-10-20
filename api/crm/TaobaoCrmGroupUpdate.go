package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmGroupUpdate 修改一个已经存在的分组
// taobao.crm.group.update
//
// 修改一个已经存在的分组，接口返回分组的修改是否成功
func TaobaoCrmGroupUpdate(clt *core.SDKClient, req *crm.TaobaoCrmGroupUpdateAPIRequest, resp *crm.TaobaoCrmGroupUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
