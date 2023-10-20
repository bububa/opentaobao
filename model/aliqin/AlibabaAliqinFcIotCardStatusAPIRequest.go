package aliqin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcIotCardStatusAPIRequest 物联卡状态查询 API请求
// alibaba.aliqin.fc.iot.cardStatus
//
// 物联卡状态查询
type AlibabaAliqinFcIotCardStatusAPIRequest struct {
	model.Params
	// SIM卡号
	_iccid string
}

// NewAlibabaAliqinFcIotCardStatusRequest 初始化AlibabaAliqinFcIotCardStatusAPIRequest对象
func NewAlibabaAliqinFcIotCardStatusRequest() *AlibabaAliqinFcIotCardStatusAPIRequest {
	return &AlibabaAliqinFcIotCardStatusAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinFcIotCardStatusAPIRequest) Reset() {
	r._iccid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIotCardStatusAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.iot.cardStatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIotCardStatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinFcIotCardStatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIccid is Iccid Setter
// SIM卡号
func (r *AlibabaAliqinFcIotCardStatusAPIRequest) SetIccid(_iccid string) error {
	r._iccid = _iccid
	r.Set("iccid", _iccid)
	return nil
}

// GetIccid Iccid Getter
func (r AlibabaAliqinFcIotCardStatusAPIRequest) GetIccid() string {
	return r._iccid
}

var poolAlibabaAliqinFcIotCardStatusAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinFcIotCardStatusRequest()
	},
}

// GetAlibabaAliqinFcIotCardStatusRequest 从 sync.Pool 获取 AlibabaAliqinFcIotCardStatusAPIRequest
func GetAlibabaAliqinFcIotCardStatusAPIRequest() *AlibabaAliqinFcIotCardStatusAPIRequest {
	return poolAlibabaAliqinFcIotCardStatusAPIRequest.Get().(*AlibabaAliqinFcIotCardStatusAPIRequest)
}

// ReleaseAlibabaAliqinFcIotCardStatusAPIRequest 将 AlibabaAliqinFcIotCardStatusAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinFcIotCardStatusAPIRequest(v *AlibabaAliqinFcIotCardStatusAPIRequest) {
	v.Reset()
	poolAlibabaAliqinFcIotCardStatusAPIRequest.Put(v)
}
