package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclNewRemoveroleAPIRequest
删除角色 API请求
alibaba.campus.acl.new.removerole

删除角色 */
type AlibabaCampusAclNewRemoveroleAPIRequest struct {
	model.Params
	// 系统入参
	_param0 *WorkBenchContext
	// 角色主键id
	_roleId int64
}

// New
