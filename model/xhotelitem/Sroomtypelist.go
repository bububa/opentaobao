package xhotelitem

// Sroomtypelist 
type Sroomtypelist struct {

    // 窗型
    
    WindowType   string `json:"window_type,omitempty" xml:"window_type,omitempty"`
    

    // 酒店图片结构化信息
    
    PicsInfo   string `json:"pics_info,omitempty" xml:"pics_info,omitempty"`
    

    // 宽带服务
    
    NetworkService   string `json:"network_service,omitempty" xml:"network_service,omitempty"`
    

    // 最后变更人
    
    LastModify   string `json:"last_modify,omitempty" xml:"last_modify,omitempty"`
    

    // includeTypes
    
    IncludeTypes   string `json:"include_types,omitempty" xml:"include_types,omitempty"`
    

    // 创建人
    
    Auditor   string `json:"auditor,omitempty" xml:"auditor,omitempty"`
    

    // 标准酒店ID
    
    Shid   int64 `json:"shid,omitempty" xml:"shid,omitempty"`
    

    // 房型原始图片
    
    OriginalPics   string `json:"original_pics,omitempty" xml:"original_pics,omitempty"`
    

    // 版本号
    
    Version   int64 `json:"version,omitempty" xml:"version,omitempty"`
    

    // 标准房型ID
    
    Srid   int64 `json:"srid,omitempty" xml:"srid,omitempty"`
    

    // 房型图片
    
    Pics   string `json:"pics,omitempty" xml:"pics,omitempty"`
    

    // 图片扩展字段
    
    PicsExt   string `json:"pics_ext,omitempty" xml:"pics_ext,omitempty"`
    

    // 面积
    
    Area   string `json:"area,omitempty" xml:"area,omitempty"`
    

    // 设施
    
    Facility   string `json:"facility,omitempty" xml:"facility,omitempty"`
    

    // 标准房型名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // propertiesDOs
    
    PropertiesDOs   string `json:"properties_d_os,omitempty" xml:"properties_d_os,omitempty"`
    

    // 是否可加床
    
    AddBed   int64 `json:"add_bed,omitempty" xml:"add_bed,omitempty"`
    

    // 标准房型的英文名
    
    NameE   string `json:"name_e,omitempty" xml:"name_e,omitempty"`
    

    // 状态
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

    // 扩展信息
    
    Extend   string `json:"extend,omitempty" xml:"extend,omitempty"`
    

    // 床型
    
    Bed   string `json:"bed,omitempty" xml:"bed,omitempty"`
    

    // transferPics
    
    TransferPics   string `json:"transfer_pics,omitempty" xml:"transfer_pics,omitempty"`
    

    // 楼层
    
    Floor   string `json:"floor,omitempty" xml:"floor,omitempty"`
    

    // 来源
    
    Source   int64 `json:"source,omitempty" xml:"source,omitempty"`
    

    // bedList
    
    BedList   string `json:"bed_list,omitempty" xml:"bed_list,omitempty"`
    

    // 最大入住人数
    
    MaxOccupancy   int64 `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty"`
    

    // 外部id
    
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    

}
