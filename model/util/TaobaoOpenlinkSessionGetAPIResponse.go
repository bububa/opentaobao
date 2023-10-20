package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenlinksessiongetAPIResponse 获取授权session信息 API返回值
// taobao.openlink.session.get
//
// 帮助第三方isv生成三方session
type TaobaoopenlinksessiongetAPIResponse struct {
	model.CommonResponse
	TaobaoopenlinksessiongetAPIResponseModel
}

// TaobaoopenlinksessiongetAPIResponseModel is 获取授权session信息 成功返回结果
type TaobaoopenlinksessiongetAPIResponseModel struct {
	XMLName xml.Name `xml:"openlink_session_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoopenlinksessiongetResult `json:"result,omitempty" xml:"result,omitempty"`
}
