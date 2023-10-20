package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbucategoryattributegetAPIRequest 类目属性获取 API请求
// alibaba.icbu.category.attribute.get
//
// 根据类目ID获取系统定义的属性
type AlibabaicbucategoryattributegetAPIRequest struct {
	model.Params
	// 发布类目id
	_catId int64
}

// NewAlibabaicbucategoryattributegetRequest 初始化AlibabaicbucategoryattributegetAPIRequest对象
func NewAlibabaicbucategoryattributegetRequest() *AlibabaicbucategoryattributegetAPIRequest {
	return &AlibabaicbucategoryattributegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbucategoryattributegetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.category.attribute.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbucategoryattributegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbucategoryattributegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCatId is CatId Setter
// 发布类目id
func (r *AlibabaicbucategoryattributegetAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r AlibabaicbucategoryattributegetAPIRequest) GetCatId() int64 {
	return r._catId
}
