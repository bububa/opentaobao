package nrt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtAssetAuthorizationAddAPIResponse 增加数据权限授权 API返回值
// tmall.nrt.asset.authorization.add
//
// 增加数据权限授权
type TmallNrtAssetAuthorizationAddAPIResponse struct {
	model.CommonResponse
	TmallNrtAssetAuthorizationAddAPIResponseModel
}

// TmallNrtAssetAuthorizationAddAPIResponseModel is 增加数据权限授权 成功返回结果
type TmallNrtAssetAuthorizationAddAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_asset_authorization_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Data *TopAssetDataAuthResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}
