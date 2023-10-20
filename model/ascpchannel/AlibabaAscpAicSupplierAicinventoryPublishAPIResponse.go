package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpAicSupplierAicinventoryPublishAPIResponse 商家仓操作aic库存发布服务 API返回值
// alibaba.ascp.aic.supplier.aicinventory.publish
//
// 商家调用这个接口来发布增加库存数据
type AlibabaAscpAicSupplierAicinventoryPublishAPIResponse struct {
	model.CommonResponse
	AlibabaAscpAicSupplierAicinventoryPublishAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpAicSupplierAicinventoryPublishAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpAicSupplierAicinventoryPublishAPIResponseModel).Reset()
}

// AlibabaAscpAicSupplierAicinventoryPublishAPIResponseModel is 商家仓操作aic库存发布服务 成功返回结果
type AlibabaAscpAicSupplierAicinventoryPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_aic_supplier_aicinventory_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	AicInventoryPublishResponse *ResultWrapper `json:"aic_inventory_publish_response,omitempty" xml:"aic_inventory_publish_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpAicSupplierAicinventoryPublishAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AicInventoryPublishResponse = nil
}

var poolAlibabaAscpAicSupplierAicinventoryPublishAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpAicSupplierAicinventoryPublishAPIResponse)
	},
}

// GetAlibabaAscpAicSupplierAicinventoryPublishAPIResponse 从 sync.Pool 获取 AlibabaAscpAicSupplierAicinventoryPublishAPIResponse
func GetAlibabaAscpAicSupplierAicinventoryPublishAPIResponse() *AlibabaAscpAicSupplierAicinventoryPublishAPIResponse {
	return poolAlibabaAscpAicSupplierAicinventoryPublishAPIResponse.Get().(*AlibabaAscpAicSupplierAicinventoryPublishAPIResponse)
}

// ReleaseAlibabaAscpAicSupplierAicinventoryPublishAPIResponse 将 AlibabaAscpAicSupplierAicinventoryPublishAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpAicSupplierAicinventoryPublishAPIResponse(v *AlibabaAscpAicSupplierAicinventoryPublishAPIResponse) {
	v.Reset()
	poolAlibabaAscpAicSupplierAicinventoryPublishAPIResponse.Put(v)
}
