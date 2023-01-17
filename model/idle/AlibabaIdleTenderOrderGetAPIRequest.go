package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleTenderOrderGetAPIRequest 暗拍读取订单 API请求
// alibaba.idle.tender.order.get
//
// 查询省心卖暗拍项目订单
type AlibabaIdleTenderOrderGetAPIRequest struct {
	model.Params
	// 查询信息
	_param0 *TenderOrderQuery
}

// NewAlibabaIdleTenderOrderGetRequest 初始化AlibabaIdleTenderOrderGetAPIRequest对象
func NewAlibabaIdleTenderOrderGetRequest() *AlibabaIdleTenderOrderGetAPIRequest {
	return &AlibabaIdleTenderOrderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleTenderOrderGetAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.tender.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleTenderOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleTenderOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 查询信息
func (r *AlibabaIdleTenderOrderGetAPIRequest) SetParam0(_param0 *TenderOrderQuery) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaIdleTenderOrderGetAPIRequest) GetParam0() *TenderOrderQuery {
	return r._param0
}
