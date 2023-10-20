package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoregionpricemanageAPIResponse 编辑区域价格 API返回值
// taobao.region.price.manage
//
// 编辑区域价格
type TaobaoregionpricemanageAPIResponse struct {
	model.CommonResponse
	TaobaoregionpricemanageAPIResponseModel
}

// TaobaoregionpricemanageAPIResponseModel is 编辑区域价格 成功返回结果
type TaobaoregionpricemanageAPIResponseModel struct {
	XMLName xml.Name `xml:"region_price_manage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
