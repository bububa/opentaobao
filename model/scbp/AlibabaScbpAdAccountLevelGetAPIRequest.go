package scbp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdAccountLevelGetAPIRequest 查询推广账户等级 API请求
// alibaba.scbp.ad.account.level.get
//
// 查询推广账户等级
type AlibabaScbpAdAccountLevelGetAPIRequest struct {
	model.Params
}

// NewAlibabaScbpAdAccountLevelGetRequest 初始化AlibabaScbpAdAccountLevelGetAPIRequest对象
func NewAlibabaScbpAdAccountLevelGetRequest() *AlibabaScbpAdAccountLevelGetAPIRequest {
	return &AlibabaScbpAdAccountLevelGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaScbpAdAccountLevelGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdAccountLevelGetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.account.level.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdAccountLevelGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaScbpAdAccountLevelGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaScbpAdAccountLevelGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaScbpAdAccountLevelGetRequest()
	},
}

// GetAlibabaScbpAdAccountLevelGetRequest 从 sync.Pool 获取 AlibabaScbpAdAccountLevelGetAPIRequest
func GetAlibabaScbpAdAccountLevelGetAPIRequest() *AlibabaScbpAdAccountLevelGetAPIRequest {
	return poolAlibabaScbpAdAccountLevelGetAPIRequest.Get().(*AlibabaScbpAdAccountLevelGetAPIRequest)
}

// ReleaseAlibabaScbpAdAccountLevelGetAPIRequest 将 AlibabaScbpAdAccountLevelGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaScbpAdAccountLevelGetAPIRequest(v *AlibabaScbpAdAccountLevelGetAPIRequest) {
	v.Reset()
	poolAlibabaScbpAdAccountLevelGetAPIRequest.Put(v)
}
