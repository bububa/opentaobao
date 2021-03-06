package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcIotQryPersoninfoAPIRequest 查询物联卡个人实人认证信息 API请求
// alibaba.aliqin.fc.iot.qry.personinfo
//
// 查询物联卡个人实人认证信息
type AlibabaAliqinFcIotQryPersoninfoAPIRequest struct {
	model.Params
	// 需要查询的iccid
	_iccid string
	// 指定查询某userid
	_userid string
	// 由系统根据业务分配
	_midPatChannel string
}

// NewAlibabaAliqinFcIotQryPersoninfoRequest 初始化AlibabaAliqinFcIotQryPersoninfoAPIRequest对象
func NewAlibabaAliqinFcIotQryPersoninfoRequest() *AlibabaAliqinFcIotQryPersoninfoAPIRequest {
	return &AlibabaAliqinFcIotQryPersoninfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIotQryPersoninfoAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.iot.qry.personinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIotQryPersoninfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetIccid is Iccid Setter
// 需要查询的iccid
func (r *AlibabaAliqinFcIotQryPersoninfoAPIRequest) SetIccid(_iccid string) error {
	r._iccid = _iccid
	r.Set("iccid", _iccid)
	return nil
}

// GetIccid Iccid Getter
func (r AlibabaAliqinFcIotQryPersoninfoAPIRequest) GetIccid() string {
	return r._iccid
}

// SetUserid is Userid Setter
// 指定查询某userid
func (r *AlibabaAliqinFcIotQryPersoninfoAPIRequest) SetUserid(_userid string) error {
	r._userid = _userid
	r.Set("userid", _userid)
	return nil
}

// GetUserid Userid Getter
func (r AlibabaAliqinFcIotQryPersoninfoAPIRequest) GetUserid() string {
	return r._userid
}

// SetMidPatChannel is MidPatChannel Setter
// 由系统根据业务分配
func (r *AlibabaAliqinFcIotQryPersoninfoAPIRequest) SetMidPatChannel(_midPatChannel string) error {
	r._midPatChannel = _midPatChannel
	r.Set("mid_pat_channel", _midPatChannel)
	return nil
}

// GetMidPatChannel MidPatChannel Getter
func (r AlibabaAliqinFcIotQryPersoninfoAPIRequest) GetMidPatChannel() string {
	return r._midPatChannel
}
