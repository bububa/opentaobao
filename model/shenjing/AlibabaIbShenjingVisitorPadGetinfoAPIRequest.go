package shenjing

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaibshenjingvisitorpadgetinfoAPIRequest 获取OSS上传参数 API请求
// alibaba.ib.shenjing.visitor.pad.getinfo
//
// PAD 端获取OSS上传参数，向OSS服务器上传图片。
type AlibabaibshenjingvisitorpadgetinfoAPIRequest struct {
	model.Params
}

// NewAlibabaibshenjingvisitorpadgetinfoRequest 初始化AlibabaibshenjingvisitorpadgetinfoAPIRequest对象
func NewAlibabaibshenjingvisitorpadgetinfoRequest() *AlibabaibshenjingvisitorpadgetinfoAPIRequest {
	return &AlibabaibshenjingvisitorpadgetinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaibshenjingvisitorpadgetinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.ib.shenjing.visitor.pad.getinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaibshenjingvisitorpadgetinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaibshenjingvisitorpadgetinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}
