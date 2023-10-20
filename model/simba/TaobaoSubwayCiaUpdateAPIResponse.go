package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayCiaUpdateAPIResponse 批量修改单元智能出价 API返回值
// taobao.subway.cia.update
//
// 批量修改直通车推广单元的智能出价配置
type TaobaoSubwayCiaUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoSubwayCiaUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSubwayCiaUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSubwayCiaUpdateAPIResponseModel).Reset()
}

// TaobaoSubwayCiaUpdateAPIResponseModel is 批量修改单元智能出价 成功返回结果
type TaobaoSubwayCiaUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_cia_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推广组Id列表
	AdgroupList []int64 `json:"adgroup_list,omitempty" xml:"adgroup_list>int64,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSubwayCiaUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AdgroupList = m.AdgroupList[:0]
}

var poolTaobaoSubwayCiaUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSubwayCiaUpdateAPIResponse)
	},
}

// GetTaobaoSubwayCiaUpdateAPIResponse 从 sync.Pool 获取 TaobaoSubwayCiaUpdateAPIResponse
func GetTaobaoSubwayCiaUpdateAPIResponse() *TaobaoSubwayCiaUpdateAPIResponse {
	return poolTaobaoSubwayCiaUpdateAPIResponse.Get().(*TaobaoSubwayCiaUpdateAPIResponse)
}

// ReleaseTaobaoSubwayCiaUpdateAPIResponse 将 TaobaoSubwayCiaUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSubwayCiaUpdateAPIResponse(v *TaobaoSubwayCiaUpdateAPIResponse) {
	v.Reset()
	poolTaobaoSubwayCiaUpdateAPIResponse.Put(v)
}
