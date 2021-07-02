package ascpchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIResponse 商家修改运单号 API返回值
// alibaba.ascp.uop.supplier.consignorder.notify.tms.change
//
// 供应商可以通过此接口，对出库回告上报的运单号进行修改，目前一次调用只能支持一个运单号的修改
type AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIResponse struct {
	model.CommonResponse
	AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIResponseModel
}

// AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIResponseModel is 商家修改运单号 成功返回结果
type AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_uop_supplier_consignorder_notify_tms_change_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	ModifyMailNoResponse *ResultWrapper `json:"modify_mail_no_response,omitempty" xml:"modify_mail_no_response,omitempty"`
}
