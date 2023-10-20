package lstbm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstbmstoreupdateAPIRequest 修改品牌商自有门店数据 API请求
// alibaba.lst.bm.store.update
//
// 修改品牌商自有门店数据
type AlibabalstbmstoreupdateAPIRequest struct {
	model.Params
	// 门店数据模型
	_openStoreDto *LstTopOpenStoreDto
}

// NewAlibabalstbmstoreupdateRequest 初始化AlibabalstbmstoreupdateAPIRequest对象
func NewAlibabalstbmstoreupdateRequest() *AlibabalstbmstoreupdateAPIRequest {
	return &AlibabalstbmstoreupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstbmstoreupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.bm.store.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstbmstoreupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstbmstoreupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenStoreDto is OpenStoreDto Setter
// 门店数据模型
func (r *AlibabalstbmstoreupdateAPIRequest) SetOpenStoreDto(_openStoreDto *LstTopOpenStoreDto) error {
	r._openStoreDto = _openStoreDto
	r.Set("open_store_dto", _openStoreDto)
	return nil
}

// GetOpenStoreDto OpenStoreDto Getter
func (r AlibabalstbmstoreupdateAPIRequest) GetOpenStoreDto() *LstTopOpenStoreDto {
	return r._openStoreDto
}
