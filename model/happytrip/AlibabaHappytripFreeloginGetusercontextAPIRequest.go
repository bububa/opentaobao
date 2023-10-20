package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahappytripfreelogingetusercontextAPIRequest 提供给外部系统的免登校验 API请求
// alibaba.happytrip.freelogin.getusercontext
//
// 免登融合，提供免登相关接口给外部供应商做登录验证
type AlibabahappytripfreelogingetusercontextAPIRequest struct {
	model.Params
	// 请求入参
	_req *SsoParamDto
}

// NewAlibabahappytripfreelogingetusercontextRequest 初始化AlibabahappytripfreelogingetusercontextAPIRequest对象
func NewAlibabahappytripfreelogingetusercontextRequest() *AlibabahappytripfreelogingetusercontextAPIRequest {
	return &AlibabahappytripfreelogingetusercontextAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahappytripfreelogingetusercontextAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.freelogin.getusercontext"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahappytripfreelogingetusercontextAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahappytripfreelogingetusercontextAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReq is Req Setter
// 请求入参
func (r *AlibabahappytripfreelogingetusercontextAPIRequest) SetReq(_req *SsoParamDto) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// GetReq Req Getter
func (r AlibabahappytripfreelogingetusercontextAPIRequest) GetReq() *SsoParamDto {
	return r._req
}
