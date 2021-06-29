package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
合理用药规则查询 APIRequest
alibaba.alihealth.druguse.query

查询用户购买的药品命中的风险规则
*/
type AlibabaAlihealthDruguseQueryRequest struct {
    model.Params

    // 入参
    command   *SafeMedicationReqCommand 

}

func NewAlibabaAlihealthDruguseQueryRequest() *AlibabaAlihealthDruguseQueryRequest{
    return &AlibabaAlihealthDruguseQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDruguseQueryRequest) GetApiMethodName() string {
    return "alibaba.alihealth.druguse.query"
}

func (r AlibabaAlihealthDruguseQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDruguseQueryRequest) SetCommand(command *SafeMedicationReqCommand) error {
    r.command = command
    r.Set("command", command)
    return nil
}

func (r AlibabaAlihealthDruguseQueryRequest) GetCommand() *SafeMedicationReqCommand {
    return r.command
}

