package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiEcSearchProjectSearchAPIRequest
大麦电商对外搜索服务 API请求
alibaba.damai.ec.search.project.search

大麦电商对外搜索服务 */
type AlibabaDamaiEcSearchProjectSearchAPIRequest struct {
	model.Params
	// 入参对象
	_param *TopSearchProjectParam
}

// New
