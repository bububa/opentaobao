package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniitemcategorygetAPIResponse 全渠道商品轻发布类目信息 API返回值
// taobao.omniitem.category.get
//
// 全渠道商品轻发布类目信息
type TaobaoomniitemcategorygetAPIResponse struct {
	model.CommonResponse
	TaobaoomniitemcategorygetAPIResponseModel
}

// TaobaoomniitemcategorygetAPIResponseModel is 全渠道商品轻发布类目信息 成功返回结果
type TaobaoomniitemcategorygetAPIResponseModel struct {
	XMLName xml.Name `xml:"omniitem_category_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoomniitemcategorygetResult `json:"result,omitempty" xml:"result,omitempty"`
}
