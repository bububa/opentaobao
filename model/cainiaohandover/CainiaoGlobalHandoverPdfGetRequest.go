package cainiaohandover

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取面单PDF文件数据 API请求
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

// 初始化CainiaoGlobalHandoverPdfGetRequest对象
func NewCainiaoGlobalHandoverPdfGetRequest() *CainiaoGlobalHandoverPdfGetRequest{
    return &CainiaoGlobalHandoverPdfGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoGlobalHandoverPdfGetRequest) GetApiMethodName() string {
    return "cainiao.global.handover.pdf.get"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoGlobalHandoverPdfGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserInfo Setter
// 用户信息
func (r *CainiaoGlobalHandoverPdfGetRequest) SetUserInfo(userInfo *UserInfoDto) error {
    r.userInfo = userInfo
    r.Set("user_info", userInfo)
    return nil
}

// UserInfo Getter
func (r CainiaoGlobalHandoverPdfGetRequest) GetUserInfo() *UserInfoDto {
    return r.userInfo
}
// Client Setter
// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoGlobalHandoverPdfGetRequest) SetClient(client string) error {
    r.client = client
    r.Set("client", client)
    return nil
}

// Client Getter
func (r CainiaoGlobalHandoverPdfGetRequest) GetClient() string {
    return r.client
}
// Locale Setter
// 多语言
func (r *CainiaoGlobalHandoverPdfGetRequest) SetLocale(locale string) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

// Locale Getter
func (r CainiaoGlobalHandoverPdfGetRequest) GetLocale() string {
    return r.locale
}
// HandoverContentId Setter
// 大包编号id
func (r *CainiaoGlobalHandoverPdfGetRequest) SetHandoverContentId(handoverContentId int64) error {
    r.handoverContentId = handoverContentId
    r.Set("handover_content_id", handoverContentId)
    return nil
}

// HandoverContentId Getter
func (r CainiaoGlobalHandoverPdfGetRequest) GetHandoverContentId() int64 {
    return r.handoverContentId
}
// Type Setter
// 打印数据类型，1：面单、4：发货标签、512：交接清单
func (r *CainiaoGlobalHandoverPdfGetRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r CainiaoGlobalHandoverPdfGetRequest) GetType() int64 {
    return r.type
}
