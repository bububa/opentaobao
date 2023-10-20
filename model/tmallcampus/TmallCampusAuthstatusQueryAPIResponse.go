package tmallcampus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallcampusauthstatusqueryAPIResponse 学生认证状态查询 API返回值
// tmall.campus.authstatus.query
//
// 学生认证状态查询
type TmallcampusauthstatusqueryAPIResponse struct {
	model.CommonResponse
	TmallcampusauthstatusqueryAPIResponseModel
}

// TmallcampusauthstatusqueryAPIResponseModel is 学生认证状态查询 成功返回结果
type TmallcampusauthstatusqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_campus_authstatus_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TmallcampusauthstatusqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
