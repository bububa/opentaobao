package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductGroupAddAPIRequest 增加商品分组 API请求
// alibaba.icbu.product.group.add
//
// 增加商品分组
type AlibabaIcbuProductGroupAddAPIRequest struct {
	model.Params
	// 分组名称
	_groupName string
	// 上级分组ID，如果建立顶级分组设为-1
	_parentId int64
	// 补充信息，如isv id
	_extraContext string
}

// NewAlibabaIcbuProductGroupAddRequest 初始化AlibabaIcbuProductGroupAddAPIRequest对象
func NewAlibabaIcbuProductGroupAddRequest() *AlibabaIcbuProductGroupAddAPIRequest {
	return &AlibabaIcbuProductGroupAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuProductGroupAddAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.product.group.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuProductGroupAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetGroupName is GroupName Setter
// 分组名称
func (r *AlibabaIcbuProductGroupAddAPIRequest) SetGroupName(_groupName string) error {
	r._groupName = _groupName
	r.Set("group_name", _groupName)
	return nil
}

// GetGroupName GroupName Getter
func (r AlibabaIcbuProductGroupAddAPIRequest) GetGroupName() string {
	return r._groupName
}

// SetParentId is ParentId Setter
// 上级分组ID，如果建立顶级分组设为-1
func (r *AlibabaIcbuProductGroupAddAPIRequest) SetParentId(_parentId int64) error {
	r._parentId = _parentId
	r.Set("parent_id", _parentId)
	return nil
}

// GetParentId ParentId Getter
func (r AlibabaIcbuProductGroupAddAPIRequest) GetParentId() int64 {
	return r._parentId
}

// SetExtraContext is ExtraContext Setter
// 补充信息，如isv id
func (r *AlibabaIcbuProductGroupAddAPIRequest) SetExtraContext(_extraContext string) error {
	r._extraContext = _extraContext
	r.Set("extra_context", _extraContext)
	return nil
}

// GetExtraContext ExtraContext Getter
func (r AlibabaIcbuProductGroupAddAPIRequest) GetExtraContext() string {
	return r._extraContext
}
