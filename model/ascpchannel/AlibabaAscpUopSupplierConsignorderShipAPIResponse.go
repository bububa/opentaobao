package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopSupplierConsignorderShipAPIResponse 履约单商家仓发货结果回传服务 API返回值
// alibaba.ascp.uop.supplier.consignorder.ship
//
// ERP通过该接口通知商家仓声明销售订单出库信息,支持履约单纬度全部发货的回传（目前不支持分批回传)
type AlibabaAscpUopSupplierConsignorderShipAPIResponse struct {
	model.CommonResponse
	AlibabaAscpUopSupplierConsignorderShipAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpUopSupplierConsignorderShipAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpUopSupplierConsignorderShipAPIResponseModel).Reset()
}

// AlibabaAscpUopSupplierConsignorderShipAPIResponseModel is 履约单商家仓发货结果回传服务 成功返回结果
type AlibabaAscpUopSupplierConsignorderShipAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_uop_supplier_consignorder_ship_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	ConsignorderShipResponse *ResultWrapper `json:"consignorder_ship_response,omitempty" xml:"consignorder_ship_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpUopSupplierConsignorderShipAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ConsignorderShipResponse = nil
}

var poolAlibabaAscpUopSupplierConsignorderShipAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpUopSupplierConsignorderShipAPIResponse)
	},
}

// GetAlibabaAscpUopSupplierConsignorderShipAPIResponse 从 sync.Pool 获取 AlibabaAscpUopSupplierConsignorderShipAPIResponse
func GetAlibabaAscpUopSupplierConsignorderShipAPIResponse() *AlibabaAscpUopSupplierConsignorderShipAPIResponse {
	return poolAlibabaAscpUopSupplierConsignorderShipAPIResponse.Get().(*AlibabaAscpUopSupplierConsignorderShipAPIResponse)
}

// ReleaseAlibabaAscpUopSupplierConsignorderShipAPIResponse 将 AlibabaAscpUopSupplierConsignorderShipAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpUopSupplierConsignorderShipAPIResponse(v *AlibabaAscpUopSupplierConsignorderShipAPIResponse) {
	v.Reset()
	poolAlibabaAscpUopSupplierConsignorderShipAPIResponse.Put(v)
}
