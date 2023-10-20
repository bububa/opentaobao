package crm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/crm"
)

// Taobaocrmgroupmove 分组移动
// taobao.crm.group.move
//
// 将一个分组下的所有会员移动到另一个分组，会员从原分组中删除&lt;br/&gt;注：移动属性为异步任务建议先调用taobao.crm.grouptask.check 确保涉及属性上没有任务。
func Taobaocrmgroupmove(clt *core.SDKClient, req *crm.TaobaocrmgroupmoveAPIRequest, session string) (*crm.TaobaocrmgroupmoveAPIResponse, error) {
	var resp crm.TaobaocrmgroupmoveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
