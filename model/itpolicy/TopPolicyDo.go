package itpolicy

// TopPolicyDo 
type TopPolicyDo struct {

    // 产品编号，内容可空，支持字母和数字,最多50个字符
    
    PolicyId   string `json:"policy_id,omitempty" xml:"policy_id,omitempty"`
    

    // 1/2RT计算方式，空：表示取严，可录入各取各或者取严，表示1/2RT佣金计算方式是各取各或者取严。
    
    RtCommissionFormula   string `json:"rt_commission_formula,omitempty" xml:"rt_commission_formula,omitempty"`
    

    // 航空公司，航空公司两字码，如CA
    
    Airline   string `json:"airline,omitempty" xml:"airline,omitempty"`
    

    // 中转类型（是否直达），空：表示不限，即直达中转都适用，可选：直达、中转
    
    TransferType   string `json:"transfer_type,omitempty" xml:"transfer_type,omitempty"`
    

    // 航程种类，可选：单程、往返、空为不限
    
    TravelType   string `json:"travel_type,omitempty" xml:"travel_type,omitempty"`
    

    // 始发地，空表示所有航线都适用可录入格式：1) 城市三代如SHA,NYC,SEL2) 国家二代如CN,US,KR3) TC区代码如TC1,TC2,TC34) 为空表示不限制允许1.2.3. 混合录入，可录入多个用，隔开表示多个最多允许录入100个多个用,分隔.可输入单个区域和多个城市,支持区域和城市同时输入，以自定义区域表为准，输入自定义名称，系统存入对应城市三字码集合最多输入100个城市
    
    DepCities   string `json:"dep_cities,omitempty" xml:"dep_cities,omitempty"`
    

    // 目的地，空表示所有航线都适用可录入格式：1) 城市三代如SHA,NYC,SEL2) 国家二代如CN,US,KR3) TC区代码如TC1,TC2,TC34) 为空表示不限制允许1.2.3. 混合录入，可录入多个用，隔开表示多个最多允许录入100个多个用,分隔.可输入单个区域和多个城市,支持区域和城市同时输入，以自定义区域表为准，输入自定义名称，系统存入对应城市三字码集合最多输入100个城市
    
    ArrCities   string `json:"arr_cities,omitempty" xml:"arr_cities,omitempty"`
    

    // 例外始发地，空表示所有航线都适用可录入格式：1) 城市三代如SHA,NYC,SEL2) 国家二代如CN,US,KR3) TC区代码如TC1,TC2,TC34) 为空表示不限制允许1.2.3. 混合录入，可录入多个用，隔开表示多个最多允许录入100个多个用,分隔.可输入单个区域和多个城市,支持区域和城市同时输入，以自定义区域表为准，输入自定义名称，系统存入对应城市三字码集合最多输入100个城市
    
    ExcludeDepCities   string `json:"exclude_dep_cities,omitempty" xml:"exclude_dep_cities,omitempty"`
    

    // 例外目的地，空表示所有航线都适用可录入格式：1) 城市三代如SHA,NYC,SEL2) 国家二代如CN,US,KR3) TC区代码如TC1,TC2,TC34) 为空表示不限制允许1.2.3. 混合录入，可录入多个用，隔开表示多个最多允许录入100个多个用,分隔.可输入单个区域和多个城市,支持区域和城市同时输入，以自定义区域表为准，输入自定义名称，系统存入对应城市三字码集合最多输入100个城市
    
    ExcludeArrCities   string `json:"exclude_arr_cities,omitempty" xml:"exclude_arr_cities,omitempty"`
    

    // 是否允许1/2RT组合销售规则，允许、不允许空表示：不允许
    
    IsSupportRt   string `json:"is_support_rt,omitempty" xml:"is_support_rt,omitempty"`
    

    // 中转点，空表示所有航线都适用可录入格式：1) 城市三代如SHA,NYC,SEL2) 国家二代如CN,US,KR3) TC区代码如TC1,TC2,TC34) 为空表示不限制允许1.2.3. 混合录入，可录入多个用，隔开表示多个最多允许录入100个多个用,分隔.可输入单个区域和多个城市,支持区域和城市同时输入，以自定义区域表为准，输入自定义名称，系统存入对应城市三字码集合最多输入100个城市
    
    TransferCities   string `json:"transfer_cities,omitempty" xml:"transfer_cities,omitempty"`
    

    // 备注信息，销售规则备注,最多300个字符
    
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    

    // 销售日期，必填项；销售日期只能输入一段;日期对为左右都是闭区间格式：可支持2013-01-01或2013/01/01格式输入；范围用~表示
    
    SaleDate   string `json:"sale_date,omitempty" xml:"sale_date,omitempty"`
    

    // 去程旅行日期，必填项；支持添加多对日期多段用,分隔；可支持2013-01-01或2013/01/01格式输入；范围用~表示
    
    DepDate   string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
    

    // 回程旅行日期，非必填，支持添加多对日期多段用,分隔；可支持2013-01-01或2013/01/01格式输入；范围用~表示
    
    RetDate   string `json:"ret_date,omitempty" xml:"ret_date,omitempty"`
    

    // 适用/例外舱位，可选：适用、例外，如果舱位有值，则必须输入适用或例外
    
    CabinRestrictType   string `json:"cabin_restrict_type,omitempty" xml:"cabin_restrict_type,omitempty"`
    

    // 舱位，空表示所有舱位都适用多个用,分隔
    
    CabinRestrict   string `json:"cabin_restrict,omitempty" xml:"cabin_restrict,omitempty"`
    

    // 服务等级，默认值空：表示所有服务等级适用头等公务超值经济经济可多选，用,分隔
    
    ServiceLevel   string `json:"service_level,omitempty" xml:"service_level,omitempty"`
    

    // 适用fareBasis，空表示所有都适用支持数字字母组合支持多个,支持通配符%不限制长度的通配符,只 支持首或末位一个%
    
    FareBasisAllowed   string `json:"fare_basis_allowed,omitempty" xml:"fare_basis_allowed,omitempty"`
    

    // 例外fareBasis，空表示所有都适用支持数字字母组合支持多个,支持通配符%不限制长度的通配符,只 支持首或末位一个%
    
    FareBasisForbidden   string `json:"fare_basis_forbidden,omitempty" xml:"fare_basis_forbidden,omitempty"`
    

    // 适用航班，同一航段之间用，隔开表示或的关系；不同航段之间用/隔开。                       1 CA001-999,CA3000-3999  表示CA001至999以及3000至3999之间航班号的航班2 MU  表示所有MU开头的航班 3 CA(LH\AZ) 表示CA开头的实际承运人为LH或AZ的航班4 CA(*)   表示CA代码共享航班/CA开头的实际承运人为其他航空公司的航班5 CA(CA)   表示CA自营航班/CA实际承运航班；6 CA(OZ)001-999 表示CA开头航班号为001-999之间且实际承运人为OZ的航班；7 为空表示无限制
    
    FlightRestrict   string `json:"flight_restrict,omitempty" xml:"flight_restrict,omitempty"`
    

    // 代码共享适用类型，空：不允许代码共享;可以选值：仅限同集团代码共享适用、代码共享适用、不允许代码共享、
    
    CodeSharingType   string `json:"code_sharing_type,omitempty" xml:"code_sharing_type,omitempty"`
    

    // 成人身份，1.不得为空2.可选：普通/学生 3多个使用,分隔
    
    PassengerType   string `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
    

    // 运价来源，空：表示平台运价; 此处为单选,输入平台运价，表示平台运价，不同的接入方可选值不一样
    
    FareSource   string `json:"fare_source,omitempty" xml:"fare_source,omitempty"`
    

    // 可适用运价渠道,可选值：公布运价,私有运价
    
    FareType   string `json:"fare_type,omitempty" xml:"fare_type,omitempty"`
    

    // 价格区间，空表示默认值为0-999999默认为正整型上限必须大于下限
    
    AllowPriceRange   string `json:"allow_price_range,omitempty" xml:"allow_price_range,omitempty"`
    

    // 预定OFFICE，空表示默认优先级最高OFFICE，可输入OFFICE，校验必须为配置中存在的OFFICE
    
    OfficeNo   string `json:"office_no,omitempty" xml:"office_no,omitempty"`
    

    // 返点，可输入负值范围-100至100之间，可保留2位小数与留钱固定金额二选一必输
    
    SaleRetention   string `json:"sale_retention,omitempty" xml:"sale_retention,omitempty"`
    

    // 留钱，返点和留钱至少二选一
    
    SaleRebase   string `json:"sale_rebase,omitempty" xml:"sale_rebase,omitempty"`
    

    // 儿童返点，非必输;可输入负值;范围-100至100之间，可保留2位小数
    
    ChildSaleRetention   string `json:"child_sale_retention,omitempty" xml:"child_sale_retention,omitempty"`
    

    // 儿童留钱，非必输;为整型;支持负数;单位元
    
    ChildSaleRebase   string `json:"child_sale_rebase,omitempty" xml:"child_sale_rebase,omitempty"`
    

    // （已废除字段）退票规定，非必输长度小于300字符请同时录入 退票规定、改签规定和行李额规定
    
    RefundRule   string `json:"refund_rule,omitempty" xml:"refund_rule,omitempty"`
    

    // （已废除字段）改签规定，非必输长度小于300字符请同时录入 退票规定、改签规定和行李额规定
    
    ReissueRule   string `json:"reissue_rule,omitempty" xml:"reissue_rule,omitempty"`
    

    // （已废除字段）误机罚金说明，非必输长度小于300字符只在退票规定不为空时才会生效
    
    NoshowRule   string `json:"noshow_rule,omitempty" xml:"noshow_rule,omitempty"`
    

    // 行李额规定，非必输长度小于300字符请同时录入 退票规定、改签规定和行李额规定
    
    LuggageRule   string `json:"luggage_rule,omitempty" xml:"luggage_rule,omitempty"`
    

    // 购票须知，非必输长度小于300字符只在退票规定不为空时才会生效
    
    BuyTicketNotice   string `json:"buy_ticket_notice,omitempty" xml:"buy_ticket_notice,omitempty"`
    

    // 商品类型，非必输；默认为普通可填写为金牌或普通
    
    ProductType   string `json:"product_type,omitempty" xml:"product_type,omitempty"`
    

    // 不同航司联运，非必输；可输入允许或不允许，空表示不允许
    
    IsAllowUnionAirline   string `json:"is_allow_union_airline,omitempty" xml:"is_allow_union_airline,omitempty"`
    

    // 非必输；09:00-18:00表示每一天的早上9点到下午6点，09:00MON-18:00FRI表示周一到周五的每天早上9点到下午6点最多录入三个时间段用逗号隔开表示或的关系可以为空，表示不限制，即工作时间为09:00-18:00
    
    WorkingTime   string `json:"working_time,omitempty" xml:"working_time,omitempty"`
    

    // 渠道名称，非必输，不同的接入方可选值不一样
    
    ChannelIdDesc   string `json:"channel_id_desc,omitempty" xml:"channel_id_desc,omitempty"`
    

    // 扩展字段，预留
    
    ExtendAttributes   string `json:"extend_attributes,omitempty" xml:"extend_attributes,omitempty"`
    

    // 全部未使用可否退票，可输入:是，否
    
    IsCanAllRefund   string `json:"is_can_all_refund,omitempty" xml:"is_can_all_refund,omitempty"`
    

    // 【全部未使用可否退票】为是时，此项为必填项。 可输入格式如：200-72-300-48-1000-0-*，表示72小时前退票手续费200；48小时到72小时，退票手续费300；飞机起飞不足48小时；退票手续费1000；飞机起飞后不予改期（输入*）； 或 10%-72-30%-48-50%-0-*，表示72小时前退票手续费为票面价的10%；48小时到72小时，退票手续费为票面价的30%；飞机起飞不足48小时；退票手续费为票面价的50%；飞机起飞后不予退票（输入*）；
    
    RefundFeeAllUnused   string `json:"refund_fee_all_unused,omitempty" xml:"refund_fee_all_unused,omitempty"`
    

    // 全部未使用退票币种，只能录入币种三字码，默认值CNY
    
    RefundCurrencyAllUnused   string `json:"refund_currency_all_unused,omitempty" xml:"refund_currency_all_unused,omitempty"`
    

    // 全部未使用退票费用收取方式，按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
    
    RefundFeeTypeAllUnused   string `json:"refund_fee_type_all_unused,omitempty" xml:"refund_fee_type_all_unused,omitempty"`
    

    // 部分未使用可否退票，可输入：是，否
    
    IsCanPartRefund   string `json:"is_can_part_refund,omitempty" xml:"is_can_part_refund,omitempty"`
    

    // 部分未使用退票费用,格式同【全部未使用退票费用】，【部分未使用可否退票】为是时，此项为必填项
    
    RefundFeePartUnused   string `json:"refund_fee_part_unused,omitempty" xml:"refund_fee_part_unused,omitempty"`
    

    // 部分未使用退票币种，可录入币种三字码，默认值CNY
    
    RefundCurrencyPartUnused   string `json:"refund_currency_part_unused,omitempty" xml:"refund_currency_part_unused,omitempty"`
    

    // 部分未使用退票费用收取方式，按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
    
    RefundFeeTypePartUnused   string `json:"refund_fee_type_part_unused,omitempty" xml:"refund_fee_type_part_unused,omitempty"`
    

    // 去程可否改期，可输入是或否
    
    CanDepChange   string `json:"can_dep_change,omitempty" xml:"can_dep_change,omitempty"`
    

    // 去程改期费用，格式说明参考【全程未使用退票费用】，注意，改期费用不能录入百分比，【去程可否改期】为是时为必填项
    
    DepChangeFee   string `json:"dep_change_fee,omitempty" xml:"dep_change_fee,omitempty"`
    

    // 去程改期币种，可录入币种三字码，默认值CNY
    
    DepChangeCurrency   string `json:"dep_change_currency,omitempty" xml:"dep_change_currency,omitempty"`
    

    // 去程改期费用收取方式,按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
    
    DepChangeFeeType   string `json:"dep_change_fee_type,omitempty" xml:"dep_change_fee_type,omitempty"`
    

    // 回程可否改期，可输入是或否
    
    CanRetChange   string `json:"can_ret_change,omitempty" xml:"can_ret_change,omitempty"`
    

    // 回程改期费用，格式同【去程改期费用】，【回程可否改期】为是时为必填
    
    RetChangeFee   string `json:"ret_change_fee,omitempty" xml:"ret_change_fee,omitempty"`
    

    // 回程改期币种，可录入币种三字码，默认值CNY
    
    RetChangeCurrency   string `json:"ret_change_currency,omitempty" xml:"ret_change_currency,omitempty"`
    

    // 回程改期费用收取方式，按每个航段收还是全程收(0:全程, 1：每个航段，默认值：全程)
    
    RetChangeFeeType   string `json:"ret_change_fee_type,omitempty" xml:"ret_change_fee_type,omitempty"`
    

    // NOSHOW是否有限制，可输入是或否
    
    NoshowRestrict   string `json:"noshow_restrict,omitempty" xml:"noshow_restrict,omitempty"`
    

    // NOSHOW时限,只能录入整数，【NOSHOW是否有限制】为是时，此项为必填项
    
    NoshowTimeRestrict   string `json:"noshow_time_restrict,omitempty" xml:"noshow_time_restrict,omitempty"`
    

    // NOSHOW时限单位(小时/天, 默认为小时)
    
    NoshowTimeRestrictUnit   string `json:"noshow_time_restrict_unit,omitempty" xml:"noshow_time_restrict_unit,omitempty"`
    

    // NOSHOW规则，可录入多个，多个用逗号分隔。可录入不可退票、不可改期，不可退票,不可改期
    
    NoshowRuleType   string `json:"noshow_rule_type,omitempty" xml:"noshow_rule_type,omitempty"`
    

    // NOSHOW金额，只能录入整数或百分比，【NOSHOW是否有限制】为是,【NOSHOW规则】不是不可退票,不可改期时，此项为必填项
    
    NoshowFee   string `json:"noshow_fee,omitempty" xml:"noshow_fee,omitempty"`
    

    // NOSHOW币种,可录入币种三字码，默认值CNY
    
    NoshowCurrency   string `json:"noshow_currency,omitempty" xml:"noshow_currency,omitempty"`
    

    // 开票大客户编码，最大长度50
    
    VipCode   string `json:"vip_code,omitempty" xml:"vip_code,omitempty"`
    

    // 国籍限制，可输入国家二字码，多个用逗号分隔，最多不超过20个
    
    Nationality   string `json:"nationality,omitempty" xml:"nationality,omitempty"`
    

    // 除外国籍限制，可输入国家二字码，多个用逗号分隔，最多不超过20个
    
    ExcludeNationality   string `json:"exclude_nationality,omitempty" xml:"exclude_nationality,omitempty"`
    

    // 年龄限制，可录入值范围12-99, 并且最低年龄不可超过70
    
    PassengerAge   string `json:"passenger_age,omitempty" xml:"passenger_age,omitempty"`
    

    // 是否适用小团，可选值是、否。空表示不限
    
    Gv2ChildRule   string `json:"gv2_child_rule,omitempty" xml:"gv2_child_rule,omitempty"`
    

    // 提前销售天数限制
    
    PresalePeriod   string `json:"presale_period,omitempty" xml:"presale_period,omitempty"`
    

    // 必填项 赋值范围: 电子行程单,旅行发票,差额行程单发票,等额行程单
    
    Receipts   string `json:"receipts,omitempty" xml:"receipts,omitempty"`
    

    // 供应来源，可为空
    
    SupplySource   string `json:"supply_source,omitempty" xml:"supply_source,omitempty"`
    

    // 是否支持缺口, 赋值范围:不支持缺口,只支持缺口,不限
    
    CanOj   string `json:"can_oj,omitempty" xml:"can_oj,omitempty"`
    

    // 例外航线,默认销售规则
    
    DefaultOdDeny   string `json:"default_od_deny,omitempty" xml:"default_od_deny,omitempty"`
    

    // 航司代码共享范围 airline_code_sharing : 格式 KA(CX/CX),CA(CZ)  英文逗号分开，/ 分隔内部，航司二字码表示
    
    AirlineCodeSharing   string `json:"airline_code_sharing,omitempty" xml:"airline_code_sharing,omitempty"`
    

    // 儿童回程改期费用，格式同成人，回程改期费用，【回程可否改期】为是时为必填
    
    ChildChangeFeeInUnused   string `json:"child_change_fee_in_unused,omitempty" xml:"child_change_fee_in_unused,omitempty"`
    

    // 儿童去程改期费用，格式同成人，格式说明参考【全程未使用退票费用】，注意，改期费用不能录入百分比，【去程可否改期】为是时为必填项
    
    ChildChangeFeeOutUnused   string `json:"child_change_fee_out_unused,omitempty" xml:"child_change_fee_out_unused,omitempty"`
    

    // 儿童部分未使用退票费用，格式同成人，【部分未使用可否退票】为是时，此项为必填项
    
    ChildRefundFeePartUnused   string `json:"child_refund_fee_part_unused,omitempty" xml:"child_refund_fee_part_unused,omitempty"`
    

    // 儿童全部未使用退票费用，格式同成人，【全部未使用可否退票】为是时，此项为必填项。 可输入格式如：200-72-300-48-1000-0-*，表示72小时前退票手续费200；48小时到72小时，退票手续费300；飞机起飞不足48小时；退票手续费1000；飞机起飞后不予改期（输入*）； 或 10%-72-30%-48-50%-0-*，表示72小时前退票手续费为票面价的10%；48小时到72小时，退票手续费为票面价的30%；飞机起飞不足48小时；退票手续费为票面价的50%；飞机起飞后不予退票（输入*）；
    
    ChildRefundFeeAllUnused   string `json:"child_refund_fee_all_unused,omitempty" xml:"child_refund_fee_all_unused,omitempty"`
    

}
