package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpUopCnReverseWarehouseorderUpdateAPIRequest
供应链中台逆向入库单修改服务 API请求
alibaba.ascp.uop.cn.reverse.warehouseorder.update

供应链中台逆向入库单修改服务 */
type AlibabaAscpUopCnReverseWarehouseorderUpdateAPIRequest struct {
	model.Params
	// 逆向入库单号
	_orderCode string
	// 是否已经退款
	_refunded bool
	// 退款原因
	_refundReason string
}

// New
