package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuproductgroupgetAPIRequest 分组信息获取 API请求
// alibaba.icbu.product.group.get
//
// 分组信息获取
type AlibabaicbuproductgroupgetAPIRequest struct {
	model.Params
	// 补充信息
	_extraContext string
	// 分组ID，传-1获得所有一级分组
	_groupId int64
}

// NewAlibabaicbuproductgroupgetRequest 初始化AlibabaicbuproductgroupgetAPIRequest对象
func NewAlibabaicbuproductgroupgetRequest() *AlibabaicbuproductgroupgetAPIRequest {
	return &AlibabaicbuproductgroupgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbuproductgroupgetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.group.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbuproductgroupgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbuproductgroupgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtraContext is ExtraContext Setter
// 补充信息
func (r *AlibabaicbuproductgroupgetAPIRequest) SetExtraContext(_extraContext string) error {
	r._extraContext = _extraContext
	r.Set("extra_context", _extraContext)
	return nil
}

// GetExtraContext ExtraContext Getter
func (r AlibabaicbuproductgroupgetAPIRequest) GetExtraContext() string {
	return r._extraContext
}

// SetGroupId is GroupId Setter
// 分组ID，传-1获得所有一级分组
func (r *AlibabaicbuproductgroupgetAPIRequest) SetGroupId(_groupId int64) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r AlibabaicbuproductgroupgetAPIRequest) GetGroupId() int64 {
	return r._groupId
}
