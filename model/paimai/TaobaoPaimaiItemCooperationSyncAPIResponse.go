package paimai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaopaimaiitemcooperationsyncAPIResponse 商品同步 API返回值
// taobao.paimai.item.cooperation.sync
//
// 商品同步
type TaobaopaimaiitemcooperationsyncAPIResponse struct {
	model.CommonResponse
	TaobaopaimaiitemcooperationsyncAPIResponseModel
}

// TaobaopaimaiitemcooperationsyncAPIResponseModel is 商品同步 成功返回结果
type TaobaopaimaiitemcooperationsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"paimai_item_cooperation_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功上传的商品ID列表
	Value []int64 `json:"value,omitempty" xml:"value>int64,omitempty"`
	// 结果描述
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
