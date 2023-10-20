package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// Youkucgamescorereport 云游戏战绩上传
// youku.cgame.score.report
//
// 云游戏战绩上传API
func Youkucgamescorereport(clt *core.SDKClient, req *cloudgame.YoukucgamescorereportAPIRequest, session string) (*cloudgame.YoukucgamescorereportAPIResponse, error) {
	var resp cloudgame.YoukucgamescorereportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
