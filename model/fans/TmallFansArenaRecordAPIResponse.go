package fans

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallFansArenaRecordAPIResponse 记录完成擂台的用户 API返回值
// tmall.fans.arena.record
//
// 记录完成擂台的用户和完成分数
type TmallFansArenaRecordAPIResponse struct {
	model.CommonResponse
	TmallFansArenaRecordAPIResponseModel
}

// Reset 清空结构体
func (m *TmallFansArenaRecordAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallFansArenaRecordAPIResponseModel).Reset()
}

// TmallFansArenaRecordAPIResponseModel is 记录完成擂台的用户 成功返回结果
type TmallFansArenaRecordAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_fans_arena_record_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	FansResult *FansResult `json:"fans_result,omitempty" xml:"fans_result,omitempty"`
}

// Reset 清空结构体
func (m *TmallFansArenaRecordAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FansResult = nil
}

var poolTmallFansArenaRecordAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallFansArenaRecordAPIResponse)
	},
}

// GetTmallFansArenaRecordAPIResponse 从 sync.Pool 获取 TmallFansArenaRecordAPIResponse
func GetTmallFansArenaRecordAPIResponse() *TmallFansArenaRecordAPIResponse {
	return poolTmallFansArenaRecordAPIResponse.Get().(*TmallFansArenaRecordAPIResponse)
}

// ReleaseTmallFansArenaRecordAPIResponse 将 TmallFansArenaRecordAPIResponse 保存到 sync.Pool
func ReleaseTmallFansArenaRecordAPIResponse(v *TmallFansArenaRecordAPIResponse) {
	v.Reset()
	poolTmallFansArenaRecordAPIResponse.Put(v)
}
