package yunos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
YUNOS生活服务群发消息 APIRequest
yunos.cloudcard.batch.opermsg.send

这个是一个群发消息接口，ISV通过该接口给订阅自己服务号的所有YUNOS终端用户发送服务号消息，目前该接口有调用频率限制，具体规则参考YUNOS开放平台文档。
*/
type YunosCloudcardBatchOpermsgSendRequest struct {
    model.Params

    // YUNOS生活服务群消息
    operBatchMsg   *OperBatchMsg 

}

func NewYunosCloudcardBatchOpermsgSendRequest() *YunosCloudcardBatchOpermsgSendRequest{
    return &YunosCloudcardBatchOpermsgSendRequest{
        Params: model.NewParams(),
    }
}

func (r YunosCloudcardBatchOpermsgSendRequest) GetApiMethodName() string {
    return "yunos.cloudcard.batch.opermsg.send"
}

func (r YunosCloudcardBatchOpermsgSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosCloudcardBatchOpermsgSendRequest) SetOperBatchMsg(operBatchMsg *OperBatchMsg) error {
    r.operBatchMsg = operBatchMsg
    r.Set("oper_batch_msg", operBatchMsg)
    return nil
}

func (r YunosCloudcardBatchOpermsgSendRequest) GetOperBatchMsg() *OperBatchMsg {
    return r.operBatchMsg
}

