package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbucategorygetAPIRequest 商品发布类目获取 API请求
// alibaba.icbu.category.get
//
// 获取商品发布类目
type AlibabaicbucategorygetAPIRequest struct {
	model.Params
	// 发布类目id,必须大于等于0， 如果为0，则查询所有一级类目
	_catId int64
}

// NewAlibabaicbucategorygetRequest 初始化AlibabaicbucategorygetAPIRequest对象
func NewAlibabaicbucategorygetRequest() *AlibabaicbucategorygetAPIRequest {
	return &AlibabaicbucategorygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbucategorygetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.category.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbucategorygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbucategorygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCatId is CatId Setter
// 发布类目id,必须大于等于0， 如果为0，则查询所有一级类目
func (r *AlibabaicbucategorygetAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r AlibabaicbucategorygetAPIRequest) GetCatId() int64 {
	return r._catId
}
