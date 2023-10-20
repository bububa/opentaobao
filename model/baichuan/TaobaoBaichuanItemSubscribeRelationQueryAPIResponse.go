package baichuan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanitemsubscriberelationqueryAPIResponse 查询单个订阅关系 API返回值
// taobao.baichuan.item.subscribe.relation.query
//
// 查询单个订阅关系
type TaobaobaichuanitemsubscriberelationqueryAPIResponse struct {
	model.CommonResponse
	TaobaobaichuanitemsubscriberelationqueryAPIResponseModel
}

// TaobaobaichuanitemsubscriberelationqueryAPIResponseModel is 查询单个订阅关系 成功返回结果
type TaobaobaichuanitemsubscriberelationqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_item_subscribe_relation_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaobaichuanitemsubscriberelationqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
