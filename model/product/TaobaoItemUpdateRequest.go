package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新商品信息 APIRequest
taobao.item.update

根据传入的num_iid更新对应的商品的数据。
传入的num_iid所对应的商品必须属于当前会话的用户 
商品的属性和sku的属性有包含的关系，商品的价格要位于sku的价格区间之中（例如，sku价格有5元、10元两种，那么商品的价格就需要大于等于5元，小于等于10元，否则更新商品会失败） 
商品的类目和商品的价格、sku的价格都有一定的相关性（具体的关系要通过类目属性查询接口获得） 
当关键属性值更新为“其他”的时候，需要输入input_pids和input_str商品才能更新成功。该接口不支持产品属性修改。
*/
type TaobaoItemUpdateRequest struct {
    model.Params

    // 针对当前商品的自定义属性值
    inputCustomCpv   string 

    // 针对当前商品的标准属性值的补充说明，让买家更加了解商品信息减少交易纠纷
    cpvMemo   string 

    // 此参数暂时不起作用
    skuSpecIds   string 

    // 此参数暂时不起作用
    skuDeliveryTimes   string 

    // 家装建材类目，商品SKU的长度，正整数，单位为cm，部分类目必选。 数据和SKU一一对应，用,分隔，如：20,30,30
    skuHdLength   string 

    // 家装建材类目，商品SKU的高度，单位为cm，部分类目必选。 天猫和淘宝格式不同。天猫：可选值为："0-15", "15-25", "25-50", "50-60", "60-80", "80-120", "120-160", "160-200"。 数据和SKU一一对应，用,分隔，格式如：15-25,25-50,25-50。 淘宝：正整数，单位为cm,格式如：20,30,30
    skuHdHeight   string 

    // 家装建材类目，商品SKU的灯头数量，正整数，大于等于3，部分类目必选。 数据和SKU一一对应，用,分隔，如：3,5,7
    skuHdLampQuantity   string 

    // 所在地省份。如浙江
    locationState   string 

    // 所在地城市。如杭州
    locationCity   string 

    // 生产许可证号
    foodSecurityPrdLicenseNo   string 

    // 产品标准号
    foodSecurityDesignCode   string 

    // 厂名
    foodSecurityFactory   string 

    // 厂址
    foodSecurityFactorySite   string 

    // 厂家联系方式
    foodSecurityContact   string 

    // 配料表
    foodSecurityMix   string 

    // 储藏方法
    foodSecurityPlanStorage   string 

    // 保质期，默认有单位，传入数字
    foodSecurityPeriod   string 

    // 食品添加剂
    foodSecurityFoodAdditive   string 

    // 供货商
    foodSecuritySupplier   string 

    // 生产开始日期，格式必须为yyyy-MM-dd
    foodSecurityProductDateStart   string 

    // 生产结束日期,格式必须为yyyy-MM-dd
    foodSecurityProductDateEnd   string 

    // 进货开始日期，要在生产日期之后，格式必须为yyyy-MM-dd
    foodSecurityStockDateStart   string 

    // 进货结束日期，要在生产日期之后，格式必须为yyyy-MM-dd
    foodSecurityStockDateEnd   string 

    // 健字号，保健品/膳食营养补充剂 这个类目下特有的信息，此类目下无需填写生产许可证编号（QS），如果填写了生产许可证编号（QS）将被忽略不保存；保存宝贝时，标题前会自动加上健字号产品名称一起作为宝贝标题；
    foodSecurityHealthProductNo   string 

    // 预约门店是否支持门店自提,1:是
    localityLifeObs   string 

    // 电子凭证版本 新版电子凭证值:1
    localityLifeVersion   string 

    // 新版电子凭证包id
    localityLifePackageid   string 

    // 订金
    msPaymentPrice   string 

    // 尾款可抵扣金额
    msPaymentVoucherPrice   string 

    // 参考价
    msPaymentReferencePrice   string 

    // 商品级别设置的发货时间。设置了商品级别的发货时间，相对发货时间，则填写相对发货时间的天数（大于3）；绝对发货时间，则填写yyyy-mm-dd格式，如2013-11-11
    deliveryTimeDeliveryTime   string 

    // 发货时间类型：绝对发货时间或者相对发货时间
    deliveryTimeDeliveryTimeType   string 

    // 设置是否使用发货时间，商品级别，sku级别
    deliveryTimeNeedDeliveryTime   string 

    // 商品数字ID，该参数必须
    numIid   int64 

    // 叶子类目id
    cid   int64 

    // 商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值。
    props   string 

    // 商品数量，取值范围:0-900000000的整数。且需要等于Sku所有数量的和 拍卖商品中增加拍只能为1，荷兰拍要在[2,500)范围内。
    num   int64 

    // 商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。 拍卖商品对应的起拍价。
    price   float64 

    // 宝贝标题. 不能超过30字符,受违禁词控制
    title   string 

    // 商品描述. 字数要大于5个字符，小于25000个字符 ，受违禁词控制
    desc   string 

    // 平邮费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:5.07，表示:5元7分, 注:post_fee,express_fee,ems_fee需一起填写
    postFee   float64 

    // 快递费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:15.07，表示:15元7分
    expressFee   float64 

    // ems费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:25.07，表示:25元7分
    emsFee   float64 

    // 上架时间。大于当前时间则宝贝会下架进入定时上架的宝贝中。
    listTime   string 

    // 加价(降价)幅度。如果为0，代表系统代理幅度。对于增价拍和荷兰拍来说是加价幅度，对于降价拍来说是降价幅度。
    increment   float64 

    // 商品图片。类型:JPG,GIF;最大长度:3M
    image   []*model.File 

    // 商品新旧程度。可选值:new（全新）,unused（闲置）,second（二手）。
    stuffStatus   string 

    // 商品的积分返点比例。如：5 表示返点比例0.5%. 注意：返点比例必须是>0的整数，而且最大是90,即为9%.B商家在发布非虚拟商品时，返点必须是 5的倍数，即0.5%的倍数。其它是1的倍数，即0.1%的倍数。无名良品商家发布商品时，复用该字段记录积分宝返点比例，返点必须是对应类目的返点步长的整数倍，默认是5，即0.5%。注意此时该字段值依旧必须是>0的整数，注意此时该字段值依旧必须是>0的整数，最高值不超过500，即50%
    auctionPoint   int64 

    // 属性值别名。如pid:vid:别名;pid1:vid1:别名1， pid:属性id vid:值id。总长度不超过800个字符，如"123:333:你好"，引号内的是10个字符。
    propertyAlias   string 

    // 重新关联商品与店铺类目，结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。
    sellerCids   string 

    // 宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.delivery.templates.get获得当前会话用户的所有邮费模板）
    postageId   int64 

    // 商家编码
    outerId   string 

    // 商品所属的产品ID(B商家发布商品需要用)
    productId   int64 

    // 商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path
    picPath   string 

    // 代充商品类型。只有少数类目下的商品可以标记上此字段，具体哪些类目可以上传可以通过taobao.itemcat.features.get获得。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型： no_mark(不做类型标记) time_card(点卡软件代充) fee_card(话费软件代充)
    autoFill   string 

    // 是否在淘宝上显示（如果传FALSE，则在淘宝主站无法显示该商品）
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

    // 货到付款运费模板ID该字段已经废弃，货到付款模板已经集成到运费模板中。
    codPostageId   int64 

    // 实物闪电发货。注意：在售的闪电发货产品不允许取消闪电发货，需要先下架商品才能取消闪电发货标记
    isLightningConsignment   bool 

    // 商品的重量(商超卖家专用字段)
    weight   int64 

    // 商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置参数就保持原有值。
    isXinpin   bool 

    // 商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改 集市卖家默认拍下减库存; 商城卖家默认付款减库存
    subStock   int64 

    // 表示商品的体积，如果需要使用按体积计费的运费模板，一定要设置这个值。该值的单位为立方米（m3），如果是其他单位，请转换成成立方米。该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）。体积和长宽高都支持小数类型。在传入体积或长宽高时候，不能带单位。体积的单位默认为立方米（m3），长宽高的单位默认为米(m)在编辑的时候，如果需要删除体积属性，请设置该值为0，如bulk:0
    itemSize   string 

    // 商品的重量，用于按重量计费的运费模板。注意：单位为kg。 只能传入数值类型（包含小数），不能带单位，单位默认为kg。 在编辑时候，如果需要在商品里删除重量的信息，就需要将值设置为0
    itemWeight   string 

    // 商品卖点信息，最长150个字符。天猫和集市都可用
    sellPoint   string 

    // 商品条形码
    barcode   string 

    // 该宝贝是否支持【7天无理由退货】，卖家选择的值只是一个因素，最终以类目和选择的属性条件来确定是否支持7天。填入字符0，表示不支持；未填写或填人字符1，表示支持7天无理由退货；<br>注意：使用该API修改商品其它属性如标题title时，如需保持商品不支持7天无理由退货状态，该字段需传入0 。
    newprepay   string 

    // 商品资质信息
    qualification   string 

    // 汽车O2O绑定线下服务标记，如不为空，表示关联服务，否则，不关联服务。
    o2oBindService   bool 

    // 忽略警告提示.
    ignorewarning   string 

    // 宝贝特征值，格式为：【key1:value1;key2:value2;key3:value3;】，key和value用【:】分隔，key&value之间用【;】分隔，只有在Top支持的特征值才能保存到宝贝上，目前支持的Key列表为：mysize_tp,是指尺码库对应的key
    features   string 

    // 售后说明模板id
    afterSaleId   int64 

    // 基础色数据，淘宝不使用
    changeProp   string 

    // 已废弃
    descModules   string 

    // 是否是线下商品。1：线上商品（默认值）；2：线上或线下商品；3：线下商品。
    isOffline   string 

    // 无线的宝贝描述
    wirelessDesc   string 

    // 手机类目spu 产品信息确认声明
    spuConfirm   bool 

    // 用户自行输入的类目属性ID串，结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。
    inputPids   string 

    // 更新的Sku的数量串，结构如：num1,num2,num3 如:2,3,4
    skuQuantities   string 

    // 更新的Sku的价格串，结构如：10.00,5.00,… 精确到2位小数;单位:元。如:200.07，表示:200元7分
    skuPrices   string 

    // 更新的sku的属性串，调用taobao.itemprops.get获取。格式:pid1:vid;pid2:vid,多个sku属性之间用逗号分隔。该字段内的属性需要在props字段同时包含。如果新增商品包含了sku，则此字段一定要传入,字段长度要控制在512个字节以内。
    skuProperties   string 

    // Sku的外部id串，结构如：1234,1342,… sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，如果没有sku_outer_ids也要写上这个参数，入参是","(这个是两个sku的示列，逗号数应该是sku个数减1)；该参数最大长度是512个字节
    skuOuterIds   string 

    // 用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5,Nike乔丹鞋;乔丹系列;乔丹鞋系列;乔丹鞋系列;json5”，多个自定义属性用','分割，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过3999字节。此处不可以使用“其他”、“其它”和“其她”这三个词。
    inputStr   string 

    // sku层面的条形码，多个SKU情况，与SKU价格库存格式类似，用逗号分隔
    skuBarcode   string 

    // 主图视频id
    videoId   int64 

    // 主图视频互动信息id，必须有主图视频id才能传互动信息id
    interactiveId   int64 

    // 淘宝租赁扩展信息
    leaseExtendsInfo   string 

    // 本地生活电子交易凭证业务，目前此字段只涉及到的信息为有效期;如果有效期为起止日期类型，此值为2012-08-06,2012-08-16如果有效期为【购买成功日 至】类型则格式为2012-08-16如果有效期为天数类型则格式为15
    localityLifeExpirydate   string 

    // 网点ID,在参数empty_fields里设置locality_life.network_id可删除网点ID
    localityLifeNetworkId   string 

    // 码商信息，格式为 码商id:nick
    localityLifeMerchant   string 

    // 核销打款,1代表核销打款 0代表非核销打款; 在参数empty_fields里设置locality_life.verification可删除核销打款
    localityLifeVerification   string 

    // 退款比例，百分比%前的数字,1-100的正整数值; 在参数empty_fields里设置locality_life.refund_ratio可删除退款比例
    localityLifeRefundRatio   int64 

    // 编辑电子凭证宝贝时候表示是否使用邮寄0: 代表不使用邮寄；1：代表使用邮寄；如果不设置这个值，代表不使用邮寄
    localityLifeChooseLogis   string 

    // 电子凭证售中自动退款比例，百分比%前的数字，介于1-100之间的整数
    localityLifeOnsaleAutoRefundRatio   int64 

    // 退款码费承担方。发布电子凭证宝贝的时候会增加“退款码费承担方”配置项，可选填：(1)s（卖家承担） (2)b(买家承担)
    localityLifeRefundmafee   string 

    // 电子凭证业务属性，数据字典是: 1、is_card:1 (暂时不用) 2、consume_way:4 （1 串码 ，4 身份证）3、consume_midmnick ：(核销放行账号:用户id-用户名，支持多个，用逗号分隔,例如 1234-测试账号35,1345-测试账号56）4、market:eticket (电子凭证商品标记) 5、has_pos:1 (1 表示商品配置线下门店，在detail上进行展示 ，没有或者其他值只不展示)格式是: k1:v2;k2:v2;........ 如：has_pos:1;market:eticket;consume_midmnick:901409638-OPPO;consume_way:4
    localityLifeEticket   string 

    // 支持宝贝信息的删除,如需删除对应的食品安全信息中的储藏方法、保质期， 则应该设置此参数的值为：food_security.plan_storage,food_security.period; 各个参数的名称之间用【,】分割, 如果对应的参数有设置过值，即使在这个列表中，也不会被删除; 目前支持此功能的宝贝信息如下：食品安全信息所有字段、电子交易凭证字段（locality_life，locality_life.verification，locality_life.refund_ratio，locality_life.network_id ，locality_life.onsale_auto_refund_ratio）。支持对全球购宝贝信息的清除（字符串中包含global_stock）
    emptyFields   string 

    // 景区门票类宝贝编辑时候，当卖家签订了支付宝代扣协议时候，需要选择支付方式：全额支付和订金支付。当scenic_ticket_pay_way为1时表示全额支付，为2时表示订金支付
    scenicTicketPayWay   int64 

    // 景区门票在选择订金支付时候，需要交的预订费。传入的值是1到20之间的数值，小数点后最多可以保留两位（多余的部分将做四舍五入的处理）。这个数值表示的是预订费的比例，最终的预订费为 scenic_ticket_book_cost乘一口价除以100
    scenicTicketBookCost   string 

    // 拍卖商品选择的拍卖类型，拍卖类型包括三种：增价拍(1)，荷兰拍(2)和降价拍(3)。
    paimaiInfoMode   int64 

    // 拍卖宝贝的保证金。对于增价拍和荷兰拍来说保证金有两种模式：淘宝默认模式（首次出价金额的10%），自定义固定保证金（固定冻结金额只能输入不超过30万的正整数），并且保证金只冻结1次。对于降价拍来说保证金只有淘宝默认的（竞拍价格的10% * 竞拍数量），并且每次出价都需要冻结保证金。对于拍卖宝贝来说，保证金是必须的，但是默认使用淘宝默认保证金模式，只有用户需要使用自定义固定保证金的时候才需要使用到这个参数。如果该参数不传或传入0则代表使用默认。
    paimaiInfoDeposit   int64 

    // 降价拍宝贝的降价周期(分钟)。降价拍宝贝的价格每隔paimai_info.interval时间会下降一次increment。
    paimaiInfoInterval   int64 

    // 降价拍宝贝的保留价。对于降价拍来说，paimai_info.reserve必须大于0，且小于price-increment，而且（price-paimai_info.reserve）/increment的计算结果必须为整数
    paimaiInfoReserve   float64 

    // 自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。
    paimaiInfoValidHour   int64 

    // 自定义销售周期的分钟数。拍卖宝贝可以自定义销售周期，这里是指定销售周期的分钟数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。
    paimaiInfoValidMinute   int64 

    // 全球购商品采购地（库存类型）全球购商品有两种库存类型：现货和代购 参数值为1时代表现货，值为2时代表代购。注意：使用时请与 全球购商品采购地（地区/国家）配合使用
    globalStockType   string 

    // 全球购商品采购地（地区/国家）,默认值只在全球购商品采购地（库存类型选择情况生效），地区国家值为（美国, 香港, 日本, 英国, 新西兰, 德国, 韩国, 荷兰, 澳洲, 法国, 意大利, 台湾, 澳门, 加拿大, 瑞士, 西班牙, 泰国, 新加坡, 马来西亚, 菲律宾, 其他）
    globalStockCountry   string 

    // 全球购商品发货地，发货地现在有两种类型：“国内”和“海外及港澳台”，参数值为1时代表“国内”，值为2时代表“海外及港澳台”
    globalStockDeliveryPlace   string 

    // 全球购商品卖家包税承诺，当值为true时，代表卖家承诺包税。
    globalStockTaxFreePromise   bool 

}

func NewTaobaoItemUpdateRequest() *TaobaoItemUpdateRequest{
    return &TaobaoItemUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoItemUpdateRequest) GetApiMethodName() string {
    return "taobao.item.update"
}

func (r TaobaoItemUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoItemUpdateRequest) SetInputCustomCpv(inputCustomCpv string) error {
    r.inputCustomCpv = inputCustomCpv
    r.Set("input_custom_cpv", inputCustomCpv)
    return nil
}

func (r TaobaoItemUpdateRequest) GetInputCustomCpv() string {
    return r.inputCustomCpv
}

func (r *TaobaoItemUpdateRequest) SetCpvMemo(cpvMemo string) error {
    r.cpvMemo = cpvMemo
    r.Set("cpv_memo", cpvMemo)
    return nil
}

func (r TaobaoItemUpdateRequest) GetCpvMemo() string {
    return r.cpvMemo
}

func (r *TaobaoItemUpdateRequest) SetSkuSpecIds(skuSpecIds string) error {
    r.skuSpecIds = skuSpecIds
    r.Set("sku_spec_ids", skuSpecIds)
    return nil
}

func (r TaobaoItemUpdateRequest) GetSkuSpecIds() string {
    return r.skuSpecIds
}

func (r *TaobaoItemUpdateRequest) SetSkuDeliveryTimes(skuDeliveryTimes string) error {
    r.skuDeliveryTimes = skuDeliveryTimes
    r.Set("sku_delivery_times", skuDeliveryTimes)
    return nil
}

func (r TaobaoItemUpdateRequest) GetSkuDeliveryTimes() string {
    return r.skuDeliveryTimes
}

func (r *TaobaoItemUpdateRequest) SetSkuHdLength(skuHdLength string) error {
    r.skuHdLength = skuHdLength
    r.Set("sku_hd_length", skuHdLength)
    return nil
}

func (r TaobaoItemUpdateRequest) GetSkuHdLength() string {
    return r.skuHdLength
}

func (r *TaobaoItemUpdateRequest) SetSkuHdHeight(skuHdHeight string) error {
    r.skuHdHeight = skuHdHeight
    r.Set("sku_hd_height", skuHdHeight)
    return nil
}

func (r TaobaoItemUpdateRequest) GetSkuHdHeight() string {
    return r.skuHdHeight
}

func (r *TaobaoItemUpdateRequest) SetSkuHdLampQuantity(skuHdLampQuantity string) error {
    r.skuHdLampQuantity = skuHdLampQuantity
    r.Set("sku_hd_lamp_quantity", skuHdLampQuantity)
    return nil
}

func (r TaobaoItemUpdateRequest) GetSkuHdLampQuantity() string {
    return r.skuHdLampQuantity
}

func (r *TaobaoItemUpdateRequest) SetLocationState(locationState string) error {
    r.locationState = locationState
    r.Set("location.state", locationState)
    return nil
}

func (r TaobaoItemUpdateRequest) GetLocationState() string {
    return r.locationState
}

func (r *TaobaoItemUpdateRequest) SetLocationCity(locationCity string) error {
    r.locationCity = locationCity
    r.Set("location.city", locationCity)
    return nil
}

func (r TaobaoItemUpdateRequest) GetLocationCity() string {
    return r.locationCity
}

func (r *TaobaoItemUpdateRequest) SetFoodSecurityPrdLicenseNo(foodSecurityPrdLicenseNo string) error {
    r.foodSecurityPrdLicenseNo = foodSecurityPrdLicenseNo
    r.Set("food_security.prd_license_no", foodSecurityPrdLicenseNo)
    return nil
}

func (r TaobaoItemUpdateRequest) GetFoodSecurityPrdLicenseNo() string {
    return r.foodSecurityPrdLicenseNo
}

func (r *TaobaoItemUpdateRequest) SetFoodSecurityDesignCode(foodSecurityDesignCode string) error {
    r.foodSecurityDesignCode = foodSecurityDesignCode
    r.Set("food_security.design_code", foodSecurityDesignCode)
    return nil
}

func (r TaobaoItemUpdateRequest) GetFoodSecurityDesignCode() string {
    return r.foodSecurityDesignCode
}

func (r *TaobaoItemUpdateRequest) SetFoodSecurityFactory(foodSecurityFactory string) error {
    r.foodSecurityFactory = foodSecurityFactory
    r.Set("food_security.factory", foodSecurityFactory)
    return nil
}

func (r TaobaoItemUpdateRequest) GetFoodSecurityFactory() string {
    return r.foodSecurityFactory
}

func (r *TaobaoItemUpdateRequest) SetFoodSecurityFactorySite(foodSecurityFactorySite string) error {
    r.foodSecurityFactorySite = foodSecurityFactorySite
    r.Set("food_security.factory_site", foodSecurityFactorySite)
    return nil
}

func (r TaobaoItemUpdateRequest) GetFoodSecurityFactorySite() string {
    return r.foodSecurityFactorySite
}

func (r *TaobaoItemUpdateRequest) SetFoodSecurityContact(foodSecurityContact string) error {
    r.foodSecurityContact = foodSecurityContact
    r.Set("food_security.contact", foodSecurityContact)
    return nil
}

func (r TaobaoItemUpdateRequest) GetFoodSecurityContact() string {
    return r.foodSecurityContact
}

func (r *TaobaoItemUpdateRequest) SetFoodSecurityMix(foodSecurityMix string) error {
    r.foodSecurityMix = foodSecurityMix
    r.Set("food_security.mix", foodSecurityMix)
    return nil
}

func (r TaobaoItemUpdateRequest) GetFoodSecurityMix() string {
    return r.foodSecurityMix
}

func (r *TaobaoItemUpdateRequest) SetFoodSecurityPlanStorage(foodSecurityPlanStorage string) error {
    r.foodSecurityPlanStorage = foodSecurityPlanStorage
    r.Set("food_security.plan_storage", foodSecurityPlanStorage)
    return nil
}

func (r TaobaoItemUpdateRequest) GetFoodSecurityPlanStorage() string {
    return r.foodSecurityPlanStorage
}

func (r *TaobaoItemUpdateRequest) SetFoodSecurityPeriod(foodSecurityPeriod string) error {
    r.foodSecurityPeriod = foodSecurityPeriod
    r.Set("food_security.period", foodSecurityPeriod)
    return nil
}

func (r TaobaoItemUpdateRequest) GetFoodSecurityPeriod() string {
    return r.foodSecurityPeriod
}

func (r *TaobaoItemUpdateRequest) SetFoodSecurityFoodAdditive(foodSecurityFoodAdditive string) error {
    r.foodSecurityFoodAdditive = foodSecurityFoodAdditive
    r.Set("food_security.food_additive", foodSecurityFoodAdditive)
    return nil
}

func (r TaobaoItemUpdateRequest) GetFoodSecurityFoodAdditive() string {
    return r.foodSecurityFoodAdditive
}

func (r *TaobaoItemUpdateRequest) SetFoodSecuritySupplier(foodSecuritySupplier string) error {
    r.foodSecuritySupplier = foodSecuritySupplier
    r.Set("food_security.supplier", foodSecuritySupplier)
    return nil
}

func (r TaobaoItemUpdateRequest) GetFoodSecuritySupplier() string {
    return r.foodSecuritySupplier
}

func (r *TaobaoItemUpdateRequest) SetFoodSecurityProductDateStart(foodSecurityProductDateStart string) error {
    r.foodSecurityProductDateStart = foodSecurityProductDateStart
    r.Set("food_security.product_date_start", foodSecurityProductDateStart)
    return nil
}

func (r TaobaoItemUpdateRequest) GetFoodSecurityProductDateStart() string {
    return r.foodSecurityProductDateStart
}

func (r *TaobaoItemUpdateRequest) SetFoodSecurityProductDateEnd(foodSecurityProductDateEnd string) error {
    r.foodSecurityProductDateEnd = foodSecurityProductDateEnd
    r.Set("food_security.product_date_end", foodSecurityProductDateEnd)
    return nil
}

func (r TaobaoItemUpdateRequest) GetFoodSecurityProductDateEnd() string {
    return r.foodSecurityProductDateEnd
}

func (r *TaobaoItemUpdateRequest) SetFoodSecurityStockDateStart(foodSecurityStockDateStart string) error {
    r.foodSecurityStockDateStart = foodSecurityStockDateStart
    r.Set("food_security.stock_date_start", foodSecurityStockDateStart)
    return nil
}

func (r TaobaoItemUpdateRequest) GetFoodSecurityStockDateStart() string {
    return r.foodSecurityStockDateStart
}

func (r *TaobaoItemUpdateRequest) SetFoodSecurityStockDateEnd(foodSecurityStockDateEnd string) error {
    r.foodSecurityStockDateEnd = foodSecurityStockDateEnd
    r.Set("food_security.stock_date_end", foodSecurityStockDateEnd)
    return nil
}

func (r TaobaoItemUpdateRequest) GetFoodSecurityStockDateEnd() string {
    return r.foodSecurityStockDateEnd
}

func (r *TaobaoItemUpdateRequest) SetFoodSecurityHealthProductNo(foodSecurityHealthProductNo string) error {
    r.foodSecurityHealthProductNo = foodSecurityHealthProductNo
    r.Set("food_security.health_product_no", foodSecurityHealthProductNo)
    return nil
}

func (r TaobaoItemUpdateRequest) GetFoodSecurityHealthProductNo() string {
    return r.foodSecurityHealthProductNo
}

func (r *TaobaoItemUpdateRequest) SetLocalityLifeObs(localityLifeObs string) error {
    r.localityLifeObs = localityLifeObs
    r.Set("locality_life.obs", localityLifeObs)
    return nil
}

func (r TaobaoItemUpdateRequest) GetLocalityLifeObs() string {
    return r.localityLifeObs
}

func (r *TaobaoItemUpdateRequest) SetLocalityLifeVersion(localityLifeVersion string) error {
    r.localityLifeVersion = localityLifeVersion
    r.Set("locality_life.version", localityLifeVersion)
    return nil
}

func (r TaobaoItemUpdateRequest) GetLocalityLifeVersion() string {
    return r.localityLifeVersion
}

func (r *TaobaoItemUpdateRequest) SetLocalityLifePackageid(localityLifePackageid string) error {
    r.localityLifePackageid = localityLifePackageid
    r.Set("locality_life.packageid", localityLifePackageid)
    return nil
}

func (r TaobaoItemUpdateRequest) GetLocalityLifePackageid() string {
    return r.localityLifePackageid
}

func (r *TaobaoItemUpdateRequest) SetMsPaymentPrice(msPaymentPrice string) error {
    r.msPaymentPrice = msPaymentPrice
    r.Set("ms_payment.price", msPaymentPrice)
    return nil
}

func (r TaobaoItemUpdateRequest) GetMsPaymentPrice() string {
    return r.msPaymentPrice
}

func (r *TaobaoItemUpdateRequest) SetMsPaymentVoucherPrice(msPaymentVoucherPrice string) error {
    r.msPaymentVoucherPrice = msPaymentVoucherPrice
    r.Set("ms_payment.voucher_price", msPaymentVoucherPrice)
    return nil
}

func (r TaobaoItemUpdateRequest) GetMsPaymentVoucherPrice() string {
    return r.msPaymentVoucherPrice
}

func (r *TaobaoItemUpdateRequest) SetMsPaymentReferencePrice(msPaymentReferencePrice string) error {
    r.msPaymentReferencePrice = msPaymentReferencePrice
    r.Set("ms_payment.reference_price", msPaymentReferencePrice)
    return nil
}

func (r TaobaoItemUpdateRequest) GetMsPaymentReferencePrice() string {
    return r.msPaymentReferencePrice
}

func (r *TaobaoItemUpdateRequest) SetDeliveryTimeDeliveryTime(deliveryTimeDeliveryTime string) error {
    r.deliveryTimeDeliveryTime = deliveryTimeDeliveryTime
    r.Set("delivery_time.delivery_time", deliveryTimeDeliveryTime)
    return nil
}

func (r TaobaoItemUpdateRequest) GetDeliveryTimeDeliveryTime() string {
    return r.deliveryTimeDeliveryTime
}

func (r *TaobaoItemUpdateRequest) SetDeliveryTimeDeliveryTimeType(deliveryTimeDeliveryTimeType string) error {
    r.deliveryTimeDeliveryTimeType = deliveryTimeDeliveryTimeType
    r.Set("delivery_time.delivery_time_type", deliveryTimeDeliveryTimeType)
    return nil
}

func (r TaobaoItemUpdateRequest) GetDeliveryTimeDeliveryTimeType() string {
    return r.deliveryTimeDeliveryTimeType
}

func (r *TaobaoItemUpdateRequest) SetDeliveryTimeNeedDeliveryTime(deliveryTimeNeedDeliveryTime string) error {
    r.deliveryTimeNeedDeliveryTime = deliveryTimeNeedDeliveryTime
    r.Set("delivery_time.need_delivery_time", deliveryTimeNeedDeliveryTime)
    return nil
}

func (r TaobaoItemUpdateRequest) GetDeliveryTimeNeedDeliveryTime() string {
    return r.deliveryTimeNeedDeliveryTime
}

func (r *TaobaoItemUpdateRequest) SetNumIid(numIid int64) error {
    r.numIid = numIid
    r.Set("num_iid", numIid)
    return nil
}

func (r TaobaoItemUpdateRequest) GetNumIid() int64 {
    return r.numIid
}

func (r *TaobaoItemUpdateRequest) SetCid(cid int64) error {
    r.cid = cid
    r.Set("cid", cid)
    return nil
}

func (r TaobaoItemUpdateRequest) GetCid() int64 {
    return r.cid
}

func (r *TaobaoItemUpdateRequest) SetProps(props string) error {
    r.props = props
    r.Set("props", props)
    return nil
}

func (r TaobaoItemUpdateRequest) GetProps() string {
    return r.props
}

func (r *TaobaoItemUpdateRequest) SetNum(num int64) error {
    r.num = num
    r.Set("num", num)
    return nil
}

func (r TaobaoItemUpdateRequest) GetNum() int64 {
    return r.num
}

func (r *TaobaoItemUpdateRequest) SetPrice(price float64) error {
    r.price = price
    r.Set("price", price)
    return nil
}

func (r TaobaoItemUpdateRequest) GetPrice() float64 {
    return r.price
}

func (r *TaobaoItemUpdateRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

func (r TaobaoItemUpdateRequest) GetTitle() string {
    return r.title
}

func (r *TaobaoItemUpdateRequest) SetDesc(desc string) error {
    r.desc = desc
    r.Set("desc", desc)
    return nil
}

func (r TaobaoItemUpdateRequest) GetDesc() string {
    return r.desc
}

func (r *TaobaoItemUpdateRequest) SetPostFee(postFee float64) error {
    r.postFee = postFee
    r.Set("post_fee", postFee)
    return nil
}

func (r TaobaoItemUpdateRequest) GetPostFee() float64 {
    return r.postFee
}

func (r *TaobaoItemUpdateRequest) SetExpressFee(expressFee float64) error {
    r.expressFee = expressFee
    r.Set("express_fee", expressFee)
    return nil
}

func (r TaobaoItemUpdateRequest) GetExpressFee() float64 {
    return r.expressFee
}

func (r *TaobaoItemUpdateRequest) SetEmsFee(emsFee float64) error {
    r.emsFee = emsFee
    r.Set("ems_fee", emsFee)
    return nil
}

func (r TaobaoItemUpdateRequest) GetEmsFee() float64 {
    return r.emsFee
}

func (r *TaobaoItemUpdateRequest) SetListTime(listTime string) error {
    r.listTime = listTime
    r.Set("list_time", listTime)
    return nil
}

func (r TaobaoItemUpdateRequest) GetListTime() string {
    return r.listTime
}

func (r *TaobaoItemUpdateRequest) SetIncrement(increment float64) error {
    r.increment = increment
    r.Set("increment", increment)
    return nil
}

func (r TaobaoItemUpdateRequest) GetIncrement() float64 {
    return r.increment
}

func (r *TaobaoItemUpdateRequest) SetImage(image []*model.File) error {
    r.image = image
    r.Set("image", image)
    return nil
}

func (r TaobaoItemUpdateRequest) GetImage() []*model.File {
    return r.image
}

func (r *TaobaoItemUpdateRequest) SetStuffStatus(stuffStatus string) error {
    r.stuffStatus = stuffStatus
    r.Set("stuff_status", stuffStatus)
    return nil
}

func (r TaobaoItemUpdateRequest) GetStuffStatus() string {
    return r.stuffStatus
}

func (r *TaobaoItemUpdateRequest) SetAuctionPoint(auctionPoint int64) error {
    r.auctionPoint = auctionPoint
    r.Set("auction_point", auctionPoint)
    return nil
}

func (r TaobaoItemUpdateRequest) GetAuctionPoint() int64 {
    return r.auctionPoint
}

func (r *TaobaoItemUpdateRequest) SetPropertyAlias(propertyAlias string) error {
    r.propertyAlias = propertyAlias
    r.Set("property_alias", propertyAlias)
    return nil
}

func (r TaobaoItemUpdateRequest) GetPropertyAlias() string {
    return r.propertyAlias
}

func (r *TaobaoItemUpdateRequest) SetSellerCids(sellerCids string) error {
    r.sellerCids = sellerCids
    r.Set("seller_cids", sellerCids)
    return nil
}

func (r TaobaoItemUpdateRequest) GetSellerCids() string {
    return r.sellerCids
}

func (r *TaobaoItemUpdateRequest) SetPostageId(postageId int64) error {
    r.postageId = postageId
    r.Set("postage_id", postageId)
    return nil
}

func (r TaobaoItemUpdateRequest) GetPostageId() int64 {
    return r.postageId
}

func (r *TaobaoItemUpdateRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoItemUpdateRequest) GetOuterId() string {
    return r.outerId
}

func (r *TaobaoItemUpdateRequest) SetProductId(productId int64) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r TaobaoItemUpdateRequest) GetProductId() int64 {
    return r.productId
}

func (r *TaobaoItemUpdateRequest) SetPicPath(picPath string) error {
    r.picPath = picPath
    r.Set("pic_path", picPath)
    return nil
}

func (r TaobaoItemUpdateRequest) GetPicPath() string {
    return r.picPath
}

func (r *TaobaoItemUpdateRequest) SetAutoFill(autoFill string) error {
    r.autoFill = autoFill
    r.Set("auto_fill", autoFill)
    return nil
}

func (r TaobaoItemUpdateRequest) GetAutoFill() string {
    return r.autoFill
}

func (r *TaobaoItemUpdateRequest) SetIsTaobao(isTaobao bool) error {
    r.isTaobao = isTaobao
    r.Set("is_taobao", isTaobao)
    return nil
}

func (r TaobaoItemUpdateRequest) GetIsTaobao() bool {
    return r.isTaobao
}

func (r *TaobaoItemUpdateRequest) SetIsEx(isEx bool) error {
    r.isEx = isEx
    r.Set("is_ex", isEx)
    return nil
}

func (r TaobaoItemUpdateRequest) GetIsEx() bool {
    return r.isEx
}

func (r *TaobaoItemUpdateRequest) SetIs3D(is3D bool) error {
    r.is3D = is3D
    r.Set("is_3D", is3D)
    return nil
}

func (r TaobaoItemUpdateRequest) GetIs3D() bool {
    return r.is3D
}

func (r *TaobaoItemUpdateRequest) SetIsReplaceSku(isReplaceSku bool) error {
    r.isReplaceSku = isReplaceSku
    r.Set("is_replace_sku", isReplaceSku)
    return nil
}

func (r TaobaoItemUpdateRequest) GetIsReplaceSku() bool {
    return r.isReplaceSku
}

func (r *TaobaoItemUpdateRequest) SetLang(lang string) error {
    r.lang = lang
    r.Set("lang", lang)
    return nil
}

func (r TaobaoItemUpdateRequest) GetLang() string {
    return r.lang
}

func (r *TaobaoItemUpdateRequest) SetHasDiscount(hasDiscount bool) error {
    r.hasDiscount = hasDiscount
    r.Set("has_discount", hasDiscount)
    return nil
}

func (r TaobaoItemUpdateRequest) GetHasDiscount() bool {
    return r.hasDiscount
}

func (r *TaobaoItemUpdateRequest) SetHasShowcase(hasShowcase bool) error {
    r.hasShowcase = hasShowcase
    r.Set("has_showcase", hasShowcase)
    return nil
}

func (r TaobaoItemUpdateRequest) GetHasShowcase() bool {
    return r.hasShowcase
}

func (r *TaobaoItemUpdateRequest) SetApproveStatus(approveStatus string) error {
    r.approveStatus = approveStatus
    r.Set("approve_status", approveStatus)
    return nil
}

func (r TaobaoItemUpdateRequest) GetApproveStatus() string {
    return r.approveStatus
}

func (r *TaobaoItemUpdateRequest) SetFreightPayer(freightPayer string) error {
    r.freightPayer = freightPayer
    r.Set("freight_payer", freightPayer)
    return nil
}

func (r TaobaoItemUpdateRequest) GetFreightPayer() string {
    return r.freightPayer
}

func (r *TaobaoItemUpdateRequest) SetValidThru(validThru int64) error {
    r.validThru = validThru
    r.Set("valid_thru", validThru)
    return nil
}

func (r TaobaoItemUpdateRequest) GetValidThru() int64 {
    return r.validThru
}

func (r *TaobaoItemUpdateRequest) SetHasInvoice(hasInvoice bool) error {
    r.hasInvoice = hasInvoice
    r.Set("has_invoice", hasInvoice)
    return nil
}

func (r TaobaoItemUpdateRequest) GetHasInvoice() bool {
    return r.hasInvoice
}

func (r *TaobaoItemUpdateRequest) SetHasWarranty(hasWarranty bool) error {
    r.hasWarranty = hasWarranty
    r.Set("has_warranty", hasWarranty)
    return nil
}

func (r TaobaoItemUpdateRequest) GetHasWarranty() bool {
    return r.hasWarranty
}

func (r *TaobaoItemUpdateRequest) SetSellPromise(sellPromise bool) error {
    r.sellPromise = sellPromise
    r.Set("sell_promise", sellPromise)
    return nil
}

func (r TaobaoItemUpdateRequest) GetSellPromise() bool {
    return r.sellPromise
}

func (r *TaobaoItemUpdateRequest) SetCodPostageId(codPostageId int64) error {
    r.codPostageId = codPostageId
    r.Set("cod_postage_id", codPostageId)
    return nil
}

func (r TaobaoItemUpdateRequest) GetCodPostageId() int64 {
    return r.codPostageId
}

func (r *TaobaoItemUpdateRequest) SetIsLightningConsignment(isLightningConsignment bool) error {
    r.isLightningConsignment = isLightningConsignment
    r.Set("is_lightning_consignment", isLightningConsignment)
    return nil
}

func (r TaobaoItemUpdateRequest) GetIsLightningConsignment() bool {
    return r.isLightningConsignment
}

func (r *TaobaoItemUpdateRequest) SetWeight(weight int64) error {
    r.weight = weight
    r.Set("weight", weight)
    return nil
}

func (r TaobaoItemUpdateRequest) GetWeight() int64 {
    return r.weight
}

func (r *TaobaoItemUpdateRequest) SetIsXinpin(isXinpin bool) error {
    r.isXinpin = isXinpin
    r.Set("is_xinpin", isXinpin)
    return nil
}

func (r TaobaoItemUpdateRequest) GetIsXinpin() bool {
    return r.isXinpin
}

func (r *TaobaoItemUpdateRequest) SetSubStock(subStock int64) error {
    r.subStock = subStock
    r.Set("sub_stock", subStock)
    return nil
}

func (r TaobaoItemUpdateRequest) GetSubStock() int64 {
    return r.subStock
}

func (r *TaobaoItemUpdateRequest) SetItemSize(itemSize string) error {
    r.itemSize = itemSize
    r.Set("item_size", itemSize)
    return nil
}

func (r TaobaoItemUpdateRequest) GetItemSize() string {
    return r.itemSize
}

func (r *TaobaoItemUpdateRequest) SetItemWeight(itemWeight string) error {
    r.itemWeight = itemWeight
    r.Set("item_weight", itemWeight)
    return nil
}

func (r TaobaoItemUpdateRequest) GetItemWeight() string {
    return r.itemWeight
}

func (r *TaobaoItemUpdateRequest) SetSellPoint(sellPoint string) error {
    r.sellPoint = sellPoint
    r.Set("sell_point", sellPoint)
    return nil
}

func (r TaobaoItemUpdateRequest) GetSellPoint() string {
    return r.sellPoint
}

func (r *TaobaoItemUpdateRequest) SetBarcode(barcode string) error {
    r.barcode = barcode
    r.Set("barcode", barcode)
    return nil
}

func (r TaobaoItemUpdateRequest) GetBarcode() string {
    return r.barcode
}

func (r *TaobaoItemUpdateRequest) SetNewprepay(newprepay string) error {
    r.newprepay = newprepay
    r.Set("newprepay", newprepay)
    return nil
}

func (r TaobaoItemUpdateRequest) GetNewprepay() string {
    return r.newprepay
}

func (r *TaobaoItemUpdateRequest) SetQualification(qualification string) error {
    r.qualification = qualification
    r.Set("qualification", qualification)
    return nil
}

func (r TaobaoItemUpdateRequest) GetQualification() string {
    return r.qualification
}

func (r *TaobaoItemUpdateRequest) SetO2oBindService(o2oBindService bool) error {
    r.o2oBindService = o2oBindService
    r.Set("o2o_bind_service", o2oBindService)
    return nil
}

func (r TaobaoItemUpdateRequest) GetO2oBindService() bool {
    return r.o2oBindService
}

func (r *TaobaoItemUpdateRequest) SetIgnorewarning(ignorewarning string) error {
    r.ignorewarning = ignorewarning
    r.Set("ignorewarning", ignorewarning)
    return nil
}

func (r TaobaoItemUpdateRequest) GetIgnorewarning() string {
    return r.ignorewarning
}

func (r *TaobaoItemUpdateRequest) SetFeatures(features string) error {
    r.features = features
    r.Set("features", features)
    return nil
}

func (r TaobaoItemUpdateRequest) GetFeatures() string {
    return r.features
}

func (r *TaobaoItemUpdateRequest) SetAfterSaleId(afterSaleId int64) error {
    r.afterSaleId = afterSaleId
    r.Set("after_sale_id", afterSaleId)
    return nil
}

func (r TaobaoItemUpdateRequest) GetAfterSaleId() int64 {
    return r.afterSaleId
}

func (r *TaobaoItemUpdateRequest) SetChangeProp(changeProp string) error {
    r.changeProp = changeProp
    r.Set("change_prop", changeProp)
    return nil
}

func (r TaobaoItemUpdateRequest) GetChangeProp() string {
    return r.changeProp
}

func (r *TaobaoItemUpdateRequest) SetDescModules(descModules string) error {
    r.descModules = descModules
    r.Set("desc_modules", descModules)
    return nil
}

func (r TaobaoItemUpdateRequest) GetDescModules() string {
    return r.descModules
}

func (r *TaobaoItemUpdateRequest) SetIsOffline(isOffline string) error {
    r.isOffline = isOffline
    r.Set("is_offline", isOffline)
    return nil
}

func (r TaobaoItemUpdateRequest) GetIsOffline() string {
    return r.isOffline
}

func (r *TaobaoItemUpdateRequest) SetWirelessDesc(wirelessDesc string) error {
    r.wirelessDesc = wirelessDesc
    r.Set("wireless_desc", wirelessDesc)
    return nil
}

func (r TaobaoItemUpdateRequest) GetWirelessDesc() string {
    return r.wirelessDesc
}

func (r *TaobaoItemUpdateRequest) SetSpuConfirm(spuConfirm bool) error {
    r.spuConfirm = spuConfirm
    r.Set("spu_confirm", spuConfirm)
    return nil
}

func (r TaobaoItemUpdateRequest) GetSpuConfirm() bool {
    return r.spuConfirm
}

func (r *TaobaoItemUpdateRequest) SetInputPids(inputPids string) error {
    r.inputPids = inputPids
    r.Set("input_pids", inputPids)
    return nil
}

func (r TaobaoItemUpdateRequest) GetInputPids() string {
    return r.inputPids
}

func (r *TaobaoItemUpdateRequest) SetSkuQuantities(skuQuantities string) error {
    r.skuQuantities = skuQuantities
    r.Set("sku_quantities", skuQuantities)
    return nil
}

func (r TaobaoItemUpdateRequest) GetSkuQuantities() string {
    return r.skuQuantities
}

func (r *TaobaoItemUpdateRequest) SetSkuPrices(skuPrices string) error {
    r.skuPrices = skuPrices
    r.Set("sku_prices", skuPrices)
    return nil
}

func (r TaobaoItemUpdateRequest) GetSkuPrices() string {
    return r.skuPrices
}

func (r *TaobaoItemUpdateRequest) SetSkuProperties(skuProperties string) error {
    r.skuProperties = skuProperties
    r.Set("sku_properties", skuProperties)
    return nil
}

func (r TaobaoItemUpdateRequest) GetSkuProperties() string {
    return r.skuProperties
}

func (r *TaobaoItemUpdateRequest) SetSkuOuterIds(skuOuterIds string) error {
    r.skuOuterIds = skuOuterIds
    r.Set("sku_outer_ids", skuOuterIds)
    return nil
}

func (r TaobaoItemUpdateRequest) GetSkuOuterIds() string {
    return r.skuOuterIds
}

func (r *TaobaoItemUpdateRequest) SetInputStr(inputStr string) error {
    r.inputStr = inputStr
    r.Set("input_str", inputStr)
    return nil
}

func (r TaobaoItemUpdateRequest) GetInputStr() string {
    return r.inputStr
}

func (r *TaobaoItemUpdateRequest) SetSkuBarcode(skuBarcode string) error {
    r.skuBarcode = skuBarcode
    r.Set("sku_barcode", skuBarcode)
    return nil
}

func (r TaobaoItemUpdateRequest) GetSkuBarcode() string {
    return r.skuBarcode
}

func (r *TaobaoItemUpdateRequest) SetVideoId(videoId int64) error {
    r.videoId = videoId
    r.Set("video_id", videoId)
    return nil
}

func (r TaobaoItemUpdateRequest) GetVideoId() int64 {
    return r.videoId
}

func (r *TaobaoItemUpdateRequest) SetInteractiveId(interactiveId int64) error {
    r.interactiveId = interactiveId
    r.Set("interactive_id", interactiveId)
    return nil
}

func (r TaobaoItemUpdateRequest) GetInteractiveId() int64 {
    return r.interactiveId
}

func (r *TaobaoItemUpdateRequest) SetLeaseExtendsInfo(leaseExtendsInfo string) error {
    r.leaseExtendsInfo = leaseExtendsInfo
    r.Set("lease_extends_info", leaseExtendsInfo)
    return nil
}

func (r TaobaoItemUpdateRequest) GetLeaseExtendsInfo() string {
    return r.leaseExtendsInfo
}

func (r *TaobaoItemUpdateRequest) SetLocalityLifeExpirydate(localityLifeExpirydate string) error {
    r.localityLifeExpirydate = localityLifeExpirydate
    r.Set("locality_life.expirydate", localityLifeExpirydate)
    return nil
}

func (r TaobaoItemUpdateRequest) GetLocalityLifeExpirydate() string {
    return r.localityLifeExpirydate
}

func (r *TaobaoItemUpdateRequest) SetLocalityLifeNetworkId(localityLifeNetworkId string) error {
    r.localityLifeNetworkId = localityLifeNetworkId
    r.Set("locality_life.network_id", localityLifeNetworkId)
    return nil
}

func (r TaobaoItemUpdateRequest) GetLocalityLifeNetworkId() string {
    return r.localityLifeNetworkId
}

func (r *TaobaoItemUpdateRequest) SetLocalityLifeMerchant(localityLifeMerchant string) error {
    r.localityLifeMerchant = localityLifeMerchant
    r.Set("locality_life.merchant", localityLifeMerchant)
    return nil
}

func (r TaobaoItemUpdateRequest) GetLocalityLifeMerchant() string {
    return r.localityLifeMerchant
}

func (r *TaobaoItemUpdateRequest) SetLocalityLifeVerification(localityLifeVerification string) error {
    r.localityLifeVerification = localityLifeVerification
    r.Set("locality_life.verification", localityLifeVerification)
    return nil
}

func (r TaobaoItemUpdateRequest) GetLocalityLifeVerification() string {
    return r.localityLifeVerification
}

func (r *TaobaoItemUpdateRequest) SetLocalityLifeRefundRatio(localityLifeRefundRatio int64) error {
    r.localityLifeRefundRatio = localityLifeRefundRatio
    r.Set("locality_life.refund_ratio", localityLifeRefundRatio)
    return nil
}

func (r TaobaoItemUpdateRequest) GetLocalityLifeRefundRatio() int64 {
    return r.localityLifeRefundRatio
}

func (r *TaobaoItemUpdateRequest) SetLocalityLifeChooseLogis(localityLifeChooseLogis string) error {
    r.localityLifeChooseLogis = localityLifeChooseLogis
    r.Set("locality_life.choose_logis", localityLifeChooseLogis)
    return nil
}

func (r TaobaoItemUpdateRequest) GetLocalityLifeChooseLogis() string {
    return r.localityLifeChooseLogis
}

func (r *TaobaoItemUpdateRequest) SetLocalityLifeOnsaleAutoRefundRatio(localityLifeOnsaleAutoRefundRatio int64) error {
    r.localityLifeOnsaleAutoRefundRatio = localityLifeOnsaleAutoRefundRatio
    r.Set("locality_life.onsale_auto_refund_ratio", localityLifeOnsaleAutoRefundRatio)
    return nil
}

func (r TaobaoItemUpdateRequest) GetLocalityLifeOnsaleAutoRefundRatio() int64 {
    return r.localityLifeOnsaleAutoRefundRatio
}

func (r *TaobaoItemUpdateRequest) SetLocalityLifeRefundmafee(localityLifeRefundmafee string) error {
    r.localityLifeRefundmafee = localityLifeRefundmafee
    r.Set("locality_life.refundmafee", localityLifeRefundmafee)
    return nil
}

func (r TaobaoItemUpdateRequest) GetLocalityLifeRefundmafee() string {
    return r.localityLifeRefundmafee
}

func (r *TaobaoItemUpdateRequest) SetLocalityLifeEticket(localityLifeEticket string) error {
    r.localityLifeEticket = localityLifeEticket
    r.Set("locality_life.eticket", localityLifeEticket)
    return nil
}

func (r TaobaoItemUpdateRequest) GetLocalityLifeEticket() string {
    return r.localityLifeEticket
}

func (r *TaobaoItemUpdateRequest) SetEmptyFields(emptyFields string) error {
    r.emptyFields = emptyFields
    r.Set("empty_fields", emptyFields)
    return nil
}

func (r TaobaoItemUpdateRequest) GetEmptyFields() string {
    return r.emptyFields
}

func (r *TaobaoItemUpdateRequest) SetScenicTicketPayWay(scenicTicketPayWay int64) error {
    r.scenicTicketPayWay = scenicTicketPayWay
    r.Set("scenic_ticket_pay_way", scenicTicketPayWay)
    return nil
}

func (r TaobaoItemUpdateRequest) GetScenicTicketPayWay() int64 {
    return r.scenicTicketPayWay
}

func (r *TaobaoItemUpdateRequest) SetScenicTicketBookCost(scenicTicketBookCost string) error {
    r.scenicTicketBookCost = scenicTicketBookCost
    r.Set("scenic_ticket_book_cost", scenicTicketBookCost)
    return nil
}

func (r TaobaoItemUpdateRequest) GetScenicTicketBookCost() string {
    return r.scenicTicketBookCost
}

func (r *TaobaoItemUpdateRequest) SetPaimaiInfoMode(paimaiInfoMode int64) error {
    r.paimaiInfoMode = paimaiInfoMode
    r.Set("paimai_info.mode", paimaiInfoMode)
    return nil
}

func (r TaobaoItemUpdateRequest) GetPaimaiInfoMode() int64 {
    return r.paimaiInfoMode
}

func (r *TaobaoItemUpdateRequest) SetPaimaiInfoDeposit(paimaiInfoDeposit int64) error {
    r.paimaiInfoDeposit = paimaiInfoDeposit
    r.Set("paimai_info.deposit", paimaiInfoDeposit)
    return nil
}

func (r TaobaoItemUpdateRequest) GetPaimaiInfoDeposit() int64 {
    return r.paimaiInfoDeposit
}

func (r *TaobaoItemUpdateRequest) SetPaimaiInfoInterval(paimaiInfoInterval int64) error {
    r.paimaiInfoInterval = paimaiInfoInterval
    r.Set("paimai_info.interval", paimaiInfoInterval)
    return nil
}

func (r TaobaoItemUpdateRequest) GetPaimaiInfoInterval() int64 {
    return r.paimaiInfoInterval
}

func (r *TaobaoItemUpdateRequest) SetPaimaiInfoReserve(paimaiInfoReserve float64) error {
    r.paimaiInfoReserve = paimaiInfoReserve
    r.Set("paimai_info.reserve", paimaiInfoReserve)
    return nil
}

func (r TaobaoItemUpdateRequest) GetPaimaiInfoReserve() float64 {
    return r.paimaiInfoReserve
}

func (r *TaobaoItemUpdateRequest) SetPaimaiInfoValidHour(paimaiInfoValidHour int64) error {
    r.paimaiInfoValidHour = paimaiInfoValidHour
    r.Set("paimai_info.valid_hour", paimaiInfoValidHour)
    return nil
}

func (r TaobaoItemUpdateRequest) GetPaimaiInfoValidHour() int64 {
    return r.paimaiInfoValidHour
}

func (r *TaobaoItemUpdateRequest) SetPaimaiInfoValidMinute(paimaiInfoValidMinute int64) error {
    r.paimaiInfoValidMinute = paimaiInfoValidMinute
    r.Set("paimai_info.valid_minute", paimaiInfoValidMinute)
    return nil
}

func (r TaobaoItemUpdateRequest) GetPaimaiInfoValidMinute() int64 {
    return r.paimaiInfoValidMinute
}

func (r *TaobaoItemUpdateRequest) SetGlobalStockType(globalStockType string) error {
    r.globalStockType = globalStockType
    r.Set("global_stock_type", globalStockType)
    return nil
}

func (r TaobaoItemUpdateRequest) GetGlobalStockType() string {
    return r.globalStockType
}

func (r *TaobaoItemUpdateRequest) SetGlobalStockCountry(globalStockCountry string) error {
    r.globalStockCountry = globalStockCountry
    r.Set("global_stock_country", globalStockCountry)
    return nil
}

func (r TaobaoItemUpdateRequest) GetGlobalStockCountry() string {
    return r.globalStockCountry
}

func (r *TaobaoItemUpdateRequest) SetGlobalStockDeliveryPlace(globalStockDeliveryPlace string) error {
    r.globalStockDeliveryPlace = globalStockDeliveryPlace
    r.Set("global_stock_delivery_place", globalStockDeliveryPlace)
    return nil
}

func (r TaobaoItemUpdateRequest) GetGlobalStockDeliveryPlace() string {
    return r.globalStockDeliveryPlace
}

func (r *TaobaoItemUpdateRequest) SetGlobalStockTaxFreePromise(globalStockTaxFreePromise bool) error {
    r.globalStockTaxFreePromise = globalStockTaxFreePromise
    r.Set("global_stock_tax_free_promise", globalStockTaxFreePromise)
    return nil
}

func (r TaobaoItemUpdateRequest) GetGlobalStockTaxFreePromise() bool {
    return r.globalStockTaxFreePromise
}

