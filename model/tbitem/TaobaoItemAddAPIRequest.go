package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemAddAPIRequest 添加一个商品 API请求
// taobao.item.add
//
// 此接口用于新增一个淘宝商品
// 商品的属性和sku的属性有包含的关系，商品的价格要位于sku的价格区间之中（例如，sku价格有5元、10元两种，那么商品的价格就需要大于等于5元，小于等于10元，否则新增商品会失败）
// 商品的类目和商品的价格、sku的价格都有一定的相关性（具体的关系要通过类目属性查询接口获得）
// 商品的运费承担方式和邮费设置有相关性，卖家承担运费不用设置邮费，买家承担运费需要设置邮费
// 当关键属性值选择了“其他”的时候，需要输入input_pids和input_str商品才能添加成功。
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=itemapi&#34; target=&#34;_blank&#34;&gt;点击查看更多商品API说明&lt;/a&gt;&lt;/strong&gt;
type TaobaoItemAddAPIRequest struct {
	model.Params
	// 此字段已经废弃，不再使用
	_imageUrls []string
	// 商品上传后的状态。可选值:onsale(出售中),instock(仓库中);默认值:onsale
	_approveStatus string
	// 发布类型。可选值:fixed(一口价),auction(拍卖)。B商家不能发布拍卖商品，而且拍卖商品是没有SKU的。拍卖商品发布时需要附加拍卖商品信息：拍卖类型(paimai_info.mode，拍卖类型包括三种：增价拍[1]，荷兰拍[2]以及降价拍[3])，商品数量(num)，起拍价(price)，价格幅度(increament)，保证金(paimai_info.deposit)。另外拍卖商品支持自定义销售周期，通过paimai_info.valid_hour和paimai_info.valid_minute来指定。对于降价拍来说需要设置降价周期(paimai_info.interval)和拍卖保留价(paimai_info.reserve)。注意：通过taobao.item.get接口获取拍卖信息时，会返回除了valid_hour和valid_minute之外的所有拍卖信息。
	_type string
	// 宝贝描述。字数要大于5个字符，小于25000个字符，受违禁词控制
	_desc string
	// 商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值
	_props string
	// 运费承担方式。可选值:seller（卖家承担）,buyer(买家承担);默认值:seller。卖家承担不用设置邮费和postage_id.买家承担的时候，必填邮费和postage_id 如果用户设置了运费模板会优先使用运费模板，否则要同步设置邮费（post_fee,express_fee,ems_fee）
	_freightPayer string
	// 商品所属的店铺类目列表。按逗号分隔。结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。
	_sellerCids string
	// 定时上架时间。(时间格式：yyyy-MM-dd HH:mm:ss)
	_listTime string
	// 新旧程度。可选值：new(新)，second(二手)。B商家不能发布二手商品。如果是二手商品，特定类目下属性里面必填新旧成色属性
	_stuffStatus string
	// 宝贝标题。不能超过30字符，受违禁词控制。天猫图书管控类目最大允许120字符；
	_title string
	// 商品外部编码，该字段的最大长度是64个字节
	_outerId string
	// 属性值别名。如pid:vid:别名;pid1:vid1:别名1 ，其中：pid是属性id vid是属性值id。总长度不超过800个字符，如"123:333:你好"，引号内的是10个字符。
	_propertyAlias string
	// 商品文字的字符集。繁体传入"zh_HK"，简体传入"zh_CN"，不传默认为简体
	_lang string
	// （推荐）商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path
	_picPath string
	// 所在地省份。如浙江
	_locationState string
	// 所在地城市。如杭州 。
	_locationCity string
	// 代充商品类型。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型： no_mark(不做类型标记) time_card(点卡软件代充) fee_card(话费软件代充)
	_autoFill string
	// 宝贝特征值，格式为：【key1:value1;key2:value2;key3:value3;】，key和value用【:】分隔，key&value之间用【;】分隔，只有在Top支持的特征值才能保存到宝贝上，目前支持的Key列表为：mysize_tp
	_features string
	// 生产许可证号
	_foodSecurityPrdLicenseNo string
	// 产品标准号
	_foodSecurityDesignCode string
	// 厂名
	_foodSecurityFactory string
	// 厂址
	_foodSecurityFactorySite string
	// 厂家联系方式
	_foodSecurityContact string
	// 配料表
	_foodSecurityMix string
	// 储藏方法
	_foodSecurityPlanStorage string
	// 保质期，默认有单位，传入数字
	_foodSecurityPeriod string
	// 食品添加剂
	_foodSecurityFoodAdditive string
	// 供货商
	_foodSecuritySupplier string
	// 生产开始日期，格式必须为yyyy-MM-dd
	_foodSecurityProductDateStart string
	// 生产结束日期,格式必须为yyyy-MM-dd
	_foodSecurityProductDateEnd string
	// 进货开始日期，要在生产日期之后，格式必须为yyyy-MM-dd
	_foodSecurityStockDateStart string
	// 进货结束日期，要在生产日期之后，格式必须为yyyy-MM-dd
	_foodSecurityStockDateEnd string
	// 健字号，保健品/膳食营养补充剂 这个类目下特有的信息，此类目下无需填写生产许可证编号（QS），如果填写了生产许可证编号（QS）将被忽略不保存；保存宝贝时，标题前会自动加上健字号产品名称一起作为宝贝标题；
	_foodSecurityHealthProductNo string
	// 景区门票在选择订金支付时候，需要交的预订费。传入的值是1到20之间的数值，小数点后最多可以保留两位（多余的部分将做四舍五入的处理）。这个数值表示的是预订费的比例，最终的预订费为 scenic_ticket_book_cost乘一口价除以100
	_scenicTicketBookCost string
	// 全球购商品采购地（库存类型），有两种库存类型：现货和代购参数值为1时代表现货，值为2时代表代购。注意：使用时请与 全球购商品采购地（地区/国家）配合使用
	_globalStockType string
	// 全球购商品采购地（地区/国家）,默认值只在全球购商品采购地（库存类型选择情况生效），地区国家值请填写法定的国家名称，类如（美国, 香港, 日本, 英国, 新西兰, 德国, 韩国, 荷兰, 法国, 意大利, 台湾, 澳门, 加拿大, 瑞士, 西班牙, 泰国, 新加坡, 马来西亚, 菲律宾），不要使用其他
	_globalStockCountry string
	// 全球购商品发货地，发货地现在有两种类型：“国内”和“海外及港澳台”，参数值为1时代表“国内”，值为2时代表“海外及港澳台”，默认为国内。注意：卖家必须已经签署并启用“海外直邮”合约，才能选择发货地为“海外及港澳台”
	_globalStockDeliveryPlace string
	// 本地生活电子交易凭证业务，目前此字段只涉及到的信息为有效期;如果有效期为起止日期类型，此值为2012-08-06,2012-08-16如果有效期为【购买成功日 至】类型则格式为2012-08-16如果有效期为天数类型则格式为15
	_localityLifeExpirydate string
	// 网点ID
	_localityLifeNetworkId string
	// 码商信息，格式为 码商id:nick
	_localityLifeMerchant string
	// 核销打款 1代表核销打款 0代表非核销打款
	_localityLifeVerification string
	// 发布电子凭证宝贝时候表示是否使用邮寄 0: 代表不使用邮寄； 1：代表使用邮寄；如果不设置这个值，代表不使用邮寄
	_localityLifeChooseLogis string
	// 退款码费承担方。发布电子凭证宝贝的时候会增加“退款码费承担方”配置项，可选填：(1)s（卖家承担） (2)b(买家承担)
	_localityLifeRefundmafee string
	// 预约门店是否支持门店自提,1:是
	_localityLifeObs string
	// 电子凭证业务属性，数据字典是: 1、is_card:1 (暂时不用) 2、consume_way:4 （1 串码 ，4 身份证）3、consume_midmnick ：(核销放行账号:用户id-用户名，支持多个，用逗号分隔,例如 1234-测试账号35,1345-测试账号56）4、market:eticket (电子凭证商品标记) 5、has_pos:1 (1 表示商品配置线下门店，在detail上进行展示 ，没有或者其他值只不展示)格式是: k1:v2;k2:v2;........ 如：has_pos:1;market:eticket;consume_midmnick:901409638-OPPO;consume_way:4
	_localityLifeEticket string
	// 新版电子凭证字段
	_localityLifeVersion string
	// 新版电子凭证包 id
	_localityLifePackageid string
	// 商品的重量，用于按重量计费的运费模板。注意：单位为kg。只能传入数值类型（包含小数），不能带单位，单位默认为kg。
	_itemWeight string
	// 表示商品的体积，如果需要使用按体积计费的运费模板，一定要设置这个值。该值的单位为立方米（m3），如果是其他单位，请转换成成立方米。该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）。体积和长宽高都支持小数类型。在传入体积或长宽高时候，不能带单位。体积的单位默认为立方米（m3），长宽高的单位默认为米(m)该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）
	_itemSize string
	// 针对当前商品的自定义属性值，目前是针对销售属性值自定义的，所以调用方需要把自定义属性值对应的虚拟属性值ID（负整数，例如例子中的 -1和-2）像标准属性值值的id一样设置到SKU上，如果自定义属性值有属性值图片，也要设置到属性图片上
	_inputCustomCpv string
	// 针对当前商品的标准属性值的补充说明，让买家更加了解商品信息减少交易纠纷
	_cpvMemo string
	// 定制工具Id如果支持定制市场，这个值不填写，就用之前的定制工具Id，之前的定制工具Id没有值就默认为-1
	_customMadeTypeId string
	// 该宝贝是否支持【7天无理由退货】，卖家选择的值只是一个因素，最终以类目和选择的属性条件来确定是否支持7天。填入字符0，表示不支持；未填写或填人字符1，表示支持7天无理由退货；
	_newprepay string
	// 商品条形码
	_barcode string
	// 商品卖点信息，最长150个字符。天猫商家和集市卖家都可用。
	_sellPoint string
	// 商品资质信息
	_qualification string
	// 忽略警告提示.
	_ignorewarning string
	// 参考价。该商品订单首次支付价格为 订金 价格，用户可根据 参考价 估算全款。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm
	_msPaymentReferencePrice string
	// 尾款可抵扣金额。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm
	_msPaymentVoucherPrice string
	// 订金。在“线上付订金线下付尾款”模式中，有订金、尾款可抵扣金额和参考价，三者需要同时填写。该商品订单首次支付价格为 订金 价格，用户可根据 参考价 估算全款。该模式有别于“一口价”付款方式，针对一个商品，只能选择两种付款方式中的一种，其适用于家装、二手车等场景。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm
	_msPaymentPrice string
	// 设置是否使用发货时间，商品级别，sku级别
	_deliveryTimeNeedDeliveryTime string
	// 发货时间类型：绝对发货时间或者相对发货时间
	_deliveryTimeDeliveryTimeType string
	// 商品级别设置的发货时间。设置了商品级别的发货时间，相对发货时间，则填写相对发货时间的天数（大于3）；绝对发货时间，则填写yyyy-mm-dd格式，如2013-11-11
	_deliveryTimeDeliveryTime string
	// 基础色数据，淘宝不使用
	_changeProp string
	// 已废弃
	_descModules string
	// 是否是线下商品。1：线上商品（默认值）；2：线上或线下商品；3：线下商品。
	_isOffline string
	// 无线的宝贝描述
	_wirelessDesc string
	// 租赁扩展信息
	_leaseExtendsInfo string
	// 仅淘小铺卖家需要。佣金比例(15.3对应的佣金比例为15.3%).只支持小数点后1位。多余的位数四舍五入(15.32会保存为15.3%
	_brokerage string
	// 业务身份编码。淘小铺编码为"taobao-taoxiaopu"
	_bizCode string
	// 用户自行输入的类目属性ID串，结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。
	_inputPids string
	// 用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5,Nike乔丹鞋;乔丹系列;乔丹鞋系列;乔丹鞋系列;json5”，多个自定义属性用','分割，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过3999字节。此处不可以使用“其他”、“其它”和“其她”这三个词。
	_inputStr string
	// 更新的sku的属性串，调用taobao.itemprops.get获取。格式:pid1:vid;pid2:vid,多个sku属性之间用逗号分隔。该字段内的属性需要在props字段同时包含。如果新增商品包含了sku，则此字段一定要传入,字段长度要控制在512个字节以内。
	_skuProperties string
	// Sku的数量串，结构如：num1,num2,num3 如：2,3
	_skuQuantities string
	// Sku的价格串，结构如：10.00,5.00,… 精确到2位小数;单位:元。如:200.07，表示:200元7分
	_skuPrices string
	// Sku的外部id串，结构如：1234,1342,… sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，如果没有sku_outer_ids也要写上这个参数，入参是","(这个是两个sku的示列，逗号数应该是sku个数减1)；该参数最大长度是512个字节
	_skuOuterIds string
	// sku层面的条形码，多个SKU情况，与SKU价格库存格式类似，用逗号分隔
	_skuBarcode string
	// 此参数暂时不起作用
	_skuSpecIds string
	// 此参数暂时不起作用
	_skuDeliveryTimes string
	// 家装建材类目，商品SKU的长度，正整数，单位为cm，部分类目必选。 数据和SKU一一对应，用,分隔，如：20,30,30
	_skuHdLength string
	// 家装建材类目，商品SKU的高度，单位为cm，部分类目必选。 天猫和淘宝格式不同。天猫：可选值为："0-15", "15-25", "25-50", "50-60", "60-80", "80-120", "120-160", "160-200"。 数据和SKU一一对应，用,分隔，格式如：15-25,25-50,25-50。 淘宝：正整数，单位为cm,格式如：20,30,30
	_skuHdHeight string
	// 家装建材类目，商品SKU的灯头数量，正整数，大于等于3，部分类目必选。天猫商家专用。 数据和SKU一一对应，用,分隔，如：3,5,7
	_skuHdLampQuantity string
	// 叶子类目id
	_cid int64
	// 有效期。可选值:7,14;单位:天;默认值:14
	_validThru int64
	// 平邮费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:5.07，表示:5元7分. 注:post_fee,express_fee,ems_fee需要一起填写
	_postFee float64
	// 快递费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:15.07，表示:15元7分
	_expressFee float64
	// ems费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:25.07，表示:25元7分
	_emsFee float64
	// 加价(降价)幅度。如果为0，代表系统代理幅度。对于增价拍和荷兰拍来说是加价幅度，对于降价拍来说是降价幅度。
	_increment float64
	// 商品数量。取值范围:0-900000000的整数。且需要等于Sku所有数量的和。拍卖商品中增加拍只能为1，荷兰拍要在[2,500)范围内。
	_num int64
	// 商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。拍卖商品对应的起拍价。
	_price float64
	// 宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.delivery.template.get获得当前会话用户的所有邮费模板）
	_postageId int64
	// 商品所属的产品ID(B商家发布商品需要用)
	_productId int64
	// 商品的积分返点比例。如:5,表示:返点比例0.5%. 注意：返点比例必须是>0的整数，而且最大是90,即为9%.B商家在发布非虚拟商品时，返点必须是 5的倍数，即0.5%的倍数。其它是1的倍数，即0.1%的倍数。无名良品商家发布商品时，复用该字段记录积分宝返点比例，返点必须是对应类目的返点步长的整数倍，默认是5，即0.5%。注意此时该字段值依旧必须是>0的整数，最高值不超过500，即50%
	_auctionPoint int64
	// 商品主图片。类型:JPG,GIF;最大长度:3M。（推荐使用pic_path字段，先把图片上传到卖家图片空间）
	_image *model.File
	// 此为货到付款运费模板的ID，对应的JAVA类型是long,如果COD卖家应用了货到付款运费模板，此值要进行设置。该字段已经废弃
	_codPostageId int64
	// 商品的重量(商超卖家专用字段)
	_weight int64
	// 商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改集市卖家默认拍下减库存;商城卖家默认付款减库存
	_subStock int64
	// 景区门票类宝贝发布时候，当卖家签订了支付宝代扣协议时候，需要选择支付方式：全额支付和订金支付。当scenic_ticket_pay_way为1时表示全额支付，为2时表示订金支付
	_scenicTicketPayWay int64
	// 退款比例，百分比%前的数字,1-100的正整数值
	_localityLifeRefundRatio int64
	// 电子凭证售中自动退款比例，百分比%前的数字，介于1-100之间的整数
	_localityLifeOnsaleAutoRefundRatio int64
	// 拍卖商品选择的拍卖类型，拍卖类型包括三种：增价拍(1)，荷兰拍(2)和降价拍(3)。
	_paimaiInfoMode int64
	// 拍卖宝贝的保证金。对于增价拍和荷兰拍来说保证金有两种模式：淘宝默认模式（首次出价金额的10%），自定义固定保证金（固定冻结金额只能输入不超过30万的正整数），并且保证金只冻结1次。对于降价拍来说保证金只有淘宝默认的（竞拍价格的10% * 竞拍数量），并且每次出价都需要冻结保证金。对于拍卖宝贝来说，保证金是必须的，但是默认使用淘宝默认保证金模式，只有用户需要使用自定义固定保证金的时候才需要使用到这个参数，如果该参数不传或传入0则代表使用默认。
	_paimaiInfoDeposit int64
	// 降价拍宝贝的降价周期(分钟)。降价拍宝贝的价格每隔paimai_info.interval时间会下降一次increment。
	_paimaiInfoInterval int64
	// 降价拍宝贝的保留价。对于降价拍来说，paimai_info.reserve必须大于0，且小于price-increment，而且（price-paimai_info.reserve）/increment的计算结果必须为整数
	_paimaiInfoReserve float64
	// 自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。
	_paimaiInfoValidHour int64
	// 自定义销售周期的分钟数。拍卖宝贝可以自定义销售周期，这里是指定销售周期的分钟数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。
	_paimaiInfoValidMinute int64
	// 售后说明模板id
	_afterSaleId int64
	// 主图视频id
	_videoId int64
	// 主图视频互动信息id，必须填写主图视频id才能有互动信息id
	_interactiveId int64
	// 是否有保修。可选值:true,false;默认值:false(不保修)
	_hasWarranty bool
	// 是否有发票。可选值:true,false (商城卖家此字段必须为true);默认值:false(无发票)
	_hasInvoice bool
	// 橱窗推荐。可选值:true,false;默认值:false(不推荐)
	_hasShowcase bool
	// 支持会员打折。可选值:true,false;默认值:false(不打折)
	_hasDiscount bool
	// 是否在淘宝上显示（如果传FALSE，则在淘宝主站无法显示该商品）
	_isTaobao bool
	// 是否在外店显示
	_isEx bool
	// 是否是3D
	_is3D bool
	// 是否承诺退换货服务!虚拟商品无须设置此项!
	_sellPromise bool
	// 实物闪电发货
	_isLightningConsignment bool
	// 商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:add-xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置该参数值或设置为false效果一致。
	_isXinpin bool
	// 全球购商品卖家包税承诺，当值为true时，代表卖家承诺包税。注意：请与“全球购商品发货地”配合使用，包税承诺必须满足：1、发货地为海外及港澳台 2、卖家已经签署并启用“包税合约”合约
	_globalStockTaxFreePromise bool
	// 是否支持定制市场 true代表支持，false代表支持,如果为空代表与之前保持不变不会修改
	_supportCustomMade bool
	// 汽车O2O绑定线下服务标记，如不为空，表示关联服务，否则，不关联服务。
	_o2oBindService bool
	// 手机类目spu 优化，信息确认字段
	_spuConfirm bool
}

// NewTaobaoItemAddRequest 初始化TaobaoItemAddAPIRequest对象
func NewTaobaoItemAddRequest() *TaobaoItemAddAPIRequest {
	return &TaobaoItemAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemAddAPIRequest) GetApiMethodName() string {
	return "taobao.item.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoItemAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetImageUrls is ImageUrls Setter
// 此字段已经废弃，不再使用
func (r *TaobaoItemAddAPIRequest) SetImageUrls(_imageUrls []string) error {
	r._imageUrls = _imageUrls
	r.Set("image_urls", _imageUrls)
	return nil
}

// GetImageUrls ImageUrls Getter
func (r TaobaoItemAddAPIRequest) GetImageUrls() []string {
	return r._imageUrls
}

// SetApproveStatus is ApproveStatus Setter
// 商品上传后的状态。可选值:onsale(出售中),instock(仓库中);默认值:onsale
func (r *TaobaoItemAddAPIRequest) SetApproveStatus(_approveStatus string) error {
	r._approveStatus = _approveStatus
	r.Set("approve_status", _approveStatus)
	return nil
}

// GetApproveStatus ApproveStatus Getter
func (r TaobaoItemAddAPIRequest) GetApproveStatus() string {
	return r._approveStatus
}

// SetType is Type Setter
// 发布类型。可选值:fixed(一口价),auction(拍卖)。B商家不能发布拍卖商品，而且拍卖商品是没有SKU的。拍卖商品发布时需要附加拍卖商品信息：拍卖类型(paimai_info.mode，拍卖类型包括三种：增价拍[1]，荷兰拍[2]以及降价拍[3])，商品数量(num)，起拍价(price)，价格幅度(increament)，保证金(paimai_info.deposit)。另外拍卖商品支持自定义销售周期，通过paimai_info.valid_hour和paimai_info.valid_minute来指定。对于降价拍来说需要设置降价周期(paimai_info.interval)和拍卖保留价(paimai_info.reserve)。注意：通过taobao.item.get接口获取拍卖信息时，会返回除了valid_hour和valid_minute之外的所有拍卖信息。
func (r *TaobaoItemAddAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoItemAddAPIRequest) GetType() string {
	return r._type
}

// SetDesc is Desc Setter
// 宝贝描述。字数要大于5个字符，小于25000个字符，受违禁词控制
func (r *TaobaoItemAddAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r TaobaoItemAddAPIRequest) GetDesc() string {
	return r._desc
}

// SetProps is Props Setter
// 商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值
func (r *TaobaoItemAddAPIRequest) SetProps(_props string) error {
	r._props = _props
	r.Set("props", _props)
	return nil
}

// GetProps Props Getter
func (r TaobaoItemAddAPIRequest) GetProps() string {
	return r._props
}

// SetFreightPayer is FreightPayer Setter
// 运费承担方式。可选值:seller（卖家承担）,buyer(买家承担);默认值:seller。卖家承担不用设置邮费和postage_id.买家承担的时候，必填邮费和postage_id 如果用户设置了运费模板会优先使用运费模板，否则要同步设置邮费（post_fee,express_fee,ems_fee）
func (r *TaobaoItemAddAPIRequest) SetFreightPayer(_freightPayer string) error {
	r._freightPayer = _freightPayer
	r.Set("freight_payer", _freightPayer)
	return nil
}

// GetFreightPayer FreightPayer Getter
func (r TaobaoItemAddAPIRequest) GetFreightPayer() string {
	return r._freightPayer
}

// SetSellerCids is SellerCids Setter
// 商品所属的店铺类目列表。按逗号分隔。结构:&#34;,cid1,cid2,...,&#34;，如果店铺类目存在二级类目，必须传入子类目cids。
func (r *TaobaoItemAddAPIRequest) SetSellerCids(_sellerCids string) error {
	r._sellerCids = _sellerCids
	r.Set("seller_cids", _sellerCids)
	return nil
}

// GetSellerCids SellerCids Getter
func (r TaobaoItemAddAPIRequest) GetSellerCids() string {
	return r._sellerCids
}

// SetListTime is ListTime Setter
// 定时上架时间。(时间格式：yyyy-MM-dd HH:mm:ss)
func (r *TaobaoItemAddAPIRequest) SetListTime(_listTime string) error {
	r._listTime = _listTime
	r.Set("list_time", _listTime)
	return nil
}

// GetListTime ListTime Getter
func (r TaobaoItemAddAPIRequest) GetListTime() string {
	return r._listTime
}

// SetStuffStatus is StuffStatus Setter
// 新旧程度。可选值：new(新)，second(二手)。B商家不能发布二手商品。如果是二手商品，特定类目下属性里面必填新旧成色属性
func (r *TaobaoItemAddAPIRequest) SetStuffStatus(_stuffStatus string) error {
	r._stuffStatus = _stuffStatus
	r.Set("stuff_status", _stuffStatus)
	return nil
}

// GetStuffStatus StuffStatus Getter
func (r TaobaoItemAddAPIRequest) GetStuffStatus() string {
	return r._stuffStatus
}

// SetTitle is Title Setter
// 宝贝标题。不能超过30字符，受违禁词控制。天猫图书管控类目最大允许120字符；
func (r *TaobaoItemAddAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaoItemAddAPIRequest) GetTitle() string {
	return r._title
}

// SetOuterId is OuterId Setter
// 商品外部编码，该字段的最大长度是64个字节
func (r *TaobaoItemAddAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoItemAddAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetPropertyAlias is PropertyAlias Setter
// 属性值别名。如pid:vid:别名;pid1:vid1:别名1 ，其中：pid是属性id vid是属性值id。总长度不超过800个字符，如&#34;123:333:你好&#34;，引号内的是10个字符。
func (r *TaobaoItemAddAPIRequest) SetPropertyAlias(_propertyAlias string) error {
	r._propertyAlias = _propertyAlias
	r.Set("property_alias", _propertyAlias)
	return nil
}

// GetPropertyAlias PropertyAlias Getter
func (r TaobaoItemAddAPIRequest) GetPropertyAlias() string {
	return r._propertyAlias
}

// SetLang is Lang Setter
// 商品文字的字符集。繁体传入&#34;zh_HK&#34;，简体传入&#34;zh_CN&#34;，不传默认为简体
func (r *TaobaoItemAddAPIRequest) SetLang(_lang string) error {
	r._lang = _lang
	r.Set("lang", _lang)
	return nil
}

// GetLang Lang Getter
func (r TaobaoItemAddAPIRequest) GetLang() string {
	return r._lang
}

// SetPicPath is PicPath Setter
// （推荐）商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path
func (r *TaobaoItemAddAPIRequest) SetPicPath(_picPath string) error {
	r._picPath = _picPath
	r.Set("pic_path", _picPath)
	return nil
}

// GetPicPath PicPath Getter
func (r TaobaoItemAddAPIRequest) GetPicPath() string {
	return r._picPath
}

// SetLocationState is LocationState Setter
// 所在地省份。如浙江
func (r *TaobaoItemAddAPIRequest) SetLocationState(_locationState string) error {
	r._locationState = _locationState
	r.Set("location.state", _locationState)
	return nil
}

// GetLocationState LocationState Getter
func (r TaobaoItemAddAPIRequest) GetLocationState() string {
	return r._locationState
}

// SetLocationCity is LocationCity Setter
// 所在地城市。如杭州 。
func (r *TaobaoItemAddAPIRequest) SetLocationCity(_locationCity string) error {
	r._locationCity = _locationCity
	r.Set("location.city", _locationCity)
	return nil
}

// GetLocationCity LocationCity Getter
func (r TaobaoItemAddAPIRequest) GetLocationCity() string {
	return r._locationCity
}

// SetAutoFill is AutoFill Setter
// 代充商品类型。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型： no_mark(不做类型标记) time_card(点卡软件代充) fee_card(话费软件代充)
func (r *TaobaoItemAddAPIRequest) SetAutoFill(_autoFill string) error {
	r._autoFill = _autoFill
	r.Set("auto_fill", _autoFill)
	return nil
}

// GetAutoFill AutoFill Getter
func (r TaobaoItemAddAPIRequest) GetAutoFill() string {
	return r._autoFill
}

// SetFeatures is Features Setter
// 宝贝特征值，格式为：【key1:value1;key2:value2;key3:value3;】，key和value用【:】分隔，key&amp;value之间用【;】分隔，只有在Top支持的特征值才能保存到宝贝上，目前支持的Key列表为：mysize_tp
func (r *TaobaoItemAddAPIRequest) SetFeatures(_features string) error {
	r._features = _features
	r.Set("features", _features)
	return nil
}

// GetFeatures Features Getter
func (r TaobaoItemAddAPIRequest) GetFeatures() string {
	return r._features
}

// SetFoodSecurityPrdLicenseNo is FoodSecurityPrdLicenseNo Setter
// 生产许可证号
func (r *TaobaoItemAddAPIRequest) SetFoodSecurityPrdLicenseNo(_foodSecurityPrdLicenseNo string) error {
	r._foodSecurityPrdLicenseNo = _foodSecurityPrdLicenseNo
	r.Set("food_security.prd_license_no", _foodSecurityPrdLicenseNo)
	return nil
}

// GetFoodSecurityPrdLicenseNo FoodSecurityPrdLicenseNo Getter
func (r TaobaoItemAddAPIRequest) GetFoodSecurityPrdLicenseNo() string {
	return r._foodSecurityPrdLicenseNo
}

// SetFoodSecurityDesignCode is FoodSecurityDesignCode Setter
// 产品标准号
func (r *TaobaoItemAddAPIRequest) SetFoodSecurityDesignCode(_foodSecurityDesignCode string) error {
	r._foodSecurityDesignCode = _foodSecurityDesignCode
	r.Set("food_security.design_code", _foodSecurityDesignCode)
	return nil
}

// GetFoodSecurityDesignCode FoodSecurityDesignCode Getter
func (r TaobaoItemAddAPIRequest) GetFoodSecurityDesignCode() string {
	return r._foodSecurityDesignCode
}

// SetFoodSecurityFactory is FoodSecurityFactory Setter
// 厂名
func (r *TaobaoItemAddAPIRequest) SetFoodSecurityFactory(_foodSecurityFactory string) error {
	r._foodSecurityFactory = _foodSecurityFactory
	r.Set("food_security.factory", _foodSecurityFactory)
	return nil
}

// GetFoodSecurityFactory FoodSecurityFactory Getter
func (r TaobaoItemAddAPIRequest) GetFoodSecurityFactory() string {
	return r._foodSecurityFactory
}

// SetFoodSecurityFactorySite is FoodSecurityFactorySite Setter
// 厂址
func (r *TaobaoItemAddAPIRequest) SetFoodSecurityFactorySite(_foodSecurityFactorySite string) error {
	r._foodSecurityFactorySite = _foodSecurityFactorySite
	r.Set("food_security.factory_site", _foodSecurityFactorySite)
	return nil
}

// GetFoodSecurityFactorySite FoodSecurityFactorySite Getter
func (r TaobaoItemAddAPIRequest) GetFoodSecurityFactorySite() string {
	return r._foodSecurityFactorySite
}

// SetFoodSecurityContact is FoodSecurityContact Setter
// 厂家联系方式
func (r *TaobaoItemAddAPIRequest) SetFoodSecurityContact(_foodSecurityContact string) error {
	r._foodSecurityContact = _foodSecurityContact
	r.Set("food_security.contact", _foodSecurityContact)
	return nil
}

// GetFoodSecurityContact FoodSecurityContact Getter
func (r TaobaoItemAddAPIRequest) GetFoodSecurityContact() string {
	return r._foodSecurityContact
}

// SetFoodSecurityMix is FoodSecurityMix Setter
// 配料表
func (r *TaobaoItemAddAPIRequest) SetFoodSecurityMix(_foodSecurityMix string) error {
	r._foodSecurityMix = _foodSecurityMix
	r.Set("food_security.mix", _foodSecurityMix)
	return nil
}

// GetFoodSecurityMix FoodSecurityMix Getter
func (r TaobaoItemAddAPIRequest) GetFoodSecurityMix() string {
	return r._foodSecurityMix
}

// SetFoodSecurityPlanStorage is FoodSecurityPlanStorage Setter
// 储藏方法
func (r *TaobaoItemAddAPIRequest) SetFoodSecurityPlanStorage(_foodSecurityPlanStorage string) error {
	r._foodSecurityPlanStorage = _foodSecurityPlanStorage
	r.Set("food_security.plan_storage", _foodSecurityPlanStorage)
	return nil
}

// GetFoodSecurityPlanStorage FoodSecurityPlanStorage Getter
func (r TaobaoItemAddAPIRequest) GetFoodSecurityPlanStorage() string {
	return r._foodSecurityPlanStorage
}

// SetFoodSecurityPeriod is FoodSecurityPeriod Setter
// 保质期，默认有单位，传入数字
func (r *TaobaoItemAddAPIRequest) SetFoodSecurityPeriod(_foodSecurityPeriod string) error {
	r._foodSecurityPeriod = _foodSecurityPeriod
	r.Set("food_security.period", _foodSecurityPeriod)
	return nil
}

// GetFoodSecurityPeriod FoodSecurityPeriod Getter
func (r TaobaoItemAddAPIRequest) GetFoodSecurityPeriod() string {
	return r._foodSecurityPeriod
}

// SetFoodSecurityFoodAdditive is FoodSecurityFoodAdditive Setter
// 食品添加剂
func (r *TaobaoItemAddAPIRequest) SetFoodSecurityFoodAdditive(_foodSecurityFoodAdditive string) error {
	r._foodSecurityFoodAdditive = _foodSecurityFoodAdditive
	r.Set("food_security.food_additive", _foodSecurityFoodAdditive)
	return nil
}

// GetFoodSecurityFoodAdditive FoodSecurityFoodAdditive Getter
func (r TaobaoItemAddAPIRequest) GetFoodSecurityFoodAdditive() string {
	return r._foodSecurityFoodAdditive
}

// SetFoodSecuritySupplier is FoodSecuritySupplier Setter
// 供货商
func (r *TaobaoItemAddAPIRequest) SetFoodSecuritySupplier(_foodSecuritySupplier string) error {
	r._foodSecuritySupplier = _foodSecuritySupplier
	r.Set("food_security.supplier", _foodSecuritySupplier)
	return nil
}

// GetFoodSecuritySupplier FoodSecuritySupplier Getter
func (r TaobaoItemAddAPIRequest) GetFoodSecuritySupplier() string {
	return r._foodSecuritySupplier
}

// SetFoodSecurityProductDateStart is FoodSecurityProductDateStart Setter
// 生产开始日期，格式必须为yyyy-MM-dd
func (r *TaobaoItemAddAPIRequest) SetFoodSecurityProductDateStart(_foodSecurityProductDateStart string) error {
	r._foodSecurityProductDateStart = _foodSecurityProductDateStart
	r.Set("food_security.product_date_start", _foodSecurityProductDateStart)
	return nil
}

// GetFoodSecurityProductDateStart FoodSecurityProductDateStart Getter
func (r TaobaoItemAddAPIRequest) GetFoodSecurityProductDateStart() string {
	return r._foodSecurityProductDateStart
}

// SetFoodSecurityProductDateEnd is FoodSecurityProductDateEnd Setter
// 生产结束日期,格式必须为yyyy-MM-dd
func (r *TaobaoItemAddAPIRequest) SetFoodSecurityProductDateEnd(_foodSecurityProductDateEnd string) error {
	r._foodSecurityProductDateEnd = _foodSecurityProductDateEnd
	r.Set("food_security.product_date_end", _foodSecurityProductDateEnd)
	return nil
}

// GetFoodSecurityProductDateEnd FoodSecurityProductDateEnd Getter
func (r TaobaoItemAddAPIRequest) GetFoodSecurityProductDateEnd() string {
	return r._foodSecurityProductDateEnd
}

// SetFoodSecurityStockDateStart is FoodSecurityStockDateStart Setter
// 进货开始日期，要在生产日期之后，格式必须为yyyy-MM-dd
func (r *TaobaoItemAddAPIRequest) SetFoodSecurityStockDateStart(_foodSecurityStockDateStart string) error {
	r._foodSecurityStockDateStart = _foodSecurityStockDateStart
	r.Set("food_security.stock_date_start", _foodSecurityStockDateStart)
	return nil
}

// GetFoodSecurityStockDateStart FoodSecurityStockDateStart Getter
func (r TaobaoItemAddAPIRequest) GetFoodSecurityStockDateStart() string {
	return r._foodSecurityStockDateStart
}

// SetFoodSecurityStockDateEnd is FoodSecurityStockDateEnd Setter
// 进货结束日期，要在生产日期之后，格式必须为yyyy-MM-dd
func (r *TaobaoItemAddAPIRequest) SetFoodSecurityStockDateEnd(_foodSecurityStockDateEnd string) error {
	r._foodSecurityStockDateEnd = _foodSecurityStockDateEnd
	r.Set("food_security.stock_date_end", _foodSecurityStockDateEnd)
	return nil
}

// GetFoodSecurityStockDateEnd FoodSecurityStockDateEnd Getter
func (r TaobaoItemAddAPIRequest) GetFoodSecurityStockDateEnd() string {
	return r._foodSecurityStockDateEnd
}

// SetFoodSecurityHealthProductNo is FoodSecurityHealthProductNo Setter
// 健字号，保健品/膳食营养补充剂 这个类目下特有的信息，此类目下无需填写生产许可证编号（QS），如果填写了生产许可证编号（QS）将被忽略不保存；保存宝贝时，标题前会自动加上健字号产品名称一起作为宝贝标题；
func (r *TaobaoItemAddAPIRequest) SetFoodSecurityHealthProductNo(_foodSecurityHealthProductNo string) error {
	r._foodSecurityHealthProductNo = _foodSecurityHealthProductNo
	r.Set("food_security.health_product_no", _foodSecurityHealthProductNo)
	return nil
}

// GetFoodSecurityHealthProductNo FoodSecurityHealthProductNo Getter
func (r TaobaoItemAddAPIRequest) GetFoodSecurityHealthProductNo() string {
	return r._foodSecurityHealthProductNo
}

// SetScenicTicketBookCost is ScenicTicketBookCost Setter
// 景区门票在选择订金支付时候，需要交的预订费。传入的值是1到20之间的数值，小数点后最多可以保留两位（多余的部分将做四舍五入的处理）。这个数值表示的是预订费的比例，最终的预订费为 scenic_ticket_book_cost乘一口价除以100
func (r *TaobaoItemAddAPIRequest) SetScenicTicketBookCost(_scenicTicketBookCost string) error {
	r._scenicTicketBookCost = _scenicTicketBookCost
	r.Set("scenic_ticket_book_cost", _scenicTicketBookCost)
	return nil
}

// GetScenicTicketBookCost ScenicTicketBookCost Getter
func (r TaobaoItemAddAPIRequest) GetScenicTicketBookCost() string {
	return r._scenicTicketBookCost
}

// SetGlobalStockType is GlobalStockType Setter
// 全球购商品采购地（库存类型），有两种库存类型：现货和代购参数值为1时代表现货，值为2时代表代购。注意：使用时请与 全球购商品采购地（地区/国家）配合使用
func (r *TaobaoItemAddAPIRequest) SetGlobalStockType(_globalStockType string) error {
	r._globalStockType = _globalStockType
	r.Set("global_stock_type", _globalStockType)
	return nil
}

// GetGlobalStockType GlobalStockType Getter
func (r TaobaoItemAddAPIRequest) GetGlobalStockType() string {
	return r._globalStockType
}

// SetGlobalStockCountry is GlobalStockCountry Setter
// 全球购商品采购地（地区/国家）,默认值只在全球购商品采购地（库存类型选择情况生效），地区国家值请填写法定的国家名称，类如（美国, 香港, 日本, 英国, 新西兰, 德国, 韩国, 荷兰, 法国, 意大利, 台湾, 澳门, 加拿大, 瑞士, 西班牙, 泰国, 新加坡, 马来西亚, 菲律宾），不要使用其他
func (r *TaobaoItemAddAPIRequest) SetGlobalStockCountry(_globalStockCountry string) error {
	r._globalStockCountry = _globalStockCountry
	r.Set("global_stock_country", _globalStockCountry)
	return nil
}

// GetGlobalStockCountry GlobalStockCountry Getter
func (r TaobaoItemAddAPIRequest) GetGlobalStockCountry() string {
	return r._globalStockCountry
}

// SetGlobalStockDeliveryPlace is GlobalStockDeliveryPlace Setter
// 全球购商品发货地，发货地现在有两种类型：“国内”和“海外及港澳台”，参数值为1时代表“国内”，值为2时代表“海外及港澳台”，默认为国内。注意：卖家必须已经签署并启用“海外直邮”合约，才能选择发货地为“海外及港澳台”
func (r *TaobaoItemAddAPIRequest) SetGlobalStockDeliveryPlace(_globalStockDeliveryPlace string) error {
	r._globalStockDeliveryPlace = _globalStockDeliveryPlace
	r.Set("global_stock_delivery_place", _globalStockDeliveryPlace)
	return nil
}

// GetGlobalStockDeliveryPlace GlobalStockDeliveryPlace Getter
func (r TaobaoItemAddAPIRequest) GetGlobalStockDeliveryPlace() string {
	return r._globalStockDeliveryPlace
}

// SetLocalityLifeExpirydate is LocalityLifeExpirydate Setter
// 本地生活电子交易凭证业务，目前此字段只涉及到的信息为有效期;如果有效期为起止日期类型，此值为2012-08-06,2012-08-16如果有效期为【购买成功日 至】类型则格式为2012-08-16如果有效期为天数类型则格式为15
func (r *TaobaoItemAddAPIRequest) SetLocalityLifeExpirydate(_localityLifeExpirydate string) error {
	r._localityLifeExpirydate = _localityLifeExpirydate
	r.Set("locality_life.expirydate", _localityLifeExpirydate)
	return nil
}

// GetLocalityLifeExpirydate LocalityLifeExpirydate Getter
func (r TaobaoItemAddAPIRequest) GetLocalityLifeExpirydate() string {
	return r._localityLifeExpirydate
}

// SetLocalityLifeNetworkId is LocalityLifeNetworkId Setter
// 网点ID
func (r *TaobaoItemAddAPIRequest) SetLocalityLifeNetworkId(_localityLifeNetworkId string) error {
	r._localityLifeNetworkId = _localityLifeNetworkId
	r.Set("locality_life.network_id", _localityLifeNetworkId)
	return nil
}

// GetLocalityLifeNetworkId LocalityLifeNetworkId Getter
func (r TaobaoItemAddAPIRequest) GetLocalityLifeNetworkId() string {
	return r._localityLifeNetworkId
}

// SetLocalityLifeMerchant is LocalityLifeMerchant Setter
// 码商信息，格式为 码商id:nick
func (r *TaobaoItemAddAPIRequest) SetLocalityLifeMerchant(_localityLifeMerchant string) error {
	r._localityLifeMerchant = _localityLifeMerchant
	r.Set("locality_life.merchant", _localityLifeMerchant)
	return nil
}

// GetLocalityLifeMerchant LocalityLifeMerchant Getter
func (r TaobaoItemAddAPIRequest) GetLocalityLifeMerchant() string {
	return r._localityLifeMerchant
}

// SetLocalityLifeVerification is LocalityLifeVerification Setter
// 核销打款 1代表核销打款 0代表非核销打款
func (r *TaobaoItemAddAPIRequest) SetLocalityLifeVerification(_localityLifeVerification string) error {
	r._localityLifeVerification = _localityLifeVerification
	r.Set("locality_life.verification", _localityLifeVerification)
	return nil
}

// GetLocalityLifeVerification LocalityLifeVerification Getter
func (r TaobaoItemAddAPIRequest) GetLocalityLifeVerification() string {
	return r._localityLifeVerification
}

// SetLocalityLifeChooseLogis is LocalityLifeChooseLogis Setter
// 发布电子凭证宝贝时候表示是否使用邮寄 0: 代表不使用邮寄； 1：代表使用邮寄；如果不设置这个值，代表不使用邮寄
func (r *TaobaoItemAddAPIRequest) SetLocalityLifeChooseLogis(_localityLifeChooseLogis string) error {
	r._localityLifeChooseLogis = _localityLifeChooseLogis
	r.Set("locality_life.choose_logis", _localityLifeChooseLogis)
	return nil
}

// GetLocalityLifeChooseLogis LocalityLifeChooseLogis Getter
func (r TaobaoItemAddAPIRequest) GetLocalityLifeChooseLogis() string {
	return r._localityLifeChooseLogis
}

// SetLocalityLifeRefundmafee is LocalityLifeRefundmafee Setter
// 退款码费承担方。发布电子凭证宝贝的时候会增加“退款码费承担方”配置项，可选填：(1)s（卖家承担） (2)b(买家承担)
func (r *TaobaoItemAddAPIRequest) SetLocalityLifeRefundmafee(_localityLifeRefundmafee string) error {
	r._localityLifeRefundmafee = _localityLifeRefundmafee
	r.Set("locality_life.refundmafee", _localityLifeRefundmafee)
	return nil
}

// GetLocalityLifeRefundmafee LocalityLifeRefundmafee Getter
func (r TaobaoItemAddAPIRequest) GetLocalityLifeRefundmafee() string {
	return r._localityLifeRefundmafee
}

// SetLocalityLifeObs is LocalityLifeObs Setter
// 预约门店是否支持门店自提,1:是
func (r *TaobaoItemAddAPIRequest) SetLocalityLifeObs(_localityLifeObs string) error {
	r._localityLifeObs = _localityLifeObs
	r.Set("locality_life.obs", _localityLifeObs)
	return nil
}

// GetLocalityLifeObs LocalityLifeObs Getter
func (r TaobaoItemAddAPIRequest) GetLocalityLifeObs() string {
	return r._localityLifeObs
}

// SetLocalityLifeEticket is LocalityLifeEticket Setter
// 电子凭证业务属性，数据字典是: 1、is_card:1 (暂时不用) 2、consume_way:4 （1 串码 ，4 身份证）3、consume_midmnick ：(核销放行账号:用户id-用户名，支持多个，用逗号分隔,例如 1234-测试账号35,1345-测试账号56）4、market:eticket (电子凭证商品标记) 5、has_pos:1 (1 表示商品配置线下门店，在detail上进行展示 ，没有或者其他值只不展示)格式是: k1:v2;k2:v2;........ 如：has_pos:1;market:eticket;consume_midmnick:901409638-OPPO;consume_way:4
func (r *TaobaoItemAddAPIRequest) SetLocalityLifeEticket(_localityLifeEticket string) error {
	r._localityLifeEticket = _localityLifeEticket
	r.Set("locality_life.eticket", _localityLifeEticket)
	return nil
}

// GetLocalityLifeEticket LocalityLifeEticket Getter
func (r TaobaoItemAddAPIRequest) GetLocalityLifeEticket() string {
	return r._localityLifeEticket
}

// SetLocalityLifeVersion is LocalityLifeVersion Setter
// 新版电子凭证字段
func (r *TaobaoItemAddAPIRequest) SetLocalityLifeVersion(_localityLifeVersion string) error {
	r._localityLifeVersion = _localityLifeVersion
	r.Set("locality_life.version", _localityLifeVersion)
	return nil
}

// GetLocalityLifeVersion LocalityLifeVersion Getter
func (r TaobaoItemAddAPIRequest) GetLocalityLifeVersion() string {
	return r._localityLifeVersion
}

// SetLocalityLifePackageid is LocalityLifePackageid Setter
// 新版电子凭证包 id
func (r *TaobaoItemAddAPIRequest) SetLocalityLifePackageid(_localityLifePackageid string) error {
	r._localityLifePackageid = _localityLifePackageid
	r.Set("locality_life.packageid", _localityLifePackageid)
	return nil
}

// GetLocalityLifePackageid LocalityLifePackageid Getter
func (r TaobaoItemAddAPIRequest) GetLocalityLifePackageid() string {
	return r._localityLifePackageid
}

// SetItemWeight is ItemWeight Setter
// 商品的重量，用于按重量计费的运费模板。注意：单位为kg。只能传入数值类型（包含小数），不能带单位，单位默认为kg。
func (r *TaobaoItemAddAPIRequest) SetItemWeight(_itemWeight string) error {
	r._itemWeight = _itemWeight
	r.Set("item_weight", _itemWeight)
	return nil
}

// GetItemWeight ItemWeight Getter
func (r TaobaoItemAddAPIRequest) GetItemWeight() string {
	return r._itemWeight
}

// SetItemSize is ItemSize Setter
// 表示商品的体积，如果需要使用按体积计费的运费模板，一定要设置这个值。该值的单位为立方米（m3），如果是其他单位，请转换成成立方米。该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）。体积和长宽高都支持小数类型。在传入体积或长宽高时候，不能带单位。体积的单位默认为立方米（m3），长宽高的单位默认为米(m)该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）
func (r *TaobaoItemAddAPIRequest) SetItemSize(_itemSize string) error {
	r._itemSize = _itemSize
	r.Set("item_size", _itemSize)
	return nil
}

// GetItemSize ItemSize Getter
func (r TaobaoItemAddAPIRequest) GetItemSize() string {
	return r._itemSize
}

// SetInputCustomCpv is InputCustomCpv Setter
// 针对当前商品的自定义属性值，目前是针对销售属性值自定义的，所以调用方需要把自定义属性值对应的虚拟属性值ID（负整数，例如例子中的 -1和-2）像标准属性值值的id一样设置到SKU上，如果自定义属性值有属性值图片，也要设置到属性图片上
func (r *TaobaoItemAddAPIRequest) SetInputCustomCpv(_inputCustomCpv string) error {
	r._inputCustomCpv = _inputCustomCpv
	r.Set("input_custom_cpv", _inputCustomCpv)
	return nil
}

// GetInputCustomCpv InputCustomCpv Getter
func (r TaobaoItemAddAPIRequest) GetInputCustomCpv() string {
	return r._inputCustomCpv
}

// SetCpvMemo is CpvMemo Setter
// 针对当前商品的标准属性值的补充说明，让买家更加了解商品信息减少交易纠纷
func (r *TaobaoItemAddAPIRequest) SetCpvMemo(_cpvMemo string) error {
	r._cpvMemo = _cpvMemo
	r.Set("cpv_memo", _cpvMemo)
	return nil
}

// GetCpvMemo CpvMemo Getter
func (r TaobaoItemAddAPIRequest) GetCpvMemo() string {
	return r._cpvMemo
}

// SetCustomMadeTypeId is CustomMadeTypeId Setter
// 定制工具Id如果支持定制市场，这个值不填写，就用之前的定制工具Id，之前的定制工具Id没有值就默认为-1
func (r *TaobaoItemAddAPIRequest) SetCustomMadeTypeId(_customMadeTypeId string) error {
	r._customMadeTypeId = _customMadeTypeId
	r.Set("custom_made_type_id", _customMadeTypeId)
	return nil
}

// GetCustomMadeTypeId CustomMadeTypeId Getter
func (r TaobaoItemAddAPIRequest) GetCustomMadeTypeId() string {
	return r._customMadeTypeId
}

// SetNewprepay is Newprepay Setter
// 该宝贝是否支持【7天无理由退货】，卖家选择的值只是一个因素，最终以类目和选择的属性条件来确定是否支持7天。填入字符0，表示不支持；未填写或填人字符1，表示支持7天无理由退货；
func (r *TaobaoItemAddAPIRequest) SetNewprepay(_newprepay string) error {
	r._newprepay = _newprepay
	r.Set("newprepay", _newprepay)
	return nil
}

// GetNewprepay Newprepay Getter
func (r TaobaoItemAddAPIRequest) GetNewprepay() string {
	return r._newprepay
}

// SetBarcode is Barcode Setter
// 商品条形码
func (r *TaobaoItemAddAPIRequest) SetBarcode(_barcode string) error {
	r._barcode = _barcode
	r.Set("barcode", _barcode)
	return nil
}

// GetBarcode Barcode Getter
func (r TaobaoItemAddAPIRequest) GetBarcode() string {
	return r._barcode
}

// SetSellPoint is SellPoint Setter
// 商品卖点信息，最长150个字符。天猫商家和集市卖家都可用。
func (r *TaobaoItemAddAPIRequest) SetSellPoint(_sellPoint string) error {
	r._sellPoint = _sellPoint
	r.Set("sell_point", _sellPoint)
	return nil
}

// GetSellPoint SellPoint Getter
func (r TaobaoItemAddAPIRequest) GetSellPoint() string {
	return r._sellPoint
}

// SetQualification is Qualification Setter
// 商品资质信息
func (r *TaobaoItemAddAPIRequest) SetQualification(_qualification string) error {
	r._qualification = _qualification
	r.Set("qualification", _qualification)
	return nil
}

// GetQualification Qualification Getter
func (r TaobaoItemAddAPIRequest) GetQualification() string {
	return r._qualification
}

// SetIgnorewarning is Ignorewarning Setter
// 忽略警告提示.
func (r *TaobaoItemAddAPIRequest) SetIgnorewarning(_ignorewarning string) error {
	r._ignorewarning = _ignorewarning
	r.Set("ignorewarning", _ignorewarning)
	return nil
}

// GetIgnorewarning Ignorewarning Getter
func (r TaobaoItemAddAPIRequest) GetIgnorewarning() string {
	return r._ignorewarning
}

// SetMsPaymentReferencePrice is MsPaymentReferencePrice Setter
// 参考价。该商品订单首次支付价格为 订金 价格，用户可根据 参考价 估算全款。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm
func (r *TaobaoItemAddAPIRequest) SetMsPaymentReferencePrice(_msPaymentReferencePrice string) error {
	r._msPaymentReferencePrice = _msPaymentReferencePrice
	r.Set("ms_payment.reference_price", _msPaymentReferencePrice)
	return nil
}

// GetMsPaymentReferencePrice MsPaymentReferencePrice Getter
func (r TaobaoItemAddAPIRequest) GetMsPaymentReferencePrice() string {
	return r._msPaymentReferencePrice
}

// SetMsPaymentVoucherPrice is MsPaymentVoucherPrice Setter
// 尾款可抵扣金额。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm
func (r *TaobaoItemAddAPIRequest) SetMsPaymentVoucherPrice(_msPaymentVoucherPrice string) error {
	r._msPaymentVoucherPrice = _msPaymentVoucherPrice
	r.Set("ms_payment.voucher_price", _msPaymentVoucherPrice)
	return nil
}

// GetMsPaymentVoucherPrice MsPaymentVoucherPrice Getter
func (r TaobaoItemAddAPIRequest) GetMsPaymentVoucherPrice() string {
	return r._msPaymentVoucherPrice
}

// SetMsPaymentPrice is MsPaymentPrice Setter
// 订金。在“线上付订金线下付尾款”模式中，有订金、尾款可抵扣金额和参考价，三者需要同时填写。该商品订单首次支付价格为 订金 价格，用户可根据 参考价 估算全款。该模式有别于“一口价”付款方式，针对一个商品，只能选择两种付款方式中的一种，其适用于家装、二手车等场景。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm
func (r *TaobaoItemAddAPIRequest) SetMsPaymentPrice(_msPaymentPrice string) error {
	r._msPaymentPrice = _msPaymentPrice
	r.Set("ms_payment.price", _msPaymentPrice)
	return nil
}

// GetMsPaymentPrice MsPaymentPrice Getter
func (r TaobaoItemAddAPIRequest) GetMsPaymentPrice() string {
	return r._msPaymentPrice
}

// SetDeliveryTimeNeedDeliveryTime is DeliveryTimeNeedDeliveryTime Setter
// 设置是否使用发货时间，商品级别，sku级别
func (r *TaobaoItemAddAPIRequest) SetDeliveryTimeNeedDeliveryTime(_deliveryTimeNeedDeliveryTime string) error {
	r._deliveryTimeNeedDeliveryTime = _deliveryTimeNeedDeliveryTime
	r.Set("delivery_time.need_delivery_time", _deliveryTimeNeedDeliveryTime)
	return nil
}

// GetDeliveryTimeNeedDeliveryTime DeliveryTimeNeedDeliveryTime Getter
func (r TaobaoItemAddAPIRequest) GetDeliveryTimeNeedDeliveryTime() string {
	return r._deliveryTimeNeedDeliveryTime
}

// SetDeliveryTimeDeliveryTimeType is DeliveryTimeDeliveryTimeType Setter
// 发货时间类型：绝对发货时间或者相对发货时间
func (r *TaobaoItemAddAPIRequest) SetDeliveryTimeDeliveryTimeType(_deliveryTimeDeliveryTimeType string) error {
	r._deliveryTimeDeliveryTimeType = _deliveryTimeDeliveryTimeType
	r.Set("delivery_time.delivery_time_type", _deliveryTimeDeliveryTimeType)
	return nil
}

// GetDeliveryTimeDeliveryTimeType DeliveryTimeDeliveryTimeType Getter
func (r TaobaoItemAddAPIRequest) GetDeliveryTimeDeliveryTimeType() string {
	return r._deliveryTimeDeliveryTimeType
}

// SetDeliveryTimeDeliveryTime is DeliveryTimeDeliveryTime Setter
// 商品级别设置的发货时间。设置了商品级别的发货时间，相对发货时间，则填写相对发货时间的天数（大于3）；绝对发货时间，则填写yyyy-mm-dd格式，如2013-11-11
func (r *TaobaoItemAddAPIRequest) SetDeliveryTimeDeliveryTime(_deliveryTimeDeliveryTime string) error {
	r._deliveryTimeDeliveryTime = _deliveryTimeDeliveryTime
	r.Set("delivery_time.delivery_time", _deliveryTimeDeliveryTime)
	return nil
}

// GetDeliveryTimeDeliveryTime DeliveryTimeDeliveryTime Getter
func (r TaobaoItemAddAPIRequest) GetDeliveryTimeDeliveryTime() string {
	return r._deliveryTimeDeliveryTime
}

// SetChangeProp is ChangeProp Setter
// 基础色数据，淘宝不使用
func (r *TaobaoItemAddAPIRequest) SetChangeProp(_changeProp string) error {
	r._changeProp = _changeProp
	r.Set("change_prop", _changeProp)
	return nil
}

// GetChangeProp ChangeProp Getter
func (r TaobaoItemAddAPIRequest) GetChangeProp() string {
	return r._changeProp
}

// SetDescModules is DescModules Setter
// 已废弃
func (r *TaobaoItemAddAPIRequest) SetDescModules(_descModules string) error {
	r._descModules = _descModules
	r.Set("desc_modules", _descModules)
	return nil
}

// GetDescModules DescModules Getter
func (r TaobaoItemAddAPIRequest) GetDescModules() string {
	return r._descModules
}

// SetIsOffline is IsOffline Setter
// 是否是线下商品。1：线上商品（默认值）；2：线上或线下商品；3：线下商品。
func (r *TaobaoItemAddAPIRequest) SetIsOffline(_isOffline string) error {
	r._isOffline = _isOffline
	r.Set("is_offline", _isOffline)
	return nil
}

// GetIsOffline IsOffline Getter
func (r TaobaoItemAddAPIRequest) GetIsOffline() string {
	return r._isOffline
}

// SetWirelessDesc is WirelessDesc Setter
// 无线的宝贝描述
func (r *TaobaoItemAddAPIRequest) SetWirelessDesc(_wirelessDesc string) error {
	r._wirelessDesc = _wirelessDesc
	r.Set("wireless_desc", _wirelessDesc)
	return nil
}

// GetWirelessDesc WirelessDesc Getter
func (r TaobaoItemAddAPIRequest) GetWirelessDesc() string {
	return r._wirelessDesc
}

// SetLeaseExtendsInfo is LeaseExtendsInfo Setter
// 租赁扩展信息
func (r *TaobaoItemAddAPIRequest) SetLeaseExtendsInfo(_leaseExtendsInfo string) error {
	r._leaseExtendsInfo = _leaseExtendsInfo
	r.Set("lease_extends_info", _leaseExtendsInfo)
	return nil
}

// GetLeaseExtendsInfo LeaseExtendsInfo Getter
func (r TaobaoItemAddAPIRequest) GetLeaseExtendsInfo() string {
	return r._leaseExtendsInfo
}

// SetBrokerage is Brokerage Setter
// 仅淘小铺卖家需要。佣金比例(15.3对应的佣金比例为15.3%).只支持小数点后1位。多余的位数四舍五入(15.32会保存为15.3%
func (r *TaobaoItemAddAPIRequest) SetBrokerage(_brokerage string) error {
	r._brokerage = _brokerage
	r.Set("brokerage", _brokerage)
	return nil
}

// GetBrokerage Brokerage Getter
func (r TaobaoItemAddAPIRequest) GetBrokerage() string {
	return r._brokerage
}

// SetBizCode is BizCode Setter
// 业务身份编码。淘小铺编码为&#34;taobao-taoxiaopu&#34;
func (r *TaobaoItemAddAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r TaobaoItemAddAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetInputPids is InputPids Setter
// 用户自行输入的类目属性ID串，结构：&#34;pid1,pid2,pid3&#34;，如：&#34;20000&#34;（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。
func (r *TaobaoItemAddAPIRequest) SetInputPids(_inputPids string) error {
	r._inputPids = _inputPids
	r.Set("input_pids", _inputPids)
	return nil
}

// GetInputPids InputPids Getter
func (r TaobaoItemAddAPIRequest) GetInputPids() string {
	return r._inputPids
}

// SetInputStr is InputStr Setter
// 用户自行输入的子属性名和属性值，结构:&#34;父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....&#34;,如：“耐克;耐克系列;科比系列;科比系列;2K5,Nike乔丹鞋;乔丹系列;乔丹鞋系列;乔丹鞋系列;json5”，多个自定义属性用&#39;,&#39;分割，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过3999字节。此处不可以使用“其他”、“其它”和“其她”这三个词。
func (r *TaobaoItemAddAPIRequest) SetInputStr(_inputStr string) error {
	r._inputStr = _inputStr
	r.Set("input_str", _inputStr)
	return nil
}

// GetInputStr InputStr Getter
func (r TaobaoItemAddAPIRequest) GetInputStr() string {
	return r._inputStr
}

// SetSkuProperties is SkuProperties Setter
// 更新的sku的属性串，调用taobao.itemprops.get获取。格式:pid1:vid;pid2:vid,多个sku属性之间用逗号分隔。该字段内的属性需要在props字段同时包含。如果新增商品包含了sku，则此字段一定要传入,字段长度要控制在512个字节以内。
func (r *TaobaoItemAddAPIRequest) SetSkuProperties(_skuProperties string) error {
	r._skuProperties = _skuProperties
	r.Set("sku_properties", _skuProperties)
	return nil
}

// GetSkuProperties SkuProperties Getter
func (r TaobaoItemAddAPIRequest) GetSkuProperties() string {
	return r._skuProperties
}

// SetSkuQuantities is SkuQuantities Setter
// Sku的数量串，结构如：num1,num2,num3 如：2,3
func (r *TaobaoItemAddAPIRequest) SetSkuQuantities(_skuQuantities string) error {
	r._skuQuantities = _skuQuantities
	r.Set("sku_quantities", _skuQuantities)
	return nil
}

// GetSkuQuantities SkuQuantities Getter
func (r TaobaoItemAddAPIRequest) GetSkuQuantities() string {
	return r._skuQuantities
}

// SetSkuPrices is SkuPrices Setter
// Sku的价格串，结构如：10.00,5.00,… 精确到2位小数;单位:元。如:200.07，表示:200元7分
func (r *TaobaoItemAddAPIRequest) SetSkuPrices(_skuPrices string) error {
	r._skuPrices = _skuPrices
	r.Set("sku_prices", _skuPrices)
	return nil
}

// GetSkuPrices SkuPrices Getter
func (r TaobaoItemAddAPIRequest) GetSkuPrices() string {
	return r._skuPrices
}

// SetSkuOuterIds is SkuOuterIds Setter
// Sku的外部id串，结构如：1234,1342,… sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，如果没有sku_outer_ids也要写上这个参数，入参是&#34;,&#34;(这个是两个sku的示列，逗号数应该是sku个数减1)；该参数最大长度是512个字节
func (r *TaobaoItemAddAPIRequest) SetSkuOuterIds(_skuOuterIds string) error {
	r._skuOuterIds = _skuOuterIds
	r.Set("sku_outer_ids", _skuOuterIds)
	return nil
}

// GetSkuOuterIds SkuOuterIds Getter
func (r TaobaoItemAddAPIRequest) GetSkuOuterIds() string {
	return r._skuOuterIds
}

// SetSkuBarcode is SkuBarcode Setter
// sku层面的条形码，多个SKU情况，与SKU价格库存格式类似，用逗号分隔
func (r *TaobaoItemAddAPIRequest) SetSkuBarcode(_skuBarcode string) error {
	r._skuBarcode = _skuBarcode
	r.Set("sku_barcode", _skuBarcode)
	return nil
}

// GetSkuBarcode SkuBarcode Getter
func (r TaobaoItemAddAPIRequest) GetSkuBarcode() string {
	return r._skuBarcode
}

// SetSkuSpecIds is SkuSpecIds Setter
// 此参数暂时不起作用
func (r *TaobaoItemAddAPIRequest) SetSkuSpecIds(_skuSpecIds string) error {
	r._skuSpecIds = _skuSpecIds
	r.Set("sku_spec_ids", _skuSpecIds)
	return nil
}

// GetSkuSpecIds SkuSpecIds Getter
func (r TaobaoItemAddAPIRequest) GetSkuSpecIds() string {
	return r._skuSpecIds
}

// SetSkuDeliveryTimes is SkuDeliveryTimes Setter
// 此参数暂时不起作用
func (r *TaobaoItemAddAPIRequest) SetSkuDeliveryTimes(_skuDeliveryTimes string) error {
	r._skuDeliveryTimes = _skuDeliveryTimes
	r.Set("sku_delivery_times", _skuDeliveryTimes)
	return nil
}

// GetSkuDeliveryTimes SkuDeliveryTimes Getter
func (r TaobaoItemAddAPIRequest) GetSkuDeliveryTimes() string {
	return r._skuDeliveryTimes
}

// SetSkuHdLength is SkuHdLength Setter
// 家装建材类目，商品SKU的长度，正整数，单位为cm，部分类目必选。 数据和SKU一一对应，用,分隔，如：20,30,30
func (r *TaobaoItemAddAPIRequest) SetSkuHdLength(_skuHdLength string) error {
	r._skuHdLength = _skuHdLength
	r.Set("sku_hd_length", _skuHdLength)
	return nil
}

// GetSkuHdLength SkuHdLength Getter
func (r TaobaoItemAddAPIRequest) GetSkuHdLength() string {
	return r._skuHdLength
}

// SetSkuHdHeight is SkuHdHeight Setter
// 家装建材类目，商品SKU的高度，单位为cm，部分类目必选。 天猫和淘宝格式不同。天猫：可选值为：&#34;0-15&#34;, &#34;15-25&#34;, &#34;25-50&#34;, &#34;50-60&#34;, &#34;60-80&#34;, &#34;80-120&#34;, &#34;120-160&#34;, &#34;160-200&#34;。 数据和SKU一一对应，用,分隔，格式如：15-25,25-50,25-50。 淘宝：正整数，单位为cm,格式如：20,30,30
func (r *TaobaoItemAddAPIRequest) SetSkuHdHeight(_skuHdHeight string) error {
	r._skuHdHeight = _skuHdHeight
	r.Set("sku_hd_height", _skuHdHeight)
	return nil
}

// GetSkuHdHeight SkuHdHeight Getter
func (r TaobaoItemAddAPIRequest) GetSkuHdHeight() string {
	return r._skuHdHeight
}

// SetSkuHdLampQuantity is SkuHdLampQuantity Setter
// 家装建材类目，商品SKU的灯头数量，正整数，大于等于3，部分类目必选。天猫商家专用。 数据和SKU一一对应，用,分隔，如：3,5,7
func (r *TaobaoItemAddAPIRequest) SetSkuHdLampQuantity(_skuHdLampQuantity string) error {
	r._skuHdLampQuantity = _skuHdLampQuantity
	r.Set("sku_hd_lamp_quantity", _skuHdLampQuantity)
	return nil
}

// GetSkuHdLampQuantity SkuHdLampQuantity Getter
func (r TaobaoItemAddAPIRequest) GetSkuHdLampQuantity() string {
	return r._skuHdLampQuantity
}

// SetCid is Cid Setter
// 叶子类目id
func (r *TaobaoItemAddAPIRequest) SetCid(_cid int64) error {
	r._cid = _cid
	r.Set("cid", _cid)
	return nil
}

// GetCid Cid Getter
func (r TaobaoItemAddAPIRequest) GetCid() int64 {
	return r._cid
}

// SetValidThru is ValidThru Setter
// 有效期。可选值:7,14;单位:天;默认值:14
func (r *TaobaoItemAddAPIRequest) SetValidThru(_validThru int64) error {
	r._validThru = _validThru
	r.Set("valid_thru", _validThru)
	return nil
}

// GetValidThru ValidThru Getter
func (r TaobaoItemAddAPIRequest) GetValidThru() int64 {
	return r._validThru
}

// SetPostFee is PostFee Setter
// 平邮费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:5.07，表示:5元7分. 注:post_fee,express_fee,ems_fee需要一起填写
func (r *TaobaoItemAddAPIRequest) SetPostFee(_postFee float64) error {
	r._postFee = _postFee
	r.Set("post_fee", _postFee)
	return nil
}

// GetPostFee PostFee Getter
func (r TaobaoItemAddAPIRequest) GetPostFee() float64 {
	return r._postFee
}

// SetExpressFee is ExpressFee Setter
// 快递费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:15.07，表示:15元7分
func (r *TaobaoItemAddAPIRequest) SetExpressFee(_expressFee float64) error {
	r._expressFee = _expressFee
	r.Set("express_fee", _expressFee)
	return nil
}

// GetExpressFee ExpressFee Getter
func (r TaobaoItemAddAPIRequest) GetExpressFee() float64 {
	return r._expressFee
}

// SetEmsFee is EmsFee Setter
// ems费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:25.07，表示:25元7分
func (r *TaobaoItemAddAPIRequest) SetEmsFee(_emsFee float64) error {
	r._emsFee = _emsFee
	r.Set("ems_fee", _emsFee)
	return nil
}

// GetEmsFee EmsFee Getter
func (r TaobaoItemAddAPIRequest) GetEmsFee() float64 {
	return r._emsFee
}

// SetIncrement is Increment Setter
// 加价(降价)幅度。如果为0，代表系统代理幅度。对于增价拍和荷兰拍来说是加价幅度，对于降价拍来说是降价幅度。
func (r *TaobaoItemAddAPIRequest) SetIncrement(_increment float64) error {
	r._increment = _increment
	r.Set("increment", _increment)
	return nil
}

// GetIncrement Increment Getter
func (r TaobaoItemAddAPIRequest) GetIncrement() float64 {
	return r._increment
}

// SetNum is Num Setter
// 商品数量。取值范围:0-900000000的整数。且需要等于Sku所有数量的和。拍卖商品中增加拍只能为1，荷兰拍要在[2,500)范围内。
func (r *TaobaoItemAddAPIRequest) SetNum(_num int64) error {
	r._num = _num
	r.Set("num", _num)
	return nil
}

// GetNum Num Getter
func (r TaobaoItemAddAPIRequest) GetNum() int64 {
	return r._num
}

// SetPrice is Price Setter
// 商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。拍卖商品对应的起拍价。
func (r *TaobaoItemAddAPIRequest) SetPrice(_price float64) error {
	r._price = _price
	r.Set("price", _price)
	return nil
}

// GetPrice Price Getter
func (r TaobaoItemAddAPIRequest) GetPrice() float64 {
	return r._price
}

// SetPostageId is PostageId Setter
// 宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.delivery.template.get获得当前会话用户的所有邮费模板）
func (r *TaobaoItemAddAPIRequest) SetPostageId(_postageId int64) error {
	r._postageId = _postageId
	r.Set("postage_id", _postageId)
	return nil
}

// GetPostageId PostageId Getter
func (r TaobaoItemAddAPIRequest) GetPostageId() int64 {
	return r._postageId
}

// SetProductId is ProductId Setter
// 商品所属的产品ID(B商家发布商品需要用)
func (r *TaobaoItemAddAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaoItemAddAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetAuctionPoint is AuctionPoint Setter
// 商品的积分返点比例。如:5,表示:返点比例0.5%. 注意：返点比例必须是&gt;0的整数，而且最大是90,即为9%.B商家在发布非虚拟商品时，返点必须是 5的倍数，即0.5%的倍数。其它是1的倍数，即0.1%的倍数。无名良品商家发布商品时，复用该字段记录积分宝返点比例，返点必须是对应类目的返点步长的整数倍，默认是5，即0.5%。注意此时该字段值依旧必须是&gt;0的整数，最高值不超过500，即50%
func (r *TaobaoItemAddAPIRequest) SetAuctionPoint(_auctionPoint int64) error {
	r._auctionPoint = _auctionPoint
	r.Set("auction_point", _auctionPoint)
	return nil
}

// GetAuctionPoint AuctionPoint Getter
func (r TaobaoItemAddAPIRequest) GetAuctionPoint() int64 {
	return r._auctionPoint
}

// SetImage is Image Setter
// 商品主图片。类型:JPG,GIF;最大长度:3M。（推荐使用pic_path字段，先把图片上传到卖家图片空间）
func (r *TaobaoItemAddAPIRequest) SetImage(_image *model.File) error {
	r._image = _image
	r.Set("image", _image)
	return nil
}

// GetImage Image Getter
func (r TaobaoItemAddAPIRequest) GetImage() *model.File {
	return r._image
}

// SetCodPostageId is CodPostageId Setter
// 此为货到付款运费模板的ID，对应的JAVA类型是long,如果COD卖家应用了货到付款运费模板，此值要进行设置。该字段已经废弃
func (r *TaobaoItemAddAPIRequest) SetCodPostageId(_codPostageId int64) error {
	r._codPostageId = _codPostageId
	r.Set("cod_postage_id", _codPostageId)
	return nil
}

// GetCodPostageId CodPostageId Getter
func (r TaobaoItemAddAPIRequest) GetCodPostageId() int64 {
	return r._codPostageId
}

// SetWeight is Weight Setter
// 商品的重量(商超卖家专用字段)
func (r *TaobaoItemAddAPIRequest) SetWeight(_weight int64) error {
	r._weight = _weight
	r.Set("weight", _weight)
	return nil
}

// GetWeight Weight Getter
func (r TaobaoItemAddAPIRequest) GetWeight() int64 {
	return r._weight
}

// SetSubStock is SubStock Setter
// 商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改集市卖家默认拍下减库存;商城卖家默认付款减库存
func (r *TaobaoItemAddAPIRequest) SetSubStock(_subStock int64) error {
	r._subStock = _subStock
	r.Set("sub_stock", _subStock)
	return nil
}

// GetSubStock SubStock Getter
func (r TaobaoItemAddAPIRequest) GetSubStock() int64 {
	return r._subStock
}

// SetScenicTicketPayWay is ScenicTicketPayWay Setter
// 景区门票类宝贝发布时候，当卖家签订了支付宝代扣协议时候，需要选择支付方式：全额支付和订金支付。当scenic_ticket_pay_way为1时表示全额支付，为2时表示订金支付
func (r *TaobaoItemAddAPIRequest) SetScenicTicketPayWay(_scenicTicketPayWay int64) error {
	r._scenicTicketPayWay = _scenicTicketPayWay
	r.Set("scenic_ticket_pay_way", _scenicTicketPayWay)
	return nil
}

// GetScenicTicketPayWay ScenicTicketPayWay Getter
func (r TaobaoItemAddAPIRequest) GetScenicTicketPayWay() int64 {
	return r._scenicTicketPayWay
}

// SetLocalityLifeRefundRatio is LocalityLifeRefundRatio Setter
// 退款比例，百分比%前的数字,1-100的正整数值
func (r *TaobaoItemAddAPIRequest) SetLocalityLifeRefundRatio(_localityLifeRefundRatio int64) error {
	r._localityLifeRefundRatio = _localityLifeRefundRatio
	r.Set("locality_life.refund_ratio", _localityLifeRefundRatio)
	return nil
}

// GetLocalityLifeRefundRatio LocalityLifeRefundRatio Getter
func (r TaobaoItemAddAPIRequest) GetLocalityLifeRefundRatio() int64 {
	return r._localityLifeRefundRatio
}

// SetLocalityLifeOnsaleAutoRefundRatio is LocalityLifeOnsaleAutoRefundRatio Setter
// 电子凭证售中自动退款比例，百分比%前的数字，介于1-100之间的整数
func (r *TaobaoItemAddAPIRequest) SetLocalityLifeOnsaleAutoRefundRatio(_localityLifeOnsaleAutoRefundRatio int64) error {
	r._localityLifeOnsaleAutoRefundRatio = _localityLifeOnsaleAutoRefundRatio
	r.Set("locality_life.onsale_auto_refund_ratio", _localityLifeOnsaleAutoRefundRatio)
	return nil
}

// GetLocalityLifeOnsaleAutoRefundRatio LocalityLifeOnsaleAutoRefundRatio Getter
func (r TaobaoItemAddAPIRequest) GetLocalityLifeOnsaleAutoRefundRatio() int64 {
	return r._localityLifeOnsaleAutoRefundRatio
}

// SetPaimaiInfoMode is PaimaiInfoMode Setter
// 拍卖商品选择的拍卖类型，拍卖类型包括三种：增价拍(1)，荷兰拍(2)和降价拍(3)。
func (r *TaobaoItemAddAPIRequest) SetPaimaiInfoMode(_paimaiInfoMode int64) error {
	r._paimaiInfoMode = _paimaiInfoMode
	r.Set("paimai_info.mode", _paimaiInfoMode)
	return nil
}

// GetPaimaiInfoMode PaimaiInfoMode Getter
func (r TaobaoItemAddAPIRequest) GetPaimaiInfoMode() int64 {
	return r._paimaiInfoMode
}

// SetPaimaiInfoDeposit is PaimaiInfoDeposit Setter
// 拍卖宝贝的保证金。对于增价拍和荷兰拍来说保证金有两种模式：淘宝默认模式（首次出价金额的10%），自定义固定保证金（固定冻结金额只能输入不超过30万的正整数），并且保证金只冻结1次。对于降价拍来说保证金只有淘宝默认的（竞拍价格的10% * 竞拍数量），并且每次出价都需要冻结保证金。对于拍卖宝贝来说，保证金是必须的，但是默认使用淘宝默认保证金模式，只有用户需要使用自定义固定保证金的时候才需要使用到这个参数，如果该参数不传或传入0则代表使用默认。
func (r *TaobaoItemAddAPIRequest) SetPaimaiInfoDeposit(_paimaiInfoDeposit int64) error {
	r._paimaiInfoDeposit = _paimaiInfoDeposit
	r.Set("paimai_info.deposit", _paimaiInfoDeposit)
	return nil
}

// GetPaimaiInfoDeposit PaimaiInfoDeposit Getter
func (r TaobaoItemAddAPIRequest) GetPaimaiInfoDeposit() int64 {
	return r._paimaiInfoDeposit
}

// SetPaimaiInfoInterval is PaimaiInfoInterval Setter
// 降价拍宝贝的降价周期(分钟)。降价拍宝贝的价格每隔paimai_info.interval时间会下降一次increment。
func (r *TaobaoItemAddAPIRequest) SetPaimaiInfoInterval(_paimaiInfoInterval int64) error {
	r._paimaiInfoInterval = _paimaiInfoInterval
	r.Set("paimai_info.interval", _paimaiInfoInterval)
	return nil
}

// GetPaimaiInfoInterval PaimaiInfoInterval Getter
func (r TaobaoItemAddAPIRequest) GetPaimaiInfoInterval() int64 {
	return r._paimaiInfoInterval
}

// SetPaimaiInfoReserve is PaimaiInfoReserve Setter
// 降价拍宝贝的保留价。对于降价拍来说，paimai_info.reserve必须大于0，且小于price-increment，而且（price-paimai_info.reserve）/increment的计算结果必须为整数
func (r *TaobaoItemAddAPIRequest) SetPaimaiInfoReserve(_paimaiInfoReserve float64) error {
	r._paimaiInfoReserve = _paimaiInfoReserve
	r.Set("paimai_info.reserve", _paimaiInfoReserve)
	return nil
}

// GetPaimaiInfoReserve PaimaiInfoReserve Getter
func (r TaobaoItemAddAPIRequest) GetPaimaiInfoReserve() float64 {
	return r._paimaiInfoReserve
}

// SetPaimaiInfoValidHour is PaimaiInfoValidHour Setter
// 自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。
func (r *TaobaoItemAddAPIRequest) SetPaimaiInfoValidHour(_paimaiInfoValidHour int64) error {
	r._paimaiInfoValidHour = _paimaiInfoValidHour
	r.Set("paimai_info.valid_hour", _paimaiInfoValidHour)
	return nil
}

// GetPaimaiInfoValidHour PaimaiInfoValidHour Getter
func (r TaobaoItemAddAPIRequest) GetPaimaiInfoValidHour() int64 {
	return r._paimaiInfoValidHour
}

// SetPaimaiInfoValidMinute is PaimaiInfoValidMinute Setter
// 自定义销售周期的分钟数。拍卖宝贝可以自定义销售周期，这里是指定销售周期的分钟数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。
func (r *TaobaoItemAddAPIRequest) SetPaimaiInfoValidMinute(_paimaiInfoValidMinute int64) error {
	r._paimaiInfoValidMinute = _paimaiInfoValidMinute
	r.Set("paimai_info.valid_minute", _paimaiInfoValidMinute)
	return nil
}

// GetPaimaiInfoValidMinute PaimaiInfoValidMinute Getter
func (r TaobaoItemAddAPIRequest) GetPaimaiInfoValidMinute() int64 {
	return r._paimaiInfoValidMinute
}

// SetAfterSaleId is AfterSaleId Setter
// 售后说明模板id
func (r *TaobaoItemAddAPIRequest) SetAfterSaleId(_afterSaleId int64) error {
	r._afterSaleId = _afterSaleId
	r.Set("after_sale_id", _afterSaleId)
	return nil
}

// GetAfterSaleId AfterSaleId Getter
func (r TaobaoItemAddAPIRequest) GetAfterSaleId() int64 {
	return r._afterSaleId
}

// SetVideoId is VideoId Setter
// 主图视频id
func (r *TaobaoItemAddAPIRequest) SetVideoId(_videoId int64) error {
	r._videoId = _videoId
	r.Set("video_id", _videoId)
	return nil
}

// GetVideoId VideoId Getter
func (r TaobaoItemAddAPIRequest) GetVideoId() int64 {
	return r._videoId
}

// SetInteractiveId is InteractiveId Setter
// 主图视频互动信息id，必须填写主图视频id才能有互动信息id
func (r *TaobaoItemAddAPIRequest) SetInteractiveId(_interactiveId int64) error {
	r._interactiveId = _interactiveId
	r.Set("interactive_id", _interactiveId)
	return nil
}

// GetInteractiveId InteractiveId Getter
func (r TaobaoItemAddAPIRequest) GetInteractiveId() int64 {
	return r._interactiveId
}

// SetHasWarranty is HasWarranty Setter
// 是否有保修。可选值:true,false;默认值:false(不保修)
func (r *TaobaoItemAddAPIRequest) SetHasWarranty(_hasWarranty bool) error {
	r._hasWarranty = _hasWarranty
	r.Set("has_warranty", _hasWarranty)
	return nil
}

// GetHasWarranty HasWarranty Getter
func (r TaobaoItemAddAPIRequest) GetHasWarranty() bool {
	return r._hasWarranty
}

// SetHasInvoice is HasInvoice Setter
// 是否有发票。可选值:true,false (商城卖家此字段必须为true);默认值:false(无发票)
func (r *TaobaoItemAddAPIRequest) SetHasInvoice(_hasInvoice bool) error {
	r._hasInvoice = _hasInvoice
	r.Set("has_invoice", _hasInvoice)
	return nil
}

// GetHasInvoice HasInvoice Getter
func (r TaobaoItemAddAPIRequest) GetHasInvoice() bool {
	return r._hasInvoice
}

// SetHasShowcase is HasShowcase Setter
// 橱窗推荐。可选值:true,false;默认值:false(不推荐)
func (r *TaobaoItemAddAPIRequest) SetHasShowcase(_hasShowcase bool) error {
	r._hasShowcase = _hasShowcase
	r.Set("has_showcase", _hasShowcase)
	return nil
}

// GetHasShowcase HasShowcase Getter
func (r TaobaoItemAddAPIRequest) GetHasShowcase() bool {
	return r._hasShowcase
}

// SetHasDiscount is HasDiscount Setter
// 支持会员打折。可选值:true,false;默认值:false(不打折)
func (r *TaobaoItemAddAPIRequest) SetHasDiscount(_hasDiscount bool) error {
	r._hasDiscount = _hasDiscount
	r.Set("has_discount", _hasDiscount)
	return nil
}

// GetHasDiscount HasDiscount Getter
func (r TaobaoItemAddAPIRequest) GetHasDiscount() bool {
	return r._hasDiscount
}

// SetIsTaobao is IsTaobao Setter
// 是否在淘宝上显示（如果传FALSE，则在淘宝主站无法显示该商品）
func (r *TaobaoItemAddAPIRequest) SetIsTaobao(_isTaobao bool) error {
	r._isTaobao = _isTaobao
	r.Set("is_taobao", _isTaobao)
	return nil
}

// GetIsTaobao IsTaobao Getter
func (r TaobaoItemAddAPIRequest) GetIsTaobao() bool {
	return r._isTaobao
}

// SetIsEx is IsEx Setter
// 是否在外店显示
func (r *TaobaoItemAddAPIRequest) SetIsEx(_isEx bool) error {
	r._isEx = _isEx
	r.Set("is_ex", _isEx)
	return nil
}

// GetIsEx IsEx Getter
func (r TaobaoItemAddAPIRequest) GetIsEx() bool {
	return r._isEx
}

// SetIs3D is Is3D Setter
// 是否是3D
func (r *TaobaoItemAddAPIRequest) SetIs3D(_is3D bool) error {
	r._is3D = _is3D
	r.Set("is_3D", _is3D)
	return nil
}

// GetIs3D Is3D Getter
func (r TaobaoItemAddAPIRequest) GetIs3D() bool {
	return r._is3D
}

// SetSellPromise is SellPromise Setter
// 是否承诺退换货服务!虚拟商品无须设置此项!
func (r *TaobaoItemAddAPIRequest) SetSellPromise(_sellPromise bool) error {
	r._sellPromise = _sellPromise
	r.Set("sell_promise", _sellPromise)
	return nil
}

// GetSellPromise SellPromise Getter
func (r TaobaoItemAddAPIRequest) GetSellPromise() bool {
	return r._sellPromise
}

// SetIsLightningConsignment is IsLightningConsignment Setter
// 实物闪电发货
func (r *TaobaoItemAddAPIRequest) SetIsLightningConsignment(_isLightningConsignment bool) error {
	r._isLightningConsignment = _isLightningConsignment
	r.Set("is_lightning_consignment", _isLightningConsignment)
	return nil
}

// GetIsLightningConsignment IsLightningConsignment Getter
func (r TaobaoItemAddAPIRequest) GetIsLightningConsignment() bool {
	return r._isLightningConsignment
}

// SetIsXinpin is IsXinpin Setter
// 商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:add-xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置该参数值或设置为false效果一致。
func (r *TaobaoItemAddAPIRequest) SetIsXinpin(_isXinpin bool) error {
	r._isXinpin = _isXinpin
	r.Set("is_xinpin", _isXinpin)
	return nil
}

// GetIsXinpin IsXinpin Getter
func (r TaobaoItemAddAPIRequest) GetIsXinpin() bool {
	return r._isXinpin
}

// SetGlobalStockTaxFreePromise is GlobalStockTaxFreePromise Setter
// 全球购商品卖家包税承诺，当值为true时，代表卖家承诺包税。注意：请与“全球购商品发货地”配合使用，包税承诺必须满足：1、发货地为海外及港澳台 2、卖家已经签署并启用“包税合约”合约
func (r *TaobaoItemAddAPIRequest) SetGlobalStockTaxFreePromise(_globalStockTaxFreePromise bool) error {
	r._globalStockTaxFreePromise = _globalStockTaxFreePromise
	r.Set("global_stock_tax_free_promise", _globalStockTaxFreePromise)
	return nil
}

// GetGlobalStockTaxFreePromise GlobalStockTaxFreePromise Getter
func (r TaobaoItemAddAPIRequest) GetGlobalStockTaxFreePromise() bool {
	return r._globalStockTaxFreePromise
}

// SetSupportCustomMade is SupportCustomMade Setter
// 是否支持定制市场 true代表支持，false代表支持,如果为空代表与之前保持不变不会修改
func (r *TaobaoItemAddAPIRequest) SetSupportCustomMade(_supportCustomMade bool) error {
	r._supportCustomMade = _supportCustomMade
	r.Set("support_custom_made", _supportCustomMade)
	return nil
}

// GetSupportCustomMade SupportCustomMade Getter
func (r TaobaoItemAddAPIRequest) GetSupportCustomMade() bool {
	return r._supportCustomMade
}

// SetO2oBindService is O2oBindService Setter
// 汽车O2O绑定线下服务标记，如不为空，表示关联服务，否则，不关联服务。
func (r *TaobaoItemAddAPIRequest) SetO2oBindService(_o2oBindService bool) error {
	r._o2oBindService = _o2oBindService
	r.Set("o2o_bind_service", _o2oBindService)
	return nil
}

// GetO2oBindService O2oBindService Getter
func (r TaobaoItemAddAPIRequest) GetO2oBindService() bool {
	return r._o2oBindService
}

// SetSpuConfirm is SpuConfirm Setter
// 手机类目spu 优化，信息确认字段
func (r *TaobaoItemAddAPIRequest) SetSpuConfirm(_spuConfirm bool) error {
	r._spuConfirm = _spuConfirm
	r.Set("spu_confirm", _spuConfirm)
	return nil
}

// GetSpuConfirm SpuConfirm Getter
func (r TaobaoItemAddAPIRequest) GetSpuConfirm() bool {
	return r._spuConfirm
}
