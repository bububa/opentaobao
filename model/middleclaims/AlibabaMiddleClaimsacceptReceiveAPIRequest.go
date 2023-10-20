package middleclaims

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMiddleClaimsacceptReceiveAPIRequest 国际化中台服务域接收保险公司理赔受理结果 API请求
// alibaba.middle.claimsaccept.receive
//
// 国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔受理结果的处理后，将该结果返回至服务域
type AlibabaMiddleClaimsacceptReceiveAPIRequest struct {
	model.Params
	// 理赔受理结果实体类
	_claimsAcceptDto *ClaimsAcceptDto
}

// NewAlibabaMiddleClaimsacceptReceiveRequest 初始化AlibabaMiddleClaimsacceptReceiveAPIRequest对象
func NewAlibabaMiddleClaimsacceptReceiveRequest() *AlibabaMiddleClaimsacceptReceiveAPIRequest {
	return &AlibabaMiddleClaimsacceptReceiveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMiddleClaimsacceptReceiveAPIRequest) Reset() {
	r._claimsAcceptDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMiddleClaimsacceptReceiveAPIRequest) GetApiMethodName() string {
	return "alibaba.middle.claimsaccept.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMiddleClaimsacceptReceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMiddleClaimsacceptReceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClaimsAcceptDto is ClaimsAcceptDto Setter
// 理赔受理结果实体类
func (r *AlibabaMiddleClaimsacceptReceiveAPIRequest) SetClaimsAcceptDto(_claimsAcceptDto *ClaimsAcceptDto) error {
	r._claimsAcceptDto = _claimsAcceptDto
	r.Set("claims_accept_dto", _claimsAcceptDto)
	return nil
}

// GetClaimsAcceptDto ClaimsAcceptDto Getter
func (r AlibabaMiddleClaimsacceptReceiveAPIRequest) GetClaimsAcceptDto() *ClaimsAcceptDto {
	return r._claimsAcceptDto
}

var poolAlibabaMiddleClaimsacceptReceiveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMiddleClaimsacceptReceiveRequest()
	},
}

// GetAlibabaMiddleClaimsacceptReceiveRequest 从 sync.Pool 获取 AlibabaMiddleClaimsacceptReceiveAPIRequest
func GetAlibabaMiddleClaimsacceptReceiveAPIRequest() *AlibabaMiddleClaimsacceptReceiveAPIRequest {
	return poolAlibabaMiddleClaimsacceptReceiveAPIRequest.Get().(*AlibabaMiddleClaimsacceptReceiveAPIRequest)
}

// ReleaseAlibabaMiddleClaimsacceptReceiveAPIRequest 将 AlibabaMiddleClaimsacceptReceiveAPIRequest 放入 sync.Pool
func ReleaseAlibabaMiddleClaimsacceptReceiveAPIRequest(v *AlibabaMiddleClaimsacceptReceiveAPIRequest) {
	v.Reset()
	poolAlibabaMiddleClaimsacceptReceiveAPIRequest.Put(v)
}
