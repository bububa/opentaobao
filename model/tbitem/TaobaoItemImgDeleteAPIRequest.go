package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemImgDeleteAPIRequest 删除商品图片 API请求
// taobao.item.img.delete
//
// 删除商品图片
type TaobaoItemImgDeleteAPIRequest struct {
	model.Params
}

// NewTaobaoItemImgDeleteRequest 初始化TaobaoItemImgDeleteAPIRequest对象
func NewTaobaoItemImgDeleteRequest() *TaobaoItemImgDeleteAPIRequest {
	return &TaobaoItemImgDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemImgDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.item.img.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemImgDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoItemImgDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}
