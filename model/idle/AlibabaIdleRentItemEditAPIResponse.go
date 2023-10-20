package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlerentitemeditAPIResponse 租赁商品编辑 API返回值
// alibaba.idle.rent.item.edit
//
// 发布闲鱼租赁商品
type AlibabaidlerentitemeditAPIResponse struct {
	model.CommonResponse
	AlibabaidlerentitemeditAPIResponseModel
}

// AlibabaidlerentitemeditAPIResponseModel is 租赁商品编辑 成功返回结果
type AlibabaidlerentitemeditAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_rent_item_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabaidlerentitemeditTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
