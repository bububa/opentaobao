package openmall

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenmallitemgetAPIResponse 获取商品详情物料 API返回值
// taobao.openmall.item.get
//
// 获取联盟开放的openmall商品
type TaobaoopenmallitemgetAPIResponse struct {
	model.CommonResponse
	TaobaoopenmallitemgetAPIResponseModel
}

// TaobaoopenmallitemgetAPIResponseModel is 获取商品详情物料 成功返回结果
type TaobaoopenmallitemgetAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_item_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoopenmallitemgetResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
