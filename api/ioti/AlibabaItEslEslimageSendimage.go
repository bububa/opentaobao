package ioti

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ioti"
)

// AlibabaItEslEslimageSendimage 下发厂测初始化图片
// alibaba.it.esl.eslimage.sendimage
//
// 工厂对生产出的电子价签进行全流程功能测试，能将出场图片通过ESL系统初始化到电子价签中
func AlibabaItEslEslimageSendimage(clt *core.SDKClient, req *ioti.AlibabaItEslEslimageSendimageAPIRequest, session string) (*ioti.AlibabaItEslEslimageSendimageAPIResponse, error) {
	var resp ioti.AlibabaItEslEslimageSendimageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
