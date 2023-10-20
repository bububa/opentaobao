package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentTableauditQuerymetaitemAPIResponse 运营位管控-查询魔盒运营位元数据列表 API返回值
// yunos.tvpubadmin.content.tableaudit.querymetaitem
//
// 运营位管控-查询魔盒运营位元数据列表
type YunosTvpubadminContentTableauditQuerymetaitemAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminContentTableauditQuerymetaitemAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTvpubadminContentTableauditQuerymetaitemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminContentTableauditQuerymetaitemAPIResponseModel).Reset()
}

// YunosTvpubadminContentTableauditQuerymetaitemAPIResponseModel is 运营位管控-查询魔盒运营位元数据列表 成功返回结果
type YunosTvpubadminContentTableauditQuerymetaitemAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_content_tableaudit_querymetaitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查看桌面坑位元数据列表（用于魔盒）
	Object string `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminContentTableauditQuerymetaitemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = ""
}

var poolYunosTvpubadminContentTableauditQuerymetaitemAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminContentTableauditQuerymetaitemAPIResponse)
	},
}

// GetYunosTvpubadminContentTableauditQuerymetaitemAPIResponse 从 sync.Pool 获取 YunosTvpubadminContentTableauditQuerymetaitemAPIResponse
func GetYunosTvpubadminContentTableauditQuerymetaitemAPIResponse() *YunosTvpubadminContentTableauditQuerymetaitemAPIResponse {
	return poolYunosTvpubadminContentTableauditQuerymetaitemAPIResponse.Get().(*YunosTvpubadminContentTableauditQuerymetaitemAPIResponse)
}

// ReleaseYunosTvpubadminContentTableauditQuerymetaitemAPIResponse 将 YunosTvpubadminContentTableauditQuerymetaitemAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminContentTableauditQuerymetaitemAPIResponse(v *YunosTvpubadminContentTableauditQuerymetaitemAPIResponse) {
	v.Reset()
	poolYunosTvpubadminContentTableauditQuerymetaitemAPIResponse.Put(v)
}
