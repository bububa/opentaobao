package youkuott

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YoukutvoperatormediapagequeryAPIResponse 运营商全量媒资分页查询 API返回值
// youku.tvoperator.media.page.query
//
// 分页获取渠道全量媒资
type YoukutvoperatormediapagequeryAPIResponse struct {
	model.CommonResponse
	YoukutvoperatormediapagequeryAPIResponseModel
}

// YoukutvoperatormediapagequeryAPIResponseModel is 运营商全量媒资分页查询 成功返回结果
type YoukutvoperatormediapagequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_tvoperator_media_page_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 数据信息
	Model *YoukutvoperatormediapagequeryModel `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功 true:成功 false:不成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
