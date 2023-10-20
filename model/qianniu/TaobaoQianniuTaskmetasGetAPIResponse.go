package qianniu

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuTaskmetasGetAPIResponse 任务元查询接口 API返回值
// taobao.qianniu.taskmetas.get
//
// 任务元查询接口
type TaobaoQianniuTaskmetasGetAPIResponse struct {
	model.CommonResponse
	TaobaoQianniuTaskmetasGetAPIResponseModel
}

// TaobaoQianniuTaskmetasGetAPIResponseModel is 任务元查询接口 成功返回结果
type TaobaoQianniuTaskmetasGetAPIResponseModel struct {
	XMLName xml.Name `xml:"qianniu_taskmetas_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// taskmetas
	Taskmetas []QtaskMetadata `json:"taskmetas,omitempty" xml:"taskmetas>qtask_metadata,omitempty"`
}
