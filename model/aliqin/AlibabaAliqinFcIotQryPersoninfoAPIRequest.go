package aliqin

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinFcIotQryPersoninfoAPIRequest) Reset() {
	r._iccid = ""
	r._userid = ""
	r._midPatChannel = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIotQryPersoninfoAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.iot.qry.personinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIotQryPersoninfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinFcIotQryPersoninfoAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAliqinFcIotQryPersoninfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinFcIotQryPersoninfoRequest()
	},
}

// GetAlibabaAliqinFcIotQryPersoninfoRequest 从 sync.Pool 获取 AlibabaAliqinFcIotQryPersoninfoAPIRequest
func GetAlibabaAliqinFcIotQryPersoninfoAPIRequest() *AlibabaAliqinFcIotQryPersoninfoAPIRequest {
	return poolAlibabaAliqinFcIotQryPersoninfoAPIRequest.Get().(*AlibabaAliqinFcIotQryPersoninfoAPIRequest)
}

// ReleaseAlibabaAliqinFcIotQryPersoninfoAPIRequest 将 AlibabaAliqinFcIotQryPersoninfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinFcIotQryPersoninfoAPIRequest(v *AlibabaAliqinFcIotQryPersoninfoAPIRequest) {
	v.Reset()
	poolAlibabaAliqinFcIotQryPersoninfoAPIRequest.Put(v)
}
