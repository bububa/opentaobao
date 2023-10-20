package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuproductgroupaddAPIRequest 增加商品分组 API请求
// alibaba.icbu.product.group.add
//
// 增加商品分组
type AlibabaicbuproductgroupaddAPIRequest struct {
	model.Params
	// 分组名称
	_groupName string
	// 补充信息，如isv id
	_extraContext string
	// 上级分组ID，如果建立顶级分组设为-1
	_parentId int64
}

// NewAlibabaicbuproductgroupaddRequest 初始化AlibabaicbuproductgroupaddAPIRequest对象
func NewAlibabaicbuproductgroupaddRequest() *AlibabaicbuproductgroupaddAPIRequest {
	return &AlibabaicbuproductgroupaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbuproductgroupaddAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.group.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbuproductgroupaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbuproductgroupaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGroupName is GroupName Setter
// 分组名称
func (r *AlibabaicbuproductgroupaddAPIRequest) SetGroupName(_groupName string) error {
	r._groupName = _groupName
	r.Set("group_name", _groupName)
	return nil
}

// GetGroupName GroupName Getter
func (r AlibabaicbuproductgroupaddAPIRequest) GetGroupName() string {
	return r._groupName
}

// SetExtraContext is ExtraContext Setter
// 补充信息，如isv id
func (r *AlibabaicbuproductgroupaddAPIRequest) SetExtraContext(_extraContext string) error {
	r._extraContext = _extraContext
	r.Set("extra_context", _extraContext)
	return nil
}

// GetExtraContext ExtraContext Getter
func (r AlibabaicbuproductgroupaddAPIRequest) GetExtraContext() string {
	return r._extraContext
}

// SetParentId is ParentId Setter
// 上级分组ID，如果建立顶级分组设为-1
func (r *AlibabaicbuproductgroupaddAPIRequest) SetParentId(_parentId int64) error {
	r._parentId = _parentId
	r.Set("parent_id", _parentId)
	return nil
}

// GetParentId ParentId Getter
func (r AlibabaicbuproductgroupaddAPIRequest) GetParentId() int64 {
	return r._parentId
}
