package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniitemskugetAPIResponse 获取全渠道门店商品sku API返回值
// taobao.omniitem.sku.get
//
// 通过skuId或者skuOutId查询全渠道门店商品sku信息
type TaobaoomniitemskugetAPIResponse struct {
	model.CommonResponse
	TaobaoomniitemskugetAPIResponseModel
}

// TaobaoomniitemskugetAPIResponseModel is 获取全渠道门店商品sku 成功返回结果
type TaobaoomniitemskugetAPIResponseModel struct {
	XMLName xml.Name `xml:"omniitem_sku_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *OmniResult `json:"result,omitempty" xml:"result,omitempty"`
}
