package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenStoreCreateAPIResponse 门店新增接口 API返回值
// taobao.qimen.store.create
//
// isv调用接口来讲线下门店同步到线上
type TaobaoQimenStoreCreateAPIResponse struct {
	model.CommonResponse
	TaobaoQimenStoreCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenStoreCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenStoreCreateAPIResponseModel).Reset()
}

// TaobaoQimenStoreCreateAPIResponseModel is 门店新增接口 成功返回结果
type TaobaoQimenStoreCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_store_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 响应标示
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应code
	QimenCode string `json:"qimen_code,omitempty" xml:"qimen_code,omitempty"`
	// 返回的门店id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenStoreCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.Flag = ""
	m.QimenCode = ""
	m.StoreId = 0
}

var poolTaobaoQimenStoreCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenStoreCreateAPIResponse)
	},
}

// GetTaobaoQimenStoreCreateAPIResponse 从 sync.Pool 获取 TaobaoQimenStoreCreateAPIResponse
func GetTaobaoQimenStoreCreateAPIResponse() *TaobaoQimenStoreCreateAPIResponse {
	return poolTaobaoQimenStoreCreateAPIResponse.Get().(*TaobaoQimenStoreCreateAPIResponse)
}

// ReleaseTaobaoQimenStoreCreateAPIResponse 将 TaobaoQimenStoreCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenStoreCreateAPIResponse(v *TaobaoQimenStoreCreateAPIResponse) {
	v.Reset()
	poolTaobaoQimenStoreCreateAPIResponse.Put(v)
}
