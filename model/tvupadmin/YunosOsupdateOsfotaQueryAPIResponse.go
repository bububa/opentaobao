package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunososupdateosfotaqueryAPIResponse 系统升级分页查询 API返回值
// yunos.osupdate.osfota.query
//
// 分页查询osoupdate系统升级列表
type YunososupdateosfotaqueryAPIResponse struct {
	model.CommonResponse
	YunososupdateosfotaqueryAPIResponseModel
}

// YunososupdateosfotaqueryAPIResponseModel is 系统升级分页查询 成功返回结果
type YunososupdateosfotaqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_osupdate_osfota_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *YunososupdateosfotaqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
