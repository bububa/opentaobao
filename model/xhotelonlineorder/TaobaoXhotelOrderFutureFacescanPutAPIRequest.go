package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelOrderFutureFacescanPutAPIRequest
未来酒店扫脸信息上传 API请求
taobao.xhotel.order.future.facescan.put

未来酒店扫脸信息上传服务，用于悉尔等厂商的扫脸设备对接 */
type TaobaoXhotelOrderFutureFacescanPutAPIRequest struct {
	model.Params
	// 扫脸参数
	_faceScanParam *FaceScanParam
}

// New
