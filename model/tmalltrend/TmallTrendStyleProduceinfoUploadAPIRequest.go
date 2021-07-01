package tmalltrend

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallTrendStyleProduceinfoUploadAPIRequest
款式生产信息同步API API请求
tmall.trend.style.produceinfo.upload

款式生产信息同步至平台 */
type TmallTrendStyleProduceinfoUploadAPIRequest struct {
	model.Params
	// 款式生产信息列表，单次同步最对1000条
	_styleProduceInfoBoList []StyleProduceInfoBO
}

// New
