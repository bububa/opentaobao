package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpUopSupplierReverseorderCreateAPIRequest
商家ERP发起创建销退单服务 API请求
alibaba.ascp.uop.supplier.reverseorder.create

商家在收到消费者实物退货后，在ERP发起创建销退单服务 */
type AlibabaAscpUopSupplierReverseorderCreateAPIRequest struct {
	model.Params
	// 逆向销退单创建请求
	_reverseCreateRequest *ReverseCreateRequest
}

// New
