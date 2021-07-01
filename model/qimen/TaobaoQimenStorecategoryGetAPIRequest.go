package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenStorecategoryGetAPIRequest
门店类目获取接口 API请求
taobao.qimen.storecategory.get

商家在ERP中调用该接口，获取门店类目 */
type TaobaoQimenStorecategoryGetAPIRequest struct {
	model.Params
	// 备注
	_remark string
}

// New
