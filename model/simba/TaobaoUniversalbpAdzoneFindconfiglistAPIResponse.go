package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpadzonefindconfiglistAPIResponse 查询所有可用资源包信息 API返回值
// taobao.universalbp.adzone.findconfiglist
//
// 查询该场景下，所有可用的资源包，可能存在数据重复。但是针对不同子场景和推广设置，可以选用的资源包有差异，建议关注补充文档，或者根据bp前端的限制，进行传参。
type TaobaouniversalbpadzonefindconfiglistAPIResponse struct {
	model.CommonResponse
	TaobaouniversalbpadzonefindconfiglistAPIResponseModel
}

// TaobaouniversalbpadzonefindconfiglistAPIResponseModel is 查询所有可用资源包信息 成功返回结果
type TaobaouniversalbpadzonefindconfiglistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_adzone_findconfiglist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaouniversalbpadzonefindconfiglistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
