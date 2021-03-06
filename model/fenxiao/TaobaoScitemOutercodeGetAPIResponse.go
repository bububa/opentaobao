package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoScitemOutercodeGetAPIResponse 根据outerCode查询商品 API返回值
// taobao.scitem.outercode.get
//
// 根据outerCode查询商品
type TaobaoScitemOutercodeGetAPIResponse struct {
	model.CommonResponse
	TaobaoScitemOutercodeGetAPIResponseModel
}

// TaobaoScitemOutercodeGetAPIResponseModel is 根据outerCode查询商品 成功返回结果
type TaobaoScitemOutercodeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"scitem_outercode_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 后台商品
	ScItem *ScItem `json:"sc_item,omitempty" xml:"sc_item,omitempty"`
}
