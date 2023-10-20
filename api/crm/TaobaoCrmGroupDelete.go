package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// Taobaocrmgroupdelete 删除分组
// taobao.crm.group.delete
//
// 将该分组下的所有会员移除出该组，同时删除该分组。注：删除分组为异步任务，必须先调用taobao.crm.grouptask.check 确保涉及属性上没有任务。
func Taobaocrmgroupdelete(clt *core.SDKClient, req *crm.TaobaocrmgroupdeleteAPIRequest, session string) (*crm.TaobaocrmgroupdeleteAPIResponse, error) {
	var resp crm.TaobaocrmgroupdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
