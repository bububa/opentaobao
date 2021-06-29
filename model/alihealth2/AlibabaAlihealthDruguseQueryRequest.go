package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
合理用药规则查询 API请求
alibaba.alihealth.druguse.query

查询用户购买的药品命中的风险规则
*/
type AlibabaAlihealthDruguseQueryRequest struct {
    model.Params
    // 入参
    command   *SafeMedicationReqCommand
}

// 初始化AlibabaAlihealthDruguseQueryRequest对象
func NewAlibabaAlihealthDruguseQueryRequest() *AlibabaAlihealthDruguseQueryRequest{
    return &AlibabaAlihealthDruguseQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDruguseQueryRequest) GetApiMethodName() string {
    return "alibaba.alihealth.druguse.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDruguseQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Command Setter
// 入参
func (r *AlibabaAlihealthDruguseQueryRequest) SetCommand(command *SafeMedicationReqCommand) error {
    r.command = command
    r.Set("command", command)
    return nil
}

// Command Getter
func (r AlibabaAlihealthDruguseQueryRequest) GetCommand() *SafeMedicationReqCommand {
    return r.command
}
