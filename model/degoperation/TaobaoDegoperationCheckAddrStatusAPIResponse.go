package degoperation

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDegoperationCheckAddrStatusAPIResponse 地址 API返回值
// taobao.degoperation.check.addr.status
//
// 激励
type TaobaoDegoperationCheckAddrStatusAPIResponse struct {
	model.CommonResponse
	TaobaoDegoperationCheckAddrStatusAPIResponseModel
}

// TaobaoDegoperationCheckAddrStatusAPIResponseModel is 地址 成功返回结果
type TaobaoDegoperationCheckAddrStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"degoperation_check_addr_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BonusResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
