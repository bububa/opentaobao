package mei

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallMeiCrmMemberSyncAPIResponse 同步推送会员信息 API返回值
// tmall.mei.crm.member.sync
//
// 品牌方通过该api主动推送会员信息。使用场景包括
// 1.用户在线上注册后，线下补充信息后，同步到线上。
// 2.其他情况的主动推送变更。
type TmallMeiCrmMemberSyncAPIResponse struct {
	model.CommonResponse
	TmallMeiCrmMemberSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TmallMeiCrmMemberSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallMeiCrmMemberSyncAPIResponseModel).Reset()
}

// TmallMeiCrmMemberSyncAPIResponseModel is 同步推送会员信息 成功返回结果
type TmallMeiCrmMemberSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_mei_crm_member_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 处理的其他信息
	MeiExtraInfo string `json:"mei_extra_info,omitempty" xml:"mei_extra_info,omitempty"`
}

// Reset 清空结构体
func (m *TmallMeiCrmMemberSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MeiExtraInfo = ""
}

var poolTmallMeiCrmMemberSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallMeiCrmMemberSyncAPIResponse)
	},
}

// GetTmallMeiCrmMemberSyncAPIResponse 从 sync.Pool 获取 TmallMeiCrmMemberSyncAPIResponse
func GetTmallMeiCrmMemberSyncAPIResponse() *TmallMeiCrmMemberSyncAPIResponse {
	return poolTmallMeiCrmMemberSyncAPIResponse.Get().(*TmallMeiCrmMemberSyncAPIResponse)
}

// ReleaseTmallMeiCrmMemberSyncAPIResponse 将 TmallMeiCrmMemberSyncAPIResponse 保存到 sync.Pool
func ReleaseTmallMeiCrmMemberSyncAPIResponse(v *TmallMeiCrmMemberSyncAPIResponse) {
	v.Reset()
	poolTmallMeiCrmMemberSyncAPIResponse.Put(v)
}
