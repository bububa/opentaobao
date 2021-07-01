package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmGroupUpdateAPIRequest
修改一个已经存在的分组 API请求
taobao.crm.group.update

修改一个已经存在的分组，接口返回分组的修改是否成功 */
type TaobaoCrmGroupUpdateAPIRequest struct {
	model.Params
	// 分组的id
	_groupId int64
	// 新的分组名，分组名称不能包含|或者：
	_newGroupName string
}

// New
