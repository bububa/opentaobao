package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenTransferorderCreateAPIResponse
调拨单创建 API返回值
taobao.qimen.transferorder.create

调拨单创建 */
type TaobaoQimenTransferorderCreateAPIResponse struct {
	model.CommonResponse
	TaobaoQimenTransferorderCreateAPIResponseModel
}

// TaobaoQimenTransferorderCreateAPIResponseModel is 调拨单创建 成功返回结果
type TaobaoQimenTransferorderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_transferorder_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenTransferorderCreateStruct `json:"response,omitempty" xml:"response,omitempty"`
}
