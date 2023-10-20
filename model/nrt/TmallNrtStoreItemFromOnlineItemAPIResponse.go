package nrt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtstoreitemfromonlineitemAPIResponse 基于新模型商品id查询摊位子品id API返回值
// tmall.nrt.store.item.from.online.item
//
// 基于新模型商品id查询摊位子品id
type TmallnrtstoreitemfromonlineitemAPIResponse struct {
	model.CommonResponse
	TmallnrtstoreitemfromonlineitemAPIResponseModel
}

// TmallnrtstoreitemfromonlineitemAPIResponseModel is 基于新模型商品id查询摊位子品id 成功返回结果
type TmallnrtstoreitemfromonlineitemAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_store_item_from_online_item_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *TmallnrtstoreitemfromonlineitemResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
