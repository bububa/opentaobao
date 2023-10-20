package nrt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtsceneactivityqueryAPIResponse 喵零场景活动查询 API返回值
// tmall.nrt.scene.activity.query
//
// 喵零场景活动查询
type TmallnrtsceneactivityqueryAPIResponse struct {
	model.CommonResponse
	TmallnrtsceneactivityqueryAPIResponseModel
}

// TmallnrtsceneactivityqueryAPIResponseModel is 喵零场景活动查询 成功返回结果
type TmallnrtsceneactivityqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_scene_activity_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	Errcode string `json:"errcode,omitempty" xml:"errcode,omitempty"`
	// 错误信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 数据
	Data *NrtSceneActivityDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}
