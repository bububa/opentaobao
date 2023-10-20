package jym

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponse 同步下架接口 API返回值
// alibaba.jym.item.external.goods.batch.synoffsale
//
// 同步下架接口
type AlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponse struct {
	model.CommonResponse
	AlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponseModel
}

// AlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponseModel is 同步下架接口 成功返回结果
type AlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_item_external_goods_batch_synoffsale_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
