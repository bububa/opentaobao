package shenjing

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIbShenjingVisitorPadGetinfoAPIRequest 获取OSS上传参数 API请求
// alibaba.ib.shenjing.visitor.pad.getinfo
//
// PAD 端获取OSS上传参数，向OSS服务器上传图片。
type AlibabaIbShenjingVisitorPadGetinfoAPIRequest struct {
	model.Params
}

// NewAlibabaIbShenjingVisitorPadGetinfoRequest 初始化AlibabaIbShenjingVisitorPadGetinfoAPIRequest对象
func NewAlibabaIbShenjingVisitorPadGetinfoRequest() *AlibabaIbShenjingVisitorPadGetinfoAPIRequest {
	return &AlibabaIbShenjingVisitorPadGetinfoAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIbShenjingVisitorPadGetinfoAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIbShenjingVisitorPadGetinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.ib.shenjing.visitor.pad.getinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIbShenjingVisitorPadGetinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIbShenjingVisitorPadGetinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaIbShenjingVisitorPadGetinfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIbShenjingVisitorPadGetinfoRequest()
	},
}

// GetAlibabaIbShenjingVisitorPadGetinfoRequest 从 sync.Pool 获取 AlibabaIbShenjingVisitorPadGetinfoAPIRequest
func GetAlibabaIbShenjingVisitorPadGetinfoAPIRequest() *AlibabaIbShenjingVisitorPadGetinfoAPIRequest {
	return poolAlibabaIbShenjingVisitorPadGetinfoAPIRequest.Get().(*AlibabaIbShenjingVisitorPadGetinfoAPIRequest)
}

// ReleaseAlibabaIbShenjingVisitorPadGetinfoAPIRequest 将 AlibabaIbShenjingVisitorPadGetinfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaIbShenjingVisitorPadGetinfoAPIRequest(v *AlibabaIbShenjingVisitorPadGetinfoAPIRequest) {
	v.Reset()
	poolAlibabaIbShenjingVisitorPadGetinfoAPIRequest.Put(v)
}
