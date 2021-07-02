package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenStoreprocessCreateAPIResponse 仓内加工单创建接口 API返回值
// taobao.qimen.storeprocess.create
//
// ERP调用奇门的接口,创建仓内加工单
type TaobaoQimenStoreprocessCreateAPIResponse struct {
	model.CommonResponse
	TaobaoQimenStoreprocessCreateAPIResponseModel
}

// TaobaoQimenStoreprocessCreateAPIResponseModel is 仓内加工单创建接口 成功返回结果
type TaobaoQimenStoreprocessCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_storeprocess_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *StoreProcessCreateResponse `json:"response,omitempty" xml:"response,omitempty"`
}
