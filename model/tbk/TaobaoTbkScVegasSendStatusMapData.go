package tbk

// TaobaotbkscvegassendstatusMapData 结构体
type TaobaotbkscvegassendstatusMapData struct {
	// 若该用户当前无待核销的红包，则返回1，若当前有待核销的红包，则返回0
	IsNewUser string `json:"is_new_user,omitempty" xml:"is_new_user,omitempty"`
}
