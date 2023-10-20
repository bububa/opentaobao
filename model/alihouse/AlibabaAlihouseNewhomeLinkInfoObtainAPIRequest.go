package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeLinkInfoObtainAPIRequest 落地页获取 API请求
// alibaba.alihouse.newhome.link.info.obtain
//
// 落地页获取
type AlibabaAlihouseNewhomeLinkInfoObtainAPIRequest struct {
	model.Params
	// 请求体
	_linkInfoReq *LinkInfoReqDto
}

// NewAlibabaAlihouseNewhomeLinkInfoObtainRequest 初始化AlibabaAlihouseNewhomeLinkInfoObtainAPIRequest对象
func NewAlibabaAlihouseNewhomeLinkInfoObtainRequest() *AlibabaAlihouseNewhomeLinkInfoObtainAPIRequest {
	return &AlibabaAlihouseNewhomeLinkInfoObtainAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeLinkInfoObtainAPIRequest) Reset() {
	r._linkInfoReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeLinkInfoObtainAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.link.info.obtain"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeLinkInfoObtainAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeLinkInfoObtainAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLinkInfoReq is LinkInfoReq Setter
// 请求体
func (r *AlibabaAlihouseNewhomeLinkInfoObtainAPIRequest) SetLinkInfoReq(_linkInfoReq *LinkInfoReqDto) error {
	r._linkInfoReq = _linkInfoReq
	r.Set("link_info_req", _linkInfoReq)
	return nil
}

// GetLinkInfoReq LinkInfoReq Getter
func (r AlibabaAlihouseNewhomeLinkInfoObtainAPIRequest) GetLinkInfoReq() *LinkInfoReqDto {
	return r._linkInfoReq
}

var poolAlibabaAlihouseNewhomeLinkInfoObtainAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeLinkInfoObtainRequest()
	},
}

// GetAlibabaAlihouseNewhomeLinkInfoObtainRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeLinkInfoObtainAPIRequest
func GetAlibabaAlihouseNewhomeLinkInfoObtainAPIRequest() *AlibabaAlihouseNewhomeLinkInfoObtainAPIRequest {
	return poolAlibabaAlihouseNewhomeLinkInfoObtainAPIRequest.Get().(*AlibabaAlihouseNewhomeLinkInfoObtainAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeLinkInfoObtainAPIRequest 将 AlibabaAlihouseNewhomeLinkInfoObtainAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeLinkInfoObtainAPIRequest(v *AlibabaAlihouseNewhomeLinkInfoObtainAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeLinkInfoObtainAPIRequest.Put(v)
}
