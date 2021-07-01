package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpUopSupplierWaybillQueryAPIRequest
ERP调用打印面单取号接口 API请求
alibaba.ascp.uop.supplier.waybill.query

ERP调用打印面单取号接口 */
type AlibabaAscpUopSupplierWaybillQueryAPIRequest struct {
	model.Params
	// 查询面单请求参数
	_waybillQueryRequest *Waybillqueryrequest
}

// New
