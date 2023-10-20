package drug

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalihealthdrugusershopgetAPIResponse 根据用户id获取店铺id API返回值
// taobao.alihealth.drug.user.shop.get
//
// 提供给千牛智能客服，获取用户当前咨询的店铺ID
type TaobaoalihealthdrugusershopgetAPIResponse struct {
	model.CommonResponse
	TaobaoalihealthdrugusershopgetAPIResponseModel
}

// TaobaoalihealthdrugusershopgetAPIResponseModel is 根据用户id获取店铺id 成功返回结果
type TaobaoalihealthdrugusershopgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alihealth_drug_user_shop_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// shopId
	ShopId int64 `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
}
