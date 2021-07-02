package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentShowGetbyshowlongid 迎客松根据节目longid获取节目元数据
// yunos.tvpubadmin.content.show.getbyshowlongid
//
// 迎客松根据节目longid获取节目元数据
func YunosTvpubadminContentShowGetbyshowlongid(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentShowGetbyshowlongidAPIRequest, session string) (*tvupadmin.YunosTvpubadminContentShowGetbyshowlongidAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminContentShowGetbyshowlongidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
