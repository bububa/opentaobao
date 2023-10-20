package idleisv

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvitemdownshelfAPIResponse 服务商闲鱼商品下架 API返回值
// alibaba.idle.isv.item.downshelf
//
// 供外部服务商ISV进行闲鱼商品下架操作
type AlibabaidleisvitemdownshelfAPIResponse struct {
	model.CommonResponse
	AlibabaidleisvitemdownshelfAPIResponseModel
}

// AlibabaidleisvitemdownshelfAPIResponseModel is 服务商闲鱼商品下架 成功返回结果
type AlibabaidleisvitemdownshelfAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_item_downshelf_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果result
	Result *AlibabaidleisvitemdownshelfTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
