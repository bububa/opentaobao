package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuPhotobankGroupListAPIRequest 图片银行分组信息获取 API请求
// alibaba.icbu.photobank.group.list
//
// 图片银行分组信息获取
type AlibabaIcbuPhotobankGroupListAPIRequest struct {
	model.Params
	// 补充信息
	_extraContext string
	// 查询图片分组信息，如果传入id，则获取当前分组和所有子分组信息，否则获取所有一级分组信息
	_id int64
}

// NewAlibabaIcbuPhotobankGroupListRequest 初始化AlibabaIcbuPhotobankGroupListAPIRequest对象
func NewAlibabaIcbuPhotobankGroupListRequest() *AlibabaIcbuPhotobankGroupListAPIRequest {
	return &AlibabaIcbuPhotobankGroupListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuPhotobankGroupListAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.photobank.group.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuPhotobankGroupListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetExtraContext is ExtraContext Setter
// 补充信息
func (r *AlibabaIcbuPhotobankGroupListAPIRequest) SetExtraContext(_extraContext string) error {
	r._extraContext = _extraContext
	r.Set("extra_context", _extraContext)
	return nil
}

// GetExtraContext ExtraContext Getter
func (r AlibabaIcbuPhotobankGroupListAPIRequest) GetExtraContext() string {
	return r._extraContext
}

// SetId is Id Setter
// 查询图片分组信息，如果传入id，则获取当前分组和所有子分组信息，否则获取所有一级分组信息
func (r *AlibabaIcbuPhotobankGroupListAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaIcbuPhotobankGroupListAPIRequest) GetId() int64 {
	return r._id
}
