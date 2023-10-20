package openim

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeGetalltribesAPIResponse 获取用户群列表 API返回值
// taobao.openim.tribe.getalltribes
//
// OPENIM群服务获取用户群列表
type TaobaoOpenimTribeGetalltribesAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimTribeGetalltribesAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenimTribeGetalltribesAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenimTribeGetalltribesAPIResponseModel).Reset()
}

// TaobaoOpenimTribeGetalltribesAPIResponseModel is 获取用户群列表 成功返回结果
type TaobaoOpenimTribeGetalltribesAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_tribe_getalltribes_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 群列表信息
	TribeInfoList []TribeInfo `json:"tribe_info_list,omitempty" xml:"tribe_info_list>tribe_info,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenimTribeGetalltribesAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TribeInfoList = m.TribeInfoList[:0]
}

var poolTaobaoOpenimTribeGetalltribesAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenimTribeGetalltribesAPIResponse)
	},
}

// GetTaobaoOpenimTribeGetalltribesAPIResponse 从 sync.Pool 获取 TaobaoOpenimTribeGetalltribesAPIResponse
func GetTaobaoOpenimTribeGetalltribesAPIResponse() *TaobaoOpenimTribeGetalltribesAPIResponse {
	return poolTaobaoOpenimTribeGetalltribesAPIResponse.Get().(*TaobaoOpenimTribeGetalltribesAPIResponse)
}

// ReleaseTaobaoOpenimTribeGetalltribesAPIResponse 将 TaobaoOpenimTribeGetalltribesAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenimTribeGetalltribesAPIResponse(v *TaobaoOpenimTribeGetalltribesAPIResponse) {
	v.Reset()
	poolTaobaoOpenimTribeGetalltribesAPIResponse.Put(v)
}
