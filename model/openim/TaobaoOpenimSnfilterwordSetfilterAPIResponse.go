package openim

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimsnfilterwordsetfilterAPIResponse 关键词过滤 API返回值
// taobao.openim.snfilterword.setfilter
//
// 设置openim关键词过滤
type TaobaoopenimsnfilterwordsetfilterAPIResponse struct {
	model.CommonResponse
	TaobaoopenimsnfilterwordsetfilterAPIResponseModel
}

// TaobaoopenimsnfilterwordsetfilterAPIResponseModel is 关键词过滤 成功返回结果
type TaobaoopenimsnfilterwordsetfilterAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_snfilterword_setfilter_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误原因
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 成功
	Errid int64 `json:"errid,omitempty" xml:"errid,omitempty"`
}
