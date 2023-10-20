package baodian

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaodianserverdategetAPIResponse 服务器时间获取 API返回值
// taobao.baodian.server.date.get
//
// 获取服务器时间
type TaobaobaodianserverdategetAPIResponse struct {
	model.CommonResponse
	TaobaobaodianserverdategetAPIResponseModel
}

// TaobaobaodianserverdategetAPIResponseModel is 服务器时间获取 成功返回结果
type TaobaobaodianserverdategetAPIResponseModel struct {
	XMLName xml.Name `xml:"baodian_server_date_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回时间为毫秒
	ServerDate int64 `json:"server_date,omitempty" xml:"server_date,omitempty"`
}
