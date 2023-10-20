package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelUpdateAPIResponse 酒店更新接口（ID不存在自动新增） API返回值
// taobao.xhotel.update
//
// 酒店更新接口
type TaobaoXhotelUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelUpdateAPIResponseModel
}

// TaobaoXhotelUpdateAPIResponseModel is 酒店更新接口（ID不存在自动新增） 成功返回结果
type TaobaoXhotelUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 酒店信息
	Xhotel *Xhotel `json:"xhotel,omitempty" xml:"xhotel,omitempty"`
}
