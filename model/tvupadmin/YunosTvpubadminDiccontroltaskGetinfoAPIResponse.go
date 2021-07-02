package tvupadmin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDiccontroltaskGetinfoAPIResponse 获取停开服任务详情 API返回值
// yunos.tvpubadmin.diccontroltask.getinfo
//
// 获取停开服任务详情
type YunosTvpubadminDiccontroltaskGetinfoAPIResponse struct {
	model.CommonResponse
	YunosTvpubadminDiccontroltaskGetinfoAPIResponseModel
}

// YunosTvpubadminDiccontroltaskGetinfoAPIResponseModel is 获取停开服任务详情 成功返回结果
type YunosTvpubadminDiccontroltaskGetinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_diccontroltask_getinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object *DicControlTaskDo `json:"object,omitempty" xml:"object,omitempty"`
}
