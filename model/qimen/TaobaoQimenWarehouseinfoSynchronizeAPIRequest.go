package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenWarehouseinfoSynchronizeAPIRequest
仓库同步接口 API请求
taobao.qimen.warehouseinfo.synchronize

仓库同步接口 */
type TaobaoQimenWarehouseinfoSynchronizeAPIRequest struct {
	model.Params
	// 请求报文
	_request *TaobaoQimenWarehouseinfoSynchronizeRequest
}

// New
