package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtMemberSynchronize 新零售会员同步接口
// tmall.nrt.member.synchronize
//
// 新零售会员上翻接口，商家的会员信息同步至阿里侧
func TmallNrtMemberSynchronize(clt *core.SDKClient, req *nrt.TmallNrtMemberSynchronizeAPIRequest, resp *nrt.TmallNrtMemberSynchronizeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
