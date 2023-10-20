package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpmaterialshopgetAPIResponse 获取店铺信息 API返回值
// taobao.universalbp.material.shop.get
//
// 获取店铺信息
type TaobaouniversalbpmaterialshopgetAPIResponse struct {
	model.CommonResponse
	TaobaouniversalbpmaterialshopgetAPIResponseModel
}

// TaobaouniversalbpmaterialshopgetAPIResponseModel is 获取店铺信息 成功返回结果
type TaobaouniversalbpmaterialshopgetAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_material_shop_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaouniversalbpmaterialshopgetTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
