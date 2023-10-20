package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuphotobankgrouplistAPIRequest 图片银行分组信息获取 API请求
// alibaba.icbu.photobank.group.list
//
// 图片银行分组信息获取
type AlibabaicbuphotobankgrouplistAPIRequest struct {
	model.Params
	// 补充信息
	_extraContext string
	// 查询图片分组信息，如果传入id，则获取当前分组和所有子分组信息，否则获取所有一级分组信息
	_id int64
}

// NewAlibabaicbuphotobankgrouplistRequest 初始化AlibabaicbuphotobankgrouplistAPIRequest对象
func NewAlibabaicbuphotobankgrouplistRequest() *AlibabaicbuphotobankgrouplistAPIRequest {
	return &AlibabaicbuphotobankgrouplistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbuphotobankgrouplistAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.photobank.group.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbuphotobankgrouplistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbuphotobankgrouplistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtraContext is ExtraContext Setter
// 补充信息
func (r *AlibabaicbuphotobankgrouplistAPIRequest) SetExtraContext(_extraContext string) error {
	r._extraContext = _extraContext
	r.Set("extra_context", _extraContext)
	return nil
}

// GetExtraContext ExtraContext Getter
func (r AlibabaicbuphotobankgrouplistAPIRequest) GetExtraContext() string {
	return r._extraContext
}

// SetId is Id Setter
// 查询图片分组信息，如果传入id，则获取当前分组和所有子分组信息，否则获取所有一级分组信息
func (r *AlibabaicbuphotobankgrouplistAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlibabaicbuphotobankgrouplistAPIRequest) GetId() int64 {
	return r._id
}
