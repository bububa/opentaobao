package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoScitemGetAPIResponse 根据id查询商品 API返回值
// taobao.scitem.get
//
// 根据id查询商品
type TaobaoScitemGetAPIResponse struct {
	model.CommonResponse
	TaobaoScitemGetAPIResponseModel
}

// TaobaoScitemGetAPIResponseModel is 根据id查询商品 成功返回结果
type TaobaoScitemGetAPIResponseModel struct {
	XMLName xml.Name `xml:"scitem_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 后端商品
	ScItem *ScItem `json:"sc_item,omitempty" xml:"sc_item,omitempty"`
}
