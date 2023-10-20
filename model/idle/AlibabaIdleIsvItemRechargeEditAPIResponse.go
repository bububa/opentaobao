package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleisvitemrechargeeditAPIResponse 闲鱼商品直充功能编辑 API返回值
// alibaba.idle.isv.item.recharge.edit
//
// 闲鱼商品直充功能编辑
type AlibabaidleisvitemrechargeeditAPIResponse struct {
	model.CommonResponse
	AlibabaidleisvitemrechargeeditAPIResponseModel
}

// AlibabaidleisvitemrechargeeditAPIResponseModel is 闲鱼商品直充功能编辑 成功返回结果
type AlibabaidleisvitemrechargeeditAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_item_recharge_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaidleisvitemrechargeeditTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
