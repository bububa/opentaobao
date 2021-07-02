package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentVideoGetauditlist 迎客松视频审核记录查询
// yunos.tvpubadmin.content.video.getauditlist
//
// 迎客松视频审核记录查询
func YunosTvpubadminContentVideoGetauditlist(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentVideoGetauditlistAPIRequest, session string) (*tvupadmin.YunosTvpubadminContentVideoGetauditlistAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminContentVideoGetauditlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
