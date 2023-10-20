package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoScitemAddAPIResponse 发布后端商品 API返回值
// taobao.scitem.add
//
// 发布后端商品
type TaobaoScitemAddAPIResponse struct {
	model.CommonResponse
	TaobaoScitemAddAPIResponseModel
}

// TaobaoScitemAddAPIResponseModel is 发布后端商品 成功返回结果
type TaobaoScitemAddAPIResponseModel struct {
	XMLName xml.Name `xml:"scitem_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 后台商品信息
	ScItem *ScItem `json:"sc_item,omitempty" xml:"sc_item,omitempty"`
}
