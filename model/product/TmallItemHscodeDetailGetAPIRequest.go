package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallItemHscodeDetailGetAPIRequest
通过hscode获取计量单位 API请求
tmall.item.hscode.detail.get

通过hscode获取计量单位和销售单位 */
type TmallItemHscodeDetailGetAPIRequest struct {
	model.Params
	// hscode
	_hscode string
}

// New
