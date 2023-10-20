package middleclaims

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMiddleClaimsresultReceiveAPIRequest 国际化中台服务域接收理赔结果 API请求
// alibaba.middle.claimsresult.receive
//
// 国际化中台服务域与保险公司交互对接一个订单在保险公司方对该订单进行理赔结果的处理后，将该结果返回至服务域
type AlibabaMiddleClaimsresultReceiveAPIRequest struct {
	model.Params
	// 理赔结果实体
	_claimsResultDTO *ClaimsResultDto
}

// NewAlibabaMiddleClaimsresultReceiveRequest 初始化AlibabaMiddleClaimsresultReceiveAPIRequest对象
func NewAlibabaMiddleClaimsresultReceiveRequest() *AlibabaMiddleClaimsresultReceiveAPIRequest {
	return &AlibabaMiddleClaimsresultReceiveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMiddleClaimsresultReceiveAPIRequest) Reset() {
	r._claimsResultDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMiddleClaimsresultReceiveAPIRequest) GetApiMethodName() string {
	return "alibaba.middle.claimsresult.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMiddleClaimsresultReceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMiddleClaimsresultReceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClaimsResultDTO is ClaimsResultDTO Setter
// 理赔结果实体
func (r *AlibabaMiddleClaimsresultReceiveAPIRequest) SetClaimsResultDTO(_claimsResultDTO *ClaimsResultDto) error {
	r._claimsResultDTO = _claimsResultDTO
	r.Set("claims_result_d_t_o", _claimsResultDTO)
	return nil
}

// GetClaimsResultDTO ClaimsResultDTO Getter
func (r AlibabaMiddleClaimsresultReceiveAPIRequest) GetClaimsResultDTO() *ClaimsResultDto {
	return r._claimsResultDTO
}

var poolAlibabaMiddleClaimsresultReceiveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMiddleClaimsresultReceiveRequest()
	},
}

// GetAlibabaMiddleClaimsresultReceiveRequest 从 sync.Pool 获取 AlibabaMiddleClaimsresultReceiveAPIRequest
func GetAlibabaMiddleClaimsresultReceiveAPIRequest() *AlibabaMiddleClaimsresultReceiveAPIRequest {
	return poolAlibabaMiddleClaimsresultReceiveAPIRequest.Get().(*AlibabaMiddleClaimsresultReceiveAPIRequest)
}

// ReleaseAlibabaMiddleClaimsresultReceiveAPIRequest 将 AlibabaMiddleClaimsresultReceiveAPIRequest 放入 sync.Pool
func ReleaseAlibabaMiddleClaimsresultReceiveAPIRequest(v *AlibabaMiddleClaimsresultReceiveAPIRequest) {
	v.Reset()
	poolAlibabaMiddleClaimsresultReceiveAPIRequest.Put(v)
}
