package cainiaolocker

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoEndpointLockerTopOrderNoticeAPIResponse
手动触发发短信 API返回值
cainiao.endpoint.locker.top.order.notice

合作公司对订单手动触发短信，有次数限制 */
type CainiaoEndpointLockerTopOrderNoticeAPIResponse struct {
	model.CommonResponse
	CainiaoEndpointLockerTopOrderNoticeAPIResponseModel
}

// CainiaoEndpointLockerTopOrderNoticeAPIResponseModel is 手动触发发短信 成功返回结果
type CainiaoEndpointLockerTopOrderNoticeAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_endpoint_locker_top_order_notice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *SingleResult `json:"result,omitempty" xml:"result,omitempty"`
}
