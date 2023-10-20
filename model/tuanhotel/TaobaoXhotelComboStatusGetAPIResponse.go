package tuanhotel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelComboStatusGetAPIResponse 酒店宝贝状态查询 API返回值
// taobao.xhotel.combo.status.get
//
// 酒店宝贝状态查询
type TaobaoXhotelComboStatusGetAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelComboStatusGetAPIResponseModel
}

// TaobaoXhotelComboStatusGetAPIResponseModel is 酒店宝贝状态查询 成功返回结果
type TaobaoXhotelComboStatusGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_combo_status_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Data *ItemInfoListResult `json:"data,omitempty" xml:"data,omitempty"`
}
