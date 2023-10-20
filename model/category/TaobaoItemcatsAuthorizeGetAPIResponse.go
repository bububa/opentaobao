package category

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitemcatsauthorizegetAPIResponse 查询商家被授权品牌列表和类目列表 API返回值
// taobao.itemcats.authorize.get
//
// 查询B商家被授权品牌列表、类目列表和 c 商家新品类目列表
type TaobaoitemcatsauthorizegetAPIResponse struct {
	model.CommonResponse
	TaobaoitemcatsauthorizegetAPIResponseModel
}

// TaobaoitemcatsauthorizegetAPIResponseModel is 查询商家被授权品牌列表和类目列表 成功返回结果
type TaobaoitemcatsauthorizegetAPIResponseModel struct {
	XMLName xml.Name `xml:"itemcats_authorize_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 里面有3个数组：Brand[]品牌列表,ItemCat[] 类目列表XinpinItemCat[] 针对于C卖家新品类目列表
	SellerAuthorize *SellerAuthorize `json:"seller_authorize,omitempty" xml:"seller_authorize,omitempty"`
}
