package xhotelitem

// SRoomType 结构体
type SRoomType struct {
	// 房型名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 楼层
	Floor string `json:"floor,omitempty" xml:"floor,omitempty"`
	// 宽带服务&lt;br/&gt;&#34;0&#34;,&#34;有线上网(免费),&lt;br/&gt;&#34;1&#34;,&#34;有线上网(无)&#34;,&lt;br/&gt;&#34;2&#34;,&#34;有线上网(收费)&#34;,&lt;br/&gt;&#34;3&#34;,&#34;有线上网(部分有且免费)&#34;,&lt;br/&gt;&#34;4&#34;,&#34;有线上网(部分有且收费)&#34;
	Internet string `json:"internet,omitempty" xml:"internet,omitempty"`
	// pic_url
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// facility
	Facility string `json:"facility,omitempty" xml:"facility,omitempty"`
	// 面积
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 扩展字段
	Extend string `json:"extend,omitempty" xml:"extend,omitempty"`
	// 创建时间
	CreatedTime string `json:"created_time,omitempty" xml:"created_time,omitempty"`
	// 修改时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 窗型，枚举类型&lt;br/&gt;0, 无窗,&lt;br/&gt;1, 有窗;
	WindowType string `json:"window_type,omitempty" xml:"window_type,omitempty"`
	// 床型。json格式：[{&#34;bedType&#34;:&#34;大床&#34;,&#34;bedSize&#34;:&#34;1.5m&#34;},{&#34;bedType&#34;:&#34;双床&#34;,&#34;bedSize&#34;:&#34;1.2m&#34;}]
	Bed string `json:"bed,omitempty" xml:"bed,omitempty"`
	// 匹配的标准房型
	Srid int64 `json:"srid,omitempty" xml:"srid,omitempty"`
	// shid
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 最大入住人数
	MaxOccupancy int64 `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty"`
	// 状态。0:正常;-1:删除
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
