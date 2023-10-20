package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoDataLogisticsCpDeliveryAgingPredictAPIResponse CP配送物流时效预测 API返回值
// cainiao.data.logistics.cp.delivery.aging.predict
//
// 时效和服务预期是商家发货时比较关注的信息，也是选择快递公司的一个重要参考（除去长期合作合同因素）。所以，在商家发货时给商家提供线路时效预估能帮助商家选择更好的快递公司，且对消费者来说也能整体提升体验。
//
// 日常时效是数值字符串
// 大促时效是数值区间字符串
// 方式1:
// 输入：发货省、市、区、详细地址，收货省、市、区、街道、详细地址，快递公司ID
// 输出：预估时效（小时数）
type CainiaoDataLogisticsCpDeliveryAgingPredictAPIResponse struct {
	model.CommonResponse
	CainiaoDataLogisticsCpDeliveryAgingPredictAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoDataLogisticsCpDeliveryAgingPredictAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoDataLogisticsCpDeliveryAgingPredictAPIResponseModel).Reset()
}

// CainiaoDataLogisticsCpDeliveryAgingPredictAPIResponseModel is CP配送物流时效预测 成功返回结果
type CainiaoDataLogisticsCpDeliveryAgingPredictAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_data_logistics_cp_delivery_aging_predict_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回数据值
	Data *DeliveryTimingDto `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoDataLogisticsCpDeliveryAgingPredictAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolCainiaoDataLogisticsCpDeliveryAgingPredictAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoDataLogisticsCpDeliveryAgingPredictAPIResponse)
	},
}

// GetCainiaoDataLogisticsCpDeliveryAgingPredictAPIResponse 从 sync.Pool 获取 CainiaoDataLogisticsCpDeliveryAgingPredictAPIResponse
func GetCainiaoDataLogisticsCpDeliveryAgingPredictAPIResponse() *CainiaoDataLogisticsCpDeliveryAgingPredictAPIResponse {
	return poolCainiaoDataLogisticsCpDeliveryAgingPredictAPIResponse.Get().(*CainiaoDataLogisticsCpDeliveryAgingPredictAPIResponse)
}

// ReleaseCainiaoDataLogisticsCpDeliveryAgingPredictAPIResponse 将 CainiaoDataLogisticsCpDeliveryAgingPredictAPIResponse 保存到 sync.Pool
func ReleaseCainiaoDataLogisticsCpDeliveryAgingPredictAPIResponse(v *CainiaoDataLogisticsCpDeliveryAgingPredictAPIResponse) {
	v.Reset()
	poolCainiaoDataLogisticsCpDeliveryAgingPredictAPIResponse.Put(v)
}
