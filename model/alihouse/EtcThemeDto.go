package alihouse

import (
	"sync"
)

// EtcThemeDto 结构体
type EtcThemeDto struct {
	// 手动导入的商品信息
	ThemeDetailExcels []EtcThemeDetailExcelDto `json:"theme_detail_excels,omitempty" xml:"theme_detail_excels>etc_theme_detail_excel_dto,omitempty"`
	// 外部主键
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 副标题
	Descr string `json:"descr,omitempty" xml:"descr,omitempty"`
	// 城市编码,逗号隔开
	CityCodes string `json:"city_codes,omitempty" xml:"city_codes,omitempty"`
	// 详情头图
	DetailImage string `json:"detail_image,omitempty" xml:"detail_image,omitempty"`
	// 是否在头图展示标题与副标题 1：不展示 2：展示 �
	DetailShowTitle string `json:"detail_show_title,omitempty" xml:"detail_show_title,omitempty"`
	// 详情色值
	DetailColorVal string `json:"detail_color_val,omitempty" xml:"detail_color_val,omitempty"`
	// 业务类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 数据来源
	DetailDataSource int64 `json:"detail_data_source,omitempty" xml:"detail_data_source,omitempty"`
	// 是否展示筛选器 �1展示 0不展示
	DetailSelectorIsShow int64 `json:"detail_selector_is_show,omitempty" xml:"detail_selector_is_show,omitempty"`
	// 主题模板id（cardType为主题详情卡片） �
	DetailThemeTemplateId int64 `json:"detail_theme_template_id,omitempty" xml:"detail_theme_template_id,omitempty"`
	// 是否展示留资组件（0: 否 1：是）
	DetailShowFetchInfoComponent int64 `json:"detail_show_fetch_info_component,omitempty" xml:"detail_show_fetch_info_component,omitempty"`
	// 分类(1:价格，2：面积，3:户型 4: 居住体验，5:人群，6:销售状态，7:优选好房，8:周边配套，9:价格+面积，10:价格+户型，11:面积+户型，12:榜单城市化，13:楼盘时刻，14:楼盘属性，15:需求类型，16:权益，17:城市化)
	Category int64 `json:"category,omitempty" xml:"category,omitempty"`
	// 主题最小内容数量。内容大于等于此值的主题才会被透出到场景中
	MinContentNum int64 `json:"min_content_num,omitempty" xml:"min_content_num,omitempty"`
	// 选品ID
	SelectionId int64 `json:"selection_id,omitempty" xml:"selection_id,omitempty"`
	// 云主题id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 上下架状态（1上架，2下架）
	OnlineStatus int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
	// 是否删除(0:未删除 1:已删除)
	IsDeleted int64 `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
}

var poolEtcThemeDto = sync.Pool{
	New: func() any {
		return new(EtcThemeDto)
	},
}

// GetEtcThemeDto() 从对象池中获取EtcThemeDto
func GetEtcThemeDto() *EtcThemeDto {
	return poolEtcThemeDto.Get().(*EtcThemeDto)
}

// ReleaseEtcThemeDto 释放EtcThemeDto
func ReleaseEtcThemeDto(v *EtcThemeDto) {
	v.ThemeDetailExcels = v.ThemeDetailExcels[:0]
	v.OuterId = ""
	v.Name = ""
	v.Title = ""
	v.Descr = ""
	v.CityCodes = ""
	v.DetailImage = ""
	v.DetailShowTitle = ""
	v.DetailColorVal = ""
	v.Type = 0
	v.DetailDataSource = 0
	v.DetailSelectorIsShow = 0
	v.DetailThemeTemplateId = 0
	v.DetailShowFetchInfoComponent = 0
	v.Category = 0
	v.MinContentNum = 0
	v.SelectionId = 0
	v.Id = 0
	v.OnlineStatus = 0
	v.IsDeleted = 0
	poolEtcThemeDto.Put(v)
}
