package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreSdtcancelAPIResponse 通知速店通取消取号 API返回值
// taobao.omniorder.store.sdtcancel
//
// 通知速店通取消取号
type TaobaoOmniorderStoreSdtcancelAPIResponse struct {
	model.CommonResponse
	TaobaoOmniorderStoreSdtcancelAPIResponseModel
}

// TaobaoOmniorderStoreSdtcancelAPIResponseModel is 通知速店通取消取号 成功返回结果
type TaobaoOmniorderStoreSdtcancelAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_store_sdtcancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoOmniorderStoreSdtcancelResult `json:"result,omitempty" xml:"result,omitempty"`
}
