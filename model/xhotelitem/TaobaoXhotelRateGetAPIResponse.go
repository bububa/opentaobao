package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelrategetAPIResponse 酒店产品库rate查询 API返回值
// taobao.xhotel.rate.get
//
// 酒店产品库rate查询
type TaobaoxhotelrategetAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelrategetAPIResponseModel
}

// TaobaoxhotelrategetAPIResponseModel is 酒店产品库rate查询 成功返回结果
type TaobaoxhotelrategetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_rate_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// rate
	Rate *Rate `json:"rate,omitempty" xml:"rate,omitempty"`
}
