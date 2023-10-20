package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSalestarAdgroupFindbycampidAPIResponse (销量明星)批量获取推广计划下的推广组信息 API返回值
// taobao.simba.salestar.adgroup.findbycampid
//
// 批量得到推广计划下的推广组
type TaobaoSimbaSalestarAdgroupFindbycampidAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaSalestarAdgroupFindbycampidAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaSalestarAdgroupFindbycampidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaSalestarAdgroupFindbycampidAPIResponseModel).Reset()
}

// TaobaoSimbaSalestarAdgroupFindbycampidAPIResponseModel is (销量明星)批量获取推广计划下的推广组信息 成功返回结果
type TaobaoSimbaSalestarAdgroupFindbycampidAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_salestar_adgroup_findbycampid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的推广组分页对象
	Adgroups *ADGroupPage `json:"adgroups,omitempty" xml:"adgroups,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaSalestarAdgroupFindbycampidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Adgroups = nil
}

var poolTaobaoSimbaSalestarAdgroupFindbycampidAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaSalestarAdgroupFindbycampidAPIResponse)
	},
}

// GetTaobaoSimbaSalestarAdgroupFindbycampidAPIResponse 从 sync.Pool 获取 TaobaoSimbaSalestarAdgroupFindbycampidAPIResponse
func GetTaobaoSimbaSalestarAdgroupFindbycampidAPIResponse() *TaobaoSimbaSalestarAdgroupFindbycampidAPIResponse {
	return poolTaobaoSimbaSalestarAdgroupFindbycampidAPIResponse.Get().(*TaobaoSimbaSalestarAdgroupFindbycampidAPIResponse)
}

// ReleaseTaobaoSimbaSalestarAdgroupFindbycampidAPIResponse 将 TaobaoSimbaSalestarAdgroupFindbycampidAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaSalestarAdgroupFindbycampidAPIResponse(v *TaobaoSimbaSalestarAdgroupFindbycampidAPIResponse) {
	v.Reset()
	poolTaobaoSimbaSalestarAdgroupFindbycampidAPIResponse.Put(v)
}
