package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclInsertroleAPIRequest
新增角色 API请求
alibaba.campus.acl.insertrole

新增角色 */
type AlibabaCampusAclInsertroleAPIRequest struct {
	model.Params
	// 公司id,不填统一为SYS_000
	_companyId int64
	// 系统id
	_systemId string
	// 园区id
	_campusId int64
	// 角色描述
	_roleDesc string
	// 角色名称
	_roleName string
	// 默认人员角色.还有device设备角色类型
	_roleType string
	// 角色唯一ID,统一ROLE_开头,不填默认生成ROLE_UUID(32位随机数)
	_roleId string
	// 操作人id(不填默认appCode)
	_userId string
}

// New
