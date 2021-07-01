package alihealth2

// DentalOuterStoreRequest 结构体
type DentalOuterStoreRequest struct {
	// 门店主图
	Logo string `json:"logo,omitempty" xml:"logo,omitempty"`
	// 门店名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 城市码
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 门店描述
	StoreDesc string `json:"store_desc,omitempty" xml:"store_desc,omitempty"`
	// 医疗执业许可证名称
	LicenseName string `json:"license_name,omitempty" xml:"license_name,omitempty"`
	// 门店号码
	StorePhone string `json:"store_phone,omitempty" xml:"store_phone,omitempty"`
	// 经度
	PointX string `json:"point_x,omitempty" xml:"point_x,omitempty"`
	// 营业执照号
	LicenseNo string `json:"license_no,omitempty" xml:"license_no,omitempty"`
	// 纬度
	PointY string `json:"point_y,omitempty" xml:"point_y,omitempty"`
	// 门店地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 门店ID
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 外部门店code
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 关键字
	KeyWords string `json:"key_words,omitempty" xml:"key_words,omitempty"`
	// 营业时间
	WorkTime string `json:"work_time,omitempty" xml:"work_time,omitempty"`
	// 交通路线
	Routes string `json:"routes,omitempty" xml:"routes,omitempty"`
	// 标记图片
	SignPic string `json:"sign_pic,omitempty" xml:"sign_pic,omitempty"`
	// 营业执照图片
	LicensePics []string `json:"license_pics,omitempty" xml:"license_pics>string,omitempty"`
	// 医疗执业许可证图片
	MedicalPics []string `json:"medical_pics,omitempty" xml:"medical_pics>string,omitempty"`
	// 门店图片
	StorePics []string `json:"store_pics,omitempty" xml:"store_pics>string,omitempty"`
}
