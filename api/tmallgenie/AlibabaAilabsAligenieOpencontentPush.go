package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaailabsaligenieopencontentpush 天猫精灵内容接入标准接口
// alibaba.ailabs.aligenie.opencontent.push
//
// 第三方内容接入天猫精灵内容库，供相关技能使用
func Alibabaailabsaligenieopencontentpush(clt *core.SDKClient, req *tmallgenie.AlibabaailabsaligenieopencontentpushAPIRequest, session string) (*tmallgenie.AlibabaailabsaligenieopencontentpushAPIResponse, error) {
	var resp tmallgenie.AlibabaailabsaligenieopencontentpushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
