package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminAdmOttQueryAPIResponse 优酷OTT端广告素材查询 API返回值
// yunos.tvpubadmin.adm.ott.query
//
// 查询广告素材
type YunosTvpubadminAdmOttQueryAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminAdmOttQueryAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminAdmOttQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminAdmOttQueryAPIResponseModel).Reset()
}

// YunosTvpubadminAdmOttQueryAPIResponseModel is 优酷OTT端广告素材查询 成功返回结果
type YunosTvpubadminAdmOttQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_adm_ott_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回数据详情，json格式
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminAdmOttQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = ""
}

var poolYunosTvpubadminAdmOttQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminAdmOttQueryAPIResponse)
	},
}

// GetYunosTvpubadminAdmOttQueryAPIResponse 从 sync.Pool 获取 YunosTvpubadminAdmOttQueryAPIResponse
func GetYunosTvpubadminAdmOttQueryAPIResponse() *YunosTvpubadminAdmOttQueryAPIResponse {
	return poolYunosTvpubadminAdmOttQueryAPIResponse.Get().(*YunosTvpubadminAdmOttQueryAPIResponse)
}

// ReleaseYunosTvpubadminAdmOttQueryAPIResponse 将 YunosTvpubadminAdmOttQueryAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminAdmOttQueryAPIResponse(v *YunosTvpubadminAdmOttQueryAPIResponse) {
	v.Reset()
	poolYunosTvpubadminAdmOttQueryAPIResponse.Put(v)
}
