package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaocntecitemchangemessageAPIResponse 商品变更消息 API返回值
// cainiao.cntec.item.change.message
//
// 供货商商品信息变更消息
type CainiaocntecitemchangemessageAPIResponse struct {
	model.CommonResponse
	CainiaocntecitemchangemessageAPIResponseModel
}

// CainiaocntecitemchangemessageAPIResponseModel is 商品变更消息 成功返回结果
type CainiaocntecitemchangemessageAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cntec_item_change_message_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用返回的result结构体
	Result *CainiaocntecitemchangemessageResult `json:"result,omitempty" xml:"result,omitempty"`
}
