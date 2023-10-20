package tvupadmin

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *YunosTvpubadminDiccontroltaskGetinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTvpubadminDiccontroltaskGetinfoAPIResponseModel).Reset()
}

// YunosTvpubadminDiccontroltaskGetinfoAPIResponseModel is 获取停开服任务详情 成功返回结果
type YunosTvpubadminDiccontroltaskGetinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_tvpubadmin_diccontroltask_getinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// object
	Object *DicControlTaskDo `json:"object,omitempty" xml:"object,omitempty"`
}

// Reset 清空结构体
func (m *YunosTvpubadminDiccontroltaskGetinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = nil
}

var poolYunosTvpubadminDiccontroltaskGetinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTvpubadminDiccontroltaskGetinfoAPIResponse)
	},
}

// GetYunosTvpubadminDiccontroltaskGetinfoAPIResponse 从 sync.Pool 获取 YunosTvpubadminDiccontroltaskGetinfoAPIResponse
func GetYunosTvpubadminDiccontroltaskGetinfoAPIResponse() *YunosTvpubadminDiccontroltaskGetinfoAPIResponse {
	return poolYunosTvpubadminDiccontroltaskGetinfoAPIResponse.Get().(*YunosTvpubadminDiccontroltaskGetinfoAPIResponse)
}

// ReleaseYunosTvpubadminDiccontroltaskGetinfoAPIResponse 将 YunosTvpubadminDiccontroltaskGetinfoAPIResponse 保存到 sync.Pool
func ReleaseYunosTvpubadminDiccontroltaskGetinfoAPIResponse(v *YunosTvpubadminDiccontroltaskGetinfoAPIResponse) {
	v.Reset()
	poolYunosTvpubadminDiccontroltaskGetinfoAPIResponse.Put(v)
}
