package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressCollectResourceTmsDeleteAPIResponse 上门取退可揽范围删除 API返回值
// taobao.logistics.express.collect.resource.tms.delete
//
// 上门取退可揽范围删除
type TaobaoLogisticsExpressCollectResourceTmsDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsExpressCollectResourceTmsDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressCollectResourceTmsDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsExpressCollectResourceTmsDeleteAPIResponseModel).Reset()
}

// TaobaoLogisticsExpressCollectResourceTmsDeleteAPIResponseModel is 上门取退可揽范围删除 成功返回结果
type TaobaoLogisticsExpressCollectResourceTmsDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_collect_resource_tms_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	CollectResourceDeleteResponse *CollectResourceDeleteResponse `json:"collect_resource_delete_response,omitempty" xml:"collect_resource_delete_response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressCollectResourceTmsDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CollectResourceDeleteResponse = nil
}

var poolTaobaoLogisticsExpressCollectResourceTmsDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsExpressCollectResourceTmsDeleteAPIResponse)
	},
}

// GetTaobaoLogisticsExpressCollectResourceTmsDeleteAPIResponse 从 sync.Pool 获取 TaobaoLogisticsExpressCollectResourceTmsDeleteAPIResponse
func GetTaobaoLogisticsExpressCollectResourceTmsDeleteAPIResponse() *TaobaoLogisticsExpressCollectResourceTmsDeleteAPIResponse {
	return poolTaobaoLogisticsExpressCollectResourceTmsDeleteAPIResponse.Get().(*TaobaoLogisticsExpressCollectResourceTmsDeleteAPIResponse)
}

// ReleaseTaobaoLogisticsExpressCollectResourceTmsDeleteAPIResponse 将 TaobaoLogisticsExpressCollectResourceTmsDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsExpressCollectResourceTmsDeleteAPIResponse(v *TaobaoLogisticsExpressCollectResourceTmsDeleteAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsExpressCollectResourceTmsDeleteAPIResponse.Put(v)
}
