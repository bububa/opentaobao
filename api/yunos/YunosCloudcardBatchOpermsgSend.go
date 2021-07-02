package yunos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/yunos"
)

// YunosCloudcardBatchOpermsgSend YUNOS生活服务群发消息
// yunos.cloudcard.batch.opermsg.send
//
// 这个是一个群发消息接口，ISV通过该接口给订阅自己服务号的所有YUNOS终端用户发送服务号消息，目前该接口有调用频率限制，具体规则参考YUNOS开放平台文档。
func YunosCloudcardBatchOpermsgSend(clt *core.SDKClient, req *yunos.YunosCloudcardBatchOpermsgSendAPIRequest, session string) (*yunos.YunosCloudcardBatchOpermsgSendAPIResponse, error) {
	var resp yunos.YunosCloudcardBatchOpermsgSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
