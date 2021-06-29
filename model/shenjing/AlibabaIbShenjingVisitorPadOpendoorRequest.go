package shenjing

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
访客发起开门 APIRequest
alibaba.ib.shenjing.visitor.pad.opendoor

访客PAD端录入完人脸后，可以点击开门按钮开门。
*/
type AlibabaIbShenjingVisitorPadOpendoorRequest struct {
    model.Params

    // 访客标识
    id   string 

    // padid
    padId   string 

}

func NewAlibabaIbShenjingVisitorPadOpendoorRequest() *AlibabaIbShenjingVisitorPadOpendoorRequest{
    return &AlibabaIbShenjingVisitorPadOpendoorRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIbShenjingVisitorPadOpendoorRequest) GetApiMethodName() string {
    return "alibaba.ib.shenjing.visitor.pad.opendoor"
}

func (r AlibabaIbShenjingVisitorPadOpendoorRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIbShenjingVisitorPadOpendoorRequest) SetId(id string) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AlibabaIbShenjingVisitorPadOpendoorRequest) GetId() string {
    return r.id
}

func (r *AlibabaIbShenjingVisitorPadOpendoorRequest) SetPadId(padId string) error {
    r.padId = padId
    r.Set("pad_id", padId)
    return nil
}

func (r AlibabaIbShenjingVisitorPadOpendoorRequest) GetPadId() string {
    return r.padId
}

