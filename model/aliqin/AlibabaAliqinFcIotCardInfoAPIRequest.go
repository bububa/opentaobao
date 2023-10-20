package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinfciotcardInfoAPIRequest 物联卡信息查询 API请求
// alibaba.aliqin.fc.iot.cardInfo
//
// 物联卡信息查询
type AlibabaaliqinfciotcardInfoAPIRequest struct {
	model.Params
	// SIM卡号
	_iccid string
}

// NewAlibabaaliqinfciotcardInfoRequest 初始化AlibabaaliqinfciotcardInfoAPIRequest对象
func NewAlibabaaliqinfciotcardInfoRequest() *AlibabaaliqinfciotcardInfoAPIRequest {
	return &AlibabaaliqinfciotcardInfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqinfciotcardInfoAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.iot.cardInfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqinfciotcardInfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqinfciotcardInfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIccid is Iccid Setter
// SIM卡号
func (r *AlibabaaliqinfciotcardInfoAPIRequest) SetIccid(_iccid string) error {
	r._iccid = _iccid
	r.Set("iccid", _iccid)
	return nil
}

// GetIccid Iccid Getter
func (r AlibabaaliqinfciotcardInfoAPIRequest) GetIccid() string {
	return r._iccid
}
