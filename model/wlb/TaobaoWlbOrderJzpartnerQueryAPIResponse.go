package wlb

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbOrderJzpartnerQueryAPIResponse 查询家装服务商列表 API返回值
// taobao.wlb.order.jzpartner.query
//
// 为支持家装类目的商家，对绑定家装物流服务的订单可以在商家的ERP中发货、批量发货，因此开发根据服务类型查询所有的服务商列表的接口
type TaobaoWlbOrderJzpartnerQueryAPIResponse struct {
	model.CommonResponse
	TaobaoWlbOrderJzpartnerQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbOrderJzpartnerQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbOrderJzpartnerQueryAPIResponseModel).Reset()
}

// TaobaoWlbOrderJzpartnerQueryAPIResponseModel is 查询家装服务商列表 成功返回结果
type TaobaoWlbOrderJzpartnerQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_order_jzpartner_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 安装服务商列表
	InstallList []Partner `json:"install_list,omitempty" xml:"install_list>partner,omitempty"`
	// 物流配送服务商对象列表
	ServerList []Partner `json:"server_list,omitempty" xml:"server_list>partner,omitempty"`
	// 查询返回信息，如果失败，存储错误信息
	ResultInfo string `json:"result_info,omitempty" xml:"result_info,omitempty"`
	// 接口查询成功或者失败
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbOrderJzpartnerQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.InstallList = m.InstallList[:0]
	m.ServerList = m.ServerList[:0]
	m.ResultInfo = ""
	m.IsSuccess = false
}

var poolTaobaoWlbOrderJzpartnerQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbOrderJzpartnerQueryAPIResponse)
	},
}

// GetTaobaoWlbOrderJzpartnerQueryAPIResponse 从 sync.Pool 获取 TaobaoWlbOrderJzpartnerQueryAPIResponse
func GetTaobaoWlbOrderJzpartnerQueryAPIResponse() *TaobaoWlbOrderJzpartnerQueryAPIResponse {
	return poolTaobaoWlbOrderJzpartnerQueryAPIResponse.Get().(*TaobaoWlbOrderJzpartnerQueryAPIResponse)
}

// ReleaseTaobaoWlbOrderJzpartnerQueryAPIResponse 将 TaobaoWlbOrderJzpartnerQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbOrderJzpartnerQueryAPIResponse(v *TaobaoWlbOrderJzpartnerQueryAPIResponse) {
	v.Reset()
	poolTaobaoWlbOrderJzpartnerQueryAPIResponse.Put(v)
}
