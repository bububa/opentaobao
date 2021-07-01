package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclQueryallemppermiitemAPIRequest
查询员工全部权限(包括角色下面的权限) API请求
alibaba.campus.acl.queryallemppermiitem

查询员工全部权限(包括角色下面的权限) */
type AlibabaCampusAclQueryallemppermiitemAPIRequest struct {
	model.Params
	// 公司id不填默认SYS_000
	_companyId int64
	// 系统id
	_systemId string
	// 园区id
	_campusId int64
	// 用户账号
	_accountId string
	// 每页多少条
	_page int64
	// 每页记录数
	_pageSize int64
}

// New
