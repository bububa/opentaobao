package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreSdtstatusAPIResponse 菜鸟裹裹运单状态查询 API返回值
// taobao.omniorder.store.sdtstatus
//
// 提供给商家查询运力单的状态。
type TaobaoOmniorderStoreSdtstatusAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderStoreSdtstatusAPIResponseModel
}

// TaobaoOmniorderStoreSdtstatusAPIResponseModel is 菜鸟裹裹运单状态查询 成功返回结果
type TaobaoOmniorderStoreSdtstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_store_sdtstatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoOmniorderStoreSdtstatusResult `json:"result,omitempty" xml:"result,omitempty"`
}
