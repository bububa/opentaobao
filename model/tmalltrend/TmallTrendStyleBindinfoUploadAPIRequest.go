package tmalltrend

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTrendStyleBindinfoUploadAPIRequest
趋势词&款式绑定信息同步API API请求
tmall.trend.style.bindinfo.upload

趋势词&款式(服饰行业)绑定信息同步至平台 */
type TmallTrendStyleBindinfoUploadAPIRequest struct {
	model.Params
	// 趋势词&款式绑定信息列表，一次最多1000条
	_trendStyleBindInfoBoList []TrendStyleBindInfoBO
}

// New
