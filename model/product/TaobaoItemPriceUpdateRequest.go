package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新商品价格 API请求
taobao.item.price.update

更新商品价格
*/
type TaobaoItemPriceUpdateRequest struct {
    model.Params
    // 所在地省份。如浙江 具体可以下载http://dl.open.taobao.com/sdk/商品城市列表.rar 取到
    locationState   string
    // 所在地城市。如杭州 具体可以下载http://dl.open.taobao.com/sdk/商品城市列表.rar 取到
    locationCity   string
    // 商品数字ID，该参数必须
    numIid   int64
    // 叶子类目id
    cid   int64
    // 商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值。
    props   string
    // 商品数量，取值范围:0-999999的整数。且需要等于Sku所有数量的和
    num   int64
    // 商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。
    price   float64
    // 宝贝标题. 不能超过60字符,受违禁词控制
    title   string
    // 商品描述. 字数要大于5个字符，小于25000个字符 ，受违禁词控制
    desc   string
    // 平邮费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:5.07，表示:5元7分, 注:post_fee,express_fee,ems_fee需一起填写
    postFee   float64
    // 快递费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:15.07，表示:15元7分
    expressFee   float64
    // ems费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:25.07，表示:25元7分
    emsFee   float64
    // 上架时间。不论是更新架下的商品还是出售中的商品，如果这个字段小于当前时间则直接上架商品，并且上架的时间为更新商品的时间，此时item.is_timing为false，如果大于当前时间则宝贝会下架进入定时上架的宝贝中。
    listTime   string
    // 加价幅度 如果为0，代表系统代理幅度
    increment   float64
    // 商品图片。类型:JPG,GIF;最大长度:500k
    image   []*model.File
    // 商品图片。类型:JPG,GIF;最大长度:500k
    image   []*model.File
    // 商品新旧程度。可选值:new（全新）,unused（闲置）,second（二手）。
    stuffStatus   string
    // 商品的积分返点比例。如：5 表示返点比例0.5%. 注意：返点比例必须是>0的整数，而且最大是90,即为9%.B商家在发布非虚拟商品时，返点必须是 5的倍数，即0.5%的倍数。其它是1的倍数，即0.1%的倍数。无名良品商家发布商品时，复用该字段记录积分宝返点比例，返点必须是对应类目的返点步长的整数倍，默认是5，即0.5%。注意此时该字段值依旧必须是>0的整数，注意此时该字段值依旧必须是>0的整数，最高值不超过500，即50%
    auctionPoint   int64
    // 属性值别名。如pid:vid:别名;pid1:vid1:别名1， pid:属性id vid:值id。总长度不超过512字节
    propertyAlias   string
    // 重新关联商品与店铺类目，结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。
    sellerCids   string
    // 宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.postages.get获得当前会话用户的所有邮费模板）
    postageId   int64
    // 商家编码
    outerId   string
    // 商品所属的产品ID(B商家发布商品需要用)
    productId   int64
    // 商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path
    picPath   string
    // 代充商品类型。只有少数类目下的商品可以标记上此字段，具体哪些类目可以上传可以通过taobao.itemcat.features.get获得。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型： <br/>no_mark(不做类型标记) <br/>time_card(点卡软件代充) <br/>fee_card(话费软件代充)
    autoFill   string
    // 是否在淘宝上显示
    isTaobao   bool
    // 是否在外店显示
    isEx   bool
    // 是否是3D
    is3D   bool
    // 是否替换sku
    isReplaceSku   bool
    // 商品文字的版本，繁体传入”zh_HK”，简体传入”zh_CN”
    lang   string
    // 支持会员打折。可选值:true,false;
    hasDiscount   bool
    // 橱窗推荐。可选值:true,false;
    hasShowcase   bool
    // 商品上传后的状态。可选值:onsale（出售中）,instock（库中），如果同时更新商品状态为出售中及list_time为将来的时间，则商品还是处于定时上架的状态, 此时item.is_timing为true
    approveStatus   string
    // 运费承担方式。运费承担方式。可选值:seller（卖家承担）,buyer(买家承担);
    freightPayer   string
    // 有效期。可选值:7,14;单位:天;
    validThru   int64
    // 是否有发票。可选值:true,false (商城卖家此字段必须为true)
    hasInvoice   bool
    // 是否有保修。可选值:true,false;
    hasWarranty   bool
    // 是否承诺退换货服务!虚拟商品无须设置此项!
    sellPromise   bool
    // 货到付款运费模板ID
    codPostageId   int64
    // 实物闪电发货。注意：在售的闪电发货产品不允许取消闪电发货，需要先下架商品才能取消闪电发货标记
    isLightningConsignment   bool
    // 商品的重量(商超卖家专用字段)
    weight   int64
    // 商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置参数就保持原有值。
    isXinpin   bool
    // 商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改 集市卖家默认拍下减库存; 商城卖家默认付款减库存
    subStock   int64
    // 忽略警告提示.
    ignorewarning   string
    // 用户自行输入的类目属性ID串，结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。
    inputPids   string
    // 更新的Sku的数量串，结构如：num1,num2,num3 如:2,3,4
    skuQuantities   string
    // 更新的Sku的价格串，结构如：10.00,5.00,… 精确到2位小数;单位:元。如:200.07，表示:200元7分
    skuPrices   string
    // 更新的Sku的属性串，调用taobao.itemprops.get获取类目属性，如果属性是销售属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid;pid:vid。该字段内的销售属性也需要在props字段填写 。如果更新时有对Sku进行操作，则Sku的properties一定要传入。
    skuProperties   string
    // Sku的外部id串，结构如：1234,1342,… sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，如果没有sku_outer_ids也要写上这个参数，入参是","(这个是两个sku的示列，逗号数应该是sku个数减1)；该参数最大长度是512个字节
    skuOuterIds   string
    // 用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5,Nike乔丹鞋;乔丹系列;乔丹鞋系列;乔丹鞋系列;json5”，多个自定义属性用','分割，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过3999字节。此处不可以使用“其他”、“其它”和“其她”这三个词。
    inputStr   string
}

// 初始化TaobaoItemPriceUpdateRequest对象
func NewTaobaoItemPriceUpdateRequest() *TaobaoItemPriceUpdateRequest{
    return &TaobaoItemPriceUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoItemPriceUpdateRequest) GetApiMethodName() string {
    return "taobao.item.price.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoItemPriceUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LocationState Setter
// 所在地省份。如浙江 具体可以下载http://dl.open.taobao.com/sdk/商品城市列表.rar 取到
func (r *TaobaoItemPriceUpdateRequest) SetLocationState(locationState string) error {
    r.locationState = locationState
    r.Set("location.state", locationState)
    return nil
}

// LocationState Getter
func (r TaobaoItemPriceUpdateRequest) GetLocationState() string {
    return r.locationState
}
// LocationCity Setter
// 所在地城市。如杭州 具体可以下载http://dl.open.taobao.com/sdk/商品城市列表.rar 取到
func (r *TaobaoItemPriceUpdateRequest) SetLocationCity(locationCity string) error {
    r.locationCity = locationCity
    r.Set("location.city", locationCity)
    return nil
}

// LocationCity Getter
func (r TaobaoItemPriceUpdateRequest) GetLocationCity() string {
    return r.locationCity
}
// NumIid Setter
// 商品数字ID，该参数必须
func (r *TaobaoItemPriceUpdateRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

// NumIid Getter
func (r TaobaoItemPriceUpdateRequest) GetNumIid() int64 {
    return r.numIid
}
// Cid Setter
// 叶子类目id
func (r *TaobaoItemPriceUpdateRequest) SetCid(cid int64) error {
    r.cid = cid
    r.Set("cid", cid)
    return nil
}

// Cid Getter
func (r TaobaoItemPriceUpdateRequest) GetCid() int64 {
    return r.cid
}
// Props Setter
// 商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值。
func (r *TaobaoItemPriceUpdateRequest) SetProps(props string) error {
    r.props = props
    r.Set("props", props)
    return nil
}

// Props Getter
func (r TaobaoItemPriceUpdateRequest) GetProps() string {
    return r.props
}
// Num Setter
// 商品数量，取值范围:0-999999的整数。且需要等于Sku所有数量的和
func (r *TaobaoItemPriceUpdateRequest) SetNum(num int64) error {
    r.num = num
    r.Set("num", num)
    return nil
}

// Num Getter
func (r TaobaoItemPriceUpdateRequest) GetNum() int64 {
    return r.num
}
// Price Setter
// 商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。
func (r *TaobaoItemPriceUpdateRequest) SetPrice(price float64) error {
    r.price = price
    r.Set("price", price)
    return nil
}

// Price Getter
func (r TaobaoItemPriceUpdateRequest) GetPrice() float64 {
    return r.price
}
// Title Setter
// 宝贝标题. 不能超过60字符,受违禁词控制
func (r *TaobaoItemPriceUpdateRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r TaobaoItemPriceUpdateRequest) GetTitle() string {
    return r.title
}
// Desc Setter
// 商品描述. 字数要大于5个字符，小于25000个字符 ，受违禁词控制
func (r *TaobaoItemPriceUpdateRequest) SetDesc(desc string) error {
    r.desc = desc
    r.Set("desc", desc)
    return nil
}

// Desc Getter
func (r TaobaoItemPriceUpdateRequest) GetDesc() string {
    return r.desc
}
// PostFee Setter
// 平邮费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:5.07，表示:5元7分, 注:post_fee,express_fee,ems_fee需一起填写
func (r *TaobaoItemPriceUpdateRequest) SetPostFee(postFee float64) error {
    r.postFee = postFee
    r.Set("post_fee", postFee)
    return nil
}

// PostFee Getter
func (r TaobaoItemPriceUpdateRequest) GetPostFee() float64 {
    return r.postFee
}
// ExpressFee Setter
// 快递费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:15.07，表示:15元7分
func (r *TaobaoItemPriceUpdateRequest) SetExpressFee(expressFee float64) error {
    r.expressFee = expressFee
    r.Set("express_fee", expressFee)
    return nil
}

// ExpressFee Getter
func (r TaobaoItemPriceUpdateRequest) GetExpressFee() float64 {
    return r.expressFee
}
// EmsFee Setter
// ems费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:25.07，表示:25元7分
func (r *TaobaoItemPriceUpdateRequest) SetEmsFee(emsFee float64) error {
    r.emsFee = emsFee
    r.Set("ems_fee", emsFee)
    return nil
}

// EmsFee Getter
func (r TaobaoItemPriceUpdateRequest) GetEmsFee() float64 {
    return r.emsFee
}
// ListTime Setter
// 上架时间。不论是更新架下的商品还是出售中的商品，如果这个字段小于当前时间则直接上架商品，并且上架的时间为更新商品的时间，此时item.is_timing为false，如果大于当前时间则宝贝会下架进入定时上架的宝贝中。
func (r *TaobaoItemPriceUpdateRequest) SetListTime(listTime string) error {
    r.listTime = listTime
    r.Set("list_time", listTime)
    return nil
}

// ListTime Getter
func (r TaobaoItemPriceUpdateRequest) GetListTime() string {
    return r.listTime
}
// Increment Setter
// 加价幅度 如果为0，代表系统代理幅度
func (r *TaobaoItemPriceUpdateRequest) SetIncrement(increment float64) error {
    r.increment = increment
    r.Set("increment", increment)
    return nil
}

// Increment Getter
func (r TaobaoItemPriceUpdateRequest) GetIncrement() float64 {
    return r.increment
}
// Image Setter
// 商品图片。类型:JPG,GIF;最大长度:500k
func (r *TaobaoItemPriceUpdateRequest) SetImage(image []*model.File) error {
    r.image = image
    r.Set("image", image)
    return nil
}

// Image Getter
func (r TaobaoItemPriceUpdateRequest) GetImage() []*model.File {
    return r.image
}
// Image Setter
// 商品图片。类型:JPG,GIF;最大长度:500k
func (r *TaobaoItemPriceUpdateRequest) SetImage(image []*model.File) error {
    r.image = image
    r.Set("image", image)
    return nil
}

// Image Getter
func (r TaobaoItemPriceUpdateRequest) GetImage() []*model.File {
    return r.image
}
// StuffStatus Setter
// 商品新旧程度。可选值:new（全新）,unused（闲置）,second（二手）。
func (r *TaobaoItemPriceUpdateRequest) SetStuffStatus(stuffStatus string) error {
    r.stuffStatus = stuffStatus
    r.Set("stuff_status", stuffStatus)
    return nil
}

// StuffStatus Getter
func (r TaobaoItemPriceUpdateRequest) GetStuffStatus() string {
    return r.stuffStatus
}
// AuctionPoint Setter
// 商品的积分返点比例。如：5 表示返点比例0.5%. 注意：返点比例必须是>0的整数，而且最大是90,即为9%.B商家在发布非虚拟商品时，返点必须是 5的倍数，即0.5%的倍数。其它是1的倍数，即0.1%的倍数。无名良品商家发布商品时，复用该字段记录积分宝返点比例，返点必须是对应类目的返点步长的整数倍，默认是5，即0.5%。注意此时该字段值依旧必须是>0的整数，注意此时该字段值依旧必须是>0的整数，最高值不超过500，即50%
func (r *TaobaoItemPriceUpdateRequest) SetAuctionPoint(auctionPoint int64) error {
    r.auctionPoint = auctionPoint
    r.Set("auction_point", auctionPoint)
    return nil
}

// AuctionPoint Getter
func (r TaobaoItemPriceUpdateRequest) GetAuctionPoint() int64 {
    return r.auctionPoint
}
// PropertyAlias Setter
// 属性值别名。如pid:vid:别名;pid1:vid1:别名1， pid:属性id vid:值id。总长度不超过512字节
func (r *TaobaoItemPriceUpdateRequest) SetPropertyAlias(propertyAlias string) error {
    r.propertyAlias = propertyAlias
    r.Set("property_alias", propertyAlias)
    return nil
}

// PropertyAlias Getter
func (r TaobaoItemPriceUpdateRequest) GetPropertyAlias() string {
    return r.propertyAlias
}
// SellerCids Setter
// 重新关联商品与店铺类目，结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。
func (r *TaobaoItemPriceUpdateRequest) SetSellerCids(sellerCids string) error {
    r.sellerCids = sellerCids
    r.Set("seller_cids", sellerCids)
    return nil
}

// SellerCids Getter
func (r TaobaoItemPriceUpdateRequest) GetSellerCids() string {
    return r.sellerCids
}
// PostageId Setter
// 宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.postages.get获得当前会话用户的所有邮费模板）
func (r *TaobaoItemPriceUpdateRequest) SetPostageId(postageId int64) error {
    r.postageId = postageId
    r.Set("postage_id", postageId)
    return nil
}

// PostageId Getter
func (r TaobaoItemPriceUpdateRequest) GetPostageId() int64 {
    return r.postageId
}
// OuterId Setter
// 商家编码
func (r *TaobaoItemPriceUpdateRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoItemPriceUpdateRequest) GetOuterId() string {
    return r.outerId
}
// ProductId Setter
// 商品所属的产品ID(B商家发布商品需要用)
func (r *TaobaoItemPriceUpdateRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r TaobaoItemPriceUpdateRequest) GetProductId() int64 {
    return r.productId
}
// PicPath Setter
// 商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path
func (r *TaobaoItemPriceUpdateRequest) SetPicPath(picPath string) error {
    r.picPath = picPath
    r.Set("pic_path", picPath)
    return nil
}

// PicPath Getter
func (r TaobaoItemPriceUpdateRequest) GetPicPath() string {
    return r.picPath
}
// AutoFill Setter
// 代充商品类型。只有少数类目下的商品可以标记上此字段，具体哪些类目可以上传可以通过taobao.itemcat.features.get获得。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型： <br/>no_mark(不做类型标记) <br/>time_card(点卡软件代充) <br/>fee_card(话费软件代充)
func (r *TaobaoItemPriceUpdateRequest) SetAutoFill(autoFill string) error {
    r.autoFill = autoFill
    r.Set("auto_fill", autoFill)
    return nil
}

// AutoFill Getter
func (r TaobaoItemPriceUpdateRequest) GetAutoFill() string {
    return r.autoFill
}
// IsTaobao Setter
// 是否在淘宝上显示
func (r *TaobaoItemPriceUpdateRequest) SetIsTaobao(isTaobao bool) error {
    r.isTaobao = isTaobao
    r.Set("is_taobao", isTaobao)
    return nil
}

// IsTaobao Getter
func (r TaobaoItemPriceUpdateRequest) GetIsTaobao() bool {
    return r.isTaobao
}
// IsEx Setter
// 是否在外店显示
func (r *TaobaoItemPriceUpdateRequest) SetIsEx(isEx bool) error {
    r.isEx = isEx
    r.Set("is_ex", isEx)
    return nil
}

// IsEx Getter
func (r TaobaoItemPriceUpdateRequest) GetIsEx() bool {
    return r.isEx
}
// Is3D Setter
// 是否是3D
func (r *TaobaoItemPriceUpdateRequest) SetIs3D(is3D bool) error {
    r.is3D = is3D
    r.Set("is_3D", is3D)
    return nil
}

// Is3D Getter
func (r TaobaoItemPriceUpdateRequest) GetIs3D() bool {
    return r.is3D
}
// IsReplaceSku Setter
// 是否替换sku
func (r *TaobaoItemPriceUpdateRequest) SetIsReplaceSku(isReplaceSku bool) error {
    r.isReplaceSku = isReplaceSku
    r.Set("is_replace_sku", isReplaceSku)
    return nil
}

// IsReplaceSku Getter
func (r TaobaoItemPriceUpdateRequest) GetIsReplaceSku() bool {
    return r.isReplaceSku
}
// Lang Setter
// 商品文字的版本，繁体传入”zh_HK”，简体传入”zh_CN”
func (r *TaobaoItemPriceUpdateRequest) SetLang(lang string) error {
    r.lang = lang
    r.Set("lang", lang)
    return nil
}

// Lang Getter
func (r TaobaoItemPriceUpdateRequest) GetLang() string {
    return r.lang
}
// HasDiscount Setter
// 支持会员打折。可选值:true,false;
func (r *TaobaoItemPriceUpdateRequest) SetHasDiscount(hasDiscount bool) error {
    r.hasDiscount = hasDiscount
    r.Set("has_discount", hasDiscount)
    return nil
}

// HasDiscount Getter
func (r TaobaoItemPriceUpdateRequest) GetHasDiscount() bool {
    return r.hasDiscount
}
// HasShowcase Setter
// 橱窗推荐。可选值:true,false;
func (r *TaobaoItemPriceUpdateRequest) SetHasShowcase(hasShowcase bool) error {
    r.hasShowcase = hasShowcase
    r.Set("has_showcase", hasShowcase)
    return nil
}

// HasShowcase Getter
func (r TaobaoItemPriceUpdateRequest) GetHasShowcase() bool {
    return r.hasShowcase
}
// ApproveStatus Setter
// 商品上传后的状态。可选值:onsale（出售中）,instock（库中），如果同时更新商品状态为出售中及list_time为将来的时间，则商品还是处于定时上架的状态, 此时item.is_timing为true
func (r *TaobaoItemPriceUpdateRequest) SetApproveStatus(approveStatus string) error {
    r.approveStatus = approveStatus
    r.Set("approve_status", approveStatus)
    return nil
}

// ApproveStatus Getter
func (r TaobaoItemPriceUpdateRequest) GetApproveStatus() string {
    return r.approveStatus
}
// FreightPayer Setter
// 运费承担方式。运费承担方式。可选值:seller（卖家承担）,buyer(买家承担);
func (r *TaobaoItemPriceUpdateRequest) SetFreightPayer(freightPayer string) error {
    r.freightPayer = freightPayer
    r.Set("freight_payer", freightPayer)
    return nil
}

// FreightPayer Getter
func (r TaobaoItemPriceUpdateRequest) GetFreightPayer() string {
    return r.freightPayer
}
// ValidThru Setter
// 有效期。可选值:7,14;单位:天;
func (r *TaobaoItemPriceUpdateRequest) SetValidThru(validThru int64) error {
    r.validThru = validThru
    r.Set("valid_thru", validThru)
    return nil
}

// ValidThru Getter
func (r TaobaoItemPriceUpdateRequest) GetValidThru() int64 {
    return r.validThru
}
// HasInvoice Setter
// 是否有发票。可选值:true,false (商城卖家此字段必须为true)
func (r *TaobaoItemPriceUpdateRequest) SetHasInvoice(hasInvoice bool) error {
    r.hasInvoice = hasInvoice
    r.Set("has_invoice", hasInvoice)
    return nil
}

// HasInvoice Getter
func (r TaobaoItemPriceUpdateRequest) GetHasInvoice() bool {
    return r.hasInvoice
}
// HasWarranty Setter
// 是否有保修。可选值:true,false;
func (r *TaobaoItemPriceUpdateRequest) SetHasWarranty(hasWarranty bool) error {
    r.hasWarranty = hasWarranty
    r.Set("has_warranty", hasWarranty)
    return nil
}

// HasWarranty Getter
func (r TaobaoItemPriceUpdateRequest) GetHasWarranty() bool {
    return r.hasWarranty
}
// SellPromise Setter
// 是否承诺退换货服务!虚拟商品无须设置此项!
func (r *TaobaoItemPriceUpdateRequest) SetSellPromise(sellPromise bool) error {
    r.sellPromise = sellPromise
    r.Set("sell_promise", sellPromise)
    return nil
}

// SellPromise Getter
func (r TaobaoItemPriceUpdateRequest) GetSellPromise() bool {
    return r.sellPromise
}
// CodPostageId Setter
// 货到付款运费模板ID
func (r *TaobaoItemPriceUpdateRequest) SetCodPostageId(codPostageId int64) error {
    r.codPostageId = codPostageId
    r.Set("cod_postage_id", codPostageId)
    return nil
}

// CodPostageId Getter
func (r TaobaoItemPriceUpdateRequest) GetCodPostageId() int64 {
    return r.codPostageId
}
// IsLightningConsignment Setter
// 实物闪电发货。注意：在售的闪电发货产品不允许取消闪电发货，需要先下架商品才能取消闪电发货标记
func (r *TaobaoItemPriceUpdateRequest) SetIsLightningConsignment(isLightningConsignment bool) error {
    r.isLightningConsignment = isLightningConsignment
    r.Set("is_lightning_consignment", isLightningConsignment)
    return nil
}

// IsLightningConsignment Getter
func (r TaobaoItemPriceUpdateRequest) GetIsLightningConsignment() bool {
    return r.isLightningConsignment
}
// Weight Setter
// 商品的重量(商超卖家专用字段)
func (r *TaobaoItemPriceUpdateRequest) SetWeight(weight int64) error {
    r.weight = weight
    r.Set("weight", weight)
    return nil
}

// Weight Getter
func (r TaobaoItemPriceUpdateRequest) GetWeight() int64 {
    return r.weight
}
// IsXinpin Setter
// 商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置参数就保持原有值。
func (r *TaobaoItemPriceUpdateRequest) SetIsXinpin(isXinpin bool) error {
    r.isXinpin = isXinpin
    r.Set("is_xinpin", isXinpin)
    return nil
}

// IsXinpin Getter
func (r TaobaoItemPriceUpdateRequest) GetIsXinpin() bool {
    return r.isXinpin
}
// SubStock Setter
// 商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改 集市卖家默认拍下减库存; 商城卖家默认付款减库存
func (r *TaobaoItemPriceUpdateRequest) SetSubStock(subStock int64) error {
    r.subStock = subStock
    r.Set("sub_stock", subStock)
    return nil
}

// SubStock Getter
func (r TaobaoItemPriceUpdateRequest) GetSubStock() int64 {
    return r.subStock
}
// Ignorewarning Setter
// 忽略警告提示.
func (r *TaobaoItemPriceUpdateRequest) SetIgnorewarning(ignorewarning string) error {
    r.ignorewarning = ignorewarning
    r.Set("ignorewarning", ignorewarning)
    return nil
}

// Ignorewarning Getter
func (r TaobaoItemPriceUpdateRequest) GetIgnorewarning() string {
    return r.ignorewarning
}
// InputPids Setter
// 用户自行输入的类目属性ID串，结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。
func (r *TaobaoItemPriceUpdateRequest) SetInputPids(inputPids string) error {
    r.inputPids = inputPids
    r.Set("input_pids", inputPids)
    return nil
}

// InputPids Getter
func (r TaobaoItemPriceUpdateRequest) GetInputPids() string {
    return r.inputPids
}
// SkuQuantities Setter
// 更新的Sku的数量串，结构如：num1,num2,num3 如:2,3,4
func (r *TaobaoItemPriceUpdateRequest) SetSkuQuantities(skuQuantities string) error {
    r.skuQuantities = skuQuantities
    r.Set("sku_quantities", skuQuantities)
    return nil
}

// SkuQuantities Getter
func (r TaobaoItemPriceUpdateRequest) GetSkuQuantities() string {
    return r.skuQuantities
}
// SkuPrices Setter
// 更新的Sku的价格串，结构如：10.00,5.00,… 精确到2位小数;单位:元。如:200.07，表示:200元7分
func (r *TaobaoItemPriceUpdateRequest) SetSkuPrices(skuPrices string) error {
    r.skuPrices = skuPrices
    r.Set("sku_prices", skuPrices)
    return nil
}

// SkuPrices Getter
func (r TaobaoItemPriceUpdateRequest) GetSkuPrices() string {
    return r.skuPrices
}
// SkuProperties Setter
// 更新的Sku的属性串，调用taobao.itemprops.get获取类目属性，如果属性是销售属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid;pid:vid。该字段内的销售属性也需要在props字段填写 。如果更新时有对Sku进行操作，则Sku的properties一定要传入。
func (r *TaobaoItemPriceUpdateRequest) SetSkuProperties(skuProperties string) error {
    r.skuProperties = skuProperties
    r.Set("sku_properties", skuProperties)
    return nil
}

// SkuProperties Getter
func (r TaobaoItemPriceUpdateRequest) GetSkuProperties() string {
    return r.skuProperties
}
// SkuOuterIds Setter
// Sku的外部id串，结构如：1234,1342,… sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，如果没有sku_outer_ids也要写上这个参数，入参是","(这个是两个sku的示列，逗号数应该是sku个数减1)；该参数最大长度是512个字节
func (r *TaobaoItemPriceUpdateRequest) SetSkuOuterIds(skuOuterIds string) error {
    r.skuOuterIds = skuOuterIds
    r.Set("sku_outer_ids", skuOuterIds)
    return nil
}

// SkuOuterIds Getter
func (r TaobaoItemPriceUpdateRequest) GetSkuOuterIds() string {
    return r.skuOuterIds
}
// InputStr Setter
// 用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5,Nike乔丹鞋;乔丹系列;乔丹鞋系列;乔丹鞋系列;json5”，多个自定义属性用','分割，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过3999字节。此处不可以使用“其他”、“其它”和“其她”这三个词。
func (r *TaobaoItemPriceUpdateRequest) SetInputStr(inputStr string) error {
    r.inputStr = inputStr
    r.Set("input_str", inputStr)
    return nil
}

// InputStr Getter
func (r TaobaoItemPriceUpdateRequest) GetInputStr() string {
    return r.inputStr
}
