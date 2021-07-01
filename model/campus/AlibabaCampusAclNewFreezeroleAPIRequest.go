package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclNewFreezeroleAPIRequest
冻结角色 API请求
alibaba.campus.acl.new.freezerole

冻结角色 */
type AlibabaCampusAclNewFreezeroleAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 角色主键id
	_roleId int64
}

// New
