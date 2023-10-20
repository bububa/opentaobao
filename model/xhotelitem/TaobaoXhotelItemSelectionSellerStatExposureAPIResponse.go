package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelitemselectionsellerstatexposureAPIResponse 选品曝光趋势 API返回值
// taobao.xhotel.item.selection.seller.stat.exposure
//
// 用于提供给商家获取选品曝光趋势
type TaobaoxhotelitemselectionsellerstatexposureAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelitemselectionsellerstatexposureAPIResponseModel
}

// TaobaoxhotelitemselectionsellerstatexposureAPIResponseModel is 选品曝光趋势 成功返回结果
type TaobaoxhotelitemselectionsellerstatexposureAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_item_selection_seller_stat_exposure_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoxhotelitemselectionsellerstatexposureResult `json:"result,omitempty" xml:"result,omitempty"`
}
