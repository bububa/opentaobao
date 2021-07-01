package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

/* YunosTvpubadminContentSingleVideoGetlist
单视频审核获取视频列表
yunos.tvpubadmin.content.single.video.getlist

牌照方在审核单视频时获取单视频列表接口 */
func YunosTvpubadminContentSingleVideoGetlist(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentSingleVideoGetlistAPIRequest, session string) (*tvupadmin.YunosTvpubadminContentSingleVideoGetlistAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminContentSingleVideoGetlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
