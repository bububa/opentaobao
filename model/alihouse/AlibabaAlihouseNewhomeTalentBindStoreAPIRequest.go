package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhometalentbindstoreAPIRequest 达人号门店关系绑定 API请求
// alibaba.alihouse.newhome.talent.bind.store
//
// 达人号门店关系绑定
type AlibabaalihousenewhometalentbindstoreAPIRequest struct {
	model.Params
	// aaa
	_storeTalentList *StoreTalentDto
}

// NewAlibabaalihousenewhometalentbindstoreRequest 初始化AlibabaalihousenewhometalentbindstoreAPIRequest对象
func NewAlibabaalihousenewhometalentbindstoreRequest() *AlibabaalihousenewhometalentbindstoreAPIRequest {
	return &AlibabaalihousenewhometalentbindstoreAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhometalentbindstoreAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.talent.bind.store"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhometalentbindstoreAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhometalentbindstoreAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreTalentList is StoreTalentList Setter
// aaa
func (r *AlibabaalihousenewhometalentbindstoreAPIRequest) SetStoreTalentList(_storeTalentList *StoreTalentDto) error {
	r._storeTalentList = _storeTalentList
	r.Set("store_talent_list", _storeTalentList)
	return nil
}

// GetStoreTalentList StoreTalentList Getter
func (r AlibabaalihousenewhometalentbindstoreAPIRequest) GetStoreTalentList() *StoreTalentDto {
	return r._storeTalentList
}
