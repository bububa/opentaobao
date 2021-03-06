package larkiot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLarkPosBasedataGetworkstationAPIResponse 根据影城id工作站和macId获取工作站 API返回值
// taobao.lark.pos.basedata.getworkstation
//
// 获取单独工作站
type TaobaoLarkPosBasedataGetworkstationAPIResponse struct {
	model.CommonResponse
	TaobaoLarkPosBasedataGetworkstationAPIResponseModel
}

// TaobaoLarkPosBasedataGetworkstationAPIResponseModel is 根据影城id工作站和macId获取工作站 成功返回结果
type TaobaoLarkPosBasedataGetworkstationAPIResponseModel struct {
	XMLName xml.Name `xml:"lark_pos_basedata_getworkstation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}
