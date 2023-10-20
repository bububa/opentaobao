package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// Alibabajymindustryvmoscommoncallback vmos游戏信息采集结果回调通知
// alibaba.jym.industry.vmos.common.callback
//
// vmos游戏信息采集结果回调通知
func Alibabajymindustryvmoscommoncallback(clt *core.SDKClient, req *jym.AlibabajymindustryvmoscommoncallbackAPIRequest, session string) (*jym.AlibabajymindustryvmoscommoncallbackAPIResponse, error) {
	var resp jym.AlibabajymindustryvmoscommoncallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
