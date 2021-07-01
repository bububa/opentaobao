package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclUpdategrantroletouserAPIRequest
修改用户到角色关系 API请求
alibaba.campus.acl.updategrantroletouser

修改用户到角色关系 */
type AlibabaCampusAclUpdategrantroletouserAPIRequest struct {
	model.Params
	// 公司id
	_companyId int64
	// 系统id
	_systemId string
	// 园区id
	_campusId int64
	// 用户账号
	_accountId string
	// 角色
	_role []RoleReq
	// 操作人id(不填默认appCode)
	_userId string
}

// New
