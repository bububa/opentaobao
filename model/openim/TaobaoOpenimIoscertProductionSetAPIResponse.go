package openim

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimIoscertProductionSetAPIResponse 设置ios证书 API返回值
// taobao.openim.ioscert.production.set
//
// 设置ios证书
type TaobaoOpenimIoscertProductionSetAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimIoscertProductionSetAPIResponseModel
}

// TaobaoOpenimIoscertProductionSetAPIResponseModel is 设置ios证书 成功返回结果
type TaobaoOpenimIoscertProductionSetAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_ioscert_production_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作成功
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}
