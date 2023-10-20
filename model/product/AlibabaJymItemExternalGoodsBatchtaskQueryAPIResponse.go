package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymitemexternalgoodsbatchtaskqueryAPIResponse 交易猫外部商家查询商品批量任务接口 API返回值
// alibaba.jym.item.external.goods.batchtask.query
//
// 供外部B端商家接入，请求查询商品批量任务，返回商品批量任务查询结果
type AlibabajymitemexternalgoodsbatchtaskqueryAPIResponse struct {
	model.CommonResponse
	AlibabajymitemexternalgoodsbatchtaskqueryAPIResponseModel
}

// AlibabajymitemexternalgoodsbatchtaskqueryAPIResponseModel is 交易猫外部商家查询商品批量任务接口 成功返回结果
type AlibabajymitemexternalgoodsbatchtaskqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_item_external_goods_batchtask_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态码
	StateCode string `json:"state_code,omitempty" xml:"state_code,omitempty"`
	// 错误信息
	ExtraErrMsg string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
	// 查询商品批量任务接口返回
	Result *GoodsBatchTaskResultDto `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Succeed bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
}
