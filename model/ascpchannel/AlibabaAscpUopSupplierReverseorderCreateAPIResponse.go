package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpUopSupplierReverseorderCreateAPIResponse
商家ERP发起创建销退单服务 API返回值
alibaba.ascp.uop.supplier.reverseorder.create

商家在收到消费者实物退货后，在ERP发起创建销退单服务 */
type AlibabaAscpUopSupplierReverseorderCreateAPIResponse struct {
	model.CommonResponse
	AlibabaAscpUopSupplierReverseorderCreateAPIResponseModel
}

// AlibabaAscpUopSupplierReverseorderCreateAPIResponseModel is 商家ERP发起创建销退单服务 成功返回结果
type AlibabaAscpUopSupplierReverseorderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_uop_supplier_reverseorder_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
