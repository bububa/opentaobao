package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest 预售商家仓接单 API请求
// alibaba.ascp.uop.taobao.presalesorder.create
//
// 预售商家仓接单
type AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest struct {
	model.Params
	// 预售商家仓接单对象
	_presalesOrderCreateRequest *PresalesordercreaterequestTest
}

// NewAlibabaAscpUopTaobaoPresalesorderCreateRequest 初始化AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest对象
func NewAlibabaAscpUopTaobaoPresalesorderCreateRequest() *AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest {
	return &AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest) Reset() {
	r._presalesOrderCreateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.taobao.presalesorder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPresalesOrderCreateRequest is PresalesOrderCreateRequest Setter
// 预售商家仓接单对象
func (r *AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest) SetPresalesOrderCreateRequest(_presalesOrderCreateRequest *PresalesordercreaterequestTest) error {
	r._presalesOrderCreateRequest = _presalesOrderCreateRequest
	r.Set("presales_order_create_request", _presalesOrderCreateRequest)
	return nil
}

// GetPresalesOrderCreateRequest PresalesOrderCreateRequest Getter
func (r AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest) GetPresalesOrderCreateRequest() *PresalesordercreaterequestTest {
	return r._presalesOrderCreateRequest
}

var poolAlibabaAscpUopTaobaoPresalesorderCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpUopTaobaoPresalesorderCreateRequest()
	},
}

// GetAlibabaAscpUopTaobaoPresalesorderCreateRequest 从 sync.Pool 获取 AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest
func GetAlibabaAscpUopTaobaoPresalesorderCreateAPIRequest() *AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest {
	return poolAlibabaAscpUopTaobaoPresalesorderCreateAPIRequest.Get().(*AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest)
}

// ReleaseAlibabaAscpUopTaobaoPresalesorderCreateAPIRequest 将 AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpUopTaobaoPresalesorderCreateAPIRequest(v *AlibabaAscpUopTaobaoPresalesorderCreateAPIRequest) {
	v.Reset()
	poolAlibabaAscpUopTaobaoPresalesorderCreateAPIRequest.Put(v)
}
