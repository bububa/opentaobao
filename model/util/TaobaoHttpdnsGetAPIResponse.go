package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaohttpdnsgetAPIResponse TOPDNS配置 API返回值
// taobao.httpdns.get
//
// 获取TOP DNS配置
type TaobaohttpdnsgetAPIResponse struct {
	model.CommonResponse
	TaobaohttpdnsgetAPIResponseModel
}

// TaobaohttpdnsgetAPIResponseModel is TOPDNS配置 成功返回结果
type TaobaohttpdnsgetAPIResponseModel struct {
	XMLName xml.Name `xml:"httpdns_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// HTTP DNS配置信息
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
