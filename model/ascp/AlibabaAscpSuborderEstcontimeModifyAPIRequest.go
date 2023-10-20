package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpSuborderEstcontimeModifyAPIRequest 向前修改发货时效 API请求
// alibaba.ascp.suborder.estcontime.modify
//
// 向前修改发货时效
type AlibabaAscpSuborderEstcontimeModifyAPIRequest struct {
	model.Params
	// 修改商品发货时效请求
	_modifyEstConTimeRequest *ModifyEstConTimeRequest
}

// NewAlibabaAscpSuborderEstcontimeModifyRequest 初始化AlibabaAscpSuborderEstcontimeModifyAPIRequest对象
func NewAlibabaAscpSuborderEstcontimeModifyRequest() *AlibabaAscpSuborderEstcontimeModifyAPIRequest {
	return &AlibabaAscpSuborderEstcontimeModifyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpSuborderEstcontimeModifyAPIRequest) Reset() {
	r._modifyEstConTimeRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpSuborderEstcontimeModifyAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.suborder.estcontime.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpSuborderEstcontimeModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpSuborderEstcontimeModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModifyEstConTimeRequest is ModifyEstConTimeRequest Setter
// 修改商品发货时效请求
func (r *AlibabaAscpSuborderEstcontimeModifyAPIRequest) SetModifyEstConTimeRequest(_modifyEstConTimeRequest *ModifyEstConTimeRequest) error {
	r._modifyEstConTimeRequest = _modifyEstConTimeRequest
	r.Set("modify_est_con_time_request", _modifyEstConTimeRequest)
	return nil
}

// GetModifyEstConTimeRequest ModifyEstConTimeRequest Getter
func (r AlibabaAscpSuborderEstcontimeModifyAPIRequest) GetModifyEstConTimeRequest() *ModifyEstConTimeRequest {
	return r._modifyEstConTimeRequest
}

var poolAlibabaAscpSuborderEstcontimeModifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpSuborderEstcontimeModifyRequest()
	},
}

// GetAlibabaAscpSuborderEstcontimeModifyRequest 从 sync.Pool 获取 AlibabaAscpSuborderEstcontimeModifyAPIRequest
func GetAlibabaAscpSuborderEstcontimeModifyAPIRequest() *AlibabaAscpSuborderEstcontimeModifyAPIRequest {
	return poolAlibabaAscpSuborderEstcontimeModifyAPIRequest.Get().(*AlibabaAscpSuborderEstcontimeModifyAPIRequest)
}

// ReleaseAlibabaAscpSuborderEstcontimeModifyAPIRequest 将 AlibabaAscpSuborderEstcontimeModifyAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpSuborderEstcontimeModifyAPIRequest(v *AlibabaAscpSuborderEstcontimeModifyAPIRequest) {
	v.Reset()
	poolAlibabaAscpSuborderEstcontimeModifyAPIRequest.Put(v)
}
