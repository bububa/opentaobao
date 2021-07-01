package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaArgusUpdateredriskAPIRequest
更新红线价格 API请求
alibaba.argus.updateredrisk

商品健康中心新增红线价格规则 */
type AlibabaArgusUpdateredriskAPIRequest struct {
	model.Params
	// 红线价格参数
	_redRiskUpdateFactor *RedRiskUpdateFactor
}

// New
