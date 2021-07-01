package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmGroupMoveAPIRequest
分组移动 API请求
taobao.crm.group.move

将一个分组下的所有会员移动到另一个分组，会员从原分组中删除<br/>注：移动属性为异步任务建议先调用taobao.crm.grouptask.check 确保涉及属性上没有任务。 */
type TaobaoCrmGroupMoveAPIRequest struct {
	model.Params
	// 需要移动的分组
	_fromGroupId int64
	// 目的分组
	_toGroupId int64
}

// New
