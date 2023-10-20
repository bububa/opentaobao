package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamjocconfpickupgoods 提货核销
// alibaba.mj.oc.confpickupgoods
//
// 此API用于在银泰商场中，消费者在提货中心提货时， 商户后台调用此接口进行提货核销操作
func Alibabamjocconfpickupgoods(clt *core.SDKClient, req *mos.AlibabamjocconfpickupgoodsAPIRequest, session string) (*mos.AlibabamjocconfpickupgoodsAPIResponse, error) {
	var resp mos.AlibabamjocconfpickupgoodsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
