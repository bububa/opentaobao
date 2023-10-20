package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpuoptaobaopresalesordercreateAPIRequest 预售商家仓接单 API请求
// alibaba.ascp.uop.taobao.presalesorder.create
//
// 预售商家仓接单
type AlibabaascpuoptaobaopresalesordercreateAPIRequest struct {
	model.Params
	// 预售商家仓接单对象
	_presalesOrderCreateRequest *PresalesordercreaterequestTest
}

// NewAlibabaascpuoptaobaopresalesordercreateRequest 初始化AlibabaascpuoptaobaopresalesordercreateAPIRequest对象
func NewAlibabaascpuoptaobaopresalesordercreateRequest() *AlibabaascpuoptaobaopresalesordercreateAPIRequest {
	return &AlibabaascpuoptaobaopresalesordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpuoptaobaopresalesordercreateAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.taobao.presalesorder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpuoptaobaopresalesordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpuoptaobaopresalesordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPresalesOrderCreateRequest is PresalesOrderCreateRequest Setter
// 预售商家仓接单对象
func (r *AlibabaascpuoptaobaopresalesordercreateAPIRequest) SetPresalesOrderCreateRequest(_presalesOrderCreateRequest *PresalesordercreaterequestTest) error {
	r._presalesOrderCreateRequest = _presalesOrderCreateRequest
	r.Set("presales_order_create_request", _presalesOrderCreateRequest)
	return nil
}

// GetPresalesOrderCreateRequest PresalesOrderCreateRequest Getter
func (r AlibabaascpuoptaobaopresalesordercreateAPIRequest) GetPresalesOrderCreateRequest() *PresalesordercreaterequestTest {
	return r._presalesOrderCreateRequest
}
