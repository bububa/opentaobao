package tblogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest 商家配送写入物流节点 API请求
// alibaba.ascp.logistics.seller.writelogisticsnode
//
// 商家配送的订单，商家写入物流节点
type AlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest struct {
	model.Params
	// 物流节点
	_nodes []LogisticsNodeTopDto
	// 物流发货单号
	_lpOrderId int64
}

// NewAlibabaAscpLogisticsSellerWritelogisticsnodeRequest 初始化AlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest对象
func NewAlibabaAscpLogisticsSellerWritelogisticsnodeRequest() *AlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest {
	return &AlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest) Reset() {
	r._nodes = r._nodes[:0]
	r._lpOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.logistics.seller.writelogisticsnode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNodes is Nodes Setter
// 物流节点
func (r *AlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest) SetNodes(_nodes []LogisticsNodeTopDto) error {
	r._nodes = _nodes
	r.Set("nodes", _nodes)
	return nil
}

// GetNodes Nodes Getter
func (r AlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest) GetNodes() []LogisticsNodeTopDto {
	return r._nodes
}

// SetLpOrderId is LpOrderId Setter
// 物流发货单号
func (r *AlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest) SetLpOrderId(_lpOrderId int64) error {
	r._lpOrderId = _lpOrderId
	r.Set("lp_order_id", _lpOrderId)
	return nil
}

// GetLpOrderId LpOrderId Getter
func (r AlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest) GetLpOrderId() int64 {
	return r._lpOrderId
}

var poolAlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpLogisticsSellerWritelogisticsnodeRequest()
	},
}

// GetAlibabaAscpLogisticsSellerWritelogisticsnodeRequest 从 sync.Pool 获取 AlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest
func GetAlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest() *AlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest {
	return poolAlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest.Get().(*AlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest)
}

// ReleaseAlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest 将 AlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest(v *AlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest) {
	v.Reset()
	poolAlibabaAscpLogisticsSellerWritelogisticsnodeAPIRequest.Put(v)
}
