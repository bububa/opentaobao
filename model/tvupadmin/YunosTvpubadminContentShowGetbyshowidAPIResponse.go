package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentshowgetbyshowidAPIResponse 迎客松根据节目id获取节目元数据 API返回值
// yunos.tvpubadmin.content.show.getbyshowid
//
// 迎客松根据节目id获取节目元数据
type YunostvpubadmincontentshowgetbyshowidAPIResponse struct {
	model.CommonResponse
	YunostvpubadmincontentshowgetbyshowidAPIResponseModel
}

// YunostvpubadmincontentshowgetbyshowidAPIResponseModel is 迎客松根据节目id获取节目元数据 成功返回结果
type YunostvpubadmincontentshowgetbyshowidAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_show_getbyshowid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 节目元数据
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}
