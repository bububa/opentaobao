package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCcfCrowdActivityuserUploadAPIResponse 品牌营销活动用户上传 API返回值
// tmall.ccf.crowd.activityuser.upload
//
// 搜集ISV的活动用户信息，将其沉淀为活动人群数据
type TmallCcfCrowdActivityuserUploadAPIResponse struct {
	model.CommonResponse
	TmallCcfCrowdActivityuserUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCcfCrowdActivityuserUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCcfCrowdActivityuserUploadAPIResponseModel).Reset()
}

// TmallCcfCrowdActivityuserUploadAPIResponseModel is 品牌营销活动用户上传 成功返回结果
type TmallCcfCrowdActivityuserUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_ccf_crowd_activityuser_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ECode string `json:"e_code,omitempty" xml:"e_code,omitempty"`
	// 错误信息
	EMsg string `json:"e_msg,omitempty" xml:"e_msg,omitempty"`
	// 返回结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否失败
	Failed bool `json:"failed,omitempty" xml:"failed,omitempty"`
	// 是否成功
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
}

// Reset 清空结构体
func (m *TmallCcfCrowdActivityuserUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ECode = ""
	m.EMsg = ""
	m.Data = false
	m.Failed = false
	m.Suc = false
}

var poolTmallCcfCrowdActivityuserUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCcfCrowdActivityuserUploadAPIResponse)
	},
}

// GetTmallCcfCrowdActivityuserUploadAPIResponse 从 sync.Pool 获取 TmallCcfCrowdActivityuserUploadAPIResponse
func GetTmallCcfCrowdActivityuserUploadAPIResponse() *TmallCcfCrowdActivityuserUploadAPIResponse {
	return poolTmallCcfCrowdActivityuserUploadAPIResponse.Get().(*TmallCcfCrowdActivityuserUploadAPIResponse)
}

// ReleaseTmallCcfCrowdActivityuserUploadAPIResponse 将 TmallCcfCrowdActivityuserUploadAPIResponse 保存到 sync.Pool
func ReleaseTmallCcfCrowdActivityuserUploadAPIResponse(v *TmallCcfCrowdActivityuserUploadAPIResponse) {
	v.Reset()
	poolTmallCcfCrowdActivityuserUploadAPIResponse.Put(v)
}
