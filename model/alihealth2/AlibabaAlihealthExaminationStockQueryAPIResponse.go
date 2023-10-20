package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationStockQueryAPIResponse 体检机构对接_体检套餐库存查询 API返回值
// alibaba.alihealth.examination.stock.query
//
// 体检机构对接_体检套餐库存查询
type AlibabaAlihealthExaminationStockQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationStockQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationStockQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthExaminationStockQueryAPIResponseModel).Reset()
}

// AlibabaAlihealthExaminationStockQueryAPIResponseModel is 体检机构对接_体检套餐库存查询 成功返回结果
type AlibabaAlihealthExaminationStockQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_stock_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 门店库存列表
	StorageList []Storage `json:"storage_list,omitempty" xml:"storage_list>storage,omitempty"`
	// 返回结果描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果编码
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 预约至少提前多少小时
	ReservationMinAheadHours string `json:"reservation_min_ahead_hours,omitempty" xml:"reservation_min_ahead_hours,omitempty"`
	// 是否支持分时能力
	TimeSharingEnable bool `json:"time_sharing_enable,omitempty" xml:"time_sharing_enable,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationStockQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.StorageList = m.StorageList[:0]
	m.Message = ""
	m.ResponseCode = ""
	m.ReservationMinAheadHours = ""
	m.TimeSharingEnable = false
}

var poolAlibabaAlihealthExaminationStockQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationStockQueryAPIResponse)
	},
}

// GetAlibabaAlihealthExaminationStockQueryAPIResponse 从 sync.Pool 获取 AlibabaAlihealthExaminationStockQueryAPIResponse
func GetAlibabaAlihealthExaminationStockQueryAPIResponse() *AlibabaAlihealthExaminationStockQueryAPIResponse {
	return poolAlibabaAlihealthExaminationStockQueryAPIResponse.Get().(*AlibabaAlihealthExaminationStockQueryAPIResponse)
}

// ReleaseAlibabaAlihealthExaminationStockQueryAPIResponse 将 AlibabaAlihealthExaminationStockQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthExaminationStockQueryAPIResponse(v *AlibabaAlihealthExaminationStockQueryAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthExaminationStockQueryAPIResponse.Put(v)
}
