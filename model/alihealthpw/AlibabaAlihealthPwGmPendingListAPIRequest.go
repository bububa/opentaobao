package alihealthpw

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPwGmPendingListAPIRequest 同情用药待审核工单查询接口 API请求
// alibaba.alihealth.pw.gm.pending.list
//
// 同情用药待审核工单查询接口，提供给合作方用来查询待处理工单列表
type AlibabaAlihealthPwGmPendingListAPIRequest struct {
	model.Params
	// 入参
	_body *PendingListReq
}

// NewAlibabaAlihealthPwGmPendingListRequest 初始化AlibabaAlihealthPwGmPendingListAPIRequest对象
func NewAlibabaAlihealthPwGmPendingListRequest() *AlibabaAlihealthPwGmPendingListAPIRequest {
	return &AlibabaAlihealthPwGmPendingListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthPwGmPendingListAPIRequest) Reset() {
	r._body = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPwGmPendingListAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.gm.pending.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPwGmPendingListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthPwGmPendingListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBody is Body Setter
// 入参
func (r *AlibabaAlihealthPwGmPendingListAPIRequest) SetBody(_body *PendingListReq) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r AlibabaAlihealthPwGmPendingListAPIRequest) GetBody() *PendingListReq {
	return r._body
}

var poolAlibabaAlihealthPwGmPendingListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthPwGmPendingListRequest()
	},
}

// GetAlibabaAlihealthPwGmPendingListRequest 从 sync.Pool 获取 AlibabaAlihealthPwGmPendingListAPIRequest
func GetAlibabaAlihealthPwGmPendingListAPIRequest() *AlibabaAlihealthPwGmPendingListAPIRequest {
	return poolAlibabaAlihealthPwGmPendingListAPIRequest.Get().(*AlibabaAlihealthPwGmPendingListAPIRequest)
}

// ReleaseAlibabaAlihealthPwGmPendingListAPIRequest 将 AlibabaAlihealthPwGmPendingListAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthPwGmPendingListAPIRequest(v *AlibabaAlihealthPwGmPendingListAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthPwGmPendingListAPIRequest.Put(v)
}
