package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车说明书元数据上传 APIRequest
tmall.aliauto.meta.receive

天猫汽车对外提供的汽车资源元数据上传接口
*/
type TmallAliautoMetaReceiveRequest struct {
    model.Params

    // 元数据参数集
    command   *ResourceMetaCommand 

}

func NewTmallAliautoMetaReceiveRequest() *TmallAliautoMetaReceiveRequest{
    return &TmallAliautoMetaReceiveRequest{
        Params: model.NewParams(),
    }
}

func (r TmallAliautoMetaReceiveRequest) GetApiMethodName() string {
    return "tmall.aliauto.meta.receive"
}

func (r TmallAliautoMetaReceiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallAliautoMetaReceiveRequest) SetCommand(command *ResourceMetaCommand) error {
    r.command = command
    r.Set("command", command)
    return nil
}

func (r TmallAliautoMetaReceiveRequest) GetCommand() *ResourceMetaCommand {
    return r.command
}

