package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDiccontroltaskQueryAPIResponse 停开服任务列表 API返回值
// yunos.tvpubadmin.diccontroltask.query
//
// 牌照方对终端设备的停开服管理
type YunosTvpubadminDiccontroltaskQueryAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDiccontroltaskQueryAPIResponseModel
}

// YunosTvpubadminDiccontroltaskQueryAPIResponseModel is 停开服任务列表 成功返回结果
type YunosTvpubadminDiccontroltaskQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_diccontroltask_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object *PaginationDo `json:"object,omitempty" xml:"object,omitempty"`
}
