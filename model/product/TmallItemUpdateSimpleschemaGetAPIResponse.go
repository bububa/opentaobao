package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemUpdateSimpleschemaGetAPIResponse 官网同购编辑商品的get接口 API返回值
// tmall.item.update.simpleschema.get
//
// 官网同购编辑商品的get接口
type TmallItemUpdateSimpleschemaGetAPIResponse struct {
	model.CommonResponse
	TmallItemUpdateSimpleschemaGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemUpdateSimpleschemaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemUpdateSimpleschemaGetAPIResponseModel).Reset()
}

// TmallItemUpdateSimpleschemaGetAPIResponseModel is 官网同购编辑商品的get接口 成功返回结果
type TmallItemUpdateSimpleschemaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_update_simpleschema_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 商品信息
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 返回结果
	Error bool `json:"error,omitempty" xml:"error,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemUpdateSimpleschemaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsg = ""
	m.Result = ""
	m.Error = false
}

var poolTmallItemUpdateSimpleschemaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemUpdateSimpleschemaGetAPIResponse)
	},
}

// GetTmallItemUpdateSimpleschemaGetAPIResponse 从 sync.Pool 获取 TmallItemUpdateSimpleschemaGetAPIResponse
func GetTmallItemUpdateSimpleschemaGetAPIResponse() *TmallItemUpdateSimpleschemaGetAPIResponse {
	return poolTmallItemUpdateSimpleschemaGetAPIResponse.Get().(*TmallItemUpdateSimpleschemaGetAPIResponse)
}

// ReleaseTmallItemUpdateSimpleschemaGetAPIResponse 将 TmallItemUpdateSimpleschemaGetAPIResponse 保存到 sync.Pool
func ReleaseTmallItemUpdateSimpleschemaGetAPIResponse(v *TmallItemUpdateSimpleschemaGetAPIResponse) {
	v.Reset()
	poolTmallItemUpdateSimpleschemaGetAPIResponse.Put(v)
}
