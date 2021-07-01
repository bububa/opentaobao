package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest
逆向销退入库单到仓结果回告 API请求
alibaba.ascp.uop.supplier.reverseorder.instorage.result

ERP回告销退入库单到仓信息回告 */
type AlibabaAscpUopSupplierReverseorderInstorageResultAPIRequest struct {
	model.Params
	// 消退入库单结果请求
	_instorageResultRequest *Instorageresultrequest
}

// New
