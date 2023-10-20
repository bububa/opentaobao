package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSalestarAdgroupUpdateAPIResponse 销量明星更新一个推广组的信息 API返回值
// taobao.simba.salestar.adgroup.update
//
// 更新一个推广组的信息，可以设置 是否上线
type TaobaoSimbaSalestarAdgroupUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaSalestarAdgroupUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaSalestarAdgroupUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaSalestarAdgroupUpdateAPIResponseModel).Reset()
}

// TaobaoSimbaSalestarAdgroupUpdateAPIResponseModel is 销量明星更新一个推广组的信息 成功返回结果
type TaobaoSimbaSalestarAdgroupUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_salestar_adgroup_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 被修改的推广组
	Adgroup *ADGroup `json:"adgroup,omitempty" xml:"adgroup,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaSalestarAdgroupUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Adgroup = nil
}

var poolTaobaoSimbaSalestarAdgroupUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaSalestarAdgroupUpdateAPIResponse)
	},
}

// GetTaobaoSimbaSalestarAdgroupUpdateAPIResponse 从 sync.Pool 获取 TaobaoSimbaSalestarAdgroupUpdateAPIResponse
func GetTaobaoSimbaSalestarAdgroupUpdateAPIResponse() *TaobaoSimbaSalestarAdgroupUpdateAPIResponse {
	return poolTaobaoSimbaSalestarAdgroupUpdateAPIResponse.Get().(*TaobaoSimbaSalestarAdgroupUpdateAPIResponse)
}

// ReleaseTaobaoSimbaSalestarAdgroupUpdateAPIResponse 将 TaobaoSimbaSalestarAdgroupUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaSalestarAdgroupUpdateAPIResponse(v *TaobaoSimbaSalestarAdgroupUpdateAPIResponse) {
	v.Reset()
	poolTaobaoSimbaSalestarAdgroupUpdateAPIResponse.Put(v)
}
