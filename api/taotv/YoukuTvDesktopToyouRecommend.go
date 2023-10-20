package taotv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/taotv"
)

// Youkutvdesktoptoyourecommend TV桌面为你推荐接口
// youku.tv.desktop.toyou.recommend
//
// 提供为你推荐数据
func Youkutvdesktoptoyourecommend(clt *core.SDKClient, req *taotv.YoukutvdesktoptoyourecommendAPIRequest, session string) (*taotv.YoukutvdesktoptoyourecommendAPIResponse, error) {
	var resp taotv.YoukutvdesktoptoyourecommendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
