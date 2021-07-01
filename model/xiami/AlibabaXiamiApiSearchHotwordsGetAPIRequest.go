package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaXiamiApiSearchHotwordsGetAPIRequest
搜索热词 API请求
alibaba.xiami.api.search.hotwords.get

搜索热词 */
type AlibabaXiamiApiSearchHotwordsGetAPIRequest struct {
	model.Params
	// 数量
	_limit int64
}

// New
