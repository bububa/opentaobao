package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpproductgroupgetAPIRequest 查询指定产品分组的下一层子分组 API请求
// alibaba.scbp.product.group.get
//
// 查询指定产品分组的下一层子分组
type AlibabascbpproductgroupgetAPIRequest struct {
	model.Params
	// 产品分组标识，null表示查询第一层分组
	_groupId string
}

// NewAlibabascbpproductgroupgetRequest 初始化AlibabascbpproductgroupgetAPIRequest对象
func NewAlibabascbpproductgroupgetRequest() *AlibabascbpproductgroupgetAPIRequest {
	return &AlibabascbpproductgroupgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpproductgroupgetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.product.group.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpproductgroupgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpproductgroupgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGroupId is GroupId Setter
// 产品分组标识，null表示查询第一层分组
func (r *AlibabascbpproductgroupgetAPIRequest) SetGroupId(_groupId string) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r AlibabascbpproductgroupgetAPIRequest) GetGroupId() string {
	return r._groupId
}
