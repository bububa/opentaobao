package c2m

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSebpOrganizationGetinviteinfoAPIRequest
淘小铺机构上下级关系 API请求
taobao.sebp.organization.getinviteinfo

机构人员获取机构上下级关系信息 */
type TaobaoSebpOrganizationGetinviteinfoAPIRequest struct {
	model.Params
	// null-请求所有，20200616-请求2020年6月16号的变更信息
	_modifyDate string
	// 第几页
	_pageNum int64
}

// New
