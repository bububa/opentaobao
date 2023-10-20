package tbk

import (
	"sync"
)

// TaobaoTbkScTpwdConvertMapData 结构体
type TaobaoTbkScTpwdConvertMapData struct {
	// 商品Id
	NumIid string `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
	// 商品淘客转链链接
	ClickUrl string `json:"click_url,omitempty" xml:"click_url,omitempty"`
	// 店铺卖家ID
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 入参淘口令对应原始链接
	OriginUrl string `json:"origin_url,omitempty" xml:"origin_url,omitempty"`
	// 入参淘口令推广链接中的pid，如果不属于当前调用的推广者则展示“0”
	OriginPid string `json:"origin_pid,omitempty" xml:"origin_pid,omitempty"`
	// 1-动态ID转链场景，2-消费者比价场景，3-商品库导购场景
	BizSceneId string `json:"biz_scene_id,omitempty" xml:"biz_scene_id,omitempty"`
	// 短口令
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 长口令
	Password string `json:"password,omitempty" xml:"password,omitempty"`
	// 短链接
	ShortUrl string `json:"short_url,omitempty" xml:"short_url,omitempty"`
	// 1-单品，2-店铺，3-活动，0-其它
	UrlType string `json:"url_type,omitempty" xml:"url_type,omitempty"`
}

var poolTaobaoTbkScTpwdConvertMapData = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScTpwdConvertMapData)
	},
}

// GetTaobaoTbkScTpwdConvertMapData() 从对象池中获取TaobaoTbkScTpwdConvertMapData
func GetTaobaoTbkScTpwdConvertMapData() *TaobaoTbkScTpwdConvertMapData {
	return poolTaobaoTbkScTpwdConvertMapData.Get().(*TaobaoTbkScTpwdConvertMapData)
}

// ReleaseTaobaoTbkScTpwdConvertMapData 释放TaobaoTbkScTpwdConvertMapData
func ReleaseTaobaoTbkScTpwdConvertMapData(v *TaobaoTbkScTpwdConvertMapData) {
	v.NumIid = ""
	v.ClickUrl = ""
	v.SellerId = ""
	v.OriginUrl = ""
	v.OriginPid = ""
	v.BizSceneId = ""
	v.Model = ""
	v.Password = ""
	v.ShortUrl = ""
	v.UrlType = ""
	poolTaobaoTbkScTpwdConvertMapData.Put(v)
}
