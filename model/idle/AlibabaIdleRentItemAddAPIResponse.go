package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlerentitemaddAPIResponse 租赁商品发布 API返回值
// alibaba.idle.rent.item.add
//
// 发布闲鱼租赁商品
type AlibabaidlerentitemaddAPIResponse struct {
	model.CommonResponse
	AlibabaidlerentitemaddAPIResponseModel
}

// AlibabaidlerentitemaddAPIResponseModel is 租赁商品发布 成功返回结果
type AlibabaidlerentitemaddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_rent_item_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabaidlerentitemaddTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
