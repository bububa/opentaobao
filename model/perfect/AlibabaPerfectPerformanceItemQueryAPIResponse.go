package perfect

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaperfectperformanceitemqueryAPIResponse 商品完美履约信息查询 API返回值
// alibaba.perfect.performance.item.query
//
// 同城零售商品完美履约信息查询
type AlibabaperfectperformanceitemqueryAPIResponse struct {
	model.CommonResponse
	AlibabaperfectperformanceitemqueryAPIResponseModel
}

// AlibabaperfectperformanceitemqueryAPIResponseModel is 商品完美履约信息查询 成功返回结果
type AlibabaperfectperformanceitemqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_perfect_performance_item_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的数据实体
	Data *ItemPerfectPerformanceQueryResp `json:"data,omitempty" xml:"data,omitempty"`
}
