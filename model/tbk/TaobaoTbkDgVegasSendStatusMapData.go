package tbk

// TaobaotbkdgvegassendstatusMapData 结构体
type TaobaotbkdgvegassendstatusMapData struct {
	// 若该用户当前无待核销的红包，则返回1，若当前有待核销的红包，则返回0
	Isnewuser string `json:"is_new_user,omitempty" xml:"is_new_user,omitempty"`
}
