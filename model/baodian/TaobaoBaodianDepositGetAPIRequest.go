package baodian

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaodianDepositGetAPIRequest
宝点用户帐户查询（已迁移） API请求
taobao.baodian.deposit.get

查询用户宝点帐户信息及当前宝点价格 */
type TaobaoBaodianDepositGetAPIRequest struct {
	model.Params
}

// New
