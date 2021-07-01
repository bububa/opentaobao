package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest
销退单上传 API请求
taobao.rdc.aligenius.warehouse.reverse.uploading

主要用于商家上传仓库销退单信息 */
type TaobaoRdcAligeniusWarehouseReverseUploadingAPIRequest struct {
	model.Params
	// 参数
	_param0 *WarehouseReverseUploadingDto
}

// New
