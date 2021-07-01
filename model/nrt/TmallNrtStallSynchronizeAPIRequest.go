package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrtStallSynchronizeAPIRequest
摊位信息同步 API请求
tmall.nrt.stall.synchronize

摊位信息同步 */
type TmallNrtStallSynchronizeAPIRequest struct {
	model.Params
	// 参数对象
	_stall *NrtStoreDto
}

// New
