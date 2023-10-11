package miniappopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappVirtualItemGetAPIResponse 小程序关联虚拟商品查询 API返回值
// taobao.miniapp.virtual.item.get
//
// 小程序关联虚拟商品查询
type TaobaoMiniappVirtualItemGetAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappVirtualItemGetAPIResponseModel
}

// TaobaoMiniappVirtualItemGetAPIResponseModel is 小程序关联虚拟商品查询 成功返回结果
type TaobaoMiniappVirtualItemGetAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_virtual_item_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回列表
	Model []MiniappItemDto `json:"model,omitempty" xml:"model>miniapp_item_dto,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 业务错误信息
	ECode int64 `json:"e_code,omitempty" xml:"e_code,omitempty"`
	// 是否成功
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
}
