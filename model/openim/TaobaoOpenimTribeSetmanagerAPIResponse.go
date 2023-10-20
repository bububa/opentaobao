package openim

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimtribesetmanagerAPIResponse OPENIM群设置管理员 API返回值
// taobao.openim.tribe.setmanager
//
// OPENIM群设置管理员
type TaobaoopenimtribesetmanagerAPIResponse struct {
	model.CommonResponse
	TaobaoopenimtribesetmanagerAPIResponseModel
}

// TaobaoopenimtribesetmanagerAPIResponseModel is OPENIM群设置管理员 成功返回结果
type TaobaoopenimtribesetmanagerAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_tribe_setmanager_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 群服务code
	TribeCode int64 `json:"tribe_code,omitempty" xml:"tribe_code,omitempty"`
}
