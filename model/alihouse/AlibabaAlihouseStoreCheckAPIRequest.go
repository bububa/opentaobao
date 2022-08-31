package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseStoreCheckAPIRequest 门店对账查询工具 API请求
// alibaba.alihouse.store.check
//
// 门店对账查询工具
type AlibabaAlihouseStoreCheckAPIRequest struct {
	model.Params
	// 外部id列表
	_outerIds []string
}

// NewAlibabaAlihouseStoreCheckRequest 初始化AlibabaAlihouseStoreCheckAPIRequest对象
func NewAlibabaAlihouseStoreCheckRequest() *AlibabaAlihouseStoreCheckAPIRequest {
	return &AlibabaAlihouseStoreCheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseStoreCheckAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.store.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseStoreCheckAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOuterIds is OuterIds Setter
// 外部id列表
func (r *AlibabaAlihouseStoreCheckAPIRequest) SetOuterIds(_outerIds []string) error {
	r._outerIds = _outerIds
	r.Set("outer_ids", _outerIds)
	return nil
}

// GetOuterIds OuterIds Getter
func (r AlibabaAlihouseStoreCheckAPIRequest) GetOuterIds() []string {
	return r._outerIds
}
