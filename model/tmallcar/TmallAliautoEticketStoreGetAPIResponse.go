package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoEticketStoreGetAPIResponse 查询电子凭证对应门店信息 API返回值
// tmall.aliauto.eticket.store.get
//
// 查询电子凭证对应门店信息
type TmallAliautoEticketStoreGetAPIResponse struct {
	model.CommonResponse
	TmallAliautoEticketStoreGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallAliautoEticketStoreGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallAliautoEticketStoreGetAPIResponseModel).Reset()
}

// TmallAliautoEticketStoreGetAPIResponseModel is 查询电子凭证对应门店信息 成功返回结果
type TmallAliautoEticketStoreGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_eticket_store_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallAliautoEticketStoreGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallAliautoEticketStoreGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallAliautoEticketStoreGetAPIResponse)
	},
}

// GetTmallAliautoEticketStoreGetAPIResponse 从 sync.Pool 获取 TmallAliautoEticketStoreGetAPIResponse
func GetTmallAliautoEticketStoreGetAPIResponse() *TmallAliautoEticketStoreGetAPIResponse {
	return poolTmallAliautoEticketStoreGetAPIResponse.Get().(*TmallAliautoEticketStoreGetAPIResponse)
}

// ReleaseTmallAliautoEticketStoreGetAPIResponse 将 TmallAliautoEticketStoreGetAPIResponse 保存到 sync.Pool
func ReleaseTmallAliautoEticketStoreGetAPIResponse(v *TmallAliautoEticketStoreGetAPIResponse) {
	v.Reset()
	poolTmallAliautoEticketStoreGetAPIResponse.Put(v)
}
