package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoProductcatAddAPIRequest
新增产品线 API请求
taobao.fenxiao.productcat.add

新增产品线 */
type TaobaoFenxiaoProductcatAddAPIRequest struct {
	model.Params
	// 产品线名称
	_name string
	// 最低零售价比例，注意：100.00%，则输入为10000
	_retailLowPercent int64
	// 最高零售价比例，注意：100.00%，则输入为10000
	_retailHighPercent int64
	// 代销默认采购价比例，注意：100.00%，则输入为10000
	_agentCostPercent int64
	// 经销默认采购价比例，注意：100.00%，则输入为10000
	_dealerCostPercent int64
}

// New
