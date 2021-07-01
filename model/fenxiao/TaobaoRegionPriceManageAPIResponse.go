package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRegionPriceManageAPIResponse
编辑区域价格 API返回值
taobao.region.price.manage

编辑区域价格 */
type TaobaoRegionPriceManageAPIResponse struct {
	model.CommonResponse
	TaobaoRegionPriceManageAPIResponseModel
}

// TaobaoRegionPriceManageAPIResponseModel is 编辑区域价格 成功返回结果
type TaobaoRegionPriceManageAPIResponseModel struct {
	XMLName xml.Name `xml:"region_price_manage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
