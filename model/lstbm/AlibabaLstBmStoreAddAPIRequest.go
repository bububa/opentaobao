package lstbm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstbmstoreaddAPIRequest 导入品牌商自有门店 API请求
// alibaba.lst.bm.store.add
//
// 导入品牌商自有门店
type AlibabalstbmstoreaddAPIRequest struct {
	model.Params
	// 门店数据模型
	_openStoreDto *LstTopOpenStoreDto
}

// NewAlibabalstbmstoreaddRequest 初始化AlibabalstbmstoreaddAPIRequest对象
func NewAlibabalstbmstoreaddRequest() *AlibabalstbmstoreaddAPIRequest {
	return &AlibabalstbmstoreaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstbmstoreaddAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.bm.store.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstbmstoreaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstbmstoreaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenStoreDto is OpenStoreDto Setter
// 门店数据模型
func (r *AlibabalstbmstoreaddAPIRequest) SetOpenStoreDto(_openStoreDto *LstTopOpenStoreDto) error {
	r._openStoreDto = _openStoreDto
	r.Set("open_store_dto", _openStoreDto)
	return nil
}

// GetOpenStoreDto OpenStoreDto Getter
func (r AlibabalstbmstoreaddAPIRequest) GetOpenStoreDto() *LstTopOpenStoreDto {
	return r._openStoreDto
}
