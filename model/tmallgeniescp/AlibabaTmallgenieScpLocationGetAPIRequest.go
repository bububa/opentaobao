package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpLocationGetAPIRequest
2-IBP查询CDC和RDC数据接口 API请求
alibaba.tmallgenie.scp.location.get

天猫精灵供应链-计划域-IBP查询CDC和RDC数据接口 */
type AlibabaTmallgenieScpLocationGetAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// New
