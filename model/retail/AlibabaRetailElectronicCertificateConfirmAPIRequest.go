package retail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailElectronicCertificateConfirmAPIRequest
确认核销接口 API请求
alibaba.retail.electronic.certificate.confirm

确认核销接口 */
type AlibabaRetailElectronicCertificateConfirmAPIRequest struct {
	model.Params
	// 核销码
	_code int64
	// 商品ID
	_itemId int64
	// 设备ID
	_deviceId string
}

// New
