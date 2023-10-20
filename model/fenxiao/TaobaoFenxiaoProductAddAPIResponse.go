package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductAddAPIResponse 添加产品 API返回值
// taobao.fenxiao.product.add
//
// 添加分销平台产品数据。业务逻辑与分销系统前台页面一致。&lt;br/&gt;&lt;br/&gt;    * 产品图片默认为空&lt;br/&gt;    * 产品发布后默认为下架状态
type TaobaoFenxiaoProductAddAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoProductAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoProductAddAPIResponseModel).Reset()
}

// TaobaoFenxiaoProductAddAPIResponseModel is 添加产品 成功返回结果
type TaobaoFenxiaoProductAddAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_product_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品创建时间 时间格式：yyyy-MM-dd HH:mm:ss
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 产品ID
	Pid int64 `json:"pid,omitempty" xml:"pid,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Created = ""
	m.Pid = 0
}

var poolTaobaoFenxiaoProductAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoProductAddAPIResponse)
	},
}

// GetTaobaoFenxiaoProductAddAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoProductAddAPIResponse
func GetTaobaoFenxiaoProductAddAPIResponse() *TaobaoFenxiaoProductAddAPIResponse {
	return poolTaobaoFenxiaoProductAddAPIResponse.Get().(*TaobaoFenxiaoProductAddAPIResponse)
}

// ReleaseTaobaoFenxiaoProductAddAPIResponse 将 TaobaoFenxiaoProductAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoProductAddAPIResponse(v *TaobaoFenxiaoProductAddAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoProductAddAPIResponse.Put(v)
}
