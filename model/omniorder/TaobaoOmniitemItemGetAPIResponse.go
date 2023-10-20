package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniitemitemgetAPIResponse 获取全渠道门店商品 API返回值
// taobao.omniitem.item.get
//
// 通过门店id/类目id/商品id单个或多个参数组合查询全渠道门店商品信息
type TaobaoomniitemitemgetAPIResponse struct {
	model.CommonResponse
	TaobaoomniitemitemgetAPIResponseModel
}

// TaobaoomniitemitemgetAPIResponseModel is 获取全渠道门店商品 成功返回结果
type TaobaoomniitemitemgetAPIResponseModel struct {
	XMLName xml.Name `xml:"omniitem_item_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *PageResult `json:"result,omitempty" xml:"result,omitempty"`
}
