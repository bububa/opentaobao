package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// Taobaocrmgroupsget 查询卖家的分组
// taobao.crm.groups.get
//
// 查询卖家的分组，返回查询到的分组列表，分页返回分组
func Taobaocrmgroupsget(clt *core.SDKClient, req *crm.TaobaocrmgroupsgetAPIRequest, session string) (*crm.TaobaocrmgroupsgetAPIResponse, error) {
	var resp crm.TaobaocrmgroupsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
