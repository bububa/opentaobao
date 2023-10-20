package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// Tmallnrtmembersynchronize 新零售会员同步接口
// tmall.nrt.member.synchronize
//
// 新零售会员上翻接口，商家的会员信息同步至阿里侧
func Tmallnrtmembersynchronize(clt *core.SDKClient, req *nrt.TmallnrtmembersynchronizeAPIRequest, session string) (*nrt.TmallnrtmembersynchronizeAPIResponse, error) {
	var resp nrt.TmallnrtmembersynchronizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
