package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimensingleitemsynchronizeAPIResponse 商品同步接口 API返回值
// taobao.qimen.singleitem.synchronize
//
// taobao.qimen.singleitem.synchronize
type TaobaoqimensingleitemsynchronizeAPIResponse struct {
	model.CommonResponse
	TaobaoqimensingleitemsynchronizeAPIResponseModel
}

// TaobaoqimensingleitemsynchronizeAPIResponseModel is 商品同步接口 成功返回结果
type TaobaoqimensingleitemsynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_singleitem_synchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoqimensingleitemsynchronizeResponse `json:"response,omitempty" xml:"response,omitempty"`
}
