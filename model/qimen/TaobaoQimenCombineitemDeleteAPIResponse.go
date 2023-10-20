package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimencombineitemdeleteAPIResponse 组合货品删除接口 API返回值
// taobao.qimen.combineitem.delete
//
// 组合货品删除
type TaobaoqimencombineitemdeleteAPIResponse struct {
	model.CommonResponse
	TaobaoqimencombineitemdeleteAPIResponseModel
}

// TaobaoqimencombineitemdeleteAPIResponseModel is 组合货品删除接口 成功返回结果
type TaobaoqimencombineitemdeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_combineitem_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *ResponseDo `json:"response,omitempty" xml:"response,omitempty"`
}
