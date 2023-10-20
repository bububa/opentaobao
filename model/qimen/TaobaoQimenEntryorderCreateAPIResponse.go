package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenentryordercreateAPIResponse 入库单创建接口 API返回值
// taobao.qimen.entryorder.create
//
// taobao.qimen.entryorder.create
type TaobaoqimenentryordercreateAPIResponse struct {
	model.CommonResponse
	TaobaoqimenentryordercreateAPIResponseModel
}

// TaobaoqimenentryordercreateAPIResponseModel is 入库单创建接口 成功返回结果
type TaobaoqimenentryordercreateAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_entryorder_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoqimenentryordercreateResponse `json:"response,omitempty" xml:"response,omitempty"`
}
