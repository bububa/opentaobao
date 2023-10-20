package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelitemselectionsellerstathotshidAPIResponse 供应链选品热销标准酒店覆盖情况 API返回值
// taobao.xhotel.item.selection.seller.stat.hotshid
//
// 供应链选品热销标准酒店覆盖情况
type TaobaoxhotelitemselectionsellerstathotshidAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelitemselectionsellerstathotshidAPIResponseModel
}

// TaobaoxhotelitemselectionsellerstathotshidAPIResponseModel is 供应链选品热销标准酒店覆盖情况 成功返回结果
type TaobaoxhotelitemselectionsellerstathotshidAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_item_selection_seller_stat_hotshid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoxhotelitemselectionsellerstathotshidResult `json:"result,omitempty" xml:"result,omitempty"`
}
