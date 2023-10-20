package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoplacestorecreateAPIResponse 商户门店创建接口 API返回值
// taobao.place.store.create
//
// 用于商家创建线下门店
type TaobaoplacestorecreateAPIResponse struct {
	model.CommonResponse
	TaobaoplacestorecreateAPIResponseModel
}

// TaobaoplacestorecreateAPIResponseModel is 商户门店创建接口 成功返回结果
type TaobaoplacestorecreateAPIResponseModel struct {
	XMLName xml.Name `xml:"place_store_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
