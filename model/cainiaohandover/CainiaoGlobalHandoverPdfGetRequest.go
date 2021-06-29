package cainiaohandover

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取面单PDF文件数据 APIRequest
cainiao.global.handover.pdf.get

返回指定大包面单的PDF文件数据
*/
type CainiaoGlobalHandoverPdfGetRequest struct {
    model.Params

    // 用户信息
    userInfo   *UserInfoDto 

    // 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
    client   string 

    // 多语言
    locale   string 

    // 大包编号id
    handoverContentId   int64 

    // 打印数据类型，1：面单、4：发货标签、512：交接清单
    type   int64 

}

func NewCainiaoGlobalHandoverPdfGetRequest() *CainiaoGlobalHandoverPdfGetRequest{
    return &CainiaoGlobalHandoverPdfGetRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoGlobalHandoverPdfGetRequest) GetApiMethodName() string {
    return "cainiao.global.handover.pdf.get"
}

func (r CainiaoGlobalHandoverPdfGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoGlobalHandoverPdfGetRequest) SetUserInfo(userInfo *UserInfoDto) error {
    r.userInfo = userInfo
    r.Set("user_info", userInfo)
    return nil
}

func (r CainiaoGlobalHandoverPdfGetRequest) GetUserInfo() *UserInfoDto {
    return r.userInfo
}

func (r *CainiaoGlobalHandoverPdfGetRequest) SetClient(client string) error {
    r.client = client
    r.Set("client", client)
    return nil
}

func (r CainiaoGlobalHandoverPdfGetRequest) GetClient() string {
    return r.client
}

func (r *CainiaoGlobalHandoverPdfGetRequest) SetLocale(locale string) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

func (r CainiaoGlobalHandoverPdfGetRequest) GetLocale() string {
    return r.locale
}

func (r *CainiaoGlobalHandoverPdfGetRequest) SetHandoverContentId(handoverContentId int64) error {
    r.handoverContentId = handoverContentId
    r.Set("handover_content_id", handoverContentId)
    return nil
}

func (r CainiaoGlobalHandoverPdfGetRequest) GetHandoverContentId() int64 {
    return r.handoverContentId
}

func (r *CainiaoGlobalHandoverPdfGetRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r CainiaoGlobalHandoverPdfGetRequest) GetType() int64 {
    return r.type
}

