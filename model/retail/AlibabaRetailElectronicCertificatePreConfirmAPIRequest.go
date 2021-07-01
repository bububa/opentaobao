package retail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailElectronicCertificatePreConfirmAPIRequest
贩卖机开始核销接口 API请求
alibaba.retail.electronic.certificate.pre.confirm

零售终端贩卖机开始核销接口,返回待领的商品ID */
type AlibabaRetailElectronicCertificatePreConfirmAPIRequest struct {
	model.Params
	// 设备ID
	_deviceId string
	// 核销码
	_code int64
}

// New
