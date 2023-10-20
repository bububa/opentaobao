package nrt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtItemMainSynchronizeAPIResponse 家装新零售主商品同步至阿里 API返回值
// tmall.nrt.item.main.synchronize
//
// 同步卖场存量线下商品到阿里
type TmallNrtItemMainSynchronizeAPIResponse struct {
	model.CommonResponse
	TmallNrtItemMainSynchronizeAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrtItemMainSynchronizeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrtItemMainSynchronizeAPIResponseModel).Reset()
}

// TmallNrtItemMainSynchronizeAPIResponseModel is 家装新零售主商品同步至阿里 成功返回结果
type TmallNrtItemMainSynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_item_main_synchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	TmallNrtItemMainSynchronize *TmallNrtItemMainSynchronizeResultDo `json:"tmall_nrt_item_main_synchronize,omitempty" xml:"tmall_nrt_item_main_synchronize,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrtItemMainSynchronizeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TmallNrtItemMainSynchronize = nil
}

var poolTmallNrtItemMainSynchronizeAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrtItemMainSynchronizeAPIResponse)
	},
}

// GetTmallNrtItemMainSynchronizeAPIResponse 从 sync.Pool 获取 TmallNrtItemMainSynchronizeAPIResponse
func GetTmallNrtItemMainSynchronizeAPIResponse() *TmallNrtItemMainSynchronizeAPIResponse {
	return poolTmallNrtItemMainSynchronizeAPIResponse.Get().(*TmallNrtItemMainSynchronizeAPIResponse)
}

// ReleaseTmallNrtItemMainSynchronizeAPIResponse 将 TmallNrtItemMainSynchronizeAPIResponse 保存到 sync.Pool
func ReleaseTmallNrtItemMainSynchronizeAPIResponse(v *TmallNrtItemMainSynchronizeAPIResponse) {
	v.Reset()
	poolTmallNrtItemMainSynchronizeAPIResponse.Put(v)
}
