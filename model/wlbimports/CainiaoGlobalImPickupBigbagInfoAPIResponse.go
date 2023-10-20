package wlbimports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalImPickupBigbagInfoAPIResponse 大包状态查询 API返回值
// cainiao.global.im.pickup.bigbag.info
//
// 大包状态查询
type CainiaoGlobalImPickupBigbagInfoAPIResponse struct {
	model.CommonResponse
	CainiaoGlobalImPickupBigbagInfoAPIResponseModel
}

// CainiaoGlobalImPickupBigbagInfoAPIResponseModel is 大包状态查询 成功返回结果
type CainiaoGlobalImPickupBigbagInfoAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_im_pickup_bigbag_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// hsfResult
	HsfResult *HsfResult `json:"hsf_result,omitempty" xml:"hsf_result,omitempty"`
}
