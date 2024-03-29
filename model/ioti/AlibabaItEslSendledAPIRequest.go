package ioti

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItEslSendledAPIRequest 厂测LED控制 API请求
// alibaba.it.esl.sendled
//
// 针对厂测生产的的价签，增加led闪灯的接口，进行led 闪灯测试
type AlibabaItEslSendledAPIRequest struct {
	model.Params
	// mac
	_macAp string
	// 0、1、2、3：关蓝绿红
	_type string
}

// NewAlibabaItEslSendledRequest 初始化AlibabaItEslSendledAPIRequest对象
func NewAlibabaItEslSendledRequest() *AlibabaItEslSendledAPIRequest {
	return &AlibabaItEslSendledAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaItEslSendledAPIRequest) Reset() {
	r._macAp = ""
	r._type = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaItEslSendledAPIRequest) GetApiMethodName() string {
	return "alibaba.it.esl.sendled"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaItEslSendledAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaItEslSendledAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMacAp is MacAp Setter
// mac
func (r *AlibabaItEslSendledAPIRequest) SetMacAp(_macAp string) error {
	r._macAp = _macAp
	r.Set("mac_ap", _macAp)
	return nil
}

// GetMacAp MacAp Getter
func (r AlibabaItEslSendledAPIRequest) GetMacAp() string {
	return r._macAp
}

// SetType is Type Setter
// 0、1、2、3：关蓝绿红
func (r *AlibabaItEslSendledAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaItEslSendledAPIRequest) GetType() string {
	return r._type
}

var poolAlibabaItEslSendledAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaItEslSendledRequest()
	},
}

// GetAlibabaItEslSendledRequest 从 sync.Pool 获取 AlibabaItEslSendledAPIRequest
func GetAlibabaItEslSendledAPIRequest() *AlibabaItEslSendledAPIRequest {
	return poolAlibabaItEslSendledAPIRequest.Get().(*AlibabaItEslSendledAPIRequest)
}

// ReleaseAlibabaItEslSendledAPIRequest 将 AlibabaItEslSendledAPIRequest 放入 sync.Pool
func ReleaseAlibabaItEslSendledAPIRequest(v *AlibabaItEslSendledAPIRequest) {
	v.Reset()
	poolAlibabaItEslSendledAPIRequest.Put(v)
}
