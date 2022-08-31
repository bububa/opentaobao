package icbulogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest 查询物流运力列表 API请求
// alibaba.onetouch.logistics.express.logistics.product.list
//
// 查询物流产品&amp;揽收仓库列表
type AlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest struct {
	model.Params
}

// NewAlibabaOnetouchLogisticsExpressLogisticsProductListRequest 初始化AlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest对象
func NewAlibabaOnetouchLogisticsExpressLogisticsProductListRequest() *AlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest {
	return &AlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.logistics.product.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOnetouchLogisticsExpressLogisticsProductListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
