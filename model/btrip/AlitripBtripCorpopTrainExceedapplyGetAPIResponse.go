package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcorpoptrainexceedapplygetAPIResponse 商旅火车票第三方超标审批单搜索接口 API返回值
// alitrip.btrip.corpop.train.exceedapply.get
//
// 商旅火车票第三方超标审批单搜索接口
type AlitripbtripcorpoptrainexceedapplygetAPIResponse struct {
	model.CommonResponse
	AlitripbtripcorpoptrainexceedapplygetAPIResponseModel
}

// AlitripbtripcorpoptrainexceedapplygetAPIResponseModel is 商旅火车票第三方超标审批单搜索接口 成功返回结果
type AlitripbtripcorpoptrainexceedapplygetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_train_exceedapply_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
