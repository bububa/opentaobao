package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenEntryorderCreateAPIResponse 入库单创建接口 API返回值
// taobao.qimen.entryorder.create
//
// ERP调用接口，创建入库单;
type TaobaoQimenEntryorderCreateAPIResponse struct {
	model.CommonResponse
	TaobaoQimenEntryorderCreateAPIResponseModel
}

// TaobaoQimenEntryorderCreateAPIResponseModel is 入库单创建接口 成功返回结果
type TaobaoQimenEntryorderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_entryorder_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenEntryorderCreateResponse `json:"response,omitempty" xml:"response,omitempty"`
}
