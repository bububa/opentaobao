package travel

// ProductBaseInfo 
type ProductBaseInfo struct {

    // 旅游日程-天数
    
    TripMaxDays   int64 `json:"trip_max_days,omitempty" xml:"trip_max_days,omitempty"`
    

    // pc端商品描述，不超过50000个字符。详情描述支持纯文本描述，也支持html格式的详情描述。html格式的详情描述中 图片链接支持外链图片（必须外网可访问， 且格式为png、jpg或jpeg，大小在500k以内）和淘宝图片空间链接。
    
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`
    

    // 手机端描述。格式： 使用xml标签（shortDesc，txt和img）进行图文混排，shortDesc标签表示添加小标题，txt标签表示添加文本描述，img标签表示添加一张图片。  注意： 1）shortDesc，txt和img三个xml标签独立，可以乱序使用，但不能嵌套； 2）图片链接支持外链图片（必须外网可访问，  且格式为png、jpg或jpeg，大小在500k以内）和淘宝图片空间链接；3）手机端描述内容以wapDesc标签标示开始和结束
    
    WapDesc   string `json:"wap_desc,omitempty" xml:"wap_desc,omitempty"`
    

    // 必填，产品类型。目前仅支持 6-国际邮轮  9-境内邮轮
    
    ItemType   int64 `json:"item_type,omitempty" xml:"item_type,omitempty"`
    

    // 目的地，所有商品必填，填写中文，以英文逗号分隔，最多可填12个，如果国家底下还有城市，则必须精确到城市。如果接口报错提示目的地不存在，则可登录商家基础信息映射管理后台（https://sell.alitrip.com/sell/basicdata/BasicDataMapping.htm）修正并生效映射关系，映射关系生效后，无法识别的目的地将自动替换为映射值 【重要提示-此字段会作为搜索/筛选条件】
    
    ToLocations   string `json:"to_locations,omitempty" xml:"to_locations,omitempty"`
    

    // 商品图片路径。最多支持5张，第一张为主图 必填，其余四张可选填（多张图片间使用英文逗号分隔）。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为png、jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。
    
    PicUrls   []string `json:"pic_urls,omitempty" xml:"pic_urls>string,omitempty"`
    

    // 旅游日程-晚数，必须大于等于0，且小于等于旅游日程-天数。 这里注意下，trip_max_days（旅游日程-天数）>=accom_nights（旅游日程-晚数）>=hotel_days（住宿晚数）
    
    AccomNights   int64 `json:"accom_nights,omitempty" xml:"accom_nights,omitempty"`
    

    // 必填，商品标题。30个字符以内  【重要-此字段会作为搜索条件】
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 商品亮点。1）目前最多支持4个亮点，超过4个的亮点描述不会被保存 2）每个亮点描述35个字符以内 3）每个亮点间用英文逗号分隔
    
    SubTitles   []string `json:"sub_titles,omitempty" xml:"sub_titles>string,omitempty"`
    

    // 可选，出行人模板id。模板id需要商家以店铺账号身份登录飞猪商家工作台，从卖家工具->出行人管理中获取。注意：如果传0则代表设置为不需要出行人模板或使用飞猪平台默认的类目模板。
    
    TravellerTemplateId   int64 `json:"traveller_template_id,omitempty" xml:"traveller_template_id,omitempty"`
    

    // 出发地（城市），填写中文，如果接口报错提示出发地不存在，则可登录商家基础信息映射管理后台（https://sell.alitrip.com/sell/basicdata/BasicDataMapping.htm）修正并生效映射关系，映射关系生效后，无法识别的出发地将自动替换为映射值 【重要提示-此属性会作为搜索/筛选条件】
    
    FromLocations   string `json:"from_locations,omitempty" xml:"from_locations,omitempty"`
    

    // 商家编码
    
    OutId   string `json:"out_id,omitempty" xml:"out_id,omitempty"`
    

}
