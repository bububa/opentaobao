package alihealthpw

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPwGmIdsListAPIRequest 同情用药根据申请单列表查询申请单 API请求
// alibaba.alihealth.pw.gm.ids.list
//
// 同情用药根据申请单列表查询申请单
type AlibabaAlihealthPwGmIdsListAPIRequest struct {
	model.Params
	// 入参
	_body *ListByApplyIdsForBReq
}

// NewAlibabaAlihealthPwGmIdsListRequest 初始化AlibabaAlihealthPwGmIdsListAPIRequest对象
func NewAlibabaAlihealthPwGmIdsListRequest() *AlibabaAlihealthPwGmIdsListAPIRequest {
	return &AlibabaAlihealthPwGmIdsListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthPwGmIdsListAPIRequest) Reset() {
	r._body = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPwGmIdsListAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.gm.ids.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPwGmIdsListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthPwGmIdsListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBody is Body Setter
// 入参
func (r *AlibabaAlihealthPwGmIdsListAPIRequest) SetBody(_body *ListByApplyIdsForBReq) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r AlibabaAlihealthPwGmIdsListAPIRequest) GetBody() *ListByApplyIdsForBReq {
	return r._body
}

var poolAlibabaAlihealthPwGmIdsListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthPwGmIdsListRequest()
	},
}

// GetAlibabaAlihealthPwGmIdsListRequest 从 sync.Pool 获取 AlibabaAlihealthPwGmIdsListAPIRequest
func GetAlibabaAlihealthPwGmIdsListAPIRequest() *AlibabaAlihealthPwGmIdsListAPIRequest {
	return poolAlibabaAlihealthPwGmIdsListAPIRequest.Get().(*AlibabaAlihealthPwGmIdsListAPIRequest)
}

// ReleaseAlibabaAlihealthPwGmIdsListAPIRequest 将 AlibabaAlihealthPwGmIdsListAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthPwGmIdsListAPIRequest(v *AlibabaAlihealthPwGmIdsListAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthPwGmIdsListAPIRequest.Put(v)
}
