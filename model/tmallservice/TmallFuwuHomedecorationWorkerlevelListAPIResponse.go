package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallfuwuhomedecorationworkerlevellistAPIResponse 查询工人分层数据接口 API返回值
// tmall.fuwu.homedecoration.workerlevel.list
//
// 查询工人分层数据接口
type TmallfuwuhomedecorationworkerlevellistAPIResponse struct {
	model.CommonResponse
	TmallfuwuhomedecorationworkerlevellistAPIResponseModel
}

// TmallfuwuhomedecorationworkerlevellistAPIResponseModel is 查询工人分层数据接口 成功返回结果
type TmallfuwuhomedecorationworkerlevellistAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_fuwu_homedecoration_workerlevel_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *TmallfuwuhomedecorationworkerlevellistResult `json:"result,omitempty" xml:"result,omitempty"`
}
