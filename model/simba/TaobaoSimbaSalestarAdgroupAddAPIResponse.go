package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSalestarAdgroupAddAPIResponse (新)创建一个推广组 API返回值
// taobao.simba.salestar.adgroup.add
//
// 创建一个推广组
type TaobaoSimbaSalestarAdgroupAddAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaSalestarAdgroupAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaSalestarAdgroupAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaSalestarAdgroupAddAPIResponseModel).Reset()
}

// TaobaoSimbaSalestarAdgroupAddAPIResponseModel is (新)创建一个推广组 成功返回结果
type TaobaoSimbaSalestarAdgroupAddAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_salestar_adgroup_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 新增加的推广组
	Adgroup *ADGroup `json:"adgroup,omitempty" xml:"adgroup,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaSalestarAdgroupAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Adgroup = nil
}

var poolTaobaoSimbaSalestarAdgroupAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaSalestarAdgroupAddAPIResponse)
	},
}

// GetTaobaoSimbaSalestarAdgroupAddAPIResponse 从 sync.Pool 获取 TaobaoSimbaSalestarAdgroupAddAPIResponse
func GetTaobaoSimbaSalestarAdgroupAddAPIResponse() *TaobaoSimbaSalestarAdgroupAddAPIResponse {
	return poolTaobaoSimbaSalestarAdgroupAddAPIResponse.Get().(*TaobaoSimbaSalestarAdgroupAddAPIResponse)
}

// ReleaseTaobaoSimbaSalestarAdgroupAddAPIResponse 将 TaobaoSimbaSalestarAdgroupAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaSalestarAdgroupAddAPIResponse(v *TaobaoSimbaSalestarAdgroupAddAPIResponse) {
	v.Reset()
	poolTaobaoSimbaSalestarAdgroupAddAPIResponse.Put(v)
}
