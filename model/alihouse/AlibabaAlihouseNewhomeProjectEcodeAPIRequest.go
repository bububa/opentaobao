package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectEcodeAPIRequest 楼盘e码更新 API请求
// alibaba.alihouse.newhome.project.ecode
//
// 更新楼盘ecode
type AlibabaAlihouseNewhomeProjectEcodeAPIRequest struct {
	model.Params
	// 楼盘外部id
	_outerId string
	// 楼盘e码
	_eCode string
}

// NewAlibabaAlihouseNewhomeProjectEcodeRequest 初始化AlibabaAlihouseNewhomeProjectEcodeAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectEcodeRequest() *AlibabaAlihouseNewhomeProjectEcodeAPIRequest {
	return &AlibabaAlihouseNewhomeProjectEcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectEcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.ecode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectEcodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOuterId is OuterId Setter
// 楼盘外部id
func (r *AlibabaAlihouseNewhomeProjectEcodeAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaAlihouseNewhomeProjectEcodeAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetECode is ECode Setter
// 楼盘e码
func (r *AlibabaAlihouseNewhomeProjectEcodeAPIRequest) SetECode(_eCode string) error {
	r._eCode = _eCode
	r.Set("e_code", _eCode)
	return nil
}

// GetECode ECode Getter
func (r AlibabaAlihouseNewhomeProjectEcodeAPIRequest) GetECode() string {
	return r._eCode
}
