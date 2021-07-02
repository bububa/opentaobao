package ioti

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaItEslEslimageSendimageAPIRequest) GetApiMethodName() string {
	return "alibaba.it.esl.eslimage.sendimage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaItEslEslimageSendimageAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
