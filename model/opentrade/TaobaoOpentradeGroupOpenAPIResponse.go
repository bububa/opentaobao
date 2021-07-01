package opentrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeGroupOpenAPIResponse
组团购场景开团 API返回值
taobao.opentrade.group.open

组团购场景下，团长开团 */
type TaobaoOpentradeGroupOpenAPIResponse struct {
	model.CommonResponse
	TaobaoOpentradeGroupOpenAPIResponseModel
}

// TaobaoOpentradeGroupOpenAPIResponseModel is 组团购场景开团 成功返回结果
type TaobaoOpentradeGroupOpenAPIResponseModel struct {
	XMLName xml.Name `xml:"opentrade_group_open_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 团信息
	Result *OpenGroupResponse `json:"result,omitempty" xml:"result,omitempty"`
}
