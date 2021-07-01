package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclGrantpermiitemstouserAPIRequest
给人直接授权 API请求
alibaba.campus.acl.grantpermiitemstouser

给人直接授权 */
type AlibabaCampusAclGrantpermiitemstouserAPIRequest struct {
	model.Params
	// 公司id不填统一默认生成SYS_000
	_companyId int64
	// 系统id
	_systemId string
	// 园区id
	_campusId int64
	// 用户id
	_empId string
	// 权限
	_priv []PermissionReq
	// 操作人id(不填默认appCode)
	_userId string
}

// New
