package lstbm

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstBmStoreAddAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.bm.store.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstBmStoreAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
