package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripFreeloginGetusercontextAPIRequest 提供给外部系统的免登校验 API请求
// alibaba.happytrip.freelogin.getusercontext
//
// 免登融合，提供免登相关接口给外部供应商做登录验证
type AlibabaHappytripFreeloginGetusercontextAPIRequest struct {
	model.Params
	// 请求入参
	_req *SsoParamDto
}

// NewAlibabaHappytripFreeloginGetusercontextRequest 初始化AlibabaHappytripFreeloginGetusercontextAPIRequest对象
func NewAlibabaHappytripFreeloginGetusercontextRequest() *AlibabaHappytripFreeloginGetusercontextAPIRequest {
	return &AlibabaHappytripFreeloginGetusercontextAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripFreeloginGetusercontextAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.freelogin.getusercontext"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripFreeloginGetusercontextAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHappytripFreeloginGetusercontextAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReq is Req Setter
// 请求入参
func (r *AlibabaHappytripFreeloginGetusercontextAPIRequest) SetReq(_req *SsoParamDto) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// GetReq Req Getter
func (r AlibabaHappytripFreeloginGetusercontextAPIRequest) GetReq() *SsoParamDto {
	return r._req
}
