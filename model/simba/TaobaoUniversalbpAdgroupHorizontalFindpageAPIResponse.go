package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpAdgroupHorizontalFindpageAPIResponse 查询单元分页列表 API返回值
// taobao.universalbp.adgroup.horizontal.findpage
//
// 查询单元分页列表
type TaobaoUniversalbpAdgroupHorizontalFindpageAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpAdgroupHorizontalFindpageAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpAdgroupHorizontalFindpageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpAdgroupHorizontalFindpageAPIResponseModel).Reset()
}

// TaobaoUniversalbpAdgroupHorizontalFindpageAPIResponseModel is 查询单元分页列表 成功返回结果
type TaobaoUniversalbpAdgroupHorizontalFindpageAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_adgroup_horizontal_findpage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpAdgroupHorizontalFindpageTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpAdgroupHorizontalFindpageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpAdgroupHorizontalFindpageAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpAdgroupHorizontalFindpageAPIResponse)
	},
}

// GetTaobaoUniversalbpAdgroupHorizontalFindpageAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpAdgroupHorizontalFindpageAPIResponse
func GetTaobaoUniversalbpAdgroupHorizontalFindpageAPIResponse() *TaobaoUniversalbpAdgroupHorizontalFindpageAPIResponse {
	return poolTaobaoUniversalbpAdgroupHorizontalFindpageAPIResponse.Get().(*TaobaoUniversalbpAdgroupHorizontalFindpageAPIResponse)
}

// ReleaseTaobaoUniversalbpAdgroupHorizontalFindpageAPIResponse 将 TaobaoUniversalbpAdgroupHorizontalFindpageAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpAdgroupHorizontalFindpageAPIResponse(v *TaobaoUniversalbpAdgroupHorizontalFindpageAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpAdgroupHorizontalFindpageAPIResponse.Put(v)
}
