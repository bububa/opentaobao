package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenstockoutcreateAPIResponse 出库单创建接口 API返回值
// taobao.qimen.stockout.create
//
// taobao.qimen.returnorder.create
type TaobaoqimenstockoutcreateAPIResponse struct {
	model.CommonResponse
	TaobaoqimenstockoutcreateAPIResponseModel
}

// TaobaoqimenstockoutcreateAPIResponseModel is 出库单创建接口 成功返回结果
type TaobaoqimenstockoutcreateAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_stockout_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoqimenstockoutcreateResponse `json:"response,omitempty" xml:"response,omitempty"`
}
