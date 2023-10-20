package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascplogisticssellerwritelogisticsnodeAPIRequest 商家配送写入物流节点 API请求
// alibaba.ascp.logistics.seller.writelogisticsnode
//
// 商家配送的订单，商家写入物流节点
type AlibabaascplogisticssellerwritelogisticsnodeAPIRequest struct {
	model.Params
	// 物流节点
	_nodes []LogisticsNodeTopDto
	// 物流发货单号
	_lpOrderId int64
}

// NewAlibabaascplogisticssellerwritelogisticsnodeRequest 初始化AlibabaascplogisticssellerwritelogisticsnodeAPIRequest对象
func NewAlibabaascplogisticssellerwritelogisticsnodeRequest() *AlibabaascplogisticssellerwritelogisticsnodeAPIRequest {
	return &AlibabaascplogisticssellerwritelogisticsnodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascplogisticssellerwritelogisticsnodeAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.logistics.seller.writelogisticsnode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascplogisticssellerwritelogisticsnodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascplogisticssellerwritelogisticsnodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNodes is Nodes Setter
// 物流节点
func (r *AlibabaascplogisticssellerwritelogisticsnodeAPIRequest) SetNodes(_nodes []LogisticsNodeTopDto) error {
	r._nodes = _nodes
	r.Set("nodes", _nodes)
	return nil
}

// GetNodes Nodes Getter
func (r AlibabaascplogisticssellerwritelogisticsnodeAPIRequest) GetNodes() []LogisticsNodeTopDto {
	return r._nodes
}

// SetLpOrderId is LpOrderId Setter
// 物流发货单号
func (r *AlibabaascplogisticssellerwritelogisticsnodeAPIRequest) SetLpOrderId(_lpOrderId int64) error {
	r._lpOrderId = _lpOrderId
	r.Set("lp_order_id", _lpOrderId)
	return nil
}

// GetLpOrderId LpOrderId Getter
func (r AlibabaascplogisticssellerwritelogisticsnodeAPIRequest) GetLpOrderId() int64 {
	return r._lpOrderId
}
