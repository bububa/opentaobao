package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

/* TaobaoCrmMemberGroupGet
获取买家身上的标签
taobao.crm.member.group.get

获取买家身上的标签，不返回标签的总人数 */
func TaobaoCrmMemberGroupGet(clt *core.SDKClient, req *crm.TaobaoCrmMemberGroupGetAPIRequest, session string) (*crm.TaobaoCrmMemberGroupGetAPIResponse, error) {
	var resp crm.TaobaoCrmMemberGroupGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
