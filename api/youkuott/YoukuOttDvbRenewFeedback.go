package youkuott

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/youkuott"
)

/* 
dvb续费之后的反馈接口 
youku.ott.dvb.renew.feedback

dvb续费之后的反馈接口
*/
func YoukuOttDvbRenewFeedback(clt *core.SDKClient, req *youkuott.YoukuOttDvbRenewFeedbackAPIRequest, session string) (*youkuott.YoukuOttDvbRenewFeedbackAPIResponse, error) {
    var resp youkuott.YoukuOttDvbRenewFeedbackAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
