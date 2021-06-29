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
    _userInfo   *UserInfoDto
    // 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
    _client   string
    // 多语言
    _locale   string
    // 大包编号id
    _handoverContentId   int64
    // 打印数据类型，1：面单、4：发货标签、512：交接清单
    _type   int64
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
func (r *CainiaoGlobalHandoverPdfGetRequest) SetUserInfo(_userInfo *UserInfoDto) error {
    r._userInfo = _userInfo
    r.Set("user_info", _userInfo)
    return nil
}

// UserInfo Getter
func (r CainiaoGlobalHandoverPdfGetRequest) GetUserInfo() *UserInfoDto {
    return r._userInfo
}
// Client Setter
// 客户端名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称
func (r *CainiaoGlobalHandoverPdfGetRequest) SetClient(_client string) error {
    r._client = _client
    r.Set("client", _client)
    return nil
}

// Client Getter
func (r CainiaoGlobalHandoverPdfGetRequest) GetClient() string {
    return r._client
}
// Locale Setter
// 多语言
func (r *CainiaoGlobalHandoverPdfGetRequest) SetLocale(_locale string) error {
    r._locale = _locale
    r.Set("locale", _locale)
    return nil
}

// Locale Getter
func (r CainiaoGlobalHandoverPdfGetRequest) GetLocale() string {
    return r._locale
}
// HandoverContentId Setter
// 大包编号id
func (r *CainiaoGlobalHandoverPdfGetRequest) SetHandoverContentId(_handoverContentId int64) error {
    r._handoverContentId = _handoverContentId
    r.Set("handover_content_id", _handoverContentId)
    return nil
}

// HandoverContentId Getter
func (r CainiaoGlobalHandoverPdfGetRequest) GetHandoverContentId() int64 {
    return r._handoverContentId
}
// Type Setter
// 打印数据类型，1：面单、4：发货标签、512：交接清单
func (r *CainiaoGlobalHandoverPdfGetRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r CainiaoGlobalHandoverPdfGetRequest) GetType() int64 {
    return r._type
}