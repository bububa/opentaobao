package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductcatUpdateAPIResponse 修改产品线 API返回值
// taobao.fenxiao.productcat.update
//
// 修改产品线
type TaobaoFenxiaoProductcatUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoProductcatUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductcatUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoProductcatUpdateAPIResponseModel).Reset()
}

// TaobaoFenxiaoProductcatUpdateAPIResponseModel is 修改产品线 成功返回结果
type TaobaoFenxiaoProductcatUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_productcat_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductcatUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoFenxiaoProductcatUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoProductcatUpdateAPIResponse)
	},
}

// GetTaobaoFenxiaoProductcatUpdateAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoProductcatUpdateAPIResponse
func GetTaobaoFenxiaoProductcatUpdateAPIResponse() *TaobaoFenxiaoProductcatUpdateAPIResponse {
	return poolTaobaoFenxiaoProductcatUpdateAPIResponse.Get().(*TaobaoFenxiaoProductcatUpdateAPIResponse)
}

// ReleaseTaobaoFenxiaoProductcatUpdateAPIResponse 将 TaobaoFenxiaoProductcatUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoProductcatUpdateAPIResponse(v *TaobaoFenxiaoProductcatUpdateAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoProductcatUpdateAPIResponse.Put(v)
}
