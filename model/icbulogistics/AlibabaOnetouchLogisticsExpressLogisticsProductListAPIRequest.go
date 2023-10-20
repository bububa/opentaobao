package icbulogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaonetouchlogisticsexpresslogisticsproductlistAPIRequest 查询物流运力列表 API请求
// alibaba.onetouch.logistics.express.logistics.product.list
//
// 查询物流产品&amp;揽收仓库列表
type AlibabaonetouchlogisticsexpresslogisticsproductlistAPIRequest struct {
	model.Params
}

// NewAlibabaonetouchlogisticsexpresslogisticsproductlistRequest 初始化AlibabaonetouchlogisticsexpresslogisticsproductlistAPIRequest对象
func NewAlibabaonetouchlogisticsexpresslogisticsproductlistRequest() *AlibabaonetouchlogisticsexpresslogisticsproductlistAPIRequest {
	return &AlibabaonetouchlogisticsexpresslogisticsproductlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaonetouchlogisticsexpresslogisticsproductlistAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.logistics.product.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaonetouchlogisticsexpresslogisticsproductlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaonetouchlogisticsexpresslogisticsproductlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}
