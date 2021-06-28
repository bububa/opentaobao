package travel

// BaseInfo 
/* model for simplify = false
type BaseInfo struct {

    // 用户指定的上下架操作：0-下架，1-上架（立即上架，定时上架）,2-从未上架。注意，已失效，默认下架状态。如需上架请调用taobao.alitrip.travel.item.shelve接口操作。
    
    ApproveStatus   int64 `json:"approve_status,omitempty"`
    

    // 商品类目id，发布商品必填，编辑选填；支持的线上类目ID： 123740001,   125038002,   50018298,   124084006,   125228016,   50011954,   50012913,   50214003,   50012917,   50134002,   50026091,   123742001,   50019817,   125210016,   124212017,   50025888,   50025831,   124142009,   123744001,   50012762,   50025880,   123166001,   50668002,   50024210,   50024208,   50024215,   50025878,   50024212,   123738001,   123164002,   50686003,   123164001,   124832008,   125408001,   50018112,   124258004,   50104001,   124730002,   124090010；支持的测试类目ID：  145732013,   145632040,   145632038,   145632039,   50018322,   145632036,   145632037,   145632034,   145632035,   145632032,   145632033,   145632031,   50019817,   146476005
    
    CategoryId   int64 `json:"category_id,omitempty"`
    

    // 宝贝所在市。不填默认设置 杭州
    
    City   string `json:"city,omitempty"`
    

    // 商品描述
    
    Desc   string `json:"desc,omitempty"`
    

    // 商品的补充描述，不同类目所设置值不一样，根据类目配置模版来决定传值；目前仅演艺类目使用，模版：145632032=showStart#演出开始时间#3#20@showEnd#演出结束时间#3#20
    
    ExtsMap   string `json:"exts_map,omitempty"`
    

    // 是否支持会员打折。可选值：true，false；默认值：false(不打折)。不传的话默认为false
    
    HasDiscount   bool `json:"has_discount,omitempty"`
    

    // 是否橱窗推荐，可选值：true，false；默认值：false(不推荐)
    
    HasShowcase   bool `json:"has_showcase,omitempty"`
    

    // 淘宝平台商品ID  商品更新时使用
    
    ItemId   int64 `json:"item_id,omitempty"`
    

    // 商品标签 提高曝光率
    
    Label   string `json:"label,omitempty"`
    

    // 上架时间。大于当前时间则宝贝会下架进入定时上架的宝贝中。
    
    OnlineTime   string `json:"online_time,omitempty"`
    

    // 商家编码
    
    OutId   string `json:"out_id,omitempty"`
    

    // 品图片路径。最多支持5张，第一张为主图 必填，其余四张可选填（多张图片间使用英文逗号分隔）。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为png、jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。
    
    PicUrls  struct {
        String  []string `json:"string,omitempty"`
    } `json:"pic_urls,omitempty"`
    

    // 商品属性列表；由类目的属性PID和VID组成，属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。
    
    Props  struct {
        CatPropInfo  []CatPropInfo `json:"cat_prop_info,omitempty"`
    } `json:"props,omitempty"`
    

    // 宝贝所在地省，后台通过用户输入城市去获取该城市对应省份
    
    Prov   string `json:"prov,omitempty"`
    

    // 商品秒杀，商品秒杀三个值：可选类型web_only(只能通过web网络秒杀)，wap_only(只能通过wap网络秒杀)，web_and_wap(既能通过web秒杀也能通过wap秒杀)
    
    SecondKill   string `json:"second_kill,omitempty"`
    

    // 宝贝店铺分类
    
    SellerCids  struct {
        String  []string `json:"string,omitempty"`
    } `json:"seller_cids,omitempty"`
    

    // 库存计数方式，默认采用拍下减库存，不支持用户选择。0拍下减库存，1付款减库存；
    
    SubStock   int64 `json:"sub_stock,omitempty"`
    

    // 商品亮点。1）目前最多支持4个亮点，超过4个的亮点描述不会被保存 2）每个亮点描述35个字符以内 3）每个亮点间用英文逗号分隔
    
    SubTitles  struct {
        String  []string `json:"string,omitempty"`
    } `json:"sub_titles,omitempty"`
    

    // 必填，商品标题。30个字符以内 【重要-此字段会作为搜索条件】
    
    Title   string `json:"title,omitempty"`
    

    // 手机描述 格式:标题 描述图片路径
    
    WapDesc   string `json:"wap_desc,omitempty"`
    

    // 出发地（城市），填写中文，如果接口报错提示出发地不存在，则可登录商家基础信息映射管理后台（https://sell.alitrip.com/sell/basicdata/BasicDataMapping.htm）修正并生效映射关系，映射关系生效后，无法识别的出发地将自动替换为映射值 【重要提示-此属性会作为搜索/筛选条件】
    
    FromLocations   string `json:"from_locations,omitempty"`
    

    // 目的地，填写中文，以英文逗号分隔，最多可填12个，如果国家底下还有城市，则必须精确到城市。如果接口报错提示目的地不存在，则可登录商家基础信息映射管理后台（https://sell.alitrip.com/sell/basicdata/BasicDataMapping.htm）修正并生效映射关系，映射关系生效后，无法识别的目的地将自动替换为映射值 【重要提示-此字段会作为搜索/筛选条件】
    
    ToLocations   string `json:"to_locations,omitempty"`
    

    // 可选，资源确认类型。1-即时确认，2-二次确认。不传默认1
    
    ConfirmType   int64 `json:"confirm_type,omitempty"`
    

    // 可选，资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认
    
    ConfirmTime   int64 `json:"confirm_time,omitempty"`
    

    // 至少提前天数，最晚成团提前天数，最小0天，最大为300天；不传则为0
    
    Duration   int64 `json:"duration,omitempty"`
    

    // 最晚收客时间:小时。仅个别类目支持
    
    ReserveDeadlineHours   int64 `json:"reserve_deadline_hours,omitempty"`
    

    // 最晚收客时间:分钟。仅个别类目支持
    
    ReserveDeadlineMinutes   int64 `json:"reserve_deadline_minutes,omitempty"`
    

}
*/

// BaseInfo 
type BaseInfo struct {

    // 用户指定的上下架操作：0-下架，1-上架（立即上架，定时上架）,2-从未上架。注意，已失效，默认下架状态。如需上架请调用taobao.alitrip.travel.item.shelve接口操作。
    ApproveStatus   int64 `json:"approve_status,omitempty"`

    // 商品类目id，发布商品必填，编辑选填；支持的线上类目ID： 123740001,   125038002,   50018298,   124084006,   125228016,   50011954,   50012913,   50214003,   50012917,   50134002,   50026091,   123742001,   50019817,   125210016,   124212017,   50025888,   50025831,   124142009,   123744001,   50012762,   50025880,   123166001,   50668002,   50024210,   50024208,   50024215,   50025878,   50024212,   123738001,   123164002,   50686003,   123164001,   124832008,   125408001,   50018112,   124258004,   50104001,   124730002,   124090010；支持的测试类目ID：  145732013,   145632040,   145632038,   145632039,   50018322,   145632036,   145632037,   145632034,   145632035,   145632032,   145632033,   145632031,   50019817,   146476005
    CategoryId   int64 `json:"category_id,omitempty"`

    // 宝贝所在市。不填默认设置 杭州
    City   string `json:"city,omitempty"`

    // 商品描述
    Desc   string `json:"desc,omitempty"`

    // 商品的补充描述，不同类目所设置值不一样，根据类目配置模版来决定传值；目前仅演艺类目使用，模版：145632032=showStart#演出开始时间#3#20@showEnd#演出结束时间#3#20
    ExtsMap   string `json:"exts_map,omitempty"`

    // 是否支持会员打折。可选值：true，false；默认值：false(不打折)。不传的话默认为false
    HasDiscount   bool `json:"has_discount,omitempty"`

    // 是否橱窗推荐，可选值：true，false；默认值：false(不推荐)
    HasShowcase   bool `json:"has_showcase,omitempty"`

    // 淘宝平台商品ID  商品更新时使用
    ItemId   int64 `json:"item_id,omitempty"`

    // 商品标签 提高曝光率
    Label   string `json:"label,omitempty"`

    // 上架时间。大于当前时间则宝贝会下架进入定时上架的宝贝中。
    OnlineTime   string `json:"online_time,omitempty"`

    // 商家编码
    OutId   string `json:"out_id,omitempty"`

    // 品图片路径。最多支持5张，第一张为主图 必填，其余四张可选填（多张图片间使用英文逗号分隔）。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为png、jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。
    PicUrls   []string `json:"pic_urls,omitempty"`

    // 商品属性列表；由类目的属性PID和VID组成，属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。
    Props   []CatPropInfo `json:"props,omitempty"`

    // 宝贝所在地省，后台通过用户输入城市去获取该城市对应省份
    Prov   string `json:"prov,omitempty"`

    // 商品秒杀，商品秒杀三个值：可选类型web_only(只能通过web网络秒杀)，wap_only(只能通过wap网络秒杀)，web_and_wap(既能通过web秒杀也能通过wap秒杀)
    SecondKill   string `json:"second_kill,omitempty"`

    // 宝贝店铺分类
    SellerCids   []string `json:"seller_cids,omitempty"`

    // 库存计数方式，默认采用拍下减库存，不支持用户选择。0拍下减库存，1付款减库存；
    SubStock   int64 `json:"sub_stock,omitempty"`

    // 商品亮点。1）目前最多支持4个亮点，超过4个的亮点描述不会被保存 2）每个亮点描述35个字符以内 3）每个亮点间用英文逗号分隔
    SubTitles   []string `json:"sub_titles,omitempty"`

    // 必填，商品标题。30个字符以内 【重要-此字段会作为搜索条件】
    Title   string `json:"title,omitempty"`

    // 手机描述 格式:标题 描述图片路径
    WapDesc   string `json:"wap_desc,omitempty"`

    // 出发地（城市），填写中文，如果接口报错提示出发地不存在，则可登录商家基础信息映射管理后台（https://sell.alitrip.com/sell/basicdata/BasicDataMapping.htm）修正并生效映射关系，映射关系生效后，无法识别的出发地将自动替换为映射值 【重要提示-此属性会作为搜索/筛选条件】
    FromLocations   string `json:"from_locations,omitempty"`

    // 目的地，填写中文，以英文逗号分隔，最多可填12个，如果国家底下还有城市，则必须精确到城市。如果接口报错提示目的地不存在，则可登录商家基础信息映射管理后台（https://sell.alitrip.com/sell/basicdata/BasicDataMapping.htm）修正并生效映射关系，映射关系生效后，无法识别的目的地将自动替换为映射值 【重要提示-此字段会作为搜索/筛选条件】
    ToLocations   string `json:"to_locations,omitempty"`

    // 可选，资源确认类型。1-即时确认，2-二次确认。不传默认1
    ConfirmType   int64 `json:"confirm_type,omitempty"`

    // 可选，资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认
    ConfirmTime   int64 `json:"confirm_time,omitempty"`

    // 至少提前天数，最晚成团提前天数，最小0天，最大为300天；不传则为0
    Duration   int64 `json:"duration,omitempty"`

    // 最晚收客时间:小时。仅个别类目支持
    ReserveDeadlineHours   int64 `json:"reserve_deadline_hours,omitempty"`

    // 最晚收客时间:分钟。仅个别类目支持
    ReserveDeadlineMinutes   int64 `json:"reserve_deadline_minutes,omitempty"`

}
