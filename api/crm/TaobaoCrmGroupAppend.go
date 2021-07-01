package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

/* TaobaoCrmGroupAppend
将一个分组添加到另外一个分组
taobao.crm.group.append

将某分组下的所有会员添加到另一个分组,注：1.该操作为异步任务，建议先调用taobao.crm.grouptask.check 确保涉及分组上没有任务；2.若分组下某会员分组数超最大限额，则该会员不会被添加到新分组，同时不影响其余会员添加分组，接口调用依然返回成功。 */
func TaobaoCrmGroupAppend(clt *core.SDKClient, req *crm.TaobaoCrmGroupAppendAPIRequest, session string) (*crm.TaobaoCrmGroupAppendAPIResponse, error) {
	var resp crm.TaobaoCrmGroupAppendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
