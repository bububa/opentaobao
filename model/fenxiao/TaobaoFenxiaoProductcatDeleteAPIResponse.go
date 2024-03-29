package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductcatDeleteAPIResponse 删除产品线 API返回值
// taobao.fenxiao.productcat.delete
//
// 删除产品线
type TaobaoFenxiaoProductcatDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoProductcatDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductcatDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoProductcatDeleteAPIResponseModel).Reset()
}

// TaobaoFenxiaoProductcatDeleteAPIResponseModel is 删除产品线 成功返回结果
type TaobaoFenxiaoProductcatDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_productcat_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoProductcatDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoFenxiaoProductcatDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoProductcatDeleteAPIResponse)
	},
}

// GetTaobaoFenxiaoProductcatDeleteAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoProductcatDeleteAPIResponse
func GetTaobaoFenxiaoProductcatDeleteAPIResponse() *TaobaoFenxiaoProductcatDeleteAPIResponse {
	return poolTaobaoFenxiaoProductcatDeleteAPIResponse.Get().(*TaobaoFenxiaoProductcatDeleteAPIResponse)
}

// ReleaseTaobaoFenxiaoProductcatDeleteAPIResponse 将 TaobaoFenxiaoProductcatDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoProductcatDeleteAPIResponse(v *TaobaoFenxiaoProductcatDeleteAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoProductcatDeleteAPIResponse.Put(v)
}
