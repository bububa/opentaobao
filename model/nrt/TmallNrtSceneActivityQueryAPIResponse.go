package nrt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtSceneActivityQueryAPIResponse 喵零场景活动查询 API返回值
// tmall.nrt.scene.activity.query
//
// 喵零场景活动查询
type TmallNrtSceneActivityQueryAPIResponse struct {
	model.CommonResponse
	TmallNrtSceneActivityQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrtSceneActivityQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrtSceneActivityQueryAPIResponseModel).Reset()
}

// TmallNrtSceneActivityQueryAPIResponseModel is 喵零场景活动查询 成功返回结果
type TmallNrtSceneActivityQueryAPIResponseModel struct {
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

// Reset 清空结构体
func (m *TmallNrtSceneActivityQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errcode = ""
	m.Errmsg = ""
	m.Data = nil
	m.Succ = false
}

var poolTmallNrtSceneActivityQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrtSceneActivityQueryAPIResponse)
	},
}

// GetTmallNrtSceneActivityQueryAPIResponse 从 sync.Pool 获取 TmallNrtSceneActivityQueryAPIResponse
func GetTmallNrtSceneActivityQueryAPIResponse() *TmallNrtSceneActivityQueryAPIResponse {
	return poolTmallNrtSceneActivityQueryAPIResponse.Get().(*TmallNrtSceneActivityQueryAPIResponse)
}

// ReleaseTmallNrtSceneActivityQueryAPIResponse 将 TmallNrtSceneActivityQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallNrtSceneActivityQueryAPIResponse(v *TmallNrtSceneActivityQueryAPIResponse) {
	v.Reset()
	poolTmallNrtSceneActivityQueryAPIResponse.Put(v)
}
