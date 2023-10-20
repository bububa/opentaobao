package nrt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtMemberSynchronizeAPIResponse 新零售会员同步接口 API返回值
// tmall.nrt.member.synchronize
//
// 新零售会员上翻接口，商家的会员信息同步至阿里侧
type TmallNrtMemberSynchronizeAPIResponse struct {
	model.CommonResponse
	TmallNrtMemberSynchronizeAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrtMemberSynchronizeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrtMemberSynchronizeAPIResponseModel).Reset()
}

// TmallNrtMemberSynchronizeAPIResponseModel is 新零售会员同步接口 成功返回结果
type TmallNrtMemberSynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_member_synchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Data *MemberSynResponse `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrtMemberSynchronizeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTmallNrtMemberSynchronizeAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrtMemberSynchronizeAPIResponse)
	},
}

// GetTmallNrtMemberSynchronizeAPIResponse 从 sync.Pool 获取 TmallNrtMemberSynchronizeAPIResponse
func GetTmallNrtMemberSynchronizeAPIResponse() *TmallNrtMemberSynchronizeAPIResponse {
	return poolTmallNrtMemberSynchronizeAPIResponse.Get().(*TmallNrtMemberSynchronizeAPIResponse)
}

// ReleaseTmallNrtMemberSynchronizeAPIResponse 将 TmallNrtMemberSynchronizeAPIResponse 保存到 sync.Pool
func ReleaseTmallNrtMemberSynchronizeAPIResponse(v *TmallNrtMemberSynchronizeAPIResponse) {
	v.Reset()
	poolTmallNrtMemberSynchronizeAPIResponse.Put(v)
}
