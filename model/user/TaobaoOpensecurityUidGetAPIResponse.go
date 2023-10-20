package user

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopensecurityuidgetAPIResponse 淘宝open security uid 获取接口 API返回值
// taobao.opensecurity.uid.get
//
// 根据明文 taobao user id 换取 app的 open_uid
type TaobaoopensecurityuidgetAPIResponse struct {
	model.CommonResponse
	TaobaoopensecurityuidgetAPIResponseModel
}

// TaobaoopensecurityuidgetAPIResponseModel is 淘宝open security uid 获取接口 成功返回结果
type TaobaoopensecurityuidgetAPIResponseModel struct {
	XMLName xml.Name `xml:"opensecurity_uid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// open security tbUserId，淘宝用户对每个Appkey会有唯一的一个open_uid
	OpenUid string `json:"open_uid,omitempty" xml:"open_uid,omitempty"`
}
