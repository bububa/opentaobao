package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupOnlineitemsvonGetAPIResponse 获取用户上架在线销售的全部宝贝 API返回值
// taobao.simba.adgroup.onlineitemsvon.get
//
// 获取用户上架在线销售的全部宝贝
type TaobaoSimbaAdgroupOnlineitemsvonGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaAdgroupOnlineitemsvonGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupOnlineitemsvonGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaAdgroupOnlineitemsvonGetAPIResponseModel).Reset()
}

// TaobaoSimbaAdgroupOnlineitemsvonGetAPIResponseModel is 获取用户上架在线销售的全部宝贝 成功返回结果
type TaobaoSimbaAdgroupOnlineitemsvonGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_adgroup_onlineitemsvon_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 带分页的淘宝商品
	PageItem *SubwayItemPartition `json:"page_item,omitempty" xml:"page_item,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupOnlineitemsvonGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.PageItem = nil
}

var poolTaobaoSimbaAdgroupOnlineitemsvonGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaAdgroupOnlineitemsvonGetAPIResponse)
	},
}

// GetTaobaoSimbaAdgroupOnlineitemsvonGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaAdgroupOnlineitemsvonGetAPIResponse
func GetTaobaoSimbaAdgroupOnlineitemsvonGetAPIResponse() *TaobaoSimbaAdgroupOnlineitemsvonGetAPIResponse {
	return poolTaobaoSimbaAdgroupOnlineitemsvonGetAPIResponse.Get().(*TaobaoSimbaAdgroupOnlineitemsvonGetAPIResponse)
}

// ReleaseTaobaoSimbaAdgroupOnlineitemsvonGetAPIResponse 将 TaobaoSimbaAdgroupOnlineitemsvonGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaAdgroupOnlineitemsvonGetAPIResponse(v *TaobaoSimbaAdgroupOnlineitemsvonGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaAdgroupOnlineitemsvonGetAPIResponse.Put(v)
}
