package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinfciotqrypersoninfoAPIRequest 查询物联卡个人实人认证信息 API请求
// alibaba.aliqin.fc.iot.qry.personinfo
//
// 查询物联卡个人实人认证信息
type AlibabaaliqinfciotqrypersoninfoAPIRequest struct {
	model.Params
	// 需要查询的iccid
	_iccid string
	// 指定查询某userid
	_userid string
	// 由系统根据业务分配
	_midPatChannel string
}

// NewAlibabaaliqinfciotqrypersoninfoRequest 初始化AlibabaaliqinfciotqrypersoninfoAPIRequest对象
func NewAlibabaaliqinfciotqrypersoninfoRequest() *AlibabaaliqinfciotqrypersoninfoAPIRequest {
	return &AlibabaaliqinfciotqrypersoninfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqinfciotqrypersoninfoAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.iot.qry.personinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqinfciotqrypersoninfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqinfciotqrypersoninfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIccid is Iccid Setter
// 需要查询的iccid
func (r *AlibabaaliqinfciotqrypersoninfoAPIRequest) SetIccid(_iccid string) error {
	r._iccid = _iccid
	r.Set("iccid", _iccid)
	return nil
}

// GetIccid Iccid Getter
func (r AlibabaaliqinfciotqrypersoninfoAPIRequest) GetIccid() string {
	return r._iccid
}

// SetUserid is Userid Setter
// 指定查询某userid
func (r *AlibabaaliqinfciotqrypersoninfoAPIRequest) SetUserid(_userid string) error {
	r._userid = _userid
	r.Set("userid", _userid)
	return nil
}

// GetUserid Userid Getter
func (r AlibabaaliqinfciotqrypersoninfoAPIRequest) GetUserid() string {
	return r._userid
}

// SetMidPatChannel is MidPatChannel Setter
// 由系统根据业务分配
func (r *AlibabaaliqinfciotqrypersoninfoAPIRequest) SetMidPatChannel(_midPatChannel string) error {
	r._midPatChannel = _midPatChannel
	r.Set("mid_pat_channel", _midPatChannel)
	return nil
}

// GetMidPatChannel MidPatChannel Getter
func (r AlibabaaliqinfciotqrypersoninfoAPIRequest) GetMidPatChannel() string {
	return r._midPatChannel
}
