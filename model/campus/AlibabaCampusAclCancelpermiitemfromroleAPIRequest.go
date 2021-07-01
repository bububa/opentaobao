package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclCancelpermiitemfromroleAPIRequest
取消角色和权限之间的关系 API请求
alibaba.campus.acl.cancelpermiitemfromrole

取消角色和权限之间的关系 */
type AlibabaCampusAclCancelpermiitemfromroleAPIRequest struct {
	model.Params
	// 公司ID
	_companyId int64
	// 系统id
	_systemId string
	// 园区id
	_campusId int64
	// 系统自动生成
	_param1 *RoleReq
	// 系统自动生成
	_param2 []PermissionReq
	// 操作人id(不填默认appCode)
	_userId string
}

// New
