package tmallnr

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallnrfulfilllogisticsqueryAPIResponse 定时送和极速达配送物流信息查询 API返回值
// tmall.nr.fulfill.logistics.query
//
// 发布定时送&amp;极速达物流信息查询服务
type TmallnrfulfilllogisticsqueryAPIResponse struct {
	model.CommonResponse
	TmallnrfulfilllogisticsqueryAPIResponseModel
}

// TmallnrfulfilllogisticsqueryAPIResponseModel is 定时送和极速达配送物流信息查询 成功返回结果
type TmallnrfulfilllogisticsqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nr_fulfill_logistics_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *NrResult `json:"result,omitempty" xml:"result,omitempty"`
}
