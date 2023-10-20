package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoOrderRemarkUpdateAPIResponse 修改采购单备注 API返回值
// taobao.fenxiao.order.remark.update
//
// 供应商修改采购单备注
type TaobaoFenxiaoOrderRemarkUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoOrderRemarkUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoOrderRemarkUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoOrderRemarkUpdateAPIResponseModel).Reset()
}

// TaobaoFenxiaoOrderRemarkUpdateAPIResponseModel is 修改采购单备注 成功返回结果
type TaobaoFenxiaoOrderRemarkUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_order_remark_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoOrderRemarkUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoFenxiaoOrderRemarkUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoOrderRemarkUpdateAPIResponse)
	},
}

// GetTaobaoFenxiaoOrderRemarkUpdateAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoOrderRemarkUpdateAPIResponse
func GetTaobaoFenxiaoOrderRemarkUpdateAPIResponse() *TaobaoFenxiaoOrderRemarkUpdateAPIResponse {
	return poolTaobaoFenxiaoOrderRemarkUpdateAPIResponse.Get().(*TaobaoFenxiaoOrderRemarkUpdateAPIResponse)
}

// ReleaseTaobaoFenxiaoOrderRemarkUpdateAPIResponse 将 TaobaoFenxiaoOrderRemarkUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoOrderRemarkUpdateAPIResponse(v *TaobaoFenxiaoOrderRemarkUpdateAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoOrderRemarkUpdateAPIResponse.Put(v)
}
