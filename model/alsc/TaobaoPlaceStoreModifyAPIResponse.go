package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoplacestoremodifyAPIResponse 商家修改线下门店 API返回值
// taobao.place.store.modify
//
// 用于商家修改线下门店信息
type TaobaoplacestoremodifyAPIResponse struct {
	model.CommonResponse
	TaobaoplacestoremodifyAPIResponseModel
}

// TaobaoplacestoremodifyAPIResponseModel is 商家修改线下门店 成功返回结果
type TaobaoplacestoremodifyAPIResponseModel struct {
	XMLName xml.Name `xml:"place_store_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否修改成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
