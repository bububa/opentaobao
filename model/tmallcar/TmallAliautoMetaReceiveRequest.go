package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车说明书元数据上传 API请求
tmall.aliauto.meta.receive

天猫汽车对外提供的汽车资源元数据上传接口
*/
type TmallAliautoMetaReceiveRequest struct {
    model.Params
    // 元数据参数集
    command   *ResourceMetaCommand
}

// 初始化TmallAliautoMetaReceiveRequest对象
func NewTmallAliautoMetaReceiveRequest() *TmallAliautoMetaReceiveRequest{
    return &TmallAliautoMetaReceiveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallAliautoMetaReceiveRequest) GetApiMethodName() string {
    return "tmall.aliauto.meta.receive"
}

// IRequest interface 方法, 获取API参数
func (r TmallAliautoMetaReceiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Command Setter
// 元数据参数集
func (r *TmallAliautoMetaReceiveRequest) SetCommand(command *ResourceMetaCommand) error {
    r.command = command
    r.Set("command", command)
    return nil
}

// Command Getter
func (r TmallAliautoMetaReceiveRequest) GetCommand() *ResourceMetaCommand {
    return r.command
}
