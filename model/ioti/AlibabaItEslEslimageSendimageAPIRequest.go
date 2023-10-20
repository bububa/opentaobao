package ioti

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItEslEslimageSendimageAPIRequest 下发厂测初始化图片 API请求
// alibaba.it.esl.eslimage.sendimage
//
// 工厂对生产出的电子价签进行全流程功能测试，能将出场图片通过ESL系统初始化到电子价签中
type AlibabaItEslEslimageSendimageAPIRequest struct {
	model.Params
	// 价签地址
	_mac string
}

// NewAlibabaItEslEslimageSendimageRequest 初始化AlibabaItEslEslimageSendimageAPIRequest对象
func NewAlibabaItEslEslimageSendimageRequest() *AlibabaItEslEslimageSendimageAPIRequest {
	return &AlibabaItEslEslimageSendimageAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaItEslEslimageSendimageAPIRequest) Reset() {
	r._mac = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaItEslEslimageSendimageAPIRequest) GetApiMethodName() string {
	return "alibaba.it.esl.eslimage.sendimage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaItEslEslimageSendimageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaItEslEslimageSendimageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMac is Mac Setter
// 价签地址
func (r *AlibabaItEslEslimageSendimageAPIRequest) SetMac(_mac string) error {
	r._mac = _mac
	r.Set("mac", _mac)
	return nil
}

// GetMac Mac Getter
func (r AlibabaItEslEslimageSendimageAPIRequest) GetMac() string {
	return r._mac
}

var poolAlibabaItEslEslimageSendimageAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaItEslEslimageSendimageRequest()
	},
}

// GetAlibabaItEslEslimageSendimageRequest 从 sync.Pool 获取 AlibabaItEslEslimageSendimageAPIRequest
func GetAlibabaItEslEslimageSendimageAPIRequest() *AlibabaItEslEslimageSendimageAPIRequest {
	return poolAlibabaItEslEslimageSendimageAPIRequest.Get().(*AlibabaItEslEslimageSendimageAPIRequest)
}

// ReleaseAlibabaItEslEslimageSendimageAPIRequest 将 AlibabaItEslEslimageSendimageAPIRequest 放入 sync.Pool
func ReleaseAlibabaItEslEslimageSendimageAPIRequest(v *AlibabaItEslEslimageSendimageAPIRequest) {
	v.Reset()
	poolAlibabaItEslEslimageSendimageAPIRequest.Put(v)
}
