package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderstorecollectqueryAPIResponse 全渠道门店自提根据核销码查询订单 API返回值
// taobao.omniorder.storecollect.query
//
// 全渠道门店自提根据核销码查询订单
type TaobaoomniorderstorecollectqueryAPIResponse struct {
	model.CommonResponse
	TaobaoomniorderstorecollectqueryAPIResponseModel
}

// TaobaoomniorderstorecollectqueryAPIResponseModel is 全渠道门店自提根据核销码查询订单 成功返回结果
type TaobaoomniorderstorecollectqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_storecollect_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoomniorderstorecollectqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
