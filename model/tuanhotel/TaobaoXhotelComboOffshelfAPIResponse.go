package tuanhotel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelcombooffshelfAPIResponse 酒店套餐下架 API返回值
// taobao.xhotel.combo.offshelf
//
// 酒店套餐下架
type TaobaoxhotelcombooffshelfAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelcombooffshelfAPIResponseModel
}

// TaobaoxhotelcombooffshelfAPIResponseModel is 酒店套餐下架 成功返回结果
type TaobaoxhotelcombooffshelfAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_combo_offshelf_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 下架状态
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}
