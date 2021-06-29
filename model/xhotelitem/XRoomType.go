package xhotelitem

// XRoomType 
type XRoomType struct {
    // 飞猪房源id
    Rid   int64 `json:"rid,omitempty" xml:"rid,omitempty"`
    // 飞猪门店id
    Hid   int64 `json:"hid,omitempty" xml:"hid,omitempty"`
    // 房源状态
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    // 匹配状态: 0：待系统匹配 1：已系统匹配，匹配成功，待卖家确认 2：已系统匹配，匹配失败，待人工匹配 3：已人工匹配，匹配成功，待卖家确认 4：已人工匹配，匹配失败 5：卖家已确认，确认&ldquo;YES&rdquo; 6：卖家已确认，确认&ldquo;NO&rdquo; 7:已系统匹配，但是匹配重复，待人工确认
    MatchStatus   int64 `json:"match_status,omitempty" xml:"match_status,omitempty"`
    // 卖家系统id
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    // 房型名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 客房在建筑的第几层，隔层为1-2层，4-5层，7-8层
    Floor   string `json:"floor,omitempty" xml:"floor,omitempty"`
    // 最大入住人数
    MaxOccupancy   int64 `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty"`
    // 面积
    Area   string `json:"area,omitempty" xml:"area,omitempty"`
    // 设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {&quot;bar&quot;:false,&quot;catv&quot;:false,&quot;ddd&quot;:false,&quot;idd&quot;:false,&quot;pubtoilet&quot;:false,&quot;toilet&quot;:false}
    Service   string `json:"service,omitempty" xml:"service,omitempty"`
    // 窗型,0：无窗/1：有窗
    WindowType   int64 `json:"window_type,omitempty" xml:"window_type,omitempty"`
    // 宽带服务。A,B,C,D。分别代表： A：无宽带，B：免费宽带，C：收费宽带，D：部分收费宽带
    Internet   string `json:"internet,omitempty" xml:"internet,omitempty"`
    // 创建时间
    CreatedTime   string `json:"created_time,omitempty" xml:"created_time,omitempty"`
    // 修改时间
    ModifiedTime   string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
    // 标准房型信息
    SRoomtype   *SRoomType `json:"s_roomtype,omitempty" xml:"s_roomtype,omitempty"`
    // 出错原因,没有匹配上标准房型时，小二拒绝的理由
    ErrorInfo   string `json:"error_info,omitempty" xml:"error_info,omitempty"`
    // 床型。按自己定义存储，比如：高低床、上下床
    BedType   string `json:"bed_type,omitempty" xml:"bed_type,omitempty"`
    // 床宽。
    BedSize   string `json:"bed_size,omitempty" xml:"bed_size,omitempty"`
    // 扩展信息的JSON。 注：此字段的值需要ISV在接入前与淘宝沟通，且确认能解析
    Extend   string `json:"extend,omitempty" xml:"extend,omitempty"`
    // 卖家房型英文名称
    NameE   string `json:"name_e,omitempty" xml:"name_e,omitempty"`
    // 酒店数据状态：匹配成功；待匹配；待审核；审核失败；疑似错误；请注意：只有状态为&ldquo;匹配成功&rdquo;才是正常状态。其他状态都不会上架商品。
    DataConfirmStr   string `json:"data_confirm_str,omitempty" xml:"data_confirm_str,omitempty"`
    // 房型维度下特殊标签含义 json: {"non-direct-roomType":1} , key:non-direct-roomType 表示非直连房型
    TagJson   string `json:"tag_json,omitempty" xml:"tag_json,omitempty"`
}
