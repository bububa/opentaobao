package moscm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosGoodsAdjustAPIResponse 调整库存 API返回值
// alibaba.mos.goods.adjust
//
// 库存调整接口
type AlibabaMosGoodsAdjustAPIResponse struct {
	model.CommonResponse
	AlibabaMosGoodsAdjustAPIResponseModel
}

// AlibabaMosGoodsAdjustAPIResponseModel is 调整库存 成功返回结果
type AlibabaMosGoodsAdjustAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_goods_adjust_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 库存调整单号
	Result *AlibabaMosGoodsAdjustResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
