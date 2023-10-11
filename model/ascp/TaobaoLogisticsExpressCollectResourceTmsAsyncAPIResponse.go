package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressCollectResourceTmsAsyncAPIResponse 配服务商揽收能力同步接口 API返回值
// taobao.logistics.express.collect.resource.tms.async
//
// 配服务商揽收能力同步接口
type TaobaoLogisticsExpressCollectResourceTmsAsyncAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsExpressCollectResourceTmsAsyncAPIResponseModel
}

// TaobaoLogisticsExpressCollectResourceTmsAsyncAPIResponseModel is 配服务商揽收能力同步接口 成功返回结果
type TaobaoLogisticsExpressCollectResourceTmsAsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_collect_resource_tms_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	CollectResourceResponse *CollectResourceResponse `json:"collect_resource_response,omitempty" xml:"collect_resource_response,omitempty"`
}
