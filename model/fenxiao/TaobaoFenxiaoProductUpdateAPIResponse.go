package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductUpdateAPIResponse 更新产品 API返回值
// taobao.fenxiao.product.update
//
// 更新分销平台产品数据，不传更新数据返回失败&lt;br&gt;&lt;br/&gt;1. 对sku进行增、删操作时，原有的sku_ids字段会被忽略，请使用sku_properties和sku_properties_del。&lt;br&gt;
type TaobaoFenxiaoProductUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoProductUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoProductUpdateAPIResponseModel).Reset()
}

// TaobaoFenxiaoProductUpdateAPIResponseModel is 更新产品 成功返回结果
type TaobaoFenxiaoProductUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_product_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 更新时间，时间格式：yyyy-MM-dd HH:mm:ss
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 产品ID
	Pid int64 `json:"pid,omitempty" xml:"pid,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Modified = ""
	m.Pid = 0
}

var poolTaobaoFenxiaoProductUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoProductUpdateAPIResponse)
	},
}

// GetTaobaoFenxiaoProductUpdateAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoProductUpdateAPIResponse
func GetTaobaoFenxiaoProductUpdateAPIResponse() *TaobaoFenxiaoProductUpdateAPIResponse {
	return poolTaobaoFenxiaoProductUpdateAPIResponse.Get().(*TaobaoFenxiaoProductUpdateAPIResponse)
}

// ReleaseTaobaoFenxiaoProductUpdateAPIResponse 将 TaobaoFenxiaoProductUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoProductUpdateAPIResponse(v *TaobaoFenxiaoProductUpdateAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoProductUpdateAPIResponse.Put(v)
}
