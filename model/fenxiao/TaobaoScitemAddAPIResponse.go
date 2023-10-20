package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoscitemaddAPIResponse 发布后端商品 API返回值
// taobao.scitem.add
//
// 发布后端商品
type TaobaoscitemaddAPIResponse struct {
	model.CommonResponse
	TaobaoscitemaddAPIResponseModel
}

// TaobaoscitemaddAPIResponseModel is 发布后端商品 成功返回结果
type TaobaoscitemaddAPIResponseModel struct {
	XMLName xml.Name `xml:"scitem_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 后台商品信息
	ScItem *ScItem `json:"sc_item,omitempty" xml:"sc_item,omitempty"`
}
