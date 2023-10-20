package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationstockqueryAPIResponse 体检机构对接_体检套餐库存查询 API返回值
// alibaba.alihealth.examination.stock.query
//
// 体检机构对接_体检套餐库存查询
type AlibabaalihealthexaminationstockqueryAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthexaminationstockqueryAPIResponseModel
}

// AlibabaalihealthexaminationstockqueryAPIResponseModel is 体检机构对接_体检套餐库存查询 成功返回结果
type AlibabaalihealthexaminationstockqueryAPIResponseModel struct {
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
