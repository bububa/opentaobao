package xiami

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiami"
)

// Alibabaxiamiapiradiomyselfget 我的电台
// alibaba.xiami.api.radio.myself.get
//
// 我的电台
func Alibabaxiamiapiradiomyselfget(clt *core.SDKClient, req *xiami.AlibabaxiamiapiradiomyselfgetAPIRequest, session string) (*xiami.AlibabaxiamiapiradiomyselfgetAPIResponse, error) {
	var resp xiami.AlibabaxiamiapiradiomyselfgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
