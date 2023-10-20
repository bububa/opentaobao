package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentchildrecoitemofflineAPIResponse 下线少儿推荐内容接口 API返回值
// yunos.tvpubadmin.content.child.recoitem.offline
//
// 下线少儿推荐内容接口
type YunostvpubadmincontentchildrecoitemofflineAPIResponse struct {
	model.CommonResponse
	YunostvpubadmincontentchildrecoitemofflineAPIResponseModel
}

// YunostvpubadmincontentchildrecoitemofflineAPIResponseModel is 下线少儿推荐内容接口 成功返回结果
type YunostvpubadmincontentchildrecoitemofflineAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_child_recoitem_offline_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 下线操作结果
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}
