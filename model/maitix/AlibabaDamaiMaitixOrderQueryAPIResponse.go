package maitix

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimaitixorderqueryAPIResponse 大麦-查询分销单 API返回值
// alibaba.damai.maitix.order.query
//
// 查询分销单
type AlibabadamaimaitixorderqueryAPIResponse struct {
	model.CommonResponse
	AlibabadamaimaitixorderqueryAPIResponseModel
}

// AlibabadamaimaitixorderqueryAPIResponseModel is 大麦-查询分销单 成功返回结果
type AlibabadamaimaitixorderqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果集
	Result *MxResult `json:"result,omitempty" xml:"result,omitempty"`
}
