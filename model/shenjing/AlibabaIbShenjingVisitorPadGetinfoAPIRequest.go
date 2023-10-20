package shenjing

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
