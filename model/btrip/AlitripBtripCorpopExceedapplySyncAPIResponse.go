package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopExceedapplySyncAPIResponse 第三方超标审批结果回传 API返回值
// alitrip.btrip.corpop.exceedapply.sync
//
// 第三方审批单推送到企业后，企业审批结束，将审批结果回传给阿里商旅
type AlitripBtripCorpopExceedapplySyncAPIResponse struct {
	model.CommonResponse
	AlitripBtripCorpopExceedapplySyncAPIResponseModel
}

// AlitripBtripCorpopExceedapplySyncAPIResponseModel is 第三方超标审批结果回传 成功返回结果
type AlitripBtripCorpopExceedapplySyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_exceedapply_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *BtripApplyResult `json:"result,omitempty" xml:"result,omitempty"`
}
