package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponse 供应链中台逆向入库单修改服务 API返回值
// alibaba.ascp.uop.cn.reverse.warehouseorder.update
//
// 供应链中台逆向入库单修改服务
type AlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponseModel).Reset()
}

// AlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponseModel is 供应链中台逆向入库单修改服务 成功返回结果
type AlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_uop_cn_reverse_warehouseorder_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	UpdateReverseWarehouseOrderResponse *ResultWrapper `json:"update_reverse_warehouse_order_response,omitempty" xml:"update_reverse_warehouse_order_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UpdateReverseWarehouseOrderResponse = nil
}

var poolAlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponse)
	},
}

// GetAlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponse 从 sync.Pool 获取 AlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponse
func GetAlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponse() *AlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponse {
	return poolAlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponse.Get().(*AlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponse)
}

// ReleaseAlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponse 将 AlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponse(v *AlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponse) {
	v.Reset()
	poolAlibabaAscpUopCnReverseWarehouseorderUpdateAPIResponse.Put(v)
}
