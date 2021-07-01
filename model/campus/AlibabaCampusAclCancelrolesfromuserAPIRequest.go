package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclCancelrolesfromuserAPIRequest
撤销用户授予的角色 API请求
alibaba.campus.acl.cancelrolesfromuser

撤销用户授予的角色 */
type AlibabaCampusAclCancelrolesfromuserAPIRequest struct {
	model.Params
	// 系统自动生成
	_role []RoleReq
	// 公司id
	_companyId int64
	// 系统id
	_systemId string
	// 园区id
	_campusId int64
	// 用户账号
	_accountId string
	// 操作人id(不填默认appCode)
	_userId string
}

// New
