package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenSingleitemSynchronizeAPIResponse 商品同步接口 API返回值
// taobao.qimen.singleitem.synchronize
//
// ERP调用奇门的接口,同步商品信息给WMS
type TaobaoQimenSingleitemSynchronizeAPIResponse struct {
	model.CommonResponse
	TaobaoQimenSingleitemSynchronizeAPIResponseModel
}

// TaobaoQimenSingleitemSynchronizeAPIResponseModel is 商品同步接口 成功返回结果
type TaobaoQimenSingleitemSynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_singleitem_synchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenSingleitemSynchronizeResponse `json:"response,omitempty" xml:"response,omitempty"`
}
