package aliqin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcIotCardInfoAPIRequest 物联卡信息查询 API请求
// alibaba.aliqin.fc.iot.cardInfo
//
// 物联卡信息查询
type AlibabaAliqinFcIotCardInfoAPIRequest struct {
	model.Params
	// SIM卡号
	_iccid string
}

// NewAlibabaAliqinFcIotCardInfoRequest 初始化AlibabaAliqinFcIotCardInfoAPIRequest对象
func NewAlibabaAliqinFcIotCardInfoRequest() *AlibabaAliqinFcIotCardInfoAPIRequest {
	return &AlibabaAliqinFcIotCardInfoAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinFcIotCardInfoAPIRequest) Reset() {
	r._iccid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIotCardInfoAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.iot.cardInfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIotCardInfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinFcIotCardInfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIccid is Iccid Setter
// SIM卡号
func (r *AlibabaAliqinFcIotCardInfoAPIRequest) SetIccid(_iccid string) error {
	r._iccid = _iccid
	r.Set("iccid", _iccid)
	return nil
}

// GetIccid Iccid Getter
func (r AlibabaAliqinFcIotCardInfoAPIRequest) GetIccid() string {
	return r._iccid
}

var poolAlibabaAliqinFcIotCardInfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinFcIotCardInfoRequest()
	},
}

// GetAlibabaAliqinFcIotCardInfoRequest 从 sync.Pool 获取 AlibabaAliqinFcIotCardInfoAPIRequest
func GetAlibabaAliqinFcIotCardInfoAPIRequest() *AlibabaAliqinFcIotCardInfoAPIRequest {
	return poolAlibabaAliqinFcIotCardInfoAPIRequest.Get().(*AlibabaAliqinFcIotCardInfoAPIRequest)
}

// ReleaseAlibabaAliqinFcIotCardInfoAPIRequest 将 AlibabaAliqinFcIotCardInfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinFcIotCardInfoAPIRequest(v *AlibabaAliqinFcIotCardInfoAPIRequest) {
	v.Reset()
	poolAlibabaAliqinFcIotCardInfoAPIRequest.Put(v)
}
