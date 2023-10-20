package legalsuit

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalstandpointscenequeryAPIResponse 查询场景 API返回值
// alibaba.legal.standpoint.scene.query
//
// 查询场景
type AlibabalegalstandpointscenequeryAPIResponse struct {
	model.CommonResponse
	AlibabalegalstandpointscenequeryAPIResponseModel
}

// AlibabalegalstandpointscenequeryAPIResponseModel is 查询场景 成功返回结果
type AlibabalegalstandpointscenequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_standpoint_scene_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 场景集合
	Content []SceneOption `json:"content,omitempty" xml:"content>scene_option,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCodeRes int64 `json:"error_code_res,omitempty" xml:"error_code_res,omitempty"`
	// 是否成功
	SuccessRes bool `json:"success_res,omitempty" xml:"success_res,omitempty"`
}
