package uscesl

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslBizStoreInsertAPIResponse 新增电子价签商家门店接口 API返回值
// taobao.uscesl.biz.store.insert
//
// 新增电子价签商家门店接口
type TaobaoUsceslBizStoreInsertAPIResponse struct {
	model.CommonResponse
	TaobaoUsceslBizStoreInsertAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUsceslBizStoreInsertAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUsceslBizStoreInsertAPIResponseModel).Reset()
}

// TaobaoUsceslBizStoreInsertAPIResponseModel is 新增电子价签商家门店接口 成功返回结果
type TaobaoUsceslBizStoreInsertAPIResponseModel struct {
	XMLName xml.Name `xml:"uscesl_biz_store_insert_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUsceslBizStoreInsertAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTaobaoUsceslBizStoreInsertAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUsceslBizStoreInsertAPIResponse)
	},
}

// GetTaobaoUsceslBizStoreInsertAPIResponse 从 sync.Pool 获取 TaobaoUsceslBizStoreInsertAPIResponse
func GetTaobaoUsceslBizStoreInsertAPIResponse() *TaobaoUsceslBizStoreInsertAPIResponse {
	return poolTaobaoUsceslBizStoreInsertAPIResponse.Get().(*TaobaoUsceslBizStoreInsertAPIResponse)
}

// ReleaseTaobaoUsceslBizStoreInsertAPIResponse 将 TaobaoUsceslBizStoreInsertAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUsceslBizStoreInsertAPIResponse(v *TaobaoUsceslBizStoreInsertAPIResponse) {
	v.Reset()
	poolTaobaoUsceslBizStoreInsertAPIResponse.Put(v)
}
