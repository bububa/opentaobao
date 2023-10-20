package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallFuwuHomedecorationWorkerlevelListAPIResponse 查询工人分层数据接口 API返回值
// tmall.fuwu.homedecoration.workerlevel.list
//
// 查询工人分层数据接口
type TmallFuwuHomedecorationWorkerlevelListAPIResponse struct {
	model.CommonResponse
	TmallFuwuHomedecorationWorkerlevelListAPIResponseModel
}

// TmallFuwuHomedecorationWorkerlevelListAPIResponseModel is 查询工人分层数据接口 成功返回结果
type TmallFuwuHomedecorationWorkerlevelListAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_fuwu_homedecoration_workerlevel_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *TmallFuwuHomedecorationWorkerlevelListResult `json:"result,omitempty" xml:"result,omitempty"`
}
