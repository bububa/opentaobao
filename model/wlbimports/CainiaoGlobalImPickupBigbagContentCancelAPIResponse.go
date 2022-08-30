package wlbimports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalImPickupBigbagContentCancelAPIResponse 进口大包取消 API返回值
// cainiao.global.im.pickup.bigbag.content.cancel
//
// 进口大包取消
type CainiaoGlobalImPickupBigbagContentCancelAPIResponse struct {
	model.CommonResponse
	CainiaoGlobalImPickupBigbagContentCancelAPIResponseModel
}

// CainiaoGlobalImPickupBigbagContentCancelAPIResponseModel is 进口大包取消 成功返回结果
type CainiaoGlobalImPickupBigbagContentCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_im_pickup_bigbag_content_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应体
	HsfResult *HsfResult `json:"hsf_result,omitempty" xml:"hsf_result,omitempty"`
}
