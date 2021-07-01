package ioti

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaItEslEslimageSendimageAPIRequest
下发厂测初始化图片 API请求
alibaba.it.esl.eslimage.sendimage

工厂对生产出的电子价签进行全流程功能测试，能将出场图片通过ESL系统初始化到电子价签中 */
type AlibabaItEslEslimageSendimageAPIRequest struct {
	model.Params
	// 价签地址
	_mac string
}

// New
