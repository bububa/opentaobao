package shop

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadataitemgetAPIResponse 获取商品 API返回值
// alibaba.data.item.get
//
// 获取商品信息，作为客户端Weex鉴权的虚拟api
type AlibabadataitemgetAPIResponse struct {
	model.CommonResponse
	AlibabadataitemgetAPIResponseModel
}

// AlibabadataitemgetAPIResponseModel is 获取商品 成功返回结果
type AlibabadataitemgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_data_item_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 获取商品信息，作为客户端Weex鉴权的虚拟api
	Unnamed string `json:"unnamed,omitempty" xml:"unnamed,omitempty"`
}
