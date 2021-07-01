package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenSupplierSynchronizeAPIRequest
供应商同步接口 API请求
taobao.qimen.supplier.synchronize

这个接口用来同步供应商信息 */
type TaobaoQimenSupplierSynchronizeAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenSupplierSynchronizeRequest
}

// New
