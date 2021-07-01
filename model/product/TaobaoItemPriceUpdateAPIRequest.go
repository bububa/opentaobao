package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemPriceUpdateAPIRequest
更新商品价格 API请求
taobao.item.price.update

更新商品价格 */
type TaobaoItemPriceUpdateAPIRequest struct {
	model.Params
	// 所在地省份。如浙江 具体可以下载http://dl.open.taobao.com/sdk/商品城市列表.rar 取到
	_locationState string
	// 所在地城市。如杭州 具体可以下载http://dl.open.taobao.com/sdk/商品城市列表.rar 取到
	_locationCity string
	// 商品数字ID，该参数必须
	_numIid int64
	// 叶子类目id
	_cid int64
	// 商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值。
	_props string
	// 商品数量，取值范围:0-999999的整数。且需要等于Sku所有数量的和
	_num int64
	// 商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。
	_price float64
	// 宝贝标题. 不能超过60字符,受违禁词控制
	_title string
	// 商品描述. 字数要大于5个字符，小于25000个字符 ，受违禁词控制
	_desc string
	// 平邮费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:5.07，表示:5元7分, 注:post_fee,express_fee,ems_fee需一起填写
	_postFee float64
	// 快递费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:15.07，表示:15元7分
	_expressFee float64
	// ems费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:25.07，表示:25元7分
	_emsFee float64
	// 上架时间。不论是更新架下的商品还是出售中的商品，如果这个字段小于当前时间则直接上架商品，并且上架的时间为更新商品的时间，此时item.is_timing为false，如果大于当前时间则宝贝会下架进入定时上架的宝贝中。
	_listTime string
	// 加价幅度 如果为0，代表系统代理幅度
	_increment float64
	// 商品图片。类型:JPG,GIF;最大长度:500k
	_image *model.File
	// 商品新旧程度。可选值:new（全新）,unused（闲置）,second（二手）。
	_stuffStatus string
	// 商品的积分返点比例。如：5 表示返点比例0.5%. 注意：返点比例必须是>0的整数，而且最大是90,即为9%.B商家在发布非虚拟商品时，返点必须是 5的倍数，即0.5%的倍数。其它是1的倍数，即0.1%的倍数。无名良品商家发布商品时，复用该字段记录积分宝返点比例，返点必须是对应类目的返点步长的整数倍，默认是5，即0.5%。注意此时该字段值依旧必须是>0的整数，注意此时该字段值依旧必须是>0的整数，最高值不超过500，即50%
	_auctionPoint int64
	// 属性值别名。如pid:vid:别名;pid1:vid1:别名1， pid:属性id vid:值id。总长度不超过512字节
	_propertyAlias string
	// 重新关联商品与店铺类目，结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。
	_sellerCids string
	// 宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.postages.get获得当前会话用户的所有邮费模板）
	_postageId int64
	// 商家编码
	_outerId string
	// 商品所属的产品ID(B商家发布商品需要用)
	_productId int64
	// 商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path
	_picPath string
	// 代充商品类型。只有少数类目下的商品可以标记上此字段，具体哪些类目可以上传可以通过taobao.itemcat.features.get获得。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型： <br/>no_mark(不做类型标记) <br/>time_card(点卡软件代充) <br/>fee_card(话费软件代充)
	_autoFill string
	// 是否在淘宝上显示
	_isTaobao bool
	// 是否在外店显示
	_isEx bool
	// 是否是3D
	_is3D bool
	// 是否替换sku
	_isReplaceSku bool
	// 商品文字的版本，繁体传入”zh_HK”，简体传入”zh_CN”
	_lang string
	// 支持会员打折。可选值:true,false;
	_hasDiscount bool
	// 橱窗推荐。可选值:true,false;
	_hasShowcase bool
	// 商品上传后的状态。可选值:onsale（出售中）,instock（库中），如果同时更新商品状态为出售中及list_time为将来的时间，则商品还是处于定时上架的状态, 此时item.is_timing为true
	_approveStatus string
	// 运费承担方式。运费承担方式。可选值:seller（卖家承担）,buyer(买家承担);
	_freightPayer string
	// 有效期。可选值:7,14;单位:天;
	_validThru int64
	// 是否有发票。可选值:true,false (商城卖家此字段必须为true)
	_hasInvoice bool
	// 是否有保修。可选值:true,false;
	_hasWarranty bool
	// 是否承诺退换货服务!虚拟商品无须设置此项!
	_sellPromise bool
	// 货到付款运费模板ID
	_codPostageId int64
	// 实物闪电发货。注意：在售的闪电发货产品不允许取消闪电发货，需要先下架商品才能取消闪电发货标记
	_isLightningConsignment bool
	// 商品的重量(商超卖家专用字段)
	_weight int64
	// 商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置参数就保持原有值。
	_isXinpin bool
	// 商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改 集市卖家默认拍下减库存; 商城卖家默认付款减库存
	_subStock int64
	// 忽略警告提示.
	_ignorewarning string
	// 用户自行输入的类目属性ID串，结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。
	_inputPids string
	// 更新的Sku的数量串，结构如：num1,num2,num3 如:2,3,4
	_skuQuantities string
	// 更新的Sku的价格串，结构如：10.00,5.00,… 精确到2位小数;单位:元。如:200.07，表示:200元7分
	_skuPrices string
	// 更新的Sku的属性串，调用taobao.itemprops.get获取类目属性，如果属性是销售属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid;pid:vid。该字段内的销售属性也需要在props字段填写 。如果更新时有对Sku进行操作，则Sku的properties一定要传入。
	_skuProperties string
	// Sku的外部id串，结构如：1234,1342,… sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，如果没有sku_outer_ids也要写上这个参数，入参是","(这个是两个sku的示列，逗号数应该是sku个数减1)；该参数最大长度是512个字节
	_skuOuterIds string
	// 用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5,Nike乔丹鞋;乔丹系列;乔丹鞋系列;乔丹鞋系列;json5”，多个自定义属性用','分割，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过3999字节。此处不可以使用“其他”、“其它”和“其她”这三个词。
	_inputStr string
}

// New
