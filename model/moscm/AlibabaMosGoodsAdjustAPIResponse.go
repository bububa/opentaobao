package moscm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosgoodsadjustAPIResponse 调整库存 API返回值
// alibaba.mos.goods.adjust
//
// 库存调整接口
type AlibabamosgoodsadjustAPIResponse struct {
	model.CommonResponse
	AlibabamosgoodsadjustAPIResponseModel
}

// AlibabamosgoodsadjustAPIResponseModel is 调整库存 成功返回结果
type AlibabamosgoodsadjustAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_goods_adjust_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 库存调整单号
	Result *AlibabamosgoodsadjustResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
