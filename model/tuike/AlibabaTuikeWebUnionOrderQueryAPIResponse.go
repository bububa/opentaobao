package tuike

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTuikeWebUnionOrderQueryAPIResponse 推客网盟合作抽佣订单查询接口 API返回值
// alibaba.tuike.web.union.order.query
//
// 推客网盟合作抽佣订单查询接口
type AlibabaTuikeWebUnionOrderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaTuikeWebUnionOrderQueryAPIResponseModel
}

// AlibabaTuikeWebUnionOrderQueryAPIResponseModel is 推客网盟合作抽佣订单查询接口 成功返回结果
type AlibabaTuikeWebUnionOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tuike_web_union_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	List []AlibabaTuikeWebUnionOrderQueryResult `json:"list,omitempty" xml:"list>alibaba_tuike_web_union_order_query_result,omitempty"`
}
