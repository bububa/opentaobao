package shenjing

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
访客发起开门 API请求
alibaba.ib.shenjing.visitor.pad.opendoor

访客PAD端录入完人脸后，可以点击开门按钮开门。
*/
type AlibabaIbShenjingVisitorPadOpendoorRequest struct {
    model.Params
    // 访客标识
    _id   string
    // padid
    _padId   string
}

// 初始化AlibabaIbShenjingVisitorPadOpendoorRequest对象
func NewAlibabaIbShenjingVisitorPadOpendoorRequest() *AlibabaIbShenjingVisitorPadOpendoorRequest{
    return &AlibabaIbShenjingVisitorPadOpendoorRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIbShenjingVisitorPadOpendoorRequest) GetApiMethodName() string {
    return "alibaba.ib.shenjing.visitor.pad.opendoor"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIbShenjingVisitorPadOpendoorRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 访客标识
func (r *AlibabaIbShenjingVisitorPadOpendoorRequest) SetId(_id string) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AlibabaIbShenjingVisitorPadOpendoorRequest) GetId() string {
    return r._id
}
// PadId Setter
// padid
func (r *AlibabaIbShenjingVisitorPadOpendoorRequest) SetPadId(_padId string) error {
    r._padId = _padId
    r.Set("pad_id", _padId)
    return nil
}

// PadId Getter
func (r AlibabaIbShenjingVisitorPadOpendoorRequest) GetPadId() string {
    return r._padId
}
