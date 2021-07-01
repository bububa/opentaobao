package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbImportsVasIdentityResultAPIRequest
集货鉴定结果 API请求
taobao.wlb.imports.vas.identity.result

集货鉴定结果查询 */
type TaobaoWlbImportsVasIdentityResultAPIRequest struct {
	model.Params
	// 物流订单编号
	_lgOrderCode string
}

// New
