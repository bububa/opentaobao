package lstbm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstBmStoreAddAPIRequest 导入品牌商自有门店 API请求
// alibaba.lst.bm.store.add
//
// 导入品牌商自有门店
type AlibabaLstBmStoreAddAPIRequest struct {
	model.Params
	// 门店数据模型
	_openStoreDto *LstTopOpenStoreDto
}

// NewAlibabaLstBmStoreAddRequest 初始化AlibabaLstBmStoreAddAPIRequest对象
func NewAlibabaLstBmStoreAddRequest() *AlibabaLstBmStoreAddAPIRequest {
	return &AlibabaLstBmStoreAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstBmStoreAddAPIRequest) Reset() {
	r._openStoreDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstBmStoreAddAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.bm.store.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstBmStoreAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstBmStoreAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenStoreDto is OpenStoreDto Setter
// 门店数据模型
func (r *AlibabaLstBmStoreAddAPIRequest) SetOpenStoreDto(_openStoreDto *LstTopOpenStoreDto) error {
	r._openStoreDto = _openStoreDto
	r.Set("open_store_dto", _openStoreDto)
	return nil
}

// GetOpenStoreDto OpenStoreDto Getter
func (r AlibabaLstBmStoreAddAPIRequest) GetOpenStoreDto() *LstTopOpenStoreDto {
	return r._openStoreDto
}

var poolAlibabaLstBmStoreAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstBmStoreAddRequest()
	},
}

// GetAlibabaLstBmStoreAddRequest 从 sync.Pool 获取 AlibabaLstBmStoreAddAPIRequest
func GetAlibabaLstBmStoreAddAPIRequest() *AlibabaLstBmStoreAddAPIRequest {
	return poolAlibabaLstBmStoreAddAPIRequest.Get().(*AlibabaLstBmStoreAddAPIRequest)
}

// ReleaseAlibabaLstBmStoreAddAPIRequest 将 AlibabaLstBmStoreAddAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstBmStoreAddAPIRequest(v *AlibabaLstBmStoreAddAPIRequest) {
	v.Reset()
	poolAlibabaLstBmStoreAddAPIRequest.Put(v)
}
