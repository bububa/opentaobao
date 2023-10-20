package ioti

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ioti"
)

// Alibabaiteslsendota 电子价签ota接口
// alibaba.it.esl.sendota
//
// 厂测接口，电子价签ota接口
func Alibabaiteslsendota(clt *core.SDKClient, req *ioti.AlibabaiteslsendotaAPIRequest, session string) (*ioti.AlibabaiteslsendotaAPIResponse, error) {
	var resp ioti.AlibabaiteslsendotaAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
