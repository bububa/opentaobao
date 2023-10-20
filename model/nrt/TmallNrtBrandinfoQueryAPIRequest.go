package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtbrandinfoqueryAPIRequest 品牌数据查询 API请求
// tmall.nrt.brandinfo.query
//
// 商家获取自己旗舰店授权的品牌id列表
type TmallnrtbrandinfoqueryAPIRequest struct {
	model.Params
}

// NewTmallnrtbrandinfoqueryRequest 初始化TmallnrtbrandinfoqueryAPIRequest对象
func NewTmallnrtbrandinfoqueryRequest() *TmallnrtbrandinfoqueryAPIRequest {
	return &TmallnrtbrandinfoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtbrandinfoqueryAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.brandinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtbrandinfoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtbrandinfoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}
