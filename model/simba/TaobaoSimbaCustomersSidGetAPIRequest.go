package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaCustomersSidGetAPIRequest
查看功能权限 API请求
taobao.simba.customers.sid.get

查询用户是否拥有某个功能权限 */
type TaobaoSimbaCustomersSidGetAPIRequest struct {
	model.Params
}

// New
