package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoMetaReceiveAPIResponse 汽车说明书元数据上传 API返回值
// tmall.aliauto.meta.receive
//
// 天猫汽车对外提供的汽车资源元数据上传接口
type TmallAliautoMetaReceiveAPIResponse struct {
	model.CommonResponse
	TmallAliautoMetaReceiveAPIResponseModel
}

// Reset 清空结构体
func (m *TmallAliautoMetaReceiveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallAliautoMetaReceiveAPIResponseModel).Reset()
}

// TmallAliautoMetaReceiveAPIResponseModel is 汽车说明书元数据上传 成功返回结果
type TmallAliautoMetaReceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_meta_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallAliautoMetaReceiveResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallAliautoMetaReceiveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallAliautoMetaReceiveAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallAliautoMetaReceiveAPIResponse)
	},
}

// GetTmallAliautoMetaReceiveAPIResponse 从 sync.Pool 获取 TmallAliautoMetaReceiveAPIResponse
func GetTmallAliautoMetaReceiveAPIResponse() *TmallAliautoMetaReceiveAPIResponse {
	return poolTmallAliautoMetaReceiveAPIResponse.Get().(*TmallAliautoMetaReceiveAPIResponse)
}

// ReleaseTmallAliautoMetaReceiveAPIResponse 将 TmallAliautoMetaReceiveAPIResponse 保存到 sync.Pool
func ReleaseTmallAliautoMetaReceiveAPIResponse(v *TmallAliautoMetaReceiveAPIResponse) {
	v.Reset()
	poolTmallAliautoMetaReceiveAPIResponse.Put(v)
}
