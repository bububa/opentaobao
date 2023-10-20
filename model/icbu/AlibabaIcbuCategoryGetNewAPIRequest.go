package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbucategorygetnewAPIRequest （新）ICBU类目树获取接口 API请求
// alibaba.icbu.category.get.new
//
// 获取商品发布类目
type AlibabaicbucategorygetnewAPIRequest struct {
	model.Params
	// 发布类目id,必须大于等于0， 如果为0，则查询所有一级类目
	_catId int64
}

// NewAlibabaicbucategorygetnewRequest 初始化AlibabaicbucategorygetnewAPIRequest对象
func NewAlibabaicbucategorygetnewRequest() *AlibabaicbucategorygetnewAPIRequest {
	return &AlibabaicbucategorygetnewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbucategorygetnewAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.category.get.new"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbucategorygetnewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbucategorygetnewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCatId is CatId Setter
// 发布类目id,必须大于等于0， 如果为0，则查询所有一级类目
func (r *AlibabaicbucategorygetnewAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r AlibabaicbucategorygetnewAPIRequest) GetCatId() int64 {
	return r._catId
}
