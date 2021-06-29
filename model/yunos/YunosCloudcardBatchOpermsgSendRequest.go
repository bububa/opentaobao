package yunos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
YUNOS生活服务群发消息 API请求
yunos.cloudcard.batch.opermsg.send

这个是一个群发消息接口，ISV通过该接口给订阅自己服务号的所有YUNOS终端用户发送服务号消息，目前该接口有调用频率限制，具体规则参考YUNOS开放平台文档。
*/
type YunosCloudcardBatchOpermsgSendRequest struct {
    model.Params
    // YUNOS生活服务群消息
    _operBatchMsg   *OperBatchMsg
}

// 初始化YunosCloudcardBatchOpermsgSendRequest对象
func NewYunosCloudcardBatchOpermsgSendRequest() *YunosCloudcardBatchOpermsgSendRequest{
    return &YunosCloudcardBatchOpermsgSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosCloudcardBatchOpermsgSendRequest) GetApiMethodName() string {
    return "yunos.cloudcard.batch.opermsg.send"
}

// IRequest interface 方法, 获取API参数
func (r YunosCloudcardBatchOpermsgSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OperBatchMsg Setter
// YUNOS生活服务群消息
func (r *YunosCloudcardBatchOpermsgSendRequest) SetOperBatchMsg(_operBatchMsg *OperBatchMsg) error {
    r._operBatchMsg = _operBatchMsg
    r.Set("oper_batch_msg", _operBatchMsg)
    return nil
}

// OperBatchMsg Getter
func (r YunosCloudcardBatchOpermsgSendRequest) GetOperBatchMsg() *OperBatchMsg {
    return r._operBatchMsg
}
