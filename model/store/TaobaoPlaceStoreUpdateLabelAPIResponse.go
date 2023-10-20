package store

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoplacestoreupdatelabelAPIResponse 商户门店标签更新接口 API返回值
// taobao.place.store.update.label
//
// 更新商户门店标签（服务、权益、标签）接口
type TaobaoplacestoreupdatelabelAPIResponse struct {
	model.CommonResponse
	TaobaoplacestoreupdatelabelAPIResponseModel
}

// TaobaoplacestoreupdatelabelAPIResponseModel is 商户门店标签更新接口 成功返回结果
type TaobaoplacestoreupdatelabelAPIResponseModel struct {
	XMLName xml.Name `xml:"place_store_update_label_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *UpdateResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
