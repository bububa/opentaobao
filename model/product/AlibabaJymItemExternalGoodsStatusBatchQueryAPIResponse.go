package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymitemexternalgoodsstatusbatchqueryAPIResponse 交易猫外部商家商品状态批量查询接口 API返回值
// alibaba.jym.item.external.goods.status.batch.query
//
// 供外部B端商家接入，请求查询商品状态，返回商品状态查询结果
type AlibabajymitemexternalgoodsstatusbatchqueryAPIResponse struct {
	model.CommonResponse
	AlibabajymitemexternalgoodsstatusbatchqueryAPIResponseModel
}

// AlibabajymitemexternalgoodsstatusbatchqueryAPIResponseModel is 交易猫外部商家商品状态批量查询接口 成功返回结果
type AlibabajymitemexternalgoodsstatusbatchqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_item_external_goods_status_batch_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态码
	StateCode string `json:"state_code,omitempty" xml:"state_code,omitempty"`
	// 错误信息
	ExtraErrMsg string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
	// 批量查询商品状态接口返回
	Result *BatchGoodsStatusResultDto `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Succeed bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
}
