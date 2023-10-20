package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallaliautoreceiptstateupdateAPIResponse 服务工单状态更新 API返回值
// tmall.aliauto.receipt.state.update
//
// 二轮车服务工单状态更新
type TmallaliautoreceiptstateupdateAPIResponse struct {
	model.CommonResponse
	TmallaliautoreceiptstateupdateAPIResponseModel
}

// TmallaliautoreceiptstateupdateAPIResponseModel is 服务工单状态更新 成功返回结果
type TmallaliautoreceiptstateupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_receipt_state_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
