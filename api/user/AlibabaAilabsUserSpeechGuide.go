package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Alibabaailabsuserspeechguide 引导语推荐接口
// alibaba.ailabs.user.speech.guide
//
// 根据用户的语音query，返回相应的引导语推荐
func Alibabaailabsuserspeechguide(clt *core.SDKClient, req *user.AlibabaailabsuserspeechguideAPIRequest, session string) (*user.AlibabaailabsuserspeechguideAPIResponse, error) {
	var resp user.AlibabaailabsuserspeechguideAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
