package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomelinkinfoobtainAPIRequest 落地页获取 API请求
// alibaba.alihouse.newhome.link.info.obtain
//
// 落地页获取
type AlibabaalihousenewhomelinkinfoobtainAPIRequest struct {
	model.Params
	// 请求体
	_linkInfoReq *LinkInfoReqDto
}

// NewAlibabaalihousenewhomelinkinfoobtainRequest 初始化AlibabaalihousenewhomelinkinfoobtainAPIRequest对象
func NewAlibabaalihousenewhomelinkinfoobtainRequest() *AlibabaalihousenewhomelinkinfoobtainAPIRequest {
	return &AlibabaalihousenewhomelinkinfoobtainAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomelinkinfoobtainAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.link.info.obtain"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomelinkinfoobtainAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomelinkinfoobtainAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLinkInfoReq is LinkInfoReq Setter
// 请求体
func (r *AlibabaalihousenewhomelinkinfoobtainAPIRequest) SetLinkInfoReq(_linkInfoReq *LinkInfoReqDto) error {
	r._linkInfoReq = _linkInfoReq
	r.Set("link_info_req", _linkInfoReq)
	return nil
}

// GetLinkInfoReq LinkInfoReq Getter
func (r AlibabaalihousenewhomelinkinfoobtainAPIRequest) GetLinkInfoReq() *LinkInfoReqDto {
	return r._linkInfoReq
}
