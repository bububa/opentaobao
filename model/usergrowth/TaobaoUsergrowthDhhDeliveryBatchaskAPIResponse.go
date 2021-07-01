package usergrowth

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsergrowthDhhDeliveryBatchaskAPIResponse
广告曝光前判定批量接口V2 API返回值
taobao.usergrowth.dhh.delivery.batchask

广告曝光前判定批量接口V2 */
type TaobaoUsergrowthDhhDeliveryBatchaskAPIResponse struct {
	model.CommonResponse
	TaobaoUsergrowthDhhDeliveryBatchaskAPIResponseModel
}

// TaobaoUsergrowthDhhDeliveryBatchaskAPIResponseModel is 广告曝光前判定批量接口V2 成功返回结果
type TaobaoUsergrowthDhhDeliveryBatchaskAPIResponseModel struct {
	XMLName xml.Name `xml:"usergrowth_dhh_delivery_batchask_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *BatchAskResultV2 `json:"result,omitempty" xml:"result,omitempty"`
}
