package openim

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeModifytribeinfoAPIResponse OPENIM群信息修改 API返回值
// taobao.openim.tribe.modifytribeinfo
//
// OPENIM群信息修改
type TaobaoOpenimTribeModifytribeinfoAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimTribeModifytribeinfoAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenimTribeModifytribeinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenimTribeModifytribeinfoAPIResponseModel).Reset()
}

// TaobaoOpenimTribeModifytribeinfoAPIResponseModel is OPENIM群信息修改 成功返回结果
type TaobaoOpenimTribeModifytribeinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_tribe_modifytribeinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 群服务code
	TribeCode int64 `json:"tribe_code,omitempty" xml:"tribe_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenimTribeModifytribeinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TribeCode = 0
}

var poolTaobaoOpenimTribeModifytribeinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenimTribeModifytribeinfoAPIResponse)
	},
}

// GetTaobaoOpenimTribeModifytribeinfoAPIResponse 从 sync.Pool 获取 TaobaoOpenimTribeModifytribeinfoAPIResponse
func GetTaobaoOpenimTribeModifytribeinfoAPIResponse() *TaobaoOpenimTribeModifytribeinfoAPIResponse {
	return poolTaobaoOpenimTribeModifytribeinfoAPIResponse.Get().(*TaobaoOpenimTribeModifytribeinfoAPIResponse)
}

// ReleaseTaobaoOpenimTribeModifytribeinfoAPIResponse 将 TaobaoOpenimTribeModifytribeinfoAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenimTribeModifytribeinfoAPIResponse(v *TaobaoOpenimTribeModifytribeinfoAPIResponse) {
	v.Reset()
	poolTaobaoOpenimTribeModifytribeinfoAPIResponse.Put(v)
}
