package youkuott

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/youkuott"
)

/* 
dvb工单反馈 
youku.ott.dvb.workorder.feedback

dvb工单处理结果反馈
*/
func YoukuOttDvbWorkorderFeedback(clt *core.SDKClient, req *youkuott.YoukuOttDvbWorkorderFeedbackAPIRequest, session string) (*youkuott.YoukuOttDvbWorkorderFeedbackAPIResponse, error) {
    var resp youkuott.YoukuOttDvbWorkorderFeedbackAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
