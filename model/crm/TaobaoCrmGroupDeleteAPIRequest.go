package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmGroupDeleteAPIRequest
删除分组 API请求
taobao.crm.group.delete

将该分组下的所有会员移除出该组，同时删除该分组。注：删除分组为异步任务，必须先调用taobao.crm.grouptask.check 确保涉及属性上没有任务。 */
type TaobaoCrmGroupDeleteAPIRequest struct {
	model.Params
	// 要删除的分组id
	_groupId int64
}

// New
