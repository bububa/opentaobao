package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// TaobaoCrmGroupsGet 查询卖家的分组
// taobao.crm.groups.get
//
// 查询卖家的分组，返回查询到的分组列表，分页返回分组
func TaobaoCrmGroupsGet(clt *core.SDKClient, req *crm.TaobaoCrmGroupsGetAPIRequest, session string) (*crm.TaobaoCrmGroupsGetAPIResponse, error) {
	var resp crm.TaobaoCrmGroupsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
