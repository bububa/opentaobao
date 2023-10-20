package yunosdm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosDmSysGetDomainAPIResponse 获取动态域名 API返回值
// yunos.dm.sys.get.domain
//
// 返回alios ucp后端域名
type YunosDmSysGetDomainAPIResponse struct {
	model.CommonResponse
	YunosDmSysGetDomainAPIResponseModel
}

// Reset 清空结构体
func (m *YunosDmSysGetDomainAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosDmSysGetDomainAPIResponseModel).Reset()
}

// YunosDmSysGetDomainAPIResponseModel is 获取动态域名 成功返回结果
type YunosDmSysGetDomainAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_dm_sys_get_domain_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// obj
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}

// Reset 清空结构体
func (m *YunosDmSysGetDomainAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Url = ""
}

var poolYunosDmSysGetDomainAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosDmSysGetDomainAPIResponse)
	},
}

// GetYunosDmSysGetDomainAPIResponse 从 sync.Pool 获取 YunosDmSysGetDomainAPIResponse
func GetYunosDmSysGetDomainAPIResponse() *YunosDmSysGetDomainAPIResponse {
	return poolYunosDmSysGetDomainAPIResponse.Get().(*YunosDmSysGetDomainAPIResponse)
}

// ReleaseYunosDmSysGetDomainAPIResponse 将 YunosDmSysGetDomainAPIResponse 保存到 sync.Pool
func ReleaseYunosDmSysGetDomainAPIResponse(v *YunosDmSysGetDomainAPIResponse) {
	v.Reset()
	poolYunosDmSysGetDomainAPIResponse.Put(v)
}
