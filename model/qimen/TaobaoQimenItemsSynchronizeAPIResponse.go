package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenitemssynchronizeAPIResponse 商品同步接口 (批量) API返回值
// taobao.qimen.items.synchronize
//
// ERP调用奇门的接口,批量同步商品信息给WMS
type TaobaoqimenitemssynchronizeAPIResponse struct {
	model.CommonResponse
	TaobaoqimenitemssynchronizeAPIResponseModel
}

// TaobaoqimenitemssynchronizeAPIResponseModel is 商品同步接口 (批量) 成功返回结果
type TaobaoqimenitemssynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_items_synchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *ItemsSynchronizeResponse `json:"response,omitempty" xml:"response,omitempty"`
}
