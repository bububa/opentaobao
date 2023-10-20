package xhotelitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelmultiplerategetAPIResponse 复杂房价查询接口 API返回值
// taobao.xhotel.multiplerate.get
//
// 查询复杂房价，支持通过入住人数，连住天数，商品信息，房价信息查询
type TaobaoxhotelmultiplerategetAPIResponse struct {
	model.CommonResponse
	TaobaoxhotelmultiplerategetAPIResponseModel
}

// TaobaoxhotelmultiplerategetAPIResponseModel is 复杂房价查询接口 成功返回结果
type TaobaoxhotelmultiplerategetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_multiplerate_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 复杂价格返回结果类
	Rates []MultipleRate `json:"rates,omitempty" xml:"rates>multiple_rate,omitempty"`
}
