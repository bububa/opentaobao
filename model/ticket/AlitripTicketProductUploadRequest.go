package ticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【门票API2.0】门票收费项目管理接口 API请求
alitrip.ticket.product.upload

航旅度假新版门票商品（门票收费项目）管理接口：支持门票商品的发布、编辑。如果在ali_product_id下没有发布过门票商品，则系统将判断为新发布商品，否则是编辑已有商品。可以通过辅助查询接口判断是否已在某个ali_product_id下发布过门票商品。
对应新发布商品的情况，有些参数是必填的，请仔细查看各字段说明。对于编辑商品的情况，ali_product_id和out_product_id至少需要填一个，其他参数都是可选，编辑情况支持增量更新（某个参数不传则使用该商品上原有值）。
*/
type AlitripTicketProductUploadRequest struct {
    model.Params
    // 新发布门票商品时必填。阿里旅行提供的收费项目编码，与商户收费项目编码进行绑定。注：一个收费项目编码对应了一个淘宝商品
    _aliProductId   int64
    // 可选，商户自定义收费项目编码。注：如果商户上传了自定义收费项目编码，则在价格库存同步接口可以使用该自定义编码更新价格库存。
    _outProductId   string
    // 可选，商家想要更新映射关系的时候，用于更新商户自定义收费项目编码。
    _updateOutProductId   string
    // 商户收费项目名称
    _outProductName   string
    // 新发布门票商品时必填。门票商品 库存类型。1、日历库存， 2、非日历库存
    _inventoryType   int64
    // inventory_type=2时必填，指定该收费项目在购买时是否需要买家指定入园日期。1、需要，2-不需要
    _needEnterDate   int64
    // 新发布门票商品时必填。门票有效期：指定入园日期后 多少天内有效。当为数字时，表示多少天内有效；当为日期时，表示到某日期有效，日期格式：yyyy-MM-dd。发布时不填则默认设置30天内有效
    _expireDate   string
    // 新发布门票商品时必填。门票 预定时间限制。1、表示无限制 购买后可立即入园，2、有限制，此时预定时间限制规则必填。
    _reserveLimitType   int64
    // 特殊选填，门票 预定时间限制规则。格式：1_18_00_3 或者 1_18_00_3_30，含义：必须提前1天拍下，且在18点00分前支付成功，订单才生效。当为提前0天时（即当日票），最后两个数字才生效，指当日票需要在出票后3小时30分钟后方可入园。
    _reserveLimitRule   string
    // 新发布门票商品时必填。门票商品发码方式
    _codeSendingInfo   *CodeSendingInfo
    // 可选，商品标题
    _title   string
    // 可选，商品主图，最多支持5张。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为png、jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。
    _picUrls   []string
    // 可选，商品详情描述，不超过50000个字符。详情描述支持纯文本描述，也支持html格式的详情描述。html格式的详情描述中 图片链接支持外链图片（必须外网可访问， 且格式为png、jpg或jpeg，大小在500k以内）和淘宝图片空间链接。
    _desc   string
}

// 初始化AlitripTicketProductUploadRequest对象
func NewAlitripTicketProductUploadRequest() *AlitripTicketProductUploadRequest{
    return &AlitripTicketProductUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTicketProductUploadRequest) GetApiMethodName() string {
    return "alitrip.ticket.product.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTicketProductUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AliProductId Setter
// 新发布门票商品时必填。阿里旅行提供的收费项目编码，与商户收费项目编码进行绑定。注：一个收费项目编码对应了一个淘宝商品
func (r *AlitripTicketProductUploadRequest) SetAliProductId(_aliProductId int64) error {
    r._aliProductId = _aliProductId
    r.Set("ali_product_id", _aliProductId)
    return nil
}

// AliProductId Getter
func (r AlitripTicketProductUploadRequest) GetAliProductId() int64 {
    return r._aliProductId
}
// OutProductId Setter
// 可选，商户自定义收费项目编码。注：如果商户上传了自定义收费项目编码，则在价格库存同步接口可以使用该自定义编码更新价格库存。
func (r *AlitripTicketProductUploadRequest) SetOutProductId(_outProductId string) error {
    r._outProductId = _outProductId
    r.Set("out_product_id", _outProductId)
    return nil
}

// OutProductId Getter
func (r AlitripTicketProductUploadRequest) GetOutProductId() string {
    return r._outProductId
}
// UpdateOutProductId Setter
// 可选，商家想要更新映射关系的时候，用于更新商户自定义收费项目编码。
func (r *AlitripTicketProductUploadRequest) SetUpdateOutProductId(_updateOutProductId string) error {
    r._updateOutProductId = _updateOutProductId
    r.Set("update_out_product_id", _updateOutProductId)
    return nil
}

// UpdateOutProductId Getter
func (r AlitripTicketProductUploadRequest) GetUpdateOutProductId() string {
    return r._updateOutProductId
}
// OutProductName Setter
// 商户收费项目名称
func (r *AlitripTicketProductUploadRequest) SetOutProductName(_outProductName string) error {
    r._outProductName = _outProductName
    r.Set("out_product_name", _outProductName)
    return nil
}

// OutProductName Getter
func (r AlitripTicketProductUploadRequest) GetOutProductName() string {
    return r._outProductName
}
// InventoryType Setter
// 新发布门票商品时必填。门票商品 库存类型。1、日历库存， 2、非日历库存
func (r *AlitripTicketProductUploadRequest) SetInventoryType(_inventoryType int64) error {
    r._inventoryType = _inventoryType
    r.Set("inventory_type", _inventoryType)
    return nil
}

// InventoryType Getter
func (r AlitripTicketProductUploadRequest) GetInventoryType() int64 {
    return r._inventoryType
}
// NeedEnterDate Setter
// inventory_type=2时必填，指定该收费项目在购买时是否需要买家指定入园日期。1、需要，2-不需要
func (r *AlitripTicketProductUploadRequest) SetNeedEnterDate(_needEnterDate int64) error {
    r._needEnterDate = _needEnterDate
    r.Set("need_enter_date", _needEnterDate)
    return nil
}

// NeedEnterDate Getter
func (r AlitripTicketProductUploadRequest) GetNeedEnterDate() int64 {
    return r._needEnterDate
}
// ExpireDate Setter
// 新发布门票商品时必填。门票有效期：指定入园日期后 多少天内有效。当为数字时，表示多少天内有效；当为日期时，表示到某日期有效，日期格式：yyyy-MM-dd。发布时不填则默认设置30天内有效
func (r *AlitripTicketProductUploadRequest) SetExpireDate(_expireDate string) error {
    r._expireDate = _expireDate
    r.Set("expire_date", _expireDate)
    return nil
}

// ExpireDate Getter
func (r AlitripTicketProductUploadRequest) GetExpireDate() string {
    return r._expireDate
}
// ReserveLimitType Setter
// 新发布门票商品时必填。门票 预定时间限制。1、表示无限制 购买后可立即入园，2、有限制，此时预定时间限制规则必填。
func (r *AlitripTicketProductUploadRequest) SetReserveLimitType(_reserveLimitType int64) error {
    r._reserveLimitType = _reserveLimitType
    r.Set("reserve_limit_type", _reserveLimitType)
    return nil
}

// ReserveLimitType Getter
func (r AlitripTicketProductUploadRequest) GetReserveLimitType() int64 {
    return r._reserveLimitType
}
// ReserveLimitRule Setter
// 特殊选填，门票 预定时间限制规则。格式：1_18_00_3 或者 1_18_00_3_30，含义：必须提前1天拍下，且在18点00分前支付成功，订单才生效。当为提前0天时（即当日票），最后两个数字才生效，指当日票需要在出票后3小时30分钟后方可入园。
func (r *AlitripTicketProductUploadRequest) SetReserveLimitRule(_reserveLimitRule string) error {
    r._reserveLimitRule = _reserveLimitRule
    r.Set("reserve_limit_rule", _reserveLimitRule)
    return nil
}

// ReserveLimitRule Getter
func (r AlitripTicketProductUploadRequest) GetReserveLimitRule() string {
    return r._reserveLimitRule
}
// CodeSendingInfo Setter
// 新发布门票商品时必填。门票商品发码方式
func (r *AlitripTicketProductUploadRequest) SetCodeSendingInfo(_codeSendingInfo *CodeSendingInfo) error {
    r._codeSendingInfo = _codeSendingInfo
    r.Set("code_sending_info", _codeSendingInfo)
    return nil
}

// CodeSendingInfo Getter
func (r AlitripTicketProductUploadRequest) GetCodeSendingInfo() *CodeSendingInfo {
    return r._codeSendingInfo
}
// Title Setter
// 可选，商品标题
func (r *AlitripTicketProductUploadRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r AlitripTicketProductUploadRequest) GetTitle() string {
    return r._title
}
// PicUrls Setter
// 可选，商品主图，最多支持5张。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为png、jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。
func (r *AlitripTicketProductUploadRequest) SetPicUrls(_picUrls []string) error {
    r._picUrls = _picUrls
    r.Set("pic_urls", _picUrls)
    return nil
}

// PicUrls Getter
func (r AlitripTicketProductUploadRequest) GetPicUrls() []string {
    return r._picUrls
}
// Desc Setter
// 可选，商品详情描述，不超过50000个字符。详情描述支持纯文本描述，也支持html格式的详情描述。html格式的详情描述中 图片链接支持外链图片（必须外网可访问， 且格式为png、jpg或jpeg，大小在500k以内）和淘宝图片空间链接。
func (r *AlitripTicketProductUploadRequest) SetDesc(_desc string) error {
    r._desc = _desc
    r.Set("desc", _desc)
    return nil
}

// Desc Getter
func (r AlitripTicketProductUploadRequest) GetDesc() string {
    return r._desc
}
