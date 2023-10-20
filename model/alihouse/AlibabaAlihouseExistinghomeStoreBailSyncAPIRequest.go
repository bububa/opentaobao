package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomestorebailsyncAPIRequest 门店保证金余额同步 API请求
// alibaba.alihouse.existinghome.store.bail.sync
//
// 门店保证金余额同步
type AlibabaalihouseexistinghomestorebailsyncAPIRequest struct {
	model.Params
	// dto
	_storeBailDto *StoreBailDto
}

// NewAlibabaalihouseexistinghomestorebailsyncRequest 初始化AlibabaalihouseexistinghomestorebailsyncAPIRequest对象
func NewAlibabaalihouseexistinghomestorebailsyncRequest() *AlibabaalihouseexistinghomestorebailsyncAPIRequest {
	return &AlibabaalihouseexistinghomestorebailsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomestorebailsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.store.bail.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomestorebailsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomestorebailsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreBailDto is StoreBailDto Setter
// dto
func (r *AlibabaalihouseexistinghomestorebailsyncAPIRequest) SetStoreBailDto(_storeBailDto *StoreBailDto) error {
	r._storeBailDto = _storeBailDto
	r.Set("store_bail_dto", _storeBailDto)
	return nil
}

// GetStoreBailDto StoreBailDto Getter
func (r AlibabaalihouseexistinghomestorebailsyncAPIRequest) GetStoreBailDto() *StoreBailDto {
	return r._storeBailDto
}
