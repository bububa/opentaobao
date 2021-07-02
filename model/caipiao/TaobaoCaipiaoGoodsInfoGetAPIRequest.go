package caipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCaipiaoGoodsInfoGetAPIRequest 根据卖家id与appkey获取商品信息 API请求
// taobao.caipiao.goods.info.get
//
// 根据卖家id与appkey获取商品信息。
type TaobaoCaipiaoGoodsInfoGetAPIRequest struct {
	model.Params
}

// NewTaobaoCaipiaoGoodsInfoGetRequest 初始化TaobaoCaipiaoGoodsInfoGetAPIRequest对象
func NewTaobaoCaipiaoGoodsInfoGetRequest() *TaobaoCaipiaoGoodsInfoGetAPIRequest {
	return &TaobaoCaipiaoGoodsInfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCaipiaoGoodsInfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.caipiao.goods.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCaipiaoGoodsInfoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
