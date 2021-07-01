package globalvirtual

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaGlobalVirtualSendcodeAPIRequest
国际虚拟商品发码服务 API请求
alibaba.global.virtual.sendcode

global virtual send code service */
type AlibabaGlobalVirtualSendcodeAPIRequest struct {
	model.Params
	// trade order id
	_tradeOrderLineId int64
	// code list
	_codeList []VirtualCertificateDo
}

// New
