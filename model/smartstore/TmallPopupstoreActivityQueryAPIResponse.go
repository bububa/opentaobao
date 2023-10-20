package smartstore

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallPopupstoreActivityQueryAPIResponse 查询某段时间内的快闪活动列表 API返回值
// tmall.popupstore.activity.query
//
// 提供给ISV查询某一时间段内包含指定appKey的活动列表
type TmallPopupstoreActivityQueryAPIResponse struct {
	model.CommonResponse
	TmallPopupstoreActivityQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallPopupstoreActivityQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallPopupstoreActivityQueryAPIResponseModel).Reset()
}

// TmallPopupstoreActivityQueryAPIResponseModel is 查询某段时间内的快闪活动列表 成功返回结果
type TmallPopupstoreActivityQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_popupstore_activity_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	ResultDto *TmallPopupstoreActivityQueryResultDto `json:"result_dto,omitempty" xml:"result_dto,omitempty"`
}

// Reset 清空结构体
func (m *TmallPopupstoreActivityQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultDto = nil
}

var poolTmallPopupstoreActivityQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallPopupstoreActivityQueryAPIResponse)
	},
}

// GetTmallPopupstoreActivityQueryAPIResponse 从 sync.Pool 获取 TmallPopupstoreActivityQueryAPIResponse
func GetTmallPopupstoreActivityQueryAPIResponse() *TmallPopupstoreActivityQueryAPIResponse {
	return poolTmallPopupstoreActivityQueryAPIResponse.Get().(*TmallPopupstoreActivityQueryAPIResponse)
}

// ReleaseTmallPopupstoreActivityQueryAPIResponse 将 TmallPopupstoreActivityQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallPopupstoreActivityQueryAPIResponse(v *TmallPopupstoreActivityQueryAPIResponse) {
	v.Reset()
	poolTmallPopupstoreActivityQueryAPIResponse.Put(v)
}
