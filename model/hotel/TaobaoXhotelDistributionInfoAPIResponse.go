package hotel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhoteldistributioninfoAPIResponse 飞猪分销通用酒店标准信息接口 API返回值
// taobao.xhotel.distribution.info
//
// 飞猪分销通用酒店标准信息接口
type TaobaoxhoteldistributioninfoAPIResponse struct {
	model.CommonResponse
	TaobaoxhoteldistributioninfoAPIResponseModel
}

// TaobaoxhoteldistributioninfoAPIResponseModel is 飞猪分销通用酒店标准信息接口 成功返回结果
type TaobaoxhoteldistributioninfoAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_distribution_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 标准酒店信息
	Hotels []ShotelInfoObject `json:"hotels,omitempty" xml:"hotels>shotel_info_object,omitempty"`
	// 酒店总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
