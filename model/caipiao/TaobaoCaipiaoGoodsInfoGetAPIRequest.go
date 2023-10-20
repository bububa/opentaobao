package caipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocaipiaogoodsinfogetAPIRequest 根据卖家id与appkey获取商品信息 API请求
// taobao.caipiao.goods.info.get
//
// 根据卖家id与appkey获取商品信息。
type TaobaocaipiaogoodsinfogetAPIRequest struct {
	model.Params
}

// NewTaobaocaipiaogoodsinfogetRequest 初始化TaobaocaipiaogoodsinfogetAPIRequest对象
func NewTaobaocaipiaogoodsinfogetRequest() *TaobaocaipiaogoodsinfogetAPIRequest {
	return &TaobaocaipiaogoodsinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocaipiaogoodsinfogetAPIRequest) GetApiMethodName() string {
	return "taobao.caipiao.goods.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocaipiaogoodsinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocaipiaogoodsinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
