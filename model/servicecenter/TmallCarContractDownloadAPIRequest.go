package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCarContractDownloadAPIRequest 合同下载 API请求
// tmall.car.contract.download
//
// 目前天猫开新车会在线上签署一份合同，协议，需要一个个在已卖出打开，另存为pdf，人工一个个下载比较麻烦，期望通过接口直接读取pdf；
// 因为比较耗时，建议一个个下载，假设并发下载，很可能限流，每天的调用量有限；
type TmallCarContractDownloadAPIRequest struct {
	model.Params
	// 天猫订单号
	_orderId int64
	// 是否下载html，true是html，false是pdf， html速度会快一点
	_html bool
}

// NewTmallCarContractDownloadRequest 初始化TmallCarContractDownloadAPIRequest对象
func NewTmallCarContractDownloadRequest() *TmallCarContractDownloadAPIRequest {
	return &TmallCarContractDownloadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarContractDownloadAPIRequest) GetApiMethodName() string {
	return "tmall.car.contract.download"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarContractDownloadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderId is OrderId Setter
// 天猫订单号
func (r *TmallCarContractDownloadAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TmallCarContractDownloadAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetHtml is Html Setter
// 是否下载html，true是html，false是pdf， html速度会快一点
func (r *TmallCarContractDownloadAPIRequest) SetHtml(_html bool) error {
	r._html = _html
	r.Set("html", _html)
	return nil
}

// GetHtml Html Getter
func (r TmallCarContractDownloadAPIRequest) GetHtml() bool {
	return r._html
}
