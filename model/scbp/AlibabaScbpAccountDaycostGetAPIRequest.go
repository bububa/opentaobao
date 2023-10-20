package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAccountDaycostGetAPIRequest 查询今日消耗 API请求
// alibaba.scbp.account.daycost.get
//
// 查询今日消耗
type AlibabaScbpAccountDaycostGetAPIRequest struct {
	model.Params
}

// NewAlibabaScbpAccountDaycostGetRequest 初始化AlibabaScbpAccountDaycostGetAPIRequest对象
func NewAlibabaScbpAccountDaycostGetRequest() *AlibabaScbpAccountDaycostGetAPIRequest {
	return &AlibabaScbpAccountDaycostGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAccountDaycostGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAccountDaycostGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.account.daycost.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAccountDaycostGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAccountDaycostGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaScbpAccountDaycostGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAccountDaycostGetRequest()
	},
}

// GetAlibabaScbpAccountDaycostGetRequest 从 sync.Pool 获取 AlibabaScbpAccountDaycostGetAPIRequest
func GetAlibabaScbpAccountDaycostGetAPIRequest() *AlibabaScbpAccountDaycostGetAPIRequest {
	return poolAlibabaScbpAccountDaycostGetAPIRequest.Get().(*AlibabaScbpAccountDaycostGetAPIRequest)
}

// ReleaseAlibabaScbpAccountDaycostGetAPIRequest 将 AlibabaScbpAccountDaycostGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAccountDaycostGetAPIRequest(v *AlibabaScbpAccountDaycostGetAPIRequest) {
	v.Reset()
	poolAlibabaScbpAccountDaycostGetAPIRequest.Put(v)
}
