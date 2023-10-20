package eleenterpriseordernew

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleEnterpriseOrdernewGetstatusAPIRequest 订单状态查询接口 API请求
// alibaba.ele.enterprise.ordernew.getstatus
//
// 订单状态查询接口
type AlibabaEleEnterpriseOrdernewGetstatusAPIRequest struct {
	model.Params
	// 订单号
	_elemeOrderId string
}

// NewAlibabaEleEnterpriseOrdernewGetstatusRequest 初始化AlibabaEleEnterpriseOrdernewGetstatusAPIRequest对象
func NewAlibabaEleEnterpriseOrdernewGetstatusRequest() *AlibabaEleEnterpriseOrdernewGetstatusAPIRequest {
	return &AlibabaEleEnterpriseOrdernewGetstatusAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEleEnterpriseOrdernewGetstatusAPIRequest) Reset() {
	r._elemeOrderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEleEnterpriseOrdernewGetstatusAPIRequest) GetApiMethodName() string {
	return "alibaba.ele.enterprise.ordernew.getstatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEleEnterpriseOrdernewGetstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEleEnterpriseOrdernewGetstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetElemeOrderId is ElemeOrderId Setter
// 订单号
func (r *AlibabaEleEnterpriseOrdernewGetstatusAPIRequest) SetElemeOrderId(_elemeOrderId string) error {
	r._elemeOrderId = _elemeOrderId
	r.Set("eleme_order_id", _elemeOrderId)
	return nil
}

// GetElemeOrderId ElemeOrderId Getter
func (r AlibabaEleEnterpriseOrdernewGetstatusAPIRequest) GetElemeOrderId() string {
	return r._elemeOrderId
}

var poolAlibabaEleEnterpriseOrdernewGetstatusAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEleEnterpriseOrdernewGetstatusRequest()
	},
}

// GetAlibabaEleEnterpriseOrdernewGetstatusRequest 从 sync.Pool 获取 AlibabaEleEnterpriseOrdernewGetstatusAPIRequest
func GetAlibabaEleEnterpriseOrdernewGetstatusAPIRequest() *AlibabaEleEnterpriseOrdernewGetstatusAPIRequest {
	return poolAlibabaEleEnterpriseOrdernewGetstatusAPIRequest.Get().(*AlibabaEleEnterpriseOrdernewGetstatusAPIRequest)
}

// ReleaseAlibabaEleEnterpriseOrdernewGetstatusAPIRequest 将 AlibabaEleEnterpriseOrdernewGetstatusAPIRequest 放入 sync.Pool
func ReleaseAlibabaEleEnterpriseOrdernewGetstatusAPIRequest(v *AlibabaEleEnterpriseOrdernewGetstatusAPIRequest) {
	v.Reset()
	poolAlibabaEleEnterpriseOrdernewGetstatusAPIRequest.Put(v)
}
