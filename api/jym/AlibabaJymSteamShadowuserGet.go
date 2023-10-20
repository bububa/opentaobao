package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// Alibabajymsteamshadowuserget 获取影子标识
// alibaba.jym.steam.shadowuser.get
//
// 交易猫Steam类目获取影子ID
func Alibabajymsteamshadowuserget(clt *core.SDKClient, req *jym.AlibabajymsteamshadowusergetAPIRequest, session string) (*jym.AlibabajymsteamshadowusergetAPIResponse, error) {
	var resp jym.AlibabajymsteamshadowusergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
