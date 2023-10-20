package lstbm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstBmStoreUpdateAPIRequest 修改品牌商自有门店数据 API请求
// alibaba.lst.bm.store.update
//
// 修改品牌商自有门店数据
type AlibabaLstBmStoreUpdateAPIRequest struct {
	model.Params
	// 门店数据模型
	_openStoreDto *LstTopOpenStoreDto
}

// NewAlibabaLstBmStoreUpdateRequest 初始化AlibabaLstBmStoreUpdateAPIRequest对象
func NewAlibabaLstBmStoreUpdateRequest() *AlibabaLstBmStoreUpdateAPIRequest {
	return &AlibabaLstBmStoreUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstBmStoreUpdateAPIRequest) Reset() {
	r._openStoreDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstBmStoreUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.bm.store.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstBmStoreUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstBmStoreUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenStoreDto is OpenStoreDto Setter
// 门店数据模型
func (r *AlibabaLstBmStoreUpdateAPIRequest) SetOpenStoreDto(_openStoreDto *LstTopOpenStoreDto) error {
	r._openStoreDto = _openStoreDto
	r.Set("open_store_dto", _openStoreDto)
	return nil
}

// GetOpenStoreDto OpenStoreDto Getter
func (r AlibabaLstBmStoreUpdateAPIRequest) GetOpenStoreDto() *LstTopOpenStoreDto {
	return r._openStoreDto
}

var poolAlibabaLstBmStoreUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstBmStoreUpdateRequest()
	},
}

// GetAlibabaLstBmStoreUpdateRequest 从 sync.Pool 获取 AlibabaLstBmStoreUpdateAPIRequest
func GetAlibabaLstBmStoreUpdateAPIRequest() *AlibabaLstBmStoreUpdateAPIRequest {
	return poolAlibabaLstBmStoreUpdateAPIRequest.Get().(*AlibabaLstBmStoreUpdateAPIRequest)
}

// ReleaseAlibabaLstBmStoreUpdateAPIRequest 将 AlibabaLstBmStoreUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstBmStoreUpdateAPIRequest(v *AlibabaLstBmStoreUpdateAPIRequest) {
	v.Reset()
	poolAlibabaLstBmStoreUpdateAPIRequest.Put(v)
}
