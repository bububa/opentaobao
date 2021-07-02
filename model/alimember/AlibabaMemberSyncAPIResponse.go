package alimember

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMemberSyncAPIResponse 会员信息同步 API返回值
// alibaba.member.sync
//
// 会员信息同步
type AlibabaMemberSyncAPIResponse struct {
	model.CommonResponse
	AlibabaMemberSyncAPIResponseModel
}

// AlibabaMemberSyncAPIResponseModel is 会员信息同步 成功返回结果
type AlibabaMemberSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_member_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaMemberSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
