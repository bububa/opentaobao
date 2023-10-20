package tuanhotel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelcomboreviewAPIResponse 套餐审核接口 API返回值
// taobao.xhotel.combo.review
//
// 套餐审核接口
type TaobaoxhotelcomboreviewAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelcomboreviewAPIResponseModel
}

// TaobaoxhotelcomboreviewAPIResponseModel is 套餐审核接口 成功返回结果
type TaobaoxhotelcomboreviewAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_combo_review_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 审核状态
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}
