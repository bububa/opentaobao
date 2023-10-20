package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsWmsOrderMaterialSyncAPIResponse 仓服务商订单包材耗材信息同步 API返回值
// taobao.logistics.wms.order.material.sync
//
// 仓服务商订单包材耗材信息同步
type TaobaoLogisticsWmsOrderMaterialSyncAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsWmsOrderMaterialSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsWmsOrderMaterialSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsWmsOrderMaterialSyncAPIResponseModel).Reset()
}

// TaobaoLogisticsWmsOrderMaterialSyncAPIResponseModel is 仓服务商订单包材耗材信息同步 成功返回结果
type TaobaoLogisticsWmsOrderMaterialSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_wms_order_material_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码描述
	BizErrorMessage string `json:"biz_error_message,omitempty" xml:"biz_error_message,omitempty"`
	// 错误码标识
	BizErrorCode string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 校验成功或者异常
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
	// 是否可重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsWmsOrderMaterialSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BizErrorMessage = ""
	m.BizErrorCode = ""
	m.Suc = false
	m.Retry = false
}

var poolTaobaoLogisticsWmsOrderMaterialSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsWmsOrderMaterialSyncAPIResponse)
	},
}

// GetTaobaoLogisticsWmsOrderMaterialSyncAPIResponse 从 sync.Pool 获取 TaobaoLogisticsWmsOrderMaterialSyncAPIResponse
func GetTaobaoLogisticsWmsOrderMaterialSyncAPIResponse() *TaobaoLogisticsWmsOrderMaterialSyncAPIResponse {
	return poolTaobaoLogisticsWmsOrderMaterialSyncAPIResponse.Get().(*TaobaoLogisticsWmsOrderMaterialSyncAPIResponse)
}

// ReleaseTaobaoLogisticsWmsOrderMaterialSyncAPIResponse 将 TaobaoLogisticsWmsOrderMaterialSyncAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsWmsOrderMaterialSyncAPIResponse(v *TaobaoLogisticsWmsOrderMaterialSyncAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsWmsOrderMaterialSyncAPIResponse.Put(v)
}
