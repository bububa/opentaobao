package cloudgame

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/cloudgame"
)

/* 
云游戏战绩上传 
youku.cgame.score.report

云游戏战绩上传API
*/
func YoukuCgameScoreReport(clt *core.SDKClient, req *cloudgame.YoukuCgameScoreReportRequest, session string) (*cloudgame.YoukuCgameScoreReportAPIResponse, error) {
    var resp cloudgame.YoukuCgameScoreReportAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
