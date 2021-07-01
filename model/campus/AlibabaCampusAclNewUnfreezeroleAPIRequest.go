package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclNewUnfreezeroleAPIRequest
解冻角色 API请求
alibaba.campus.acl.new.unfreezerole

解冻角色 */
type AlibabaCampusAclNewUnfreezeroleAPIRequest struct {
	model.Params
	// 系统参数
	_workbenchcontext *WorkBenchContext
	// 角色主键id
	_roleId int64
}

// New
