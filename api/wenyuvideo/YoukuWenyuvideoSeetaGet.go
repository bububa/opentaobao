package wenyuvideo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wenyuvideo"
)

// Youkuwenyuvideoseetaget 只看TA
// youku.wenyuvideo.seeta.get
//
// 只看Ta对外输出
func Youkuwenyuvideoseetaget(clt *core.SDKClient, req *wenyuvideo.YoukuwenyuvideoseetagetAPIRequest, session string) (*wenyuvideo.YoukuwenyuvideoseetagetAPIResponse, error) {
	var resp wenyuvideo.YoukuwenyuvideoseetagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
