package qianniu

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuTaskmetaUpdateAPIResponse 更新任务元数据 API返回值
// taobao.qianniu.taskmeta.update
//
// 由任务发起者调用
type TaobaoQianniuTaskmetaUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoQianniuTaskmetaUpdateAPIResponseModel
}

// TaobaoQianniuTaskmetaUpdateAPIResponseModel is 更新任务元数据 成功返回结果
type TaobaoQianniuTaskmetaUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"qianniu_taskmeta_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
