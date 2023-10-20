package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkumsordergetAPIResponse 查询店仓作业单据清单 （库存对账辅助）-回流单 API返回值
// alibaba.wdk.ums.order.get
//
// 查询店仓作业单据清单 （库存对账辅助）-回流单
type AlibabawdkumsordergetAPIResponse struct {
	model.CommonResponse
	AlibabawdkumsordergetAPIResponseModel
}

// AlibabawdkumsordergetAPIResponseModel is 查询店仓作业单据清单 （库存对账辅助）-回流单 成功返回结果
type AlibabawdkumsordergetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_ums_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *UtmsPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
