package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbpromoBindAPIResponse 自促活动绑定接口 API返回值
// taobao.xhotel.bnbpromo.bind
//
// 自促活动绑定接口
type TaobaoXhotelBnbpromoBindAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelBnbpromoBindAPIResponseModel
}

// TaobaoXhotelBnbpromoBindAPIResponseModel is 自促活动绑定接口 成功返回结果
type TaobaoXhotelBnbpromoBindAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbpromo_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 营销绑定返回对象
	Module *PromoBindResult `json:"module,omitempty" xml:"module,omitempty"`
}
