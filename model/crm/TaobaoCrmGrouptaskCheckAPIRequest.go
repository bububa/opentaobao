package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmGrouptaskCheckAPIRequest
查询分组任务是否完成 API请求
taobao.crm.grouptask.check

检查一个分组上是否有异步任务,异步任务包括1.将一个分组下的所有用户添加到另外一个分组2.将一个分组下的所有用户移动到另外一个分组3.删除某个分组<br/>若分组上有任务则该属性不能被操作。 */
type TaobaoCrmGrouptaskCheckAPIRequest struct {
	model.Params
	// 分组id
	_groupId int64
}

// New
