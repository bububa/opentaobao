package ioti

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaItEslSendledAPIRequest) GetApiMethodName() string {
	return "alibaba.it.esl.sendled"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaItEslSendledAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
