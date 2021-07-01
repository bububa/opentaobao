package youkuott

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/youkuott"
)

/* 
dvb首次安装订单反馈 
youku.ott.dvb.firstorder.feedback

dvb首次安装订单反馈
*/
func YoukuOttDvbFirstorderFeedback(clt *core.SDKClient, req *youkuott.YoukuOttDvbFirstorderFeedbackAPIRequest, session string) (*youkuott.YoukuOttDvbFirstorderFeedbackAPIResponse, error) {
    var resp youkuott.YoukuOttDvbFirstorderFeedbackAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
