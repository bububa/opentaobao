package paimai

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPaimaiItempropsGetAPIResponse 拍卖相关类目属性 API返回值
// taobao.paimai.itemprops.get
//
// 读取拍卖相关类目属性
type TaobaoPaimaiItempropsGetAPIResponse struct {
	model.CommonResponse
	TaobaoPaimaiItempropsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPaimaiItempropsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPaimaiItempropsGetAPIResponseModel).Reset()
}

// TaobaoPaimaiItempropsGetAPIResponseModel is 拍卖相关类目属性 成功返回结果
type TaobaoPaimaiItempropsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"paimai_itemprops_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 类目属性信息(如果是取全量或者增量，不包括属性值),根据fields传入的参数返回相应的结果
	ItemProps []ItemProp `json:"item_props,omitempty" xml:"item_props>item_prop,omitempty"`
	// lastModified
	LastModified string `json:"last_modified,omitempty" xml:"last_modified,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPaimaiItempropsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ItemProps = m.ItemProps[:0]
	m.LastModified = ""
}

var poolTaobaoPaimaiItempropsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPaimaiItempropsGetAPIResponse)
	},
}

// GetTaobaoPaimaiItempropsGetAPIResponse 从 sync.Pool 获取 TaobaoPaimaiItempropsGetAPIResponse
func GetTaobaoPaimaiItempropsGetAPIResponse() *TaobaoPaimaiItempropsGetAPIResponse {
	return poolTaobaoPaimaiItempropsGetAPIResponse.Get().(*TaobaoPaimaiItempropsGetAPIResponse)
}

// ReleaseTaobaoPaimaiItempropsGetAPIResponse 将 TaobaoPaimaiItempropsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPaimaiItempropsGetAPIResponse(v *TaobaoPaimaiItempropsGetAPIResponse) {
	v.Reset()
	poolTaobaoPaimaiItempropsGetAPIResponse.Put(v)
}
